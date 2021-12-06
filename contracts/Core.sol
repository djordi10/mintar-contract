// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.2;

import "@openzeppelin/contracts-upgradeable/token/ERC1155/extensions/ERC1155SupplyUpgradeable.sol";
import "./Sale.sol";

contract MintarCore is ERC1155SupplyUpgradeable {
    // Mapping from token ID to creator address
    mapping (uint256 => address) private _creator;
    
    // Mapping from token ID to previous owners addresses
    mapping (uint256 => uint16) private _royaltyFee;

    function initialize() public initializer {
        _setURI("url/{id}");
    }

    // id should be a uuid.v4
    // - royaltyFee is a percentage denoted by uint with two decimals, 100% would be 10000, 50% would be 5000, etc.
    // This function also emits TransferSingle event
    function create(bytes memory id, uint64 amount, uint16 __royaltyFee) external returns (uint256) {
        uint256 _tokenId = uint256(keccak256(abi.encode(id)));
        _mint(msg.sender, _tokenId, amount, "");
        _creator[_tokenId] = msg.sender;
        _royaltyFee[_tokenId] = __royaltyFee;

        return _tokenId;
    }

    function createFromContract(bytes memory id, uint64 amount, uint16 __royaltyFee, address __creator) external returns (uint256) {
        uint256 _tokenId = uint256(keccak256(abi.encode(id)));
        _mint(msg.sender, _tokenId, amount, "");
        _creator[_tokenId] = __creator;
        _royaltyFee[_tokenId] = __royaltyFee;

        return _tokenId;
    }
    
    function creator(uint256 _tokenId) external view returns (address) {
        return _creator[_tokenId];
    }

    function royaltyFee(uint256 _tokenId) external view returns (uint16) {
        return _royaltyFee[_tokenId];
    }
}