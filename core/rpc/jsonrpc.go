package rpc

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
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
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/syndtr/goleveldb/leveldb"

	"noorchain-evm-l1/core/node"
	"noorchain-evm-l1/core/txpool"
)

type Server struct {
	addr    string
	log     *log.Logger
	chainID string
	n       *node.Node
	evm     *EvmMock

	httpSrv *http.Server
	ln      net.Listener
}

func New(addr, chainID string, n *node.Node, db *leveldb.DB, l *log.Logger) *Server {
	if l == nil {
		l = log.New(ioDiscard{}, "[rpc] ", log.LstdFlags)
	}
	s := &Server{
		addr:    addr,
		log:     l,
		chainID: chainID,
		n:       n,
		evm:     NewEvmMock(db),
	}
	return s
}

func (s *Server) Start(ctx context.Context) error {

        assertRoutingTableStatic()

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
	Result  any             `json:"result"`
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

// Routing classes (M10.6)
//
// routeClass defines dispatcher routing behavior.
//
//   routeLeaderOnly      : write / leader-only (proxy if follower)
//   routeLocalThenProxy  : read local, fallback proxy leader
//   routeLocal           : local or safe stub
//
type routeClass uint8

const (
        routeLocal routeClass = iota
        routeLocalThenProxy
        routeLeaderOnly
)

// Canonical routing table (eth_*) — declarative, not yet enforced
var ethRouting = map[string]routeClass{
        "eth_sendRawTransaction":     routeLeaderOnly,
        "eth_getTransactionReceipt":  routeLocalThenProxy,
        "eth_getTransactionByHash":   routeLocalThenProxy,
        "eth_chainId":                routeLocal,
        "eth_blockNumber":            routeLocal,
        "eth_accounts":               routeLocal,
        "eth_getTransactionCount":    routeLocal,
        "eth_gasPrice":               routeLocal,
        "eth_estimateGas":            routeLocal,
        "eth_getBalance":             routeLocal,
        "eth_call":                   routeLocal,
        "eth_getBlockByNumber":       routeLocal,
}

func (s *Server) dispatch(req *rpcReq) rpcResp {
        resp := rpcResp{JSONRPC: "2.0", ID: req.ID}

        if req.JSONRPC != "2.0" {
                resp.Error = &rpcError{Code: -32600, Message: "invalid jsonrpc version"}
                return resp
        }


        // Dispatcher routing (M10.6)
        // M10.6 safety: assert routing table coverage (log-only)
        if strings.HasPrefix(req.Method, "eth_") {
                if _, ok := ethRouting[req.Method]; !ok {
                        s.log.Println("rpc warning: eth method not in routing table:", req.Method)
                }
        }


        // M10.6 effective routing (leader/follower)
        if cls, ok := ethRouting[req.Method]; ok {
                if cls == routeLeaderOnly && s.n != nil {
                        cfg := s.n.Config()
                        if strings.EqualFold(strings.TrimSpace(cfg.Role), "follower") &&
                           strings.TrimSpace(cfg.FollowRPC) != "" {
                                return s.proxyToLeader(req)
                        }
                }
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
                if s.n != nil {
                        resp.Result = toHexUint(s.n.Height())
                        return resp
                }
                resp.Result = toHexUint(0)
                return resp

        case "eth_syncing":
                resp.Result = false
                return resp

        case "eth_mining":
                resp.Result = true
                return resp

        case "net_listening":
                resp.Result = true
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

        case "eth_feeHistory":
                // Minimal EIP-1559 feeHistory for wallet compatibility (dev-only)
                resp.Result = map[string]any{
                        "oldestBlock": toHexUint(0),
                        "baseFeePerGas": []string{"0x1", "0x1"},
                        "gasUsedRatio": []float64{0},
                        "reward": [][]string{},
                }
                return resp

        case "eth_getLogs":
                resp.Result = []any{}
                return resp

        case "eth_newFilter":
                resp.Result = "0x1"
                return resp

        case "eth_newBlockFilter":
                resp.Result = "0x1"
                return resp

        case "eth_uninstallFilter":
                resp.Result = true
                return resp

        case "eth_getFilterChanges":
                resp.Result = []any{}
                return resp

        case "eth_getFilterLogs":
                resp.Result = []any{}
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
		// M8.A: accept signed raw tx, hash it, and enqueue into node txpool (no EVM execution yet).
		var params []string
		if err := json.Unmarshal(req.Params, &params); err != nil || len(params) < 1 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp


		}
		if s.n == nil {
			resp.Error = &rpcError{Code: -32000, Message: "node not attached"}
			return resp


		}
		rawHex := strings.TrimPrefix(strings.TrimSpace(params[0]), "0x")
		rawBytes, err := hex.DecodeString(rawHex)
		if err != nil {
			resp.Error = &rpcError{Code: -32602, Message: "invalid hex"}
			return resp


		}

		// Validate basic tx encoding (RLP) using geth decoder.
		var tx types.Transaction
		if err := tx.UnmarshalBinary(rawBytes); err != nil {
			resp.Error = &rpcError{Code: -32602, Message: "invalid raw tx"}
			return resp


		}

		h := crypto.Keccak256Hash(rawBytes)
		s.n.TxPool().AddPending(txpool.Tx{Hash: h.Hex(), Raw: rawBytes})
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

		// M9: prefer persisted receipts from LevelDB (rcpt/v1/<hash>) if available.
		if s.evm != nil && s.evm.db != nil {
			key := []byte("rcpt/v1/" + strings.ToLower(strings.TrimPrefix(hashStr, "0x")))
			if b, err := s.evm.db.Get(key, nil); err == nil && len(b) > 0 {
				var anyRcpt any
				if err := json.Unmarshal(b, &anyRcpt); err == nil {
					resp.Result = anyRcpt
					return resp


				}
			}
		}

		// Fallback: in-memory receipt store

                // M10: follower mode proxy — if receipt not found locally, ask leader RPC
                if s.n != nil {
                        cfg := s.n.Config()
                        if strings.EqualFold(strings.TrimSpace(cfg.Role), "follower") && strings.TrimSpace(cfg.FollowRPC) != "" {
                                // best-effort proxy to leader
                                body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","id":1,"method":"eth_getTransactionReceipt","params":["%s"]}`, hashStr))
                                resp2, err := http.Post(strings.TrimRight(cfg.FollowRPC, "/"), "application/json", bytes.NewReader(body))
                                if err == nil && resp2 != nil {
                                        b2, _ := ioReadAllLimit(resp2.Body, 2<<20)
                                        _ = resp2.Body.Close()
                                        // Parse minimal: expect {result: ...} or {error: ...}
                                        var pr struct { Result any `json:"result"`; Error *rpcError `json:"error"` }
                                        if err := json.Unmarshal(b2, &pr); err == nil {
                                                if pr.Error == nil {
                                                        // could be null or object
                                                        resp.Result = pr.Result
                                                        return resp


                                                }
                                        }
                                }
                        }
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

		// M8.A: check node txpool first (pending / minimal mined)
		if s.n != nil {
			if ptx, ok := s.n.TxPool().Get(hashStr); ok {
				var blockNumber any = nil
				var txIndex any = nil
				if bn, ok := s.n.TxIndex().Get(hashStr); ok {
					blockNumber = toHexUint(bn)
					txIndex = "0x0"
				}
				resp.Result = map[string]any{
					"hash":             ptx.Hash,
					"blockHash":        nil,
					"blockNumber":      blockNumber,
					"transactionIndex": txIndex,
					"from":             "0x" + strings.Repeat("0", 40),
					"to":               nil,
					"nonce":            "0x0",
					"value":            "0x0",
					"gas":              "0x0",
					"gasPrice":         "0x0",
					"input":            "0x",
					"type":             "0x0",
					"chainId":          chainIDToHex(s.chainID),
				}
				return resp


			}
		}

		// Fallback: legacy dev mock store (pre-M8.A)
		if !strings.HasPrefix(hashStr, "0x") || len(hashStr) != 66 {
			resp.Result = nil
			return resp


		}
		h := common.HexToHash(hashStr)

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

		v, rSig, sSig := tx.RawSignatureValues()

		var blockHash any = nil
		var blockNumber any = nil
		var idx any = nil
		if rcpt != nil {
			blockHash = rcpt.BlockHash.Hex()
			blockNumber = rcpt.BlockNumber
			idx = "0x0"
		}

		gasPrice := "0x0"
		maxFee := "0x0"
		maxPrio := "0x0"
		if tx.Type() == 2 {
			if tx.GasFeeCap() != nil {
				maxFee = "0x" + tx.GasFeeCap().Text(16)
			}
			if tx.GasTipCap() != nil {
				maxPrio = "0x" + tx.GasTipCap().Text(16)
			}
			gasPrice = maxFee
		} else {
			if tx.GasPrice() != nil {
				gasPrice = "0x" + tx.GasPrice().Text(16)
			}
		}

		var to any = nil
		if tx.To() != nil {
			to = tx.To().Hex()
		}

		resp.Result = map[string]any{
			"hash":                 tx.Hash().Hex(),
			"nonce":                toHexUint(tx.Nonce()),
			"blockHash":            blockHash,
			"blockNumber":          blockNumber,
			"transactionIndex":     idx,
			"from":                 from.Hex(),
			"to":                   to,
			"value":                "0x" + tx.Value().Text(16),
			"gas":                  toHexUint(tx.Gas()),
			"gasPrice":             gasPrice,
			"maxFeePerGas":         maxFee,
			"maxPriorityFeePerGas": maxPrio,
			"input":                "0x" + common.Bytes2Hex(tx.Data()),
			"type":                 toHexUint(uint64(tx.Type())),
			"chainId":              "0x" + chainBig.Text(16),
			"v":                    "0x" + v.Text(16),
			"r":                    "0x" + rSig.Text(16),
			"s":                    "0x" + sSig.Text(16),
		}
		return resp



        case "eth_getBlockByHash":
                // Minimal: ignore hash input, return latest block (dev-compatible)
                resp.Result = nil
                if s.n != nil {
                        n := s.n.Height()
                        // reuse the same shape as eth_getBlockByNumber
                        resp.Result = map[string]any{
                                "number":     toHexUint(n),
                                "hash":       pseudoBlockHash(n).Hex(),
                                "parentHash": pseudoBlockHash(n - 1).Hex(),
                                "nonce":      "0x0000000000000000",
                                "sha3Uncles":  "0x" + strings.Repeat("0", 64),
                                "logsBloom":   "0x" + strings.Repeat("0", 512),
                                "transactionsRoot": "0x" + strings.Repeat("0", 64),
                                "stateRoot":        "0x" + strings.Repeat("0", 64),
                                "receiptsRoot":     "0x" + strings.Repeat("0", 64),
                                "miner":            "0x" + strings.Repeat("0", 40),
                                "difficulty":       "0x0",
                                "totalDifficulty":  "0x0",
                                "extraData":        "0x",
                                "size":             "0x0",
                                "gasLimit":         "0x1c9c380",
                                "gasUsed":          "0x0",
                                "timestamp":        toHexUint(uint64(time.Now().Unix())),
                                "mixHash":          "0x" + strings.Repeat("0", 64),
                                "baseFeePerGas":    "0x1",
                                "transactions":     []any{},
                                "uncles":           []any{},
                        }
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
		n := uint64(0)
		if s.n != nil {
			n = s.n.Height()
		}
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

// M10.6 leader proxy (minimal)
func (s *Server) proxyToLeader(req *rpcReq) rpcResp {
        cfg := s.n.Config()
        body, _ := json.Marshal(req)
        respHTTP, err := http.Post(strings.TrimRight(cfg.FollowRPC, "/"), "application/json", bytes.NewReader(body))
        if err != nil || respHTTP == nil {
                return rpcResp{JSONRPC: "2.0", ID: req.ID, Error: &rpcError{Code: -32000, Message: "leader unreachable"}}
        }
        defer respHTTP.Body.Close()
        b, _ := ioReadAllLimit(respHTTP.Body, 2<<20)
        var out rpcResp
        if err := json.Unmarshal(b, &out); err != nil {
                return rpcResp{JSONRPC: "2.0", ID: req.ID, Error: &rpcError{Code: -32603, Message: "invalid leader response"}}
        }
        return out
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


// M10.6 static safety: ensure routing table entries exist in dispatcher
func assertRoutingTableStatic() {
        // list of methods handled by switch (manually maintained)
        handled := map[string]struct{}{
                "eth_chainId": {},
                "eth_blockNumber": {},
                "eth_accounts": {},
                "eth_getTransactionCount": {},
                "eth_gasPrice": {},
                "eth_estimateGas": {},
                "eth_getBalance": {},
                "eth_call": {},
                "eth_sendRawTransaction": {},
                "eth_getTransactionReceipt": {},
                "eth_getTransactionByHash": {},
                "eth_getBlockByNumber": {},
        }

        for m := range ethRouting {
                if _, ok := handled[m]; !ok {
                        panic("rpc: routing table references unhandled method: " + m)
                }
        }
}




type ioDiscard struct{}

func (ioDiscard) Write(p []byte) (int, error) { return len(p), nil }
