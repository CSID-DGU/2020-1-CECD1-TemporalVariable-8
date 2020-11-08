// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
// import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.2.2-solc-0.7/contracts/access/Ownable.sol";
// import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v3.2.2-solc-0.7/contracts/access/AccessControl.sol";

contract MakItGo is AccessControl, Ownable {
    struct Business {
        address businessOwner;
        string name;
        string url;
        uint64 expired;
    }
    
    event BusinessCreated(string indexed url, string indexed name);
    event BusinessTransfer(address indexed fromOwner, address indexed toOwner);
    event BusinessUpdate(string indexed url, uint256 expired);


    constructor(){
        _setupRole(DEFAULT_ADMIN_ROLE, owner());
    }

    mapping(string => Business) public datamap;
    string[] public urls;

    function createBusiness(string memory url, string memory name) public {
        require(datamap[url].expired == 0, "url is assigned");
        
        datamap[url] = Business({
            businessOwner : msg.sender,
            name : name,
            url : url,
            expired : uint64(block.timestamp)
        });
        urls.push(url);
    }

    function transferBusiness(string memory url, address toOwner) public {
        require(datamap[url].expired != 0, "url is unassigned");
        require(datamap[url].businessOwner == msg.sender, "transfer only allowed for businessOwner");
        
        address fromOwner = datamap[url].businessOwner;
        emit BusinessTransfer(fromOwner, toOwner);
        datamap[url].businessOwner = toOwner;
    }
    
    function weeklyUpdate(string memory url) public {
        require(datamap[url].expired != 0, "url is unassigned");
        require(datamap[url].businessOwner == msg.sender, "update only allowed for businessOwner");
        
        datamap[url].expired = uint64(block.timestamp + 7 days);
    }
    

    function listURLCount() public view returns(uint64){
        return uint64(urls.length);
    }
}
