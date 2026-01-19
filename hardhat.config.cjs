require("@nomicfoundation/hardhat-ethers");
const { task } = require("hardhat/config");

/** Debug task: verifies HRE + ethers injection */
task("check-hre", "Prints HRE injection info").setAction(async (_, hre) => {
  console.log("HRE injected:", !!hre);
  console.log("network name:", hre.network?.name);
  console.log("ethers exists:", typeof hre.ethers);
  if (hre.ethers) {
    const signers = await hre.ethers.getSigners();
    console.log("signers:", signers.map((s) => s.address));
  }
});

/** @type import("hardhat/config").HardhatUserConfig */
module.exports = {
  solidity: {
    version: "0.8.20",
    settings: {
      viaIR: true,
      optimizer: { enabled: true, runs: 200 },
    },
  },
  networks: {
    noorcore: {
      type: "http",
      url: "http://127.0.0.1:8545",
    },
  },
};
