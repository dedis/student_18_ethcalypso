pragma solidity ^0.5.1;

contract Policy {
    address private maker;
    address[] private owners;
    mapping(address => bool) private ownersMap;
    constructor(address[] memory a) public {
        maker = msg.sender;
        owners = a;
        for(uint64 i = 0; i < a.length;i++) {
            ownersMap[a[i]] = true;
        }
    }

    function CanRead(address a) public view returns (bool) {
        return ownersMap[a];
    }

    function AddOwner(address a) public {
        require(msg.sender == maker, "You are not part of the ownership of this poicy");
        ownersMap[a] = true;
    }

    function RemoveFromPolicy(address a) public {
        require(msg.sender == maker, "You are not part of the ownership of this poicy");
        ownersMap[a] = false;
    }
}