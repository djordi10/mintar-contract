const { v4: uuidv4 } = require("uuid");

const chai = require('chai');
const chaiAsPromised = require('chai-as-promised');
chai.use(chaiAsPromised);
chai.should();
const { expect } = chai;


describe("utility function", function () {
    it("token_id hash function", async function () {
        const id = ethers.utils.toUtf8Bytes("0b3dfee2-a96b-4245-9f8f-1bb5c516b2b9"); 
        const encoded = ethers.utils.defaultAbiCoder.encode(["bytes"], [id]);
        const tokenId = ethers.utils.keccak256(encoded);
        expect(ethers.BigNumber.from(tokenId).toBigInt().toString()).to.equal("34687554082616307266410521928283957738637311823920088884725935543632266907710");
    });
})

describe("core contract", function () {
    let core;
    let owner;

    before(async function(){
        [owner] = await ethers.getSigners();
        const Core = await ethers.getContractFactory("MintarCore");
        core = await Core.deploy();
    });

    it("create simple token and check balance", async function () {
        const id = ethers.utils.toUtf8Bytes(uuidv4()); 
        await core.create(id, 10, 100);
        const encoded = ethers.utils.defaultAbiCoder.encode(["bytes"], [id]);
        const tokenId = ethers.utils.keccak256(encoded);

        const b = await core.balanceOf(owner.address, ethers.BigNumber.from(tokenId).toBigInt());
        expect(ethers.BigNumber.from(b).toNumber()).to.equal(10);
    });

    it("create simple token and check creator", async function () {
        const id = ethers.utils.toUtf8Bytes(uuidv4()); 
        await core.create(id, 10, 100);
        const encoded = ethers.utils.defaultAbiCoder.encode(["bytes"], [id]);
        const tokenId = ethers.utils.keccak256(encoded);

        const creator = await core.creator(ethers.BigNumber.from(tokenId).toBigInt());
        expect(creator).to.equal(owner.address);
    });
});

describe("sale contract", function () {
    let core;
    let sale;
    let owner;
    let seller;
    let buyer;

    before(async function(){
        [owner, seller, buyer] = await ethers.getSigners();
        const Core = await ethers.getContractFactory("MintarCore");
        core = await Core.deploy();
        const Sale = await ethers.getContractFactory("MintarSale");
        sale = await Sale.deploy();
    });

    it("createAndSell simple token and check if it's on sale", async function () {
        const saleContract = await sale.deployed();
        const coreContract = await core.deployed();
        const id = ethers.utils.toUtf8Bytes(uuidv4()); 
        await core.connect(seller).setApprovalForAll(saleContract.address, true);
        await core.connect(seller).createAndSell(id, 10, 100, ethers.utils.parseEther("10"), saleContract.address);
        const encoded = ethers.utils.defaultAbiCoder.encode(["bytes"], [id]);
        const tokenId = ethers.utils.keccak256(encoded);

        const onSale = await sale.connect(seller).isOnSale(coreContract.address, ethers.BigNumber.from(tokenId).toBigInt());
        expect(onSale).to.equal(true);
    });
    it("create sale", async function () {  
        const id = ethers.utils.toUtf8Bytes(uuidv4()); 
        await core.connect(seller).create(id, 10, 100);
        const encoded = ethers.utils.defaultAbiCoder.encode(["bytes"], [id]);
        const tokenId = ethers.utils.keccak256(encoded);


        const coreContract = await core.deployed();
        const saleContract = await sale.deployed();
        await core.connect(seller).setApprovalForAll(saleContract.address, true);
        await sale.connect(seller).createSale(coreContract.address, ethers.BigNumber.from(tokenId).toBigInt(), 250000, 1);
        const onSale = await sale.connect(seller).isOnSale(coreContract.address, ethers.BigNumber.from(tokenId).toBigInt());
        expect(onSale).to.equal(true);
    });
    it("create sale and cancel it", async function () {  
        const id = ethers.utils.toUtf8Bytes(uuidv4()); 
        await core.connect(seller).create(id, 10, 100);
        const encoded = ethers.utils.defaultAbiCoder.encode(["bytes"], [id]);
        const tokenId = ethers.utils.keccak256(encoded);

        const coreContract = await core.deployed();
        const saleContract = await sale.deployed();
        await core.connect(seller).setApprovalForAll(saleContract.address, true);
        await sale.connect(seller).createSale(coreContract.address, ethers.BigNumber.from(tokenId).toBigInt(), 250000, 1);
        await sale.connect(seller).cancelSale(coreContract.address, ethers.BigNumber.from(tokenId).toBigInt(), 1);
        const onSale = await sale.connect(seller).isOnSale(coreContract.address, ethers.BigNumber.from(tokenId).toBigInt());
        expect(onSale).to.equal(false);
    });
    it("create sale and purchase it", async function () {  
        const id = ethers.utils.toUtf8Bytes(uuidv4()); 
        await core.connect(seller).create(id, 10, 100);
        const encoded = ethers.utils.defaultAbiCoder.encode(["bytes"], [id]);
        const tokenId = ethers.utils.keccak256(encoded);

        const coreContract = await core.deployed();
        const saleContract = await sale.deployed();
        await core.connect(seller).setApprovalForAll(saleContract.address, true);
        await sale.connect(seller).createSale(coreContract.address, ethers.BigNumber.from(tokenId).toBigInt(), ethers.utils.parseEther("10"), 1);
        const onSale = await sale.connect(seller).isOnSale(coreContract.address, ethers.BigNumber.from(tokenId).toBigInt());
        expect(onSale).to.equal(true);

        await sale.connect(buyer).purchase(coreContract.address, ethers.BigNumber.from(tokenId).toBigInt(), 1, { value: ethers.utils.parseEther("10") });
        // balanceBefore is always 10,000.0, balanceAfter should be 10,000.0 - 10.0 - gas fee
        const balanceAfter = await ethers.provider.getBalance(buyer.address);
        expect(balanceAfter.lt(ethers.utils.parseEther("9990"))).to.be.equal(true);

        const stillOnSale = await sale.connect(seller).isOnSale(coreContract.address, ethers.BigNumber.from(tokenId).toBigInt());
        expect(stillOnSale).to.equal(false);
    });
    it("create sale and purchase it with insufficient ether", async function () {  
        const id = ethers.utils.toUtf8Bytes(uuidv4()); 
        await core.connect(seller).create(id, 10, 100);
        const encoded = ethers.utils.defaultAbiCoder.encode(["bytes"], [id]);
        const tokenId = ethers.utils.keccak256(encoded);

        const coreContract = await core.deployed();
        const saleContract = await sale.deployed();
        await core.connect(seller).setApprovalForAll(saleContract.address, true);
        await sale.connect(seller).createSale(coreContract.address, ethers.BigNumber.from(tokenId).toBigInt(), ethers.utils.parseEther("10"), 1);
        const onSale = await sale.connect(seller).isOnSale(coreContract.address, ethers.BigNumber.from(tokenId).toBigInt());
        expect(onSale).to.equal(true);

        const purchase = sale.connect(buyer).purchase(coreContract.address, ethers.BigNumber.from(tokenId).toBigInt(), 1, { value: ethers.utils.parseEther("5") });
        expect(purchase).eventually.to.rejectedWith(Error, "VM Exception while processing transaction: reverted with reason string 'insufficient ether'");
    });
});