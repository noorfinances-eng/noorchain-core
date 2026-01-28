package rpc

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/triedb"
	"github.com/gorilla/websocket"
	"github.com/holiman/uint256"
	"github.com/syndtr/goleveldb/leveldb/util"
	"io"
	"log"
	"math/big"
	"net"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"

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
	// M15: JSON-RPC filters (in-memory, leader-only; followers proxy to leader)
	filtersMu   sync.Mutex
	filtersNext uint64
	filters     map[uint64]*rpcFilter
	// M16: WebSocket subscriptions (newHeads)
	wsMu   sync.Mutex
	wsNext uint64
	wsSubs map[uint64]*wsSub
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

		filtersNext: 1,
		filters:     make(map[uint64]*rpcFilter),

		wsNext: 1,
		wsSubs: make(map[uint64]*wsSub),
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

	// M16: WS newHeads broadcaster
	s.startWSNewHeads(ctx)

	// M17: WS logs broadcaster (leader-only; uses logrec via eth_getLogs)
	s.startWSLogs(ctx)
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

// ---- JSON-RPC filters (M15) ----
//
// In-memory filters, leader-only. Followers proxy to leader via routeLeaderOnly.
// This is sufficient for mainnet-like wallet tooling (viem, ethers) in a controlled environment.
type rpcFilterKind uint8

const (
	filterLogs rpcFilterKind = iota
	filterNewBlocks
)

type rpcFilter struct {
	kind rpcFilterKind

	// createdAtHead is the node head height at creation time (used by getFilterLogs semantics).
	createdAtHead uint64

	// lastSeenHead is the last head height returned by getFilterChanges (cursor).
	lastSeenHead uint64

	// lastAccess is for TTL/GC.
	lastAccess time.Time

	// raw filter object as received by eth_newFilter (parsed once, reused).
	// It follows eth_getLogs semantics: fromBlock/toBlock/address/topics.
	logsFilter map[string]any
}

const (
	// M15: filter lifetime / resource guardrails.
	// TTL is evaluated opportunistically on filter API calls.
	rpcFilterTTL = 60 * time.Second

	// Hard cap to prevent unbounded growth.
	rpcFilterMax = 1024
)

// filtersGCLocked evicts expired filters and enforces the cap.
// Caller MUST hold s.filtersMu.
func (s *Server) filtersGCLocked(now time.Time) {
	if s.filters == nil || len(s.filters) == 0 {
		return
	}

	// TTL eviction
	if rpcFilterTTL > 0 {
		for id, f := range s.filters {
			if f == nil {
				delete(s.filters, id)
				continue
			}
			if !f.lastAccess.IsZero() && now.Sub(f.lastAccess) > rpcFilterTTL {
				delete(s.filters, id)
			}
		}
	}

	// Cap eviction (least-recently-accessed)
	if rpcFilterMax > 0 && len(s.filters) > rpcFilterMax {
		for len(s.filters) > rpcFilterMax {
			var oldestID uint64
			var oldest time.Time
			first := true
			for id, f := range s.filters {
				t := time.Unix(0, 0)
				if f != nil && !f.lastAccess.IsZero() {
					t = f.lastAccess
				}
				if first || t.Before(oldest) {
					oldest = t
					oldestID = id
					first = false
				}
			}
			delete(s.filters, oldestID)
			if first {
				break
			}
		}
	}
}

// ---- WebSocket (M16) ----
//
// WS is served on the same addr/path as HTTP JSON-RPC (Upgrade on "/").
// Until M18, WS subscribe/unsubscribe are leader-only (followers return error).

type wsConn struct {
	c  *websocket.Conn
	mu sync.Mutex
}

func (wc *wsConn) sendJSON(v any) error {
	wc.mu.Lock()
	defer wc.mu.Unlock()
	return wc.c.WriteJSON(v)
}

type wsSubKind uint8

const (
	wsSubNewHeads wsSubKind = iota
	wsSubLogs
)

type wsSub struct {
	id         uint64
	kind       wsSubKind
	conn       *wsConn
	lastSeen   uint64
	logsFilter map[string]any
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		// Controlled environment: allow all origins.
		return true
	},
}

func (s *Server) handleWS(w http.ResponseWriter, r *http.Request) {
	c, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	wc := &wsConn{c: c}
	defer func() {
		_ = c.Close()
		s.wsUnsubscribeAll(wc)
	}()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			return
		}
		msg = bytes.TrimSpace(msg)
		if len(msg) == 0 {
			continue
		}

		// Batch: [...]
		if msg[0] == '[' {
			var reqs []rpcReq
			if err := json.Unmarshal(msg, &reqs); err != nil {
				_ = wc.sendJSON(rpcResp{JSONRPC: "2.0", ID: nil, Error: &rpcError{Code: -32700, Message: "parse error"}})
				continue
			}
			resps := make([]rpcResp, 0, len(reqs))
			for i := range reqs {
				resps = append(resps, s.dispatchWS(&reqs[i], wc))
			}
			_ = wc.sendJSON(resps)
			continue
		}

		// Single: {...}
		var req rpcReq
		if err := json.Unmarshal(msg, &req); err != nil {
			_ = wc.sendJSON(rpcResp{JSONRPC: "2.0", ID: nil, Error: &rpcError{Code: -32700, Message: "parse error"}})
			continue
		}
		resp := s.dispatchWS(&req, wc)
		_ = wc.sendJSON(resp)
	}
}

func (s *Server) dispatchWS(req *rpcReq, wc *wsConn) rpcResp {
	// Until M18: subscribe/unsubscribe are leader-only.
	if s.n != nil {
		cfg := s.n.Config()
		if strings.TrimSpace(cfg.FollowRPC) != "" {
			if req.Method == "eth_subscribe" || req.Method == "eth_unsubscribe" {
				return rpcResp{JSONRPC: "2.0", ID: req.ID, Error: &rpcError{Code: -32000, Message: "leader-only"}}
			}
		}
	}

	switch req.Method {
	case "eth_subscribe":
		return s.wsSubscribe(req, wc)
	case "eth_unsubscribe":
		return s.wsUnsubscribe(req, wc)
	default:
		return s.dispatch(req)
	}
}

