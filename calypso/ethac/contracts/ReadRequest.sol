pragma solidity ^0.5.1;

contract ReadRequest {
    address public writeRequest;
    bytes public xc;
    address public owner;
    //Each Read request also keeps track of the 
    bytes public WriteHash;

    constructor(address a, bytes memory x, bytes memory wrHash) public {
        owner = msg.sender;
        writeRequest = a;
        xc = x;
        WriteHash = wrHash;
    }

    function CompareReadWrite(address wr) public view returns (bool) {
        return wr == writeRequest;
    }
}