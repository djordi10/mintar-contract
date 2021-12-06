const { ethers, upgrades } = require('hardhat');

async function main () {
  const Core = await ethers.getContractFactory('MintarCore');
  const Sale = await ethers.getContractFactory('MintarSale')
  console.log('Deploying Contracts...');
  const core = await upgrades.deployProxy(Core);
  const sale = await upgrades.deployProxy(Sale);
  await sale.deployed();
  await core.deployed();
  console.log('MintarCore deployed to:', core.address);
  console.log('MintarSale deployed to:', sale.address);
}

main();