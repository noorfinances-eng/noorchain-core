import {
  createPublicClient,
  http,
  keccak256,
  stringToHex,
  toHex,
} from "viem";

const RPC = "http://127.0.0.1:8545";
const REG = "0xe7f1725e7734ce288f8367e1bb143e90bb3f0512";

const abi = [
  {
    type: "function",
    name: "snapshotCount",
    stateMutability: "view",
    inputs: [],
    outputs: [{ type: "uint256" }],
  },
  {
    type: "function",
    name: "latestSnapshotId",
    stateMutability: "view",
    inputs: [],
    outputs: [{ type: "uint256" }],
  },
  {
    type: "function",
    name: "getSnapshot",
    stateMutability: "view",
    inputs: [{ name: "id", type: "uint256" }],
    outputs: [
      {
        name: "",
        type: "tuple",
        components: [
          { name: "snapshotHash", type: "bytes32" },
          { name: "uri", type: "string" },
          { name: "periodStart", type: "uint64" },
          { name: "periodEnd", type: "uint64" },
          { name: "publishedAt", type: "uint64" },
          { name: "version", type: "uint32" },
          { name: "publisher", type: "address" },
        ],
      },
    ],
  },
];

const client = createPublicClient({ transport: http(RPC) });

console.log("RPC:", RPC);
console.log("PoSSRegistry:", REG);

const count = await client.readContract({ address: REG, abi, functionName: "snapshotCount" });
const latest = await client.readContract({ address: REG, abi, functionName: "latestSnapshotId" });

console.log("snapshotCount:", count);
console.log("latestSnapshotId:", latest);

// Compute selector + calldata precisely (no guessing)
const sig = "getSnapshot(uint256)";
const selector = keccak256(stringToHex(sig)).slice(0, 10);
const arg2 = toHex(2n, { size: 32 }); // 32-byte ABI word
const calldata = selector + arg2.slice(2);

console.log("selector(getSnapshot):", selector);
console.log("calldata getSnapshot(2):", calldata);

// Raw eth_call (so we can compare with curl)
const raw = await client.request({
  method: "eth_call",
  params: [{ to: REG, data: calldata }, "latest"],
});

console.log("eth_call raw:", raw);

// Decoded readContract (if raw is not 0x)
try {
  const snap2 = await client.readContract({
    address: REG,
    abi,
    functionName: "getSnapshot",
    args: [2n],
  });
  console.log("getSnapshot(2) decoded:", snap2);
} catch (e) {
  console.error("getSnapshot(2) readContract failed:", e?.message ?? e);
  process.exit(1);
}
