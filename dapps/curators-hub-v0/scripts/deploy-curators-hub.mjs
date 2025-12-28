import hre from "hardhat";
import { defineChain, http, createPublicClient, createWalletClient } from "viem";
import { privateKeyToAccount } from "viem/accounts";

const RPC_URL = "http://127.0.0.1:8545";

const noorchain2121 = defineChain({
  id: 2121,
  name: "NOORCHAIN 2.1 Local",
  nativeCurrency: { name: "NUR", symbol: "NUR", decimals: 18 },
  rpcUrls: { default: { http: [RPC_URL] } },
});

// Tooling deployer key (local only). Override with env var if you want.
const DEPLOYER_PK =
  process.env.NOOR_DEPLOYER_PK ??
  "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"; // hardhat default dev key

async function main() {
  console.log("Deploying Curators Hub v0 contracts...");

  const account = privateKeyToAccount(DEPLOYER_PK);

  const publicClient = createPublicClient({
    chain: noorchain2121,
    transport: http(RPC_URL),
  });

  const walletClient = createWalletClient({
    account,
    chain: noorchain2121,
    transport: http(RPC_URL),
  });

  // Safety: ensure we are on the expected chain
  const liveChainId = await publicClient.getChainId();
  if (liveChainId !== 2121) {
    throw new Error(`Connected to unexpected chainId=${liveChainId} (expected 2121)`);
  }
  console.log("Connected chainId:", liveChainId);
  console.log("Deployer:", account.address);

  // Read artifacts (ABI + bytecode)
  const CuratorSetArt = await hre.artifacts.readArtifact("CuratorSet");
  const PoSSRegistryArt = await hre.artifacts.readArtifact("PoSSRegistry");

  // --- Deploy CuratorSet ---
  const curatorSetHash = await walletClient.deployContract({
    abi: CuratorSetArt.abi,
    bytecode: CuratorSetArt.bytecode,
    args: [],
  });
  const curatorSetRcpt = await publicClient.waitForTransactionReceipt({ hash: curatorSetHash });
  console.log("CuratorSet tx:", curatorSetHash);
  console.log("CuratorSet deployed at:", curatorSetRcpt.contractAddress);

  // --- Deploy PoSSRegistry(curatorSet) ---
  const possRegistryHash = await walletClient.deployContract({
    abi: PoSSRegistryArt.abi,
    bytecode: PoSSRegistryArt.bytecode,
    args: [curatorSetRcpt.contractAddress],
  });
  const possRegistryRcpt = await publicClient.waitForTransactionReceipt({ hash: possRegistryHash });
  console.log("PoSSRegistry tx:", possRegistryHash);
  console.log("PoSSRegistry deployed at:", possRegistryRcpt.contractAddress);
}

main().catch((err) => {
  console.error(err);
  process.exitCode = 1;
});
