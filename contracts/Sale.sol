// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.2;

import "@openzeppelin/contracts-upgradeable/token/ERC1155/extensions/ERC1155SupplyUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC1155/IERC1155Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC1155/IERC1155ReceiverUpgradeable.sol";
import "./Core.sol";

contract MintarSale is IERC1155ReceiverUpgradeable {
    event SaleCreated(address indexed _contractAddress, uint256 indexed _tokenId, uint128 _price, address _seller, uint256 _amount);
    event SaleSuccessful(address indexed _contractAddress, uint256 indexed _tokenId, uint128 _price, address _buyer, address _seller, uint256 _amount);
    event SaleCanceled(address indexed _contractAddress, uint256 _tokenId);
    
    // Represents Item on sake
    struct Sale {
        // current owner of item
        address seller;
        // Price (in wei)
        uint128 price;
        // Amount min 1
        uint256 amount;
    }
    
    // Maps of items on Sale
    mapping (address => mapping (uint256 => Sale)) public sales;
    
    // Service Fee
    // Values 0-10,000 map to 0%-100%
    uint256 public serviceFee;

    // to statisfy ERC1155ReceiverUpgradeable interface
    function onERC1155Received(
        address, address, uint256, uint256, bytes calldata
    ) override external pure returns (bytes4) {
        return bytes4(keccak256("onERC1155Received(address,address,uint256,uint256,bytes)"));
    }
    function onERC1155BatchReceived(
        address, address, uint256[] calldata, uint256[] calldata, bytes calldata
    ) override external pure returns (bytes4) {
        return bytes4(keccak256("onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)"));
    }
    function supportsInterface(bytes4 interfaceId) override external view returns (bool){
        return true;
    }

    // Start here
    function createSaleFromContract(
        address _contractAddress,
        address _seller,
        uint256 _tokenId,
        uint128 _price,
        uint256 _amount
    ) external {
        require(_owns(_contractAddress, _seller, _tokenId, _amount), "invalid token ownership");
        _escrow(_contractAddress, _seller, _tokenId, _amount);
        
        Sale memory sale = Sale(
            _seller,
            _price,
            _amount
        );
        
        sales[_contractAddress][_tokenId] = sale;
        emit SaleCreated(_contractAddress, _tokenId, _price, _seller, _amount);
    }
    
    function createSale(
        address _contractAddress,
        uint256 _tokenId,
        uint128 _price,
        uint256 _amount
    ) external {
        require(_owns(_contractAddress, msg.sender, _tokenId, _amount), "invalid token ownership");
        _escrow(_contractAddress, msg.sender, _tokenId, _amount);
        
        Sale memory sale = Sale(
            msg.sender,
            _price,
            _amount
        );
        
        sales[_contractAddress][_tokenId] = sale;
        emit SaleCreated(_contractAddress, _tokenId, _price, msg.sender, _amount);
    }

    function createAndSell(address _contractAddress, bytes memory id, uint64 _amount, uint16 __royaltyFee, uint128 _price) public virtual returns (uint256) {
        MintarCore core = MintarCore(_contractAddress);
        uint256 _tokenId = core.createFromContract(id, _amount, __royaltyFee, msg.sender);

        Sale memory sale = Sale(
            msg.sender,
            _price,
            _amount
        );
        
        sales[_contractAddress][_tokenId] = sale;
        emit SaleCreated(_contractAddress, _tokenId, _price, msg.sender, _amount);
        return _tokenId;
    }

    function isOnSale(address _contractAddress, uint256 _tokenId) external view returns (bool) {
        Sale storage _sale = sales[_contractAddress][_tokenId];
        return (_sale.seller != address(0) && _sale.price != 0 && _sale.amount != 0);
    }
    
    function purchase(address _contractAddress, uint256 _tokenId, uint256 _amount) external payable {
        Sale storage _sale = sales[_contractAddress][_tokenId];
        address payable _seller = payable(_sale.seller);
        require(_sale.amount >= _amount, "insufficient amount");
        require(msg.value >= _sale.price, "insufficient ether");
        
        IERC1155Upgradeable contractAddress = IERC1155Upgradeable(_contractAddress);
        MintarCore core = MintarCore(_contractAddress);
        contractAddress.safeTransferFrom(address(this), msg.sender, _tokenId, _amount, "");

        _sale.amount = _sale.amount - _amount;
        if (_sale.amount <= 0) {
            delete sales[_contractAddress][_tokenId];
        }
        address payable _creator = payable(core.creator(_tokenId));
        uint16 _royaltyFee = core.royaltyFee(_tokenId);
        uint256 _creatorReceipt = msg.value * (_royaltyFee/10000);
        uint256 _serviceReceipt = msg.value * (serviceFee/10000);
        uint256 _sellerReceipt = msg.value - _creatorReceipt - _serviceReceipt;
        _creator.transfer(_creatorReceipt);
        _seller.transfer(_sellerReceipt);
        
        emit SaleSuccessful(_contractAddress, _tokenId, 0, msg.sender, _seller, _amount);
    }
    
    function cancelSale(address _contractAddress, uint256 _tokenId, uint256 _amount) external {
        Sale storage _sale = sales[_contractAddress][_tokenId];
        require(msg.sender == _sale.seller, "requester not authorized");

        IERC1155Upgradeable contractAddress = IERC1155Upgradeable(_contractAddress);
        contractAddress.safeTransferFrom(address(this), _sale.seller, _tokenId, _amount, "");

        delete sales[_contractAddress][_tokenId];
        emit SaleCanceled(_contractAddress, _tokenId);
    }
    
    /* UTILS */
    
    function _owns(address _contractAddress, address _claimant, uint256 _tokenId, uint256 _amount) internal view returns (bool) {
        ERC1155SupplyUpgradeable contractAddress = ERC1155SupplyUpgradeable(_contractAddress);
        uint256 balance = contractAddress.balanceOf(_claimant, _tokenId);
        return (balance >= _amount && balance != 0);
    }
    
    function _escrow(address _contractAddress, address _owner, uint256 _tokenId, uint256 _amount) internal {
        IERC1155Upgradeable contractAddress = IERC1155Upgradeable(_contractAddress);
        contractAddress.safeTransferFrom(_owner, address(this), _tokenId, _amount, "");
    }
}