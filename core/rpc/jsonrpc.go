package rpc

import (
	"encoding/hex"
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Server struct {
	addr    string
	log     *log.Logger
	chainID string

	block atomic.Uint64
	evm   *EvmMock

	httpSrv *http.Server
	ln      net.Listener
}

func New(addr, chainID string, l *log.Logger) *Server {
	if l == nil {
		l = log.New(ioDiscard{}, "[rpc] ", log.LstdFlags)
	}
	s := &Server{
		addr:    addr,
		log:     l,
		chainID: chainID,
		evm:     NewEvmMock(),
	}
	s.block.Store(0)
	return s
}

func (s *Server) Start(ctx context.Context) error {
	if strings.TrimSpace(s.addr) == "" {
		return errors.New("rpc: empty addr")
	}

	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("rpc: listen failed: %w", err)
	}
	s.ln = ln

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handleJSONRPC)

	s.httpSrv = &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Temporary dev block ticker.
	go func() {
		t := time.NewTicker(2 * time.Second)
		defer t.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				s.block.Add(1)
			}
		}
	}()

	go func() {
		s.log.Println("listening on", s.addr)
		if err := s.httpSrv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Println("serve error:", err)
		}
	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if s.httpSrv == nil {
		return nil
	}
	return s.httpSrv.Shutdown(ctx)
}

// ---- JSON-RPC types ----

type rpcReq struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

type rpcResp struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Result  any             `json:"result,omitempty"`
	Error   *rpcError       `json:"error,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ---- HTTP handler (single + batch) ----

func (s *Server) handleJSONRPC(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioReadAllLimit(r.Body, 2<<20)
	if err != nil {
		s.writeErrorRaw(w, nil, -32700, "parse error")
		return
	}
	body = bytes.TrimSpace(body)
	if len(body) == 0 {
		s.writeErrorRaw(w, nil, -32700, "parse error")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Batch: [...]
	if body[0] == '[' {
		var reqs []rpcReq
		if err := json.Unmarshal(body, &reqs); err != nil {
			s.writeErrorRaw(w, nil, -32700, "parse error")
			return
		}
		resps := make([]rpcResp, 0, len(reqs))
		for i := range reqs {
			resps = append(resps, s.dispatch(&reqs[i]))
		}
		_ = json.NewEncoder(w).Encode(resps)
		return
	}

	// Single: {...}
	var req rpcReq
	if err := json.Unmarshal(body, &req); err != nil {
		s.writeErrorRaw(w, nil, -32700, "parse error")
		return
	}
	resp := s.dispatch(&req)
	_ = json.NewEncoder(w).Encode(resp)
}

func (s *Server) dispatch(req *rpcReq) rpcResp {
	resp := rpcResp{JSONRPC: "2.0", ID: req.ID}

	if req.JSONRPC != "2.0" {
		resp.Error = &rpcError{Code: -32600, Message: "invalid jsonrpc version"}
		return resp
	}

	switch req.Method {

	// ---- base ----
	case "web3_clientVersion":
		resp.Result = "noorcore/2.1 (minimal-jsonrpc)"
		return resp
	case "eth_chainId":
		resp.Result = chainIDToHex(s.chainID)
		return resp
	case "net_version":
		resp.Result = chainIDToNetVersion(s.chainID)
		return resp
	case "eth_blockNumber":
		resp.Result = toHexUint(s.block.Load())
		return resp

	// ---- minimal EVM tooling surface (dev-only mock) ----
	case "eth_accounts":
		resp.Result = s.evm.Accounts()
		return resp

	case "eth_getTransactionCount":
		var params []any
		if err := json.Unmarshal(req.Params, &params); err != nil || len(params) < 1 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp
		}
		addrStr, _ := params[0].(string)
		if !common.IsHexAddress(addrStr) {
			resp.Error = &rpcError{Code: -32602, Message: "invalid address"}
			return resp
		}
		addr := common.HexToAddress(addrStr)
		resp.Result = toHexUint(s.evm.GetTransactionCount(addr))
		return resp

	case "eth_gasPrice":
		resp.Result = "0x1"
		return resp

	case "eth_estimateGas":
		resp.Result = "0x2dc6c0" // 3,000,000
		return resp

	case "eth_getBalance":
		resp.Result = "0x0"
		return resp

	case "eth_call":
		// Minimal eth_call support for PoSS view methods (dev-only).
		// params: [ {to: "0x..", data: "0x.."}, "latest" ]
		var params []any
		if err := json.Unmarshal(req.Params, &params); err != nil || len(params) < 1 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp
		}
		callObj, ok := params[0].(map[string]any)
		if !ok {
			resp.Error = &rpcError{Code: -32602, Message: "invalid call object"}
			return resp
		}
		toStr, _ := callObj["to"].(string)
		dataStr, _ := callObj["data"].(string)
		to := common.HexToAddress(toStr)
		dataStr = strings.TrimPrefix(strings.TrimSpace(dataStr), "0x")
		data, err := hex.DecodeString(dataStr)
		if err != nil {
			resp.Error = &rpcError{Code: -32602, Message: "invalid data hex"}
			return resp
		}

		if out, ok := s.evm.Call(to, data); ok {
			resp.Result = "0x" + hex.EncodeToString(out)
			return resp
		}

		// unknown call => empty result
		resp.Result = "0x"
		return resp

	case "eth_sendRawTransaction":
		var params []string
		if err := json.Unmarshal(req.Params, &params); err != nil || len(params) < 1 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp
		}
		raw := params[0]
		h, err := s.evm.SendRawTransaction(raw, chainIDToBigInt(s.chainID), s.block.Load())
		if err != nil {
			resp.Error = &rpcError{Code: -32000, Message: err.Error()}
			return resp
		}
		resp.Result = h.Hex()
		return resp

	case "eth_getTransactionReceipt":
		var params []string
		if err := json.Unmarshal(req.Params, &params); err != nil || len(params) < 1 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp
		}
		hashStr := params[0]
		if !strings.HasPrefix(hashStr, "0x") || len(hashStr) != 66 {
			resp.Result = nil
			return resp
		}
		rcpt := s.evm.GetTransactionReceipt(common.HexToHash(hashStr))
		if rcpt == nil {
			resp.Result = nil
			return resp
		}
		resp.Result = rcpt
		return resp

	case "eth_getTransactionByHash":
		// Needed by ethers.js polling (waitForTransaction)
		var params []string
		if err := json.Unmarshal(req.Params, &params); err != nil || len(params) < 1 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp
		}
		hashStr := params[0]
		if !strings.HasPrefix(hashStr, "0x") || len(hashStr) != 66 {
			resp.Result = nil
			return resp
		}
		h := common.HexToHash(hashStr)

		// Lookup tx in dev mock store
		s.evm.mu.Lock()
		tx := s.evm.txs[h]
		rcpt := s.evm.receipts[h]
		s.evm.mu.Unlock()

		if tx == nil {
			resp.Result = nil
			return resp
		}

		chainBig := chainIDToBigInt(s.chainID)
		signer := types.LatestSignerForChainID(chainBig)
		from, _ := types.Sender(signer, tx)

		// Signature values (r,s,v)
		v, rSig, sSig := tx.RawSignatureValues()

		// Optional mined fields if we have a receipt
		var blockHash any = nil
		var blockNumber any = nil
		var txIndex any = nil
		if rcpt != nil {
			blockHash = rcpt.BlockHash.Hex()
			blockNumber = rcpt.BlockNumber
			txIndex = "0x0"
		}

		// Legacy gasPrice or EIP-1559 fields
		gasPrice := "0x0"
		maxFee := "0x0"
		maxPrio := "0x0"
		if tx.Type() == 2 {
			// dynamic fee
			if tx.GasFeeCap() != nil {
				maxFee = "0x" + tx.GasFeeCap().Text(16)
			}
			if tx.GasTipCap() != nil {
				maxPrio = "0x" + tx.GasTipCap().Text(16)
			}
			// provide something for gasPrice too
			gasPrice = maxFee
		} else {
			if tx.GasPrice() != nil {
				gasPrice = "0x" + tx.GasPrice().Text(16)
			}
		}

		// Build ethers-compatible tx object
		type txResp struct {
			Hash                 string `json:"hash"`
			Nonce                string `json:"nonce"`
			BlockHash            any    `json:"blockHash"`
			BlockNumber          any    `json:"blockNumber"`
			TransactionIndex     any    `json:"transactionIndex"`
			From                 string `json:"from"`
			To                   any    `json:"to"`
			Value                string `json:"value"`
			Gas                  string `json:"gas"`
			GasPrice             string `json:"gasPrice"`
			MaxFeePerGas         string `json:"maxFeePerGas"`
			MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas"`
			Input                string `json:"input"`
			Type                 string `json:"type"`
			ChainId              string `json:"chainId"`
			V                    string `json:"v"`
			R                    string `json:"r"`
			S                    string `json:"s"`
		}

		var to any = nil
		if tx.To() != nil {
			to = tx.To().Hex()
		}

		resp.Result = txResp{
			Hash:                 tx.Hash().Hex(),
			Nonce:                toHexUint(tx.Nonce()),
			BlockHash:            blockHash,
			BlockNumber:          blockNumber,
			TransactionIndex:     txIndex,
			From:                 from.Hex(),
			To:                   to,
			Value:                "0x" + tx.Value().Text(16),
			Gas:                  toHexUint(tx.Gas()),
			GasPrice:             gasPrice,
			MaxFeePerGas:         maxFee,
			MaxPriorityFeePerGas: maxPrio,
			Input:                "0x" + common.Bytes2Hex(tx.Data()),
			Type:                 toHexUint(uint64(tx.Type())),
			ChainId:              "0x" + chainBig.Text(16),
			V:                    "0x" + v.Text(16),
			R:                    "0x" + rSig.Text(16),
			S:                    "0x" + sSig.Text(16),
		}
		return resp

	case "eth_getBlockByNumber":
		// Ethers.js expects many fields to be present. Provide a dev-compatible block shape.
		type blockResp struct {
			Number           string `json:"number"`
			Hash             string `json:"hash"`
			ParentHash       string `json:"parentHash"`
			Nonce            string `json:"nonce"`
			Sha3Uncles       string `json:"sha3Uncles"`
			LogsBloom        string `json:"logsBloom"`
			TransactionsRoot string `json:"transactionsRoot"`
			StateRoot        string `json:"stateRoot"`
			ReceiptsRoot     string `json:"receiptsRoot"`
			Miner            string `json:"miner"`
			Difficulty       string `json:"difficulty"`
			TotalDifficulty  string `json:"totalDifficulty"`
			ExtraData        string `json:"extraData"`
			Size             string `json:"size"`
			GasLimit         string `json:"gasLimit"`
			GasUsed          string `json:"gasUsed"`
			Timestamp        string `json:"timestamp"`
			MixHash          string `json:"mixHash"`
			BaseFeePerGas    string `json:"baseFeePerGas"`
			Transactions     []any  `json:"transactions"`
			Uncles           []any  `json:"uncles"`
		}

		n := s.block.Load()
		resp.Result = blockResp{
			Number:           toHexUint(n),
			Hash:             pseudoBlockHash(n).Hex(),
			ParentHash:       pseudoBlockHash(n - 1).Hex(),
			Nonce:            "0x0000000000000000",
			Sha3Uncles:       "0x" + strings.Repeat("0", 64),
			LogsBloom:        "0x" + strings.Repeat("0", 512),
			TransactionsRoot: "0x" + strings.Repeat("0", 64),
			StateRoot:        "0x" + strings.Repeat("0", 64),
			ReceiptsRoot:     "0x" + strings.Repeat("0", 64),
			Miner:            "0x" + strings.Repeat("0", 40),
			Difficulty:       "0x0",
			TotalDifficulty:  "0x0",
			ExtraData:        "0x",
			Size:             "0x0",
			GasLimit:         "0x1c9c380", // 30,000,000
			GasUsed:          "0x0",
			Timestamp:        toHexUint(uint64(time.Now().Unix())),
			MixHash:          "0x" + strings.Repeat("0", 64),
			BaseFeePerGas:    "0x1",
			Transactions:     []any{},
			Uncles:           []any{},
		}
		return resp

	default:
		resp.Error = &rpcError{Code: -32601, Message: "method not found"}
		return resp
	}
}

