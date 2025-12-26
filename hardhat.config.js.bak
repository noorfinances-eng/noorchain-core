import "@nomicfoundation/hardhat-ethers";

/** @type {import("hardhat/config").HardhatUserConfig} */
const config = {
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

export default config;
