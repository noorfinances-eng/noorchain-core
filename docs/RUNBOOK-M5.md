 NOORCHAIN 2.1 — PHASE 2 — M5 Runbook (Local Ops)

Document ID: RUNBOOK-M5  
Scope: Local, mainnet-like operator routine for two-node P2P connectivity (Node1 + Node2)  
Network Model: Minimal TCP-based P2P session persistence with bootstrap peers  
Safety: Local-only endpoints; no public exposure

---

## 0. Operator Rules

### Terminal Discipline
- **T1**: Node1 (RUNNING, do not use Ctrl+C unless explicitly stopping Node1)
- **T2**: Node2 (RUNNING, allowed to stop/restart)
- **T3**: Observation (ss/ps/grep/curl)

### Forbidden Commands (in this build)
These modes are not reliable and may start listeners unexpectedly:
- `./noorcore tendermint show-node-id ...`
- `./noorcore init ...` (if it opens listeners in your current binary)

Only use the explicit runtime flags:
- `-chain-id`
- `-data-dir`
- `-p2p-addr`
- `-boot-peers`

---

## 1. Paths and Ports (Current Local Standard)

### Node Homes (examples)
- Node1 data: `$HOME/.noor-core-fresh-1766680463`
- Node2 data: `$HOME/.noor-core-node2`

### P2P Ports
- Node1: `127.0.0.1:30303`
- Node2: `127.0.0.1:30304`
- Node2 bootstrap: `-boot-peers 127.0.0.1:30303`

---

## 2. Build and Format

From repo root:

```bash
gofmt -w core/main.go core/network/network.go
go build -o noorcore ./core
Expected:

Build completes without errors.

Binary ./noorcore is produced/updated.

3. Start Procedure (Two Nodes)
3.1 Start Node1 (T1)
bash
Copier le code
./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir "$HOME/.noor-core-fresh-1766680463" \
  -p2p-addr 127.0.0.1:30303
Expected logs:

[main] ... booting

[network] ... listening on 127.0.0.1:30303

3.2 Start Node2 (T2)
bash
Copier le code
./noorcore \
  -chain-id noorchain-2-1-local \
  -data-dir "$HOME/.noor-core-node2" \
  -p2p-addr 127.0.0.1:30304 \
  -boot-peers 127.0.0.1:30303
Expected logs:

[network] ... listening on 127.0.0.1:30304

[p2p] ... dialing peer 127.0.0.1:30303

[p2p] ... peer registered ... (peers=1)

4. Validation Gates (M5-A / M5-B)
4.1 Confirm listeners (T3)
bash
Copier le code
ss -ltnp | egrep ':30303|:30304' || true
Expected:

LISTEN on 30303 (Node1)

LISTEN on 30304 (Node2)

4.2 Confirm P2P session is established (T3)
bash
Copier le code
ss -tnp | egrep ':30303|:30304' || true
Pass criteria:

At least one ESTAB line between Node1 and Node2.

4.3 60-second stability check (T3)
bash
Copier le code
sleep 60; ss -tnp | egrep ':30303|:30304' || true
Pass criteria:

ESTAB remains present.

4.4 Restart Node2 resilience test (M5-B.2)
In T2, stop Node2 with Ctrl+C.

Immediately relaunch Node2 with the same command in section 3.2.

In T3:

bash
Copier le code
ss -tnp | egrep ':30303|:30304' || true
Pass criteria:

New ESTAB session appears (it may use a different ephemeral client port).

4.5 TCP hygiene snapshot (T3)
bash
Copier le code
ss -tan | awk '{print $1}' | sort | uniq -c
Interpretation:

A few TIME-WAIT are normal after restart.

CLOSE-WAIT should not continuously grow.

If CLOSE-WAIT persists, wait 20 seconds and re-check:

bash
Copier le code
sleep 20; ss -tan | awk '{print $1}' | sort | uniq -c
5. Stop Procedure
5.1 Stop Node2 (T2)
Press Ctrl+C in T2.

5.2 Stop Node1 (T1)
Press Ctrl+C in T1.

5.3 Emergency cleanup (use only if ports are stuck)
bash
Copier le code
pkill -INT noorcore || true
sleep 1
ss -ltnp | egrep ':30303|:30304' || true
Expected:

No listeners remain on 30303/30304.

6. Common Incidents
Incident A: “address already in use”
Meaning: a process is still bound to the port.
Actions (T3):

bash
Copier le code
ss -ltnp | egrep ':30303|:30304' || true
pgrep -a noorcore || true
Then stop the offending PID or run emergency cleanup.

Incident B: FIN-WAIT-2 / CLOSE-WAIT after restart
Meaning: previous TCP session is closing; normal if you recently stopped a node.
Action: wait 10–30 seconds; confirm a new ESTAB exists.

Incident C: Bootstrap logs show “peer registered” but no ESTAB
Meaning: the session did not persist (node was interrupted, or accept loop not active).
Action:

Ensure Node1 is running with the correct binary.

Ensure you did not interrupt Node2.

Re-check with section 4.2 and 4.4.

7. Change Control (M5 Discipline)
Allowed changes in M5:

Bootstrap flag and minimal P2P session management

Stability and restart resilience

Runbook updates

Logging and diagnostics

Not allowed in M5:

New features unrelated to hardening

Architectural refactors

Renaming interfaces without necessity