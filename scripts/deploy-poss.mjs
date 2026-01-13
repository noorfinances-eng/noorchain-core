import { network } from "hardhat";

async function main() {
  const conn = await network.connect();
  const { ethers } = conn;

  const [deployer] = await ethers.getSigners();
  console.log("Deployer:", deployer.address);

  const CuratorSet = await ethers.getContractFactory("CuratorSet");
  const curatorSet = await CuratorSet.deploy(
  deployer.address,
  [deployer.address],
  1,
  { gasLimit: 15_000_000 }
);
  await curatorSet.waitForDeployment();
  const curatorSetAddr = await curatorSet.getAddress();
  console.log("CuratorSet:", curatorSetAddr);

  const PoSSRegistry = await ethers.getContractFactory("PoSSRegistry");
 const registry = await PoSSRegistry.deploy(
  curatorSetAddr,
  { gasLimit: 15_000_000 }
);
  await registry.waitForDeployment();
  const registryAddr = await registry.getAddress();
  console.log("PoSSRegistry:", registryAddr);

  const t = await curatorSet.getThreshold();
  const c = await curatorSet.getCuratorCount();
  console.log("Threshold:", t.toString(), "CuratorCount:", c.toString());
  console.log("Done.");
}

main().catch((e) => {
  console.error(e);
  process.exitCode = 1;
});
