// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract TreasureBox {
   
    address public owner;

    mapping(address => uint32) public registered;
    mapping(address => bool) public unlocked;

    constructor() {
        owner = msg.sender;
    }
    
    event Log(uint32 id, uint64 code, bytes text, bytes32 hash);
    
    function register(uint32 id) public {
        require(id != 0, "Invalid id.");
        
        if (registered[msg.sender] != 0) {
            require(registered[msg.sender] == id,
                "You can only register one id per account.");
        }
        else {
            registered[msg.sender] = id;
        }
        
        require(!unlocked[msg.sender],
            "You have already unlocked the box.");
    }
    
    function unlock(uint64 code) public returns (bool) {
        uint32 id = registered[msg.sender];
        require(id != 0, "You haven't registered yet.");

        bytes memory text = abi.encodePacked(id, code);
        bytes32 hash = sha256(text);
        emit Log(id, code, text, hash);
        
        if (bytes4(text)&0xfffffff0 == bytes4(hash)&0xfffffff0) {
            unlocked[msg.sender] = true;
            return true;
        }
        else {
            return false;
        }
    }
}