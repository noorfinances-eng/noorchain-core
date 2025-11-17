# NOORCHAIN Core — Developer Setup (Draft)

This document explains how developers can work with the NOORCHAIN Core
repository. It will be expanded as the project evolves.

---

## 1. Requirements (for developers working locally)

Although the project owner works entirely online (GitHub + CI only),  
developers who want to run the code locally will need:

- Go `1.21+`
- Git

Later (not required now):
- Protobuf compiler (`protoc`)
- `buf` tool
- Cosmos SDK codegen tools

---

## 2. Clone the Repository

```bash
git clone https://github.com/noorfinances-eng/noorchain-core.git
cd noorchain-core
