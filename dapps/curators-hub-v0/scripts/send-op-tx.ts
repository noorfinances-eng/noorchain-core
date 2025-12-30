import fs from "node:fs";

import {
  createPublicClient,
  createWalletClient,
  http,
  encodeFunctionData,
  parseAbi,
  keccak256,
  toHex,
} from "viem";
import { privateKeyToAccount } from "viem/accounts";

const DEPLOY_PATH = new URL("../deployments/noorchain-2.1-local.json", import.meta.url);

function loadDeploy() {
  const raw = fs.readFileSync(DEPLOY_PATH, "utf8");
  return JSON.parse(raw);
}

function requireEnv(name: string) {
  const v = process.env[name];
  if (!v || !v.startsWith("0x")) {
    throw new Error(`ERROR: ${name} env not set. Example: export ${name}=0x... `);
  }
  return v as `0x${string}`;
}

const PK = requireEnv("PK");
const deploy = loadDeploy();

const RPC = (deploy.rpc ?? "http://127.0.0.1:8545") as string;
const CHAIN_ID = (deploy.chainId ?? 2121) as number;

const possRegistry = deploy?.contracts?.PoSSRegistry?.address as string;
if (!possRegistry || !possRegistry.startsWith("0x")) {
  throw new Error("ERROR: deployments file missing contracts.PoSSRegistry.address");
}

const chain = {
  id: CHAIN_ID,
  name: deploy.network ?? "noorchain-2.1-local",
  nativeCurrency: { name: "NUR", symbol: "NUR", decimals: 18 },
  rpcUrls: { default: { http: [RPC] } },
} as const;

const account = privateKeyToAccount(PK);

const publicClient = createPublicClient({ chain, transport: http(RPC) });
const walletClient = createWalletClient({ chain, transport: http(RPC), account });

// Minimal ABI we need
const abi = parseAbi([
  "function submitSnapshot((bytes32 snapshotHash,string uri,uint64 periodStart,uint64 periodEnd,uint32 version) meta,(uint8 v,bytes32 r,bytes32 s)[] sigs)",
  "function snapshotCount() view returns (uint256)",
  "function latestSnapshotId() view returns (uint256)",
]);

// Build a deterministic snapshotHash
const now = BigInt(Math.floor(Date.now() / 1000));
const snapshotHash = keccak256(toHex(now));

// Minimal meta (the chain mock/shim accepts this path)
const meta = {
  snapshotHash,
  uri: "ipfs://noorchain/poss/snapshot/demo",
  periodStart: Number(now - 60n),
  periodEnd: Number(now),
  version: 1,
};

// Empty sigs (dev-only path)
const sigs: Array<{ v: number; r: `0x${string}`; s: `0x${string}` }> = [];

const data = encodeFunctionData({
  abi,
  functionName: "submitSnapshot",
  args: [meta, sigs],
});

console.log("NOORCHAIN submitSnapshot");
console.log("RPC:", RPC);
console.log("chainId:", CHAIN_ID);
console.log("from:", account.address);
console.log("to (PoSSRegistry):", possRegistry);

const hash = await walletClient.sendTransaction({
  to: possRegistry as `0x${string}`,
  data,
});

console.log("tx:", hash);

const rcpt = await publicClient.waitForTransactionReceipt({ hash });
console.log("receipt.status:", rcpt.status);
console.log("receipt.blockNumber:", rcpt.blockNumber);

