import { network } from "hardhat";

function mustEnv(name) {
  const v = process.env[name];
  if (!v || v === "NOT_SET") throw new Error(`missing env ${name}`);
  return v;
}

async function main() {
  const conn = await network.connect();
  const { ethers } = conn;

  // Uses NOOR_PRIVATE_KEY just like your existing scripts.
  const pk = mustEnv("NOOR_PRIVATE_KEY");
  const wallet = new ethers.Wallet(pk, ethers.provider);

  const bal = await ethers.provider.getBalance(wallet.address);
  console.log("Signer:", wallet.address);
  console.log("Balance:", bal.toString());

  const F = await ethers.getContractFactory("MultiEmit", wallet);
  const c = await F.deploy();
  await c.waitForDeployment();
  const addr = await c.getAddress();
  console.log("MultiEmit:", addr);

  const tx = await c.emitTwo(111, 222);
  console.log("emitTwo tx:", tx.hash);
  const rcpt = await tx.wait();
  console.log("receipt status:", rcpt.status);

  // Print logs (ethers v6 format)
  console.log("receipt logs:", rcpt.logs.length);
  for (let i = 0; i < rcpt.logs.length; i++) {
    const l = rcpt.logs[i];
    console.log(`#${i}`, {
      address: l.address,
      topics0: l.topics?.[0],
      logIndex: l.index,
      txIndex: l.transactionIndex,
      blockNumber: l.blockNumber,
    });
  }
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
