import hre from "hardhat";

async function main() {
  const [deployer] = await hre.ethers.getSigners();
  console.log("Deployer:", deployer.address);

  const CuratorSet = await hre.ethers.getContractFactory("CuratorSet");
  const curatorSet = await CuratorSet.deploy(deployer.address, [deployer.address], 1);
  await curatorSet.waitForDeployment();
  const curatorSetAddr = await curatorSet.getAddress();
  console.log("CuratorSet:", curatorSetAddr);

  const PoSSRegistry = await hre.ethers.getContractFactory("PoSSRegistry");
  const registry = await PoSSRegistry.deploy(curatorSetAddr);
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
