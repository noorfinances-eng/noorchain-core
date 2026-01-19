import { readFileSync } from "node:fs";
import { JsonRpcProvider, Wallet, ContractFactory } from "ethers";

const RPC_URL = "http://127.0.0.1:8545";

// DEV ONLY private key (Hardhat default #0)
const DEV_PRIVKEY = "0x59c6995e998f97a5a0044966f094538f2f0f9b8677183762b2f279b9da3c8e8b";

function loadArtifact(relPath) {
  const raw = readFileSync(relPath, "utf8");
  const j = JSON.parse(raw);
  return { abi: j.abi, bytecode: j.bytecode };
}

async function main() {
  const provider = new JsonRpcProvider(RPC_URL);
  const wallet = new Wallet(DEV_PRIVKEY, provider);

  const chainIdHex = await provider.send("eth_chainId", []);
  console.log("RPC chainId:", chainIdHex);

  const deployerAddr = await wallet.getAddress();
  console.log("Deployer:", deployerAddr);

  // IMPORTANT: force nonces explicitly to avoid provider caching/coalescing
  let nonce = await provider.getTransactionCount(deployerAddr, "pending");

  const curatorArtifact = loadArtifact("artifacts/contracts/poss/CuratorSet.sol/CuratorSet.json");
  const registryArtifact = loadArtifact("artifacts/contracts/poss/PoSSRegistry.sol/PoSSRegistry.json");

  // 1) Deploy CuratorSet(admin=deployer, initialCurators=[deployer], threshold=1)
  const CuratorFactory = new ContractFactory(curatorArtifact.abi, curatorArtifact.bytecode, wallet);
  const curatorTx = await CuratorFactory.getDeployTransaction(deployerAddr, [deployerAddr], 1);

  curatorTx.nonce = nonce++;
  curatorTx.gasLimit = 3_000_000n;
  curatorTx.gasPrice = 1n;

  const sent1 = await wallet.sendTransaction(curatorTx);
  console.log("CuratorSet tx:", sent1.hash);

  const rcpt1 = await provider.waitForTransaction(sent1.hash, 1, 30_000);
  if (!rcpt1 || !rcpt1.contractAddress) throw new Error("No receipt/contractAddress for CuratorSet");
  const curatorAddr = rcpt1.contractAddress;
  console.log("CuratorSet:", curatorAddr);

  // 2) Deploy PoSSRegistry(curatorSet)
  const RegistryFactory = new ContractFactory(registryArtifact.abi, registryArtifact.bytecode, wallet);
  const regTx = await RegistryFactory.getDeployTransaction(curatorAddr);

  regTx.nonce = nonce++;
  regTx.gasLimit = 3_000_000n;
  regTx.gasPrice = 1n;

  const sent2 = await wallet.sendTransaction(regTx);
  console.log("PoSSRegistry tx:", sent2.hash);

  const rcpt2 = await provider.waitForTransaction(sent2.hash, 1, 30_000);
  if (!rcpt2 || !rcpt2.contractAddress) throw new Error("No receipt/contractAddress for PoSSRegistry");
  const registryAddr = rcpt2.contractAddress;
  console.log("PoSSRegistry:", registryAddr);

  console.log("DEPLOY_OK");
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
