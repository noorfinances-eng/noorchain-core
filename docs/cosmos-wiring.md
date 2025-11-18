# NOORCHAIN Core — Cosmos Wiring Overview (Draft)

This document explains how the NOORCHAIN application is wired internally
using Cosmos SDK and the BaseApp abstraction.

> Status: draft — mirrors the current code structure and will be updated
> as modules, encoding and Ethermint are integrated.

---

## 1. High-Level Flow

The NOORCHAIN node binary (`noord`) will, in its full form, follow this flow:

1. Configure Cosmos SDK global settings (Bech32 prefixes, etc.)
2. Create a logger
3. Open a database
4. Build the NOORCHAIN Cosmos application (BaseApp + modules)
5. Start the ABCI server with CometBFT

In code, this is split into several layers:

- `cmd/noord/main.go`
- `app.ConfigureSDK()`
- `app.NewNoorchainAppWithCosmos(...)`
- `app.NewAppBuilder(...)`
- `(*AppBuilder).BuildBaseApp()`
- `App` struct embedding `*baseapp.BaseApp`

---

## 2. Entry Point: `cmd/noord/main.go` (current + future)

Current minimal flow:

- call `app.ConfigureSDK()`
- create a simple placeholder app via `app.NewNoorchainApp()`
- call `Start()` (currently prints a message)

Future flow (Cosmos-based):

- call `app.ConfigureSDK()`
- create a logger via `app.NewNoorchainLogger()`
- open the application database
- call `app.NewNoorchainAppWithCosmos(...)`
- start the ABCI server and CometBFT

---

## 3. The `App` struct

Defined in `app/app.go`:

- wraps a `*baseapp.BaseApp`
- holds basic metadata: `Name`, `Version`
- will eventually embed all keepers and module management logic

```go
type App struct {
    *baseapp.BaseApp

    Name    string
    Version string
}
