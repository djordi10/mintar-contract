// scripts/upgrade_box.js
const { ethers, upgrades } = require('hardhat');

async function main () {
  const core = await ethers.getContractFactory('MintarCore');
  const sale = await ethers.getContractFactory('MintarSale');
  console.log('Upgrading Contract...');
  await upgrades.upgradeProxy('0x2E8747Ae062a417f6dA677DfBC58D1D99804a6F5', core);
  await upgrades.upgradeProxy('0xaD1d1b4fc23c9ac14eD3eb4E1c2c1a201CBE0c7d', sale);
  console.log('Contract upgraded');
}

main();