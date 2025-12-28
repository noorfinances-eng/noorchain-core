import { defineConfig } from "hardhat/config";
import hardhatViem from "@nomicfoundation/hardhat-viem";

const NOORCHAIN_2_1 = {
  id: 2121,
  name: "NOORCHAIN 2.1 Local",
  nativeCurrency: { name: "NUR", symbol: "NUR", decimals: 18 },
  rpcUrls: { default: { http: ["http://127.0.0.1:8545"] } },
};

export default defineConfig({
  plugins: [hardhatViem],

  solidity: {
    version: "0.8.28",
    settings: { optimizer: { enabled: true, runs: 200 }, viaIR: true },
  },

  networks: {
    localhost: {
      type: "http",
      chainType: "l1",
      url: "http://127.0.0.1:8545",
      chainId: 2121,
    },
  },

  viem: {
    chains: [NOORCHAIN_2_1],
  },
});
