pragma solidity ^0.5.1;


import "./WriteRequest.sol";

contract WriteRequestHolder {

    mapping(address => WriteRequest) public wrmap;
    mapping(address => bool) public hasRequest;
    
    constructor() public {

    }

    function addWriteRequest(address a) public {
        WriteRequest wr = WriteRequest(a);
        wrmap[a] = wr;
        hasRequest[a] = true;
    }

    function getRequestByID(address a) public view returns (WriteRequest) {
        return wrmap[a];
    }

    function canRead(address a) public returns (bool) {
        return hasRequest[a];
    }
}