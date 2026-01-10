import { defineConfig } from "hardhat/config";
import hardhatEthers from "@nomicfoundation/hardhat-ethers";

export default defineConfig({
  plugins: [hardhatEthers],
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
      accounts: process.env.NOOR_PRIVATE_KEY ? [process.env.NOOR_PRIVATE_KEY] : [],
    },
  },
});