func (s *Server) writeErrorRaw(w http.ResponseWriter, id json.RawMessage, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	resp := rpcResp{JSONRPC: "2.0", ID: id, Error: &rpcError{Code: code, Message: msg}}
	_ = json.NewEncoder(w).Encode(resp)
}

// ---- helpers ----

func toHexUint(v uint64) string {
	return "0x" + strconv.FormatUint(v, 16)
}

func chainIDToHex(chainID string) string {
	if n, err := strconv.ParseUint(chainID, 10, 64); err == nil {
		return toHexUint(n)
	}
	sum := sha256.Sum256([]byte(chainID))
	u := uint64(sum[0])<<56 | uint64(sum[1])<<48 | uint64(sum[2])<<40 | uint64(sum[3])<<32 |
		uint64(sum[4])<<24 | uint64(sum[5])<<16 | uint64(sum[6])<<8 | uint64(sum[7])
	if u == 0 {
		u = 1
	}
	return toHexUint(u)
}

func chainIDToNetVersion(chainID string) string {
	if _, err := strconv.ParseUint(chainID, 10, 64); err == nil {
		return chainID
	}
	sum := sha256.Sum256([]byte(chainID))
	u := uint32(sum[0])<<24 | uint32(sum[1])<<16 | uint32(sum[2])<<8 | uint32(sum[3])
	if u == 0 {
		u = 1
	}
	return strconv.FormatUint(uint64(u), 10)
}

func chainIDToBigInt(chainID string) *big.Int {
	if n, err := strconv.ParseUint(chainID, 10, 64); err == nil {
		return new(big.Int).SetUint64(n)
	}
	sum := sha256.Sum256([]byte(chainID))
	u := uint64(sum[0])<<56 | uint64(sum[1])<<48 | uint64(sum[2])<<40 | uint64(sum[3])<<32 |
		uint64(sum[4])<<24 | uint64(sum[5])<<16 | uint64(sum[6])<<8 | uint64(sum[7])
	if u == 0 {
		u = 1
	}
	return new(big.Int).SetUint64(u)
}

func ioReadAllLimit(rc io.ReadCloser, limit int64) ([]byte, error) {
	defer rc.Close()
	var b bytes.Buffer
	if _, err := b.ReadFrom(io.LimitReader(rc, limit)); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

type ioDiscard struct{}

func (ioDiscard) Write(p []byte) (int, error) { return len(p), nil }
