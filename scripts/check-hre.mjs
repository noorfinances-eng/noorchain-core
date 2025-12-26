export default async function (_taskArgs, hre) {
  console.log("HRE injected:", !!hre);
  console.log("network name:", hre.network?.name);
  console.log("ethers exists:", typeof hre.ethers);

  if (hre.ethers) {
    const signers = await hre.ethers.getSigners();
    console.log("signers:", signers.map(s => s.address));
  }
}