func (s *Server) wsSubscribe(req *rpcReq, wc *wsConn) rpcResp {
	resp := rpcResp{JSONRPC: "2.0", ID: req.ID}
	if req.JSONRPC != "2.0" {
		resp.Error = &rpcError{Code: -32600, Message: "invalid jsonrpc version"}
		return resp
	}

	var params []any
	if err := json.Unmarshal(req.Params, &params); err != nil || len(params) < 1 {
		resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
		return resp
	}

	typ, _ := params[0].(string)
	typ = strings.TrimSpace(typ)

	head := uint64(0)
	if s.n != nil {
		head = s.n.Height()
	}

	var kind wsSubKind
	var logsFilter map[string]any

	switch typ {
	case "newHeads":
		kind = wsSubNewHeads
	case "logs":
		kind = wsSubLogs
		lf := map[string]any{}
		if len(params) >= 2 && params[1] != nil {
			m, ok := params[1].(map[string]any)
			if !ok {
				resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
				return resp
			}
			for k, v := range m {
				lf[k] = v
			}
		}
		logsFilter = lf
	default:
		resp.Error = &rpcError{Code: -32601, Message: "subscription type not supported"}
		return resp
	}

	s.wsMu.Lock()
	id := s.wsNext
	s.wsNext++
	if s.wsSubs == nil {
		s.wsSubs = make(map[uint64]*wsSub)
	}
	sub := &wsSub{id: id, kind: kind, conn: wc, lastSeen: head, logsFilter: logsFilter}
	if kind == wsSubNewHeads {
		sub.lastSeen = 0
		sub.logsFilter = nil
	}
	s.wsSubs[id] = sub
	s.wsMu.Unlock()

	resp.Result = fmt.Sprintf("0x%x", id)
	return resp
}

func (s *Server) wsUnsubscribe(req *rpcReq, wc *wsConn) rpcResp {
	resp := rpcResp{JSONRPC: "2.0", ID: req.ID}
	if req.JSONRPC != "2.0" {
		resp.Error = &rpcError{Code: -32600, Message: "invalid jsonrpc version"}
		return resp
	}

	var params []any
	if err := json.Unmarshal(req.Params, &params); err != nil || len(params) < 1 {
		resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
		return resp
	}

	idStr, _ := params[0].(string)
	idStr = strings.TrimSpace(idStr)
	if !strings.HasPrefix(idStr, "0x") {
		resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
		return resp
	}

	u, err := strconv.ParseUint(strings.TrimPrefix(idStr, "0x"), 16, 64)
	if err != nil {
		resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
		return resp
	}

	ok := false
	s.wsMu.Lock()
	if sub, exists := s.wsSubs[u]; exists && sub != nil && sub.conn == wc {
		delete(s.wsSubs, u)
		ok = true
	}
	s.wsMu.Unlock()

	resp.Result = ok
	return resp
}

func (s *Server) wsUnsubscribeAll(wc *wsConn) {
	s.wsMu.Lock()
	defer s.wsMu.Unlock()
	for id, sub := range s.wsSubs {
		if sub != nil && sub.conn == wc {
			delete(s.wsSubs, id)
		}
	}
}

func (s *Server) startWSNewHeads(ctx context.Context) {
	if s.n == nil {
		return
	}
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		last := s.n.Height()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				h := s.n.Height()
				if h <= last {
					continue
				}
				for i := last + 1; i <= h; i++ {
					s.wsBroadcastNewHead(i)
				}
				last = h
			}
		}
	}()
}

func (s *Server) wsBroadcastNewHead(height uint64) {
	// Reuse eth_getBlockByNumber result shape (dev-compatible, includes roots/bloom if blkmeta exists).
	params := json.RawMessage(fmt.Sprintf(`["%s", false]`, toHexUint(height)))
	req := rpcReq{JSONRPC: "2.0", ID: json.RawMessage("null"), Method: "eth_getBlockByNumber", Params: params}
	out := s.dispatch(&req)
	if out.Error != nil || out.Result == nil {
		return
	}

	s.wsMu.Lock()
	subs := make([]*wsSub, 0, len(s.wsSubs))
	for _, sub := range s.wsSubs {
		if sub != nil && sub.kind == wsSubNewHeads && sub.conn != nil {
			subs = append(subs, sub)
		}
	}
	s.wsMu.Unlock()

	for _, sub := range subs {
		msg := map[string]any{
			"jsonrpc": "2.0",
			"method":  "eth_subscription",
			"params": map[string]any{
				"subscription": fmt.Sprintf("0x%x", sub.id),
				"result":       out.Result,
			},
		}
		_ = sub.conn.sendJSON(msg)
	}
}

func (s *Server) startWSLogs(ctx context.Context) {
	if s.n == nil {
		return
	}
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		type snap struct {
			id       uint64
			conn     *wsConn
			lastSeen uint64
			filter   map[string]any
		}

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				head := s.n.Height()

				// Snapshot subs (avoid holding lock while calling dispatch / writing WS)
				s.wsMu.Lock()
				subs := make([]snap, 0, len(s.wsSubs))
				for _, sub := range s.wsSubs {
					if sub == nil || sub.kind != wsSubLogs || sub.conn == nil {
						continue
					}
					lf := map[string]any{}
					if sub.logsFilter != nil {
						for k, v := range sub.logsFilter {
							lf[k] = v
						}
					}
					subs = append(subs, snap{id: sub.id, conn: sub.conn, lastSeen: sub.lastSeen, filter: lf})
				}
				s.wsMu.Unlock()

				for _, sub := range subs {
					if head <= sub.lastSeen {
						continue
					}
					from := sub.lastSeen + 1

					fo := make(map[string]any, len(sub.filter)+2)
					for k, v := range sub.filter {
						fo[k] = v
					}
					fo["fromBlock"] = toHexUint(from)
					fo["toBlock"] = toHexUint(head)

					pb, _ := json.Marshal([]any{fo})
					subReq := rpcReq{JSONRPC: "2.0", ID: json.RawMessage("null"), Method: "eth_getLogs", Params: pb}
					subResp := s.dispatch(&subReq)
					if subResp.Error != nil || subResp.Result == nil {
						// do not advance lastSeen on errors (avoid log loss)
						continue
					}

					// One notification per log (mainnet-like)
					arr, ok := subResp.Result.([]any)
					if ok {
						for _, it := range arr {
							msg := map[string]any{
								"jsonrpc": "2.0",
								"method":  "eth_subscription",
								"params": map[string]any{
									"subscription": fmt.Sprintf("0x%x", sub.id),
									"result":       it,
								},
							}
							_ = sub.conn.sendJSON(msg)
						}
					}

					// Advance lastSeen to head if subscription still exists (avoid racing unsubscribe)
					s.wsMu.Lock()
					if cur, ok := s.wsSubs[sub.id]; ok && cur != nil && cur.kind == wsSubLogs && cur.conn == sub.conn {
						cur.lastSeen = head
					}
					s.wsMu.Unlock()
				}
			}
		}
	}()
}

// ---- HTTP handler (single + batch) ----

