import { JsonRpcProvider, Wallet, Contract, keccak256, toUtf8Bytes, getBytes, Signature } from "ethers";
import { readFileSync } from "node:fs";

const RPC_URL = "http://127.0.0.1:8545";

// DEV ONLY private key (Hardhat default #0)
const DEV_PRIVKEY = "0x59c6995e998f97a5a0044966f094538f2f0f9b8677183762b2f279b9da3c8e8b";

// IMPORTANT: update this if you redeployed
const REGISTRY = "0x4716234A82C888198a6BD0059a437A423bDE5eD1";

function loadABI(relPath) {
  const raw = readFileSync(relPath, "utf8");
  return JSON.parse(raw).abi;
}

async function main() {
  const provider = new JsonRpcProvider(RPC_URL);
  const wallet = new Wallet(DEV_PRIVKEY, provider);

  const signer = await wallet.getAddress();
  console.log("Signer(curator):", signer);

  const abi = loadABI("artifacts/contracts/poss/PoSSRegistry.sol/PoSSRegistry.json");
  const registry = new Contract(REGISTRY, abi, wallet);

  // ---- meta (V0 minimal) ----
  const uri = "ipfs://example-epoch-1";
  const periodStart = BigInt(Math.floor(Date.now() / 1000));   // uint64
  const periodEnd = periodStart + 3600n;                       // +1h (uint64)
  const version = 1;                                           // uint32

  // Deterministic snapshot hash (bytes32)
  const snapshotHash = keccak256(
    toUtf8Bytes(`noorchain-poss-v0|${uri}|${periodStart}|${periodEnd}|${version}`)
  );

  const meta = {
    snapshotHash,
    uri,
    periodStart,
    periodEnd,
    version,
  };

  console.log("meta.snapshotHash:", snapshotHash);
  console.log("meta.uri:", uri);
  console.log("meta.periodStart:", periodStart.toString());
  console.log("meta.periodEnd:", periodEnd.toString());
  console.log("meta.version:", String(version));

  // ---- signature: sign snapshotHash bytes (dev choice) ----
  const sigHex = await wallet.signMessage(getBytes(snapshotHash));
  const sig = Signature.from(sigHex);

  const sigs = [{
    v: sig.v,     // uint8
    r: sig.r,     // bytes32
    s: sig.s,     // bytes32
  }];

  console.log("sig.v:", sig.v);
  console.log("sig.r:", sig.r);
  console.log("sig.s:", sig.s);

  // Force nonce explicitly
  const nonce = await provider.getTransactionCount(signer, "pending");

  const tx = await registry.submitSnapshot(meta, sigs, {
    nonce,
    gasLimit: 3_000_000n,
    gasPrice: 1n,
  });
  console.log("submitSnapshot tx:", tx.hash);

  const rcpt = await provider.waitForTransaction(tx.hash, 1, 30_000);
  console.log("receipt status:", rcpt?.status);
  console.log("SUBMIT_OK");
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
