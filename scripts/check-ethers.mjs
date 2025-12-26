import hre from "hardhat/hre";

async function main() {
  console.log("HRE OK");
  console.log("network name:", hre.network?.name);
  console.log("ethers exists:", typeof hre.ethers);

  if (hre.ethers) {
    const signers = await hre.ethers.getSigners();
    console.log("signers:", signers.map(s => s.address));
  }
}

main().catch((e) => {
  console.error(e);
  process.exitCode = 1;
});