func (s *Server) handleJSONRPC(w http.ResponseWriter, r *http.Request) {
	if websocket.IsWebSocketUpgrade(r) {
		s.handleWS(w, r)
		return
	}

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
//	routeLeaderOnly      : write / leader-only (proxy if follower)
//	routeLocalThenProxy  : read local, fallback proxy leader
//	routeLocal           : local or safe stub
type routeClass uint8

const (
	routeLocal routeClass = iota
	routeLocalThenProxy
	routeLeaderOnly
)

// Canonical routing table (eth_*) — declarative, not yet enforced
var ethRouting = map[string]routeClass{
	"eth_sendRawTransaction":    routeLeaderOnly,
	"eth_getTransactionReceipt": routeLocalThenProxy,
	"eth_getTransactionByHash":  routeLocalThenProxy,
	"eth_chainId":               routeLocal,
	"eth_blockNumber":           routeLeaderOnly,
	"eth_accounts":              routeLocal,
	"eth_getTransactionCount":   routeLeaderOnly,
	"eth_gasPrice":              routeLocal,
	"eth_estimateGas":           routeLocal,
	"eth_getBalance":            routeLeaderOnly,
	"eth_getCode":               routeLeaderOnly,
	"eth_getStorageAt":          routeLeaderOnly,
	"eth_call":                  routeLeaderOnly,
	"eth_getBlockByNumber":      routeLeaderOnly,
	"eth_getLogs":               routeLeaderOnly,

	// M15: filters — leader-only (follower proxies to leader to keep filter state consistent).
	"eth_newFilter":        routeLeaderOnly,
	"eth_newBlockFilter":   routeLeaderOnly,
	"eth_getFilterChanges": routeLeaderOnly,
	"eth_getFilterLogs":    routeLeaderOnly,
	"eth_uninstallFilter":  routeLeaderOnly,
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
			if strings.TrimSpace(cfg.FollowRPC) != "" {
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
		resp.Result = toHexUint(evmChainID)
		return resp

	case "net_version":
		resp.Result = strconv.FormatUint(evmChainID, 10)
		return resp

	case "eth_protocolVersion":
		resp.Result = "0x0"
		return resp

	case "net_peerCount":
		resp.Result = "0x0"
		return resp

	case "web3_sha3":
		var params []string
		if err := json.Unmarshal(req.Params, &params); err != nil || len(params) < 1 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp
		}
		inHex := strings.TrimPrefix(strings.TrimSpace(params[0]), "0x")
		in, err := hex.DecodeString(inHex)
		if err != nil {
			resp.Error = &rpcError{Code: -32602, Message: "invalid hex"}
			return resp
		}
		h := crypto.Keccak256Hash(in)
		resp.Result = h.Hex()
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

	case "eth_estimateGas":
		// Minimal estimate for tooling compatibility (Hardhat/Viem/ethers).
		// Deterministic heuristic (not a real EVM simulation).
		var params []any
		if err := json.Unmarshal(req.Params, &params); err != nil || len(params) < 1 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp
		}

		// params[0] is a tx object: {from,to,data,value,...}
		txo, _ := params[0].(map[string]any)
		toStr, _ := txo["to"].(string)
		dataStr, _ := txo["data"].(string)
		dataStr = strings.TrimSpace(dataStr)

		// Contract creation (no "to") => allocate higher gas.
		if strings.TrimSpace(toStr) == "" {
			resp.Result = toHexUint(5_000_000)
			return resp
		}

		// If calldata present => moderate estimate, else plain transfer.
		if dataStr != "" && dataStr != "0x" {
			resp.Result = toHexUint(300_000)
			return resp
		}

		resp.Result = toHexUint(21_000)
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

		nonce := uint64(0)

		// Prefer geth StateDB (world-state) if available
		if s.n != nil && s.n.EVMStore() != nil && s.n.EVMStore().DB() != nil {
			root := s.n.StateRootHead()

			diskdb := rawdb.NewDatabase(s.n.EVMStore().DB())
			tdb := triedb.NewDatabase(diskdb, nil)
			sdbCache := state.NewDatabase(tdb, nil)

			if st, err := state.New(root, sdbCache); err == nil {
				nonce = st.GetNonce(addr)
			} else {
				s.log.Println("rpc: state.New failed (txCount) | err:", err)
			}
		} else if s.evm != nil {
			// Fallback to legacy mock
			nonce = s.evm.GetTransactionCount(addr)
		}

		resp.Result = toHexUint(nonce)
		return resp

	case "eth_getBalance":
		// Minimal wallet/tooling compatibility (dev-only).
		// params: [ "0x..address..", "latest"|"pending"|... ]
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

		bal := uint256.NewInt(0) // default 0

		if s.n != nil && s.n.EVMStore() != nil && s.n.EVMStore().DB() != nil {
			root := s.n.StateRootHead()

			diskdb := rawdb.NewDatabase(s.n.EVMStore().DB())
			tdb := triedb.NewDatabase(diskdb, nil)
			sdbCache := state.NewDatabase(tdb, nil)

			if st, err := state.New(root, sdbCache); err == nil {
				bal = st.GetBalance(addr)
			} else {
				s.log.Println("rpc: state.New failed (balance) | err:", err)
			}
		}

		resp.Result = "0x" + bal.ToBig().Text(16)
		return resp

	case "eth_getCode":
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
		code := []byte{}
		if s.n != nil && s.n.EVMStore() != nil && s.n.EVMStore().DB() != nil {
			// params: [address, blockTag?] où blockTag = "latest"|"pending"|"earliest"| "0xHEIGHT"
			root := s.n.StateRootHead()
			callN := uint64(0)
			reqN := uint64(0)

			if s.n != nil {
				reqN = s.n.Height()
				callN = reqN
			}

			if len(params) >= 2 {
				if t, ok := params[1].(string); ok {
					switch t {
					case "latest", "pending":
						callN = reqN
					case "earliest":
						callN = 0
					default:
						if strings.HasPrefix(t, "0x") {
							tt := strings.TrimPrefix(t, "0x")
							if tt != "" {
								if v, err := strconv.ParseUint(tt, 16, 64); err == nil {
									callN = v
								}
							}
						} else {
							if v, err := strconv.ParseUint(t, 10, 64); err == nil {
								callN = v
							}
						}
					}
				}
			}

			if s.n != nil && callN > reqN {
				resp.Error = &rpcError{Code: -32000, Message: "state unavailable"}
				return resp
			}

			if s.n != nil && s.n.DB() != nil {
				key := []byte("blkmeta/v1/" + strings.TrimPrefix(toHexUint(callN), "0x"))
				if b, err := s.n.DB().Get(key, nil); err == nil && len(b) > 0 {
					var bm struct {
						StateRoot string `json:"stateRoot"`
					}
					if err := json.Unmarshal(b, &bm); err == nil {
						if strings.HasPrefix(bm.StateRoot, "0x") && len(bm.StateRoot) == 66 {
							root = common.HexToHash(bm.StateRoot)
						}
					}
				}
			}

			diskdb := rawdb.NewDatabase(s.n.EVMStore().DB())
			tdb := triedb.NewDatabase(diskdb, nil)
			sdbCache := state.NewDatabase(tdb, nil)
			if st, err := state.New(root, sdbCache); err == nil {
				code = st.GetCode(addr)
			} else {
				s.log.Println("rpc: state.New failed (getCode) | err:", err)
			}
		}

		if len(code) == 0 {
			resp.Result = "0x"
			return resp
		}
		resp.Result = "0x" + hex.EncodeToString(code)
		return resp

	case "eth_getStorageAt":
		var params []any
		if err := json.Unmarshal(req.Params, &params); err != nil || len(params) < 2 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp
		}
		addrStr, _ := params[0].(string)
		slotStr, _ := params[1].(string)
		if !common.IsHexAddress(addrStr) {
			resp.Error = &rpcError{Code: -32602, Message: "invalid address"}
			return resp
		}
		addr := common.HexToAddress(addrStr)
		slotHex := strings.TrimPrefix(strings.TrimSpace(slotStr), "0x")
		if len(slotHex) > 64 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid slot"}
			return resp
		}
		if len(slotHex)%2 == 1 {
			slotHex = "0" + slotHex
		}
		b, err := hex.DecodeString(slotHex)
		if err != nil {
			resp.Error = &rpcError{Code: -32602, Message: "invalid slot"}
			return resp
		}
		var slot common.Hash
		copy(slot[32-len(b):], b)
		val := common.Hash{}
		if s.n != nil && s.n.EVMStore() != nil && s.n.EVMStore().DB() != nil {
			root := s.n.StateRootHead()
			callN := uint64(0)
			reqN := uint64(0)

			if s.n != nil {
				reqN = s.n.Height()
				callN = reqN
			}

			if len(params) >= 3 {
				if t, ok := params[2].(string); ok {
					switch t {
					case "latest", "pending":
						callN = reqN
					case "earliest":
						callN = 0
					default:
						if strings.HasPrefix(t, "0x") {
							tt := strings.TrimPrefix(t, "0x")
							if tt != "" {
								if v, err := strconv.ParseUint(tt, 16, 64); err == nil {
									callN = v
								}
							}
						} else {
							if v, err := strconv.ParseUint(t, 10, 64); err == nil {
								callN = v
							}
						}
					}
				}
			}

			if s.n != nil && callN > reqN {
				resp.Error = &rpcError{Code: -32000, Message: "state unavailable"}
				return resp
			}

			if s.n != nil && s.n.DB() != nil {
				key := []byte("blkmeta/v1/" + strings.TrimPrefix(toHexUint(callN), "0x"))
				if b, err := s.n.DB().Get(key, nil); err == nil && len(b) > 0 {
					var bm struct {
						StateRoot string `json:"stateRoot"`
					}
					if err := json.Unmarshal(b, &bm); err == nil {
						if strings.HasPrefix(bm.StateRoot, "0x") && len(bm.StateRoot) == 66 {
							root = common.HexToHash(bm.StateRoot)
						}
					}
				}
			}

			diskdb := rawdb.NewDatabase(s.n.EVMStore().DB())
			tdb := triedb.NewDatabase(diskdb, nil)
			sdbCache := state.NewDatabase(tdb, nil)
			if st, err := state.New(root, sdbCache); err == nil {
				val = st.GetState(addr, slot)
			} else {
				s.log.Println("rpc: state.New failed (getStorageAt) | err:", err)
			}
		}

		resp.Result = val.Hex()
		return resp

	case "eth_gasPrice":
		resp.Result = "0x1"
		return resp

	case "eth_feeHistory":
		// Minimal EIP-1559 feeHistory for wallet compatibility (dev-only)
		resp.Result = map[string]any{
			"oldestBlock":   toHexUint(0),
			"baseFeePerGas": []string{"0x1", "0x1"},
			"gasUsedRatio":  []float64{0},
			"reward":        [][]string{},
		}
		return resp

	case "eth_getLogs":
		// Minimal eth_getLogs: scan persisted receipts (rcpt/v1/*) and filter logs in-memory.
		// NOTE: O(n) over receipts; acceptable for local/mainnet-like pack until a proper log index is added.
		var paramsAny []any
		if err := json.Unmarshal(req.Params, &paramsAny); err != nil || len(paramsAny) < 1 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp
		}

		filter, ok := paramsAny[0].(map[string]any)
		if !ok {
			resp.Error = &rpcError{Code: -32602, Message: "invalid filter"}
			return resp
		}

		// Helpers
		parseQty := func(v any, fallback uint64) uint64 {
			s, _ := v.(string)
			s = strings.TrimSpace(s)
			if s == "" || s == "latest" || s == "pending" {
				return fallback
			}
			if s == "earliest" {
				return 0
			}
			s = strings.TrimPrefix(s, "0x")
			if s == "" {
				return fallback
			}
			n, err := strconv.ParseUint(s, 16, 64)
			if err != nil {
				return fallback
			}
			return n
		}
		parseHexQty0 := func(v any) uint64 {
			s, _ := v.(string)
			s = strings.TrimSpace(s)
			if s == "" {
				return 0
			}
			s = strings.TrimPrefix(s, "0x")
			if s == "" {
				return 0
			}
			n, err := strconv.ParseUint(s, 16, 64)
			if err != nil {
				return 0
			}
			return n
		}

		// Range
		latest := uint64(0)
		if s.n != nil {
			latest = s.n.Height()
		}
		fromN := parseQty(filter["fromBlock"], latest)
		toN := parseQty(filter["toBlock"], latest)

		// Deterministic mainnet-like semantics:
		// - if fromBlock > head => empty set
		// - if toBlock > head   => clamp to head
		if fromN > latest {
			resp.Result = []any{}
			return resp
		}
		if toN > latest {
			toN = latest
		}

		if toN < fromN {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params: fromBlock > toBlock"}
			return resp
		}

		// Hard cap to avoid unbounded scans (client should chunk ranges).
		const maxBlockRange uint64 = 16384
		if (toN - fromN + 1) > maxBlockRange {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params: block range too large (max 16384 blocks)"}
			return resp
		}

		// Address filter: string or []string
		addrSet := map[string]bool{}
		if av, ok := filter["address"]; ok && av != nil {
			switch t := av.(type) {
			case string:
				a := strings.ToLower(strings.TrimSpace(t))
				if a != "" {
					addrSet[a] = true
				}
			case []any:
				for _, it := range t {
					if s2, ok := it.(string); ok {
						a := strings.ToLower(strings.TrimSpace(s2))
						if a != "" {
							addrSet[a] = true
						}
					}
				}
			}
		}

		// Topics filter: [] (each element can be null, "0x..", or [] "0x.." for OR)
		var topicsFilter []any
		if tv, ok := filter["topics"]; ok {
			if arr, ok := tv.([]any); ok {
				topicsFilter = arr
			}
		}
		matchTopics := func(logTopics []string) bool {
			if len(topicsFilter) == 0 {
				return true
			}
			for i := 0; i < len(topicsFilter); i++ {
				if i >= len(logTopics) {
					return false
				}
				f := topicsFilter[i]
				if f == nil {
					continue
				}
				switch x := f.(type) {
				case string:
					want := strings.ToLower(strings.TrimSpace(x))
					if want == "" {
						continue
					}
					if strings.ToLower(logTopics[i]) != want {
						return false
					}
				case []any:
					// OR set
					okAny := false
					for _, oi := range x {
						if s3, ok := oi.(string); ok {
							want := strings.ToLower(strings.TrimSpace(s3))
							if want != "" && strings.ToLower(logTopics[i]) == want {
								okAny = true
								break
							}
						}
					}
					if !okAny {
						return false
					}
				default:
					// unknown filter shape -> fail closed
					return false
				}
			}
			return true
		}

		if s.n == nil || s.n.DB() == nil {
			resp.Error = &rpcError{Code: -32000, Message: "db not available"}
			return resp
		}

		// Read log index (logrec/v1/<heightBE><txIndexBE><logIndexBE>) and filter in-memory.
		// This is range-based (no rcpt/v1 scan), mainnet-like for performance.
		type logItem struct {
			blockN   uint64
			txIndex  uint64
			logIndex uint64
			txHash   string
			m        map[string]any
		}

		items := make([]logItem, 0, 8)

		prefix := []byte("logrec/v1/")

		putU64BE := func(dst []byte, v uint64) []byte {
			return append(dst,
				byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32),
				byte(v>>24), byte(v>>16), byte(v>>8), byte(v),
			)
		}
		readU64BE := func(b []byte) uint64 {
			_ = b[7] // bounds-check hint
			return uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 |
				uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7])
		}
		seekKey := func(h uint64) []byte {
			k := make([]byte, 0, len(prefix)+24)
			k = append(k, prefix...)
			k = putU64BE(k, h)
			k = putU64BE(k, 0)
			k = putU64BE(k, 0)
			return k
		}

		it := s.n.DB().NewIterator(util.BytesPrefix(prefix), nil)
		defer it.Release()

		for ok := it.Seek(seekKey(fromN)); ok; ok = it.Next() {
			k := it.Key()
			if len(k) < len(prefix)+8 {
				continue
			}

			// height is the first 8 bytes after "logrec/v1/"
			h := readU64BE(k[len(prefix) : len(prefix)+8])
			if h > toN {
				break
			}

			var lm map[string]any
			if err := json.Unmarshal(it.Value(), &lm); err != nil {
				continue
			}

			// address filter
			if len(addrSet) > 0 {
				as, _ := lm["address"].(string)
				as = strings.ToLower(strings.TrimSpace(as))
				if as == "" || !addrSet[as] {
					continue
				}
			}

			// topics filter
			ltAny, _ := lm["topics"].([]any)
			lt := make([]string, 0, len(ltAny))
			for _, t := range ltAny {
				if ts, ok := t.(string); ok {
					lt = append(lt, ts)
				}
			}
			if !matchTopics(lt) {
				continue
			}

			txIdx := parseHexQty0(lm["transactionIndex"])
			logIdx := parseHexQty0(lm["logIndex"])
			txh, _ := lm["transactionHash"].(string)
			txh = strings.ToLower(strings.TrimSpace(txh))

			items = append(items, logItem{
				blockN:   h,
				txIndex:  txIdx,
				logIndex: logIdx,
				txHash:   txh,
				m:        lm,
			})
		}

		sort.SliceStable(items, func(i, j int) bool {
			a, b := items[i], items[j]
			if a.blockN != b.blockN {
				return a.blockN < b.blockN
			}
			if a.txIndex != b.txIndex {
				return a.txIndex < b.txIndex
			}
			if a.logIndex != b.logIndex {
				return a.logIndex < b.logIndex
			}
			return a.txHash < b.txHash
		})

		out := make([]any, 0, len(items))
		for _, it := range items {
			out = append(out, it.m)
		}

		resp.Result = out
		return resp

	case "eth_newFilter":
		// params: [ filterObject ]
		var paramsAny []any
		if err := json.Unmarshal(req.Params, &paramsAny); err != nil || len(paramsAny) < 1 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp
		}
		filterObj, ok := paramsAny[0].(map[string]any)
		if !ok {
			resp.Error = &rpcError{Code: -32602, Message: "invalid filter object"}
			return resp
		}

		head := uint64(0)
		if s.n != nil {
			head = s.n.Height()
		}

		saved := make(map[string]any, len(filterObj))
		for k, v := range filterObj {
			saved[k] = v
		}

		s.filtersMu.Lock()
		s.filtersGCLocked(time.Now())
		id := s.filtersNext
		s.filtersNext++
		s.filters[id] = &rpcFilter{
			kind:          filterLogs,
			createdAtHead: head,
			lastSeenHead:  head,
			lastAccess:    time.Now(),
			logsFilter:    saved,
		}
		s.filtersMu.Unlock()

		resp.Result = toHexUint(id)
		return resp

	case "eth_newBlockFilter":
		head := uint64(0)
		if s.n != nil {
			head = s.n.Height()
		}

		s.filtersMu.Lock()
		s.filtersGCLocked(time.Now())
		id := s.filtersNext
		s.filtersNext++
		s.filters[id] = &rpcFilter{
			kind:          filterNewBlocks,
			createdAtHead: head,
			lastSeenHead:  head,
			lastAccess:    time.Now(),
		}
		s.filtersMu.Unlock()

		resp.Result = toHexUint(id)
		return resp

	case "eth_getFilterChanges":
		// params: [ filterId ]
		var paramsAny []any
		if err := json.Unmarshal(req.Params, &paramsAny); err != nil || len(paramsAny) < 1 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp
		}
		idStr, ok := paramsAny[0].(string)
		if !ok {
			resp.Error = &rpcError{Code: -32602, Message: "invalid filter id"}
			return resp
		}
		idHex := strings.TrimPrefix(strings.ToLower(strings.TrimSpace(idStr)), "0x")
		if idHex == "" {
			resp.Error = &rpcError{Code: -32602, Message: "invalid filter id"}
			return resp
		}
		idU, err := strconv.ParseUint(idHex, 16, 64)
		if err != nil {
			resp.Error = &rpcError{Code: -32602, Message: "invalid filter id"}
			return resp
		}

		head := uint64(0)
		if s.n != nil {
			head = s.n.Height()
		}

		// Snapshot filter state under lock (do NOT hold lock while calling dispatch).
		s.filtersMu.Lock()
		s.filtersGCLocked(time.Now())
		f, ok := s.filters[idU]
		if !ok || f == nil {
			s.filtersMu.Unlock()
			resp.Error = &rpcError{Code: -32000, Message: "filter not found"}
			return resp
		}
		f.lastAccess = time.Now()
		kind := f.kind
		lastSeen := f.lastSeenHead
		logsFilter := f.logsFilter
		s.filtersMu.Unlock()

		// Clamp (safety)
		if head < lastSeen {
			lastSeen = head
		}
		from := lastSeen + 1

		// Block filter: return new block hashes since lastSeen.
		if kind == filterNewBlocks {
			out := make([]any, 0, 8)
			if from <= head {
				for h := from; h <= head; h++ {
					// Deterministic mainnet-like behavior in our model:
					// our canonical block hash is pseudoBlockHash(height), same value exposed by eth_getBlockByNumber.
					out = append(out, pseudoBlockHash(h).Hex())
				}
			}

			s.filtersMu.Lock()
			if f2, ok := s.filters[idU]; ok && f2 != nil {
				f2.lastSeenHead = head
				f2.lastAccess = time.Now()
			}
			s.filtersMu.Unlock()

			resp.Result = out
			return resp
		}

		// Log filter: return new logs since lastSeen via eth_getLogs(logrec).
		if kind != filterLogs {
			resp.Error = &rpcError{Code: -32602, Message: "invalid filter type"}
			return resp
		}

		if from > head {
			s.filtersMu.Lock()
			if f2, ok := s.filters[idU]; ok && f2 != nil {
				f2.lastSeenHead = head
				f2.lastAccess = time.Now()
			}
			s.filtersMu.Unlock()

			resp.Result = []any{}
			return resp
		}

		fo := make(map[string]any, len(logsFilter)+2)
		for k, v := range logsFilter {
			fo[k] = v
		}
		fo["fromBlock"] = toHexUint(from)
		fo["toBlock"] = toHexUint(head)

		pb, _ := json.Marshal([]any{fo})
		subReq := rpcReq{JSONRPC: "2.0", ID: req.ID, Method: "eth_getLogs", Params: pb}
		subResp := s.dispatch(&subReq)
		if subResp.Error != nil {
			return subResp
		}

		s.filtersMu.Lock()
		if f2, ok := s.filters[idU]; ok && f2 != nil {
			f2.lastSeenHead = head
			f2.lastAccess = time.Now()
		}
		s.filtersMu.Unlock()

		resp.Result = subResp.Result
		return resp

	case "eth_getFilterLogs":
		// params: [ filterId ] — returns full logs per stored filter object
		var paramsAny []any
		if err := json.Unmarshal(req.Params, &paramsAny); err != nil || len(paramsAny) < 1 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp
		}
		idStr, ok := paramsAny[0].(string)
		if !ok {
			resp.Error = &rpcError{Code: -32602, Message: "invalid filter id"}
			return resp
		}
		idHex := strings.TrimPrefix(strings.ToLower(strings.TrimSpace(idStr)), "0x")
		if idHex == "" {
			resp.Error = &rpcError{Code: -32602, Message: "invalid filter id"}
			return resp
		}
		idU, err := strconv.ParseUint(idHex, 16, 64)
		if err != nil {
			resp.Error = &rpcError{Code: -32602, Message: "invalid filter id"}
			return resp
		}

		// Snapshot under lock
		s.filtersMu.Lock()
		s.filtersGCLocked(time.Now())
		f, ok := s.filters[idU]
		if !ok || f == nil {
			s.filtersMu.Unlock()
			resp.Error = &rpcError{Code: -32000, Message: "filter not found"}
			return resp
		}
		f.lastAccess = time.Now()
		kind := f.kind
		logsFilter := f.logsFilter
		s.filtersMu.Unlock()

		if kind != filterLogs {
			resp.Error = &rpcError{Code: -32602, Message: "invalid filter type"}
			return resp
		}

		fo := make(map[string]any, len(logsFilter))
		for k, v := range logsFilter {
			fo[k] = v
		}

		pb, _ := json.Marshal([]any{fo})
		subReq := rpcReq{JSONRPC: "2.0", ID: req.ID, Method: "eth_getLogs", Params: pb}
		subResp := s.dispatch(&subReq)
		if subResp.Error != nil {
			return subResp
		}

		resp.Result = subResp.Result
		return resp

	case "eth_uninstallFilter":
		// params: [ filterId ]
		var paramsAny []any
		if err := json.Unmarshal(req.Params, &paramsAny); err != nil || len(paramsAny) < 1 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp
		}
		idStr, ok := paramsAny[0].(string)
		if !ok {
			resp.Error = &rpcError{Code: -32602, Message: "invalid filter id"}
			return resp
		}
		idHex := strings.TrimPrefix(strings.ToLower(strings.TrimSpace(idStr)), "0x")
		if idHex == "" {
			resp.Error = &rpcError{Code: -32602, Message: "invalid filter id"}
			return resp
		}
		idU, err := strconv.ParseUint(idHex, 16, 64)
		if err != nil {
			resp.Error = &rpcError{Code: -32602, Message: "invalid filter id"}
			return resp
		}

		s.filtersMu.Lock()
		s.filtersGCLocked(time.Now())
		_, existed := s.filters[idU]
		if existed {
			delete(s.filters, idU)
		}
		s.filtersMu.Unlock()

		resp.Result = existed
		return resp

	case "eth_call":
		// M12.5: real eth_call against geth world-state (read-only).
		// params: [ {to:"0x..", data:"0x..", from?:"0x.."}, "latest"|blockTag ]
		var paramsAny []any
		if err := json.Unmarshal(req.Params, &paramsAny); err != nil || len(paramsAny) < 1 {
			resp.Error = &rpcError{Code: -32602, Message: "invalid params"}
			return resp
		}
		callObj, ok := paramsAny[0].(map[string]any)
		if !ok {
			resp.Error = &rpcError{Code: -32602, Message: "invalid call object"}
			return resp
		}

		toStr, _ := callObj["to"].(string)
		dataStr, _ := callObj["data"].(string)
		fromStr, _ := callObj["from"].(string)
		to := common.HexToAddress(toStr)
		from := common.HexToAddress(fromStr)

		dataStr = strings.TrimPrefix(strings.TrimSpace(dataStr), "0x")
		data, err := hex.DecodeString(dataStr)
		if err != nil {
			resp.Error = &rpcError{Code: -32602, Message: "invalid data hex"}
			return resp
		}

		if s.n == nil || s.n.EVMStore() == nil {
			resp.Error = &rpcError{Code: -32000, Message: "evm store not available"}
			return resp
		}

		// Resolve root from optional blockTag (params[1]) using persisted blkmeta when possible.
		reqN := uint64(0)
		if s.n != nil {
			reqN = s.n.Height()
		}

		// default: latest
		callN := reqN
		if len(paramsAny) >= 2 {
			if tag, ok := paramsAny[1].(string); ok {
				t := strings.TrimSpace(strings.ToLower(tag))
				switch t {
				case "latest", "pending":
					callN = reqN
				case "earliest":
					callN = 0
				default:
					if strings.HasPrefix(t, "0x") {
						tt := strings.TrimPrefix(t, "0x")
						if tt != "" {
							if v, err := strconv.ParseUint(tt, 16, 64); err == nil {
								callN = v
							}
						}
					} else {
						if v, err := strconv.ParseUint(t, 10, 64); err == nil {
							callN = v
						}
					}
				}
			}
		}

		// If asked height is above current, fail (header not found style).
		if s.n != nil && callN > reqN {
			resp.Error = &rpcError{Code: -32000, Message: "state unavailable"}
			return resp
		}

		// Prefer blkmeta stateRoot for the requested height when available; else fallback to head.
		root := s.n.StateRootHead()
		if s.n != nil && s.n.DB() != nil {
			key := []byte("blkmeta/v1/" + strings.TrimPrefix(toHexUint(callN), "0x"))
			if b, err := s.n.DB().Get(key, nil); err == nil && len(b) > 0 {
				var bm struct {
					StateRoot string `json:"stateRoot"`
				}
				if err := json.Unmarshal(b, &bm); err == nil {
					if strings.HasPrefix(bm.StateRoot, "0x") && len(bm.StateRoot) == 66 {
						root = common.HexToHash(bm.StateRoot)
					}
				}
			}
		}

		// Build StateDB at selected root (latest or historical).
		evmdb := s.n.EVMStore().DB()
		ethdb := rawdb.NewDatabase(evmdb)
		tdb := triedb.NewDatabase(ethdb, nil)
		statedb, err := state.New(root, state.NewDatabase(tdb, nil))
		if err != nil {
			resp.Error = &rpcError{Code: -32000, Message: "state unavailable"}
			return resp
		}

		// Read-only copy for eth_call safety.
		statedbRO := statedb.Copy()

		// Minimal EVM context for STATICCALL.
		blockCtx := vm.BlockContext{
			CanTransfer: coreCanTransfer,
			Transfer:    coreTransfer,
			GetHash:     func(uint64) common.Hash { return common.Hash{} },
			Coinbase:    common.Address{},
			GasLimit:    30_000_000,
			BlockNumber: new(big.Int).SetUint64(callN),
			Time:        uint64(time.Now().Unix()),
			Difficulty:  big.NewInt(0),
			BaseFee:     big.NewInt(1),
			Random:      new(common.Hash),
		}
		txCtx := vm.TxContext{Origin: from, GasPrice: big.NewInt(1)}

		cc := *params.AllDevChainProtocolChanges
		cc.ChainID = new(big.Int).SetUint64(evmChainID)
		evm := vm.NewEVM(blockCtx, txCtx, statedbRO, &cc, vm.Config{ExtraEips: []int{3855}})

		// Execute as STATICCALL (no state mutation).
		gas := uint64(3_000_000)
		ret, _, err := evm.StaticCall(vm.AccountRef(from), to, data, gas)
		if err != nil {
			resp.Error = &rpcError{Code: -32000, Message: err.Error()}
			return resp
		}
		resp.Result = "0x" + hex.EncodeToString(ret)
		return resp

	case "debug_traceTransaction":
		resp.Error = &rpcError{Code: -32601, Message: "not supported"}
		return resp

	case "debug_traceCall":
		resp.Error = &rpcError{Code: -32601, Message: "not supported"}
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

		// Persist raw tx for eth_getTransactionByHash (wallet+tooling compatibility)
		if s.evm != nil && s.evm.db != nil {
			k := "tx/v1/" + strings.ToLower(strings.TrimPrefix(h.Hex(), "0x"))
			_ = s.evm.db.Put([]byte(k), rawBytes, nil)
		}

		if s.evm != nil {
			chainBig := new(big.Int).SetUint64(evmChainID)
			signer := types.LatestSignerForChainID(chainBig)
			if from, err := types.Sender(signer, &tx); err == nil {
				s.evm.BumpNonce(from, tx.Nonce())
			}
		}

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
					if m, ok := anyRcpt.(map[string]any); ok {
						if _, has := m["effectiveGasPrice"]; !has {
							m["effectiveGasPrice"] = "0x1"
						}
						anyRcpt = m
					}
					resp.Result = anyRcpt
					return resp
				}
			}
		}

		// Fallback: in-memory receipt store

		// M10: follower mode proxy — if receipt not found locally, ask leader RPC
		if s.n != nil {
			cfg := s.n.Config()
			if strings.TrimSpace(cfg.FollowRPC) != "" {
				// best-effort proxy to leader
				body := []byte(fmt.Sprintf(`{"jsonrpc":"2.0","id":1,"method":"eth_getTransactionReceipt","params":["%s"]}`, hashStr))
				resp2, err := http.Post(strings.TrimRight(cfg.FollowRPC, "/"), "application/json", bytes.NewReader(body))
				if err == nil && resp2 != nil {
					b2, _ := ioReadAllLimit(resp2.Body, 2<<20)
					_ = resp2.Body.Close()
					// Parse minimal: expect {result: ...} or {error: ...}
					var pr struct {
						Result any       `json:"result"`
						Error  *rpcError `json:"error"`
					}
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

		hashStr := strings.TrimSpace(params[0])
		if !strings.HasPrefix(hashStr, "0x") || len(hashStr) != 66 {
			resp.Result = nil
			return resp
		}
		// M12.3: fast-path — read persisted raw tx (tx/v1/<hash>) and return real fields.
		if s.n != nil {
			cfg := s.n.Config()
			if strings.TrimSpace(cfg.FollowRPC) != "" {
				return s.proxyToLeader(req)
			}
		}
		if s.evm != nil && s.evm.db != nil {
			k := "tx/v1/" + strings.ToLower(strings.TrimPrefix(hashStr, "0x"))
			raw, err := s.evm.db.Get([]byte(k), nil)
			if err == nil && len(raw) > 0 {
				var tx types.Transaction
				if err := tx.UnmarshalBinary(raw); err == nil {
					chainBig := new(big.Int).SetUint64(evmChainID)
					signer := types.LatestSignerForChainID(chainBig)
					from, _ := types.Sender(signer, &tx)

					vSig, rSig, sSig := tx.RawSignatureValues()
					rHex := "0x" + fmt.Sprintf("%064x", rSig)
					sHex := "0x" + fmt.Sprintf("%064x", sSig)
					vHex := toHexUint(vSig.Uint64())

					var blockNumber any = nil
					var blockHash any = nil
					var txIndex any = nil
					if bn, ok := s.n.TxIndex().Get(hashStr); ok {
						blockNumber = toHexUint(bn)
						blockHash = pseudoBlockHash(bn).Hex()
						txIndex = "0x0"
					}

					toAny := any(nil)
					if tx.To() != nil {
						toAny = tx.To().Hex()
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

					resp.Result = map[string]any{
						"hash":                 hashStr,
						"blockHash":            blockHash,
						"blockNumber":          blockNumber,
						"transactionIndex":     txIndex,
						"from":                 from.Hex(),
						"to":                   toAny,
						"nonce":                toHexUint(tx.Nonce()),
						"value":                "0x" + tx.Value().Text(16),
						"gas":                  toHexUint(tx.Gas()),
						"gasPrice":             gasPrice,
						"maxFeePerGas":         maxFee,
						"maxPriorityFeePerGas": maxPrio,
						"input":                "0x" + common.Bytes2Hex(tx.Data()),
						"type":                 toHexUint(uint64(tx.Type())),
						"chainId":              "0x" + chainBig.Text(16),

						// --- CRITICAL for ethers v6 ---
						"r": rHex,
						"s": sHex,
						"v": vHex,
					}

					return resp
				}
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

		chainBig := new(big.Int).SetUint64(evmChainID)
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
				"number":           toHexUint(n),
				"hash":             pseudoBlockHash(n).Hex(),
				"parentHash":       pseudoBlockHash(n - 1).Hex(),
				"nonce":            "0x0000000000000000",
				"sha3Uncles":       "0x" + strings.Repeat("0", 64),
				"logsBloom":        "0x" + strings.Repeat("0", 512),
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

		// params: [ blockNumberOrTag, fullTxObjects ]
		reqN := uint64(0)
		if s.n != nil {
			reqN = s.n.Height()
		}
		// default: latest
		n := reqN
		var params []any
		if err := json.Unmarshal(req.Params, &params); err == nil && len(params) >= 1 {
			if tag, ok := params[0].(string); ok {
				t := strings.TrimSpace(strings.ToLower(tag))
				switch t {
				case "latest", "pending":
					n = reqN
				case "earliest":
					n = 0
				default:
					if strings.HasPrefix(t, "0x") {
						tt := strings.TrimPrefix(t, "0x")
						if tt != "" {
							if v, err := strconv.ParseUint(tt, 16, 64); err == nil {
								n = v
							}
						}
					} else {
						if v, err := strconv.ParseUint(t, 10, 64); err == nil {
							n = v
						}
					}
				}
			}
		}
		// If asked height is above current, return null (Ethereum-compatible)
		if s.n != nil && n > reqN {
			resp.Result = nil
			return resp
		}

		// M12: if persisted block metadata exists, expose real roots/bloom
		logsBloom := "0x" + strings.Repeat("0", 512)
		stateRoot := "0x" + strings.Repeat("0", 64)
		receiptsRoot := "0x" + strings.Repeat("0", 64)
		if s.n != nil && s.n.DB() != nil {
			key := []byte("blkmeta/v1/" + strings.TrimPrefix(toHexUint(n), "0x"))
			if b, err := s.n.DB().Get(key, nil); err == nil {
				var bm struct {
					LogsBloomHex string `json:"logsBloom"`
					StateRoot    string `json:"stateRoot"`
					ReceiptsRoot string `json:"receiptsRoot"`
				}
				if err := json.Unmarshal(b, &bm); err == nil {
					if strings.HasPrefix(bm.LogsBloomHex, "0x") && len(bm.LogsBloomHex) == 514 {
						logsBloom = bm.LogsBloomHex
					}
					if strings.HasPrefix(bm.StateRoot, "0x") && len(bm.StateRoot) == 66 {
						stateRoot = bm.StateRoot
					}
					if strings.HasPrefix(bm.ReceiptsRoot, "0x") && len(bm.ReceiptsRoot) == 66 {
						receiptsRoot = bm.ReceiptsRoot
					}
				}
			}
		}

		resp.Result = blockResp{
			Number:           toHexUint(n),
			Hash:             pseudoBlockHash(n).Hex(),
			ParentHash:       pseudoBlockHash(n - 1).Hex(),
			Nonce:            "0x0000000000000000",
			Sha3Uncles:       "0x" + strings.Repeat("0", 64),
			LogsBloom:        logsBloom,
			TransactionsRoot: "0x" + strings.Repeat("0", 64),
			StateRoot:        stateRoot,
			ReceiptsRoot:     receiptsRoot,
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

// EVM Chain ID (EIP-155) must be tooling/wallet compatible.
// NOTE: JavaScript tooling (Hardhat/Viem/ethers/MetaMask) requires chainId <= Number.MAX_SAFE_INTEGER.
// For NOORCHAIN 2.1 local dev, we pin this to 2121.
const evmChainID uint64 = 2121

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
		"eth_chainId":               {},
		"eth_blockNumber":           {},
		"eth_accounts":              {},
		"eth_getTransactionCount":   {},
		"eth_gasPrice":              {},
		"eth_estimateGas":           {},
		"eth_getBalance":            {},
		"eth_getCode":               {},
		"eth_getStorageAt":          {},
		"eth_call":                  {},
		"eth_sendRawTransaction":    {},
		"eth_getTransactionReceipt": {},
		"eth_getTransactionByHash":  {},
		"eth_getBlockByNumber":      {},
		"eth_getLogs":               {},

		// M15: filters
		"eth_newFilter":        {},
		"eth_newBlockFilter":   {},
		"eth_getFilterChanges": {},
		"eth_getFilterLogs":    {},
		"eth_uninstallFilter":  {},
	}

	for m := range ethRouting {
		if _, ok := handled[m]; !ok {
			panic("rpc: routing table references unhandled method: " + m)
		}
	}
}

type ioDiscard struct{}

func (ioDiscard) Write(p []byte) (int, error) { return len(p), nil }

// ---- geth vm helpers (M12.5) ----

// coreCanTransfer reports whether the account has enough balance for the transfer.
// Minimal implementation for read-only eth_call context.
func coreCanTransfer(db vm.StateDB, addr common.Address, amount *uint256.Int) bool {
	if amount == nil || amount.Sign() == 0 {
		return true
	}
	bal := db.GetBalance(addr)
	return bal.Cmp(amount) >= 0
}

// coreTransfer performs a balance transfer in the StateDB.
// Note: eth_call uses STATICCALL, so mutations should not persist; this is still required by vm.BlockContext.
func coreTransfer(db vm.StateDB, sender, recipient common.Address, amount *uint256.Int) {
	if amount == nil || amount.Sign() == 0 {
		return
	}
	db.SubBalance(sender, amount, 0)
	db.AddBalance(recipient, amount, 0)
}
