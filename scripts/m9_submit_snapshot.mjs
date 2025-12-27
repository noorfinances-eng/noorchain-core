import { ethers } from "ethers";

// RPC
const RPC = "http://127.0.0.1:8545";

// DEV KEY (funding not needed in our shim). Deterministic for repeatability.
const PRIV = "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d";

// Addresses (set these once you have them)
const POSS_REGISTRY = process.env.POSS_REGISTRY || ""; // 0x...

// Function selector snapshotCount(): 0x098ab6a1
const SEL_SNAPSHOT_COUNT = "0x098ab6a1";

async function main() {
  const provider = new ethers.JsonRpcProvider(RPC);
  const wallet = new ethers.Wallet(PRIV, provider);

  const chainId = await provider.send("eth_chainId", []);
  console.log("chainId =", chainId);
  console.log("from   =", await wallet.getAddress());

  if (!POSS_REGISTRY || !ethers.isAddress(POSS_REGISTRY)) {
    console.log("ERR: set POSS_REGISTRY env var to PoSSRegistry address, e.g.");
    console.log("  POSS_REGISTRY=0x... node scripts/m9_submit_snapshot.mjs");
    process.exit(2);
  }

  // Minimal submitSnapshot calldata according to your contract ABI encoding:
  // We'll use ethers Interface to encode it.
  const abi = [
    "function submitSnapshot((bytes32 snapshotHash,string uri,uint64 periodStart,uint64 periodEnd,uint32 version) meta, (uint8 v,bytes32 r,bytes32 s)[] sigs) returns (uint256)",
    "function snapshotCount() view returns (uint256)"
  ];
  const iface = new ethers.Interface(abi);

  const snapshotHash = ethers.keccak256(ethers.toUtf8Bytes("m9-test-snapshot-" + Date.now()));
  const meta = {
    snapshotHash,
    uri: "ipfs://m9-test",
    periodStart: 1n,
    periodEnd: 2n,
    version: 1
  };
  const sigs = []; // shim doesn't enforce signatures (it just parses/records)

  const data = iface.encodeFunctionData("submitSnapshot", [meta, sigs]);

  const tx = await wallet.sendTransaction({
    to: POSS_REGISTRY,
    data,
    gasLimit: 3_000_000n
  });

  console.log("txHash =", tx.hash);

  // Wait a bit for your node to "mine" (tick loop)
  for (let i = 0; i < 20; i++) {
    const rcpt = await provider.send("eth_getTransactionReceipt", [tx.hash]);
    if (rcpt) {
      console.log("receipt.status =", rcpt.status);
      console.log("receipt.blockNumber =", rcpt.blockNumber);
      break;
    }
    await new Promise(r => setTimeout(r, 500));
  }

  const countHex = await provider.send("eth_call", [{ to: POSS_REGISTRY, data: SEL_SNAPSHOT_COUNT }, "latest"]);
  console.log("snapshotCount(hex) =", countHex);
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
