pragma solidity ^0.5.1;
import "./ReadRequest.sol";

contract ReadRequestHolder {
    mapping(address => bool) public hasReq;
    mapping(address => ReadRequest) public requests;

    constructor() public {
    }

    function addReadRequest(address a) public {
        ReadRequest rr = ReadRequest(a);
        hasReq[a] = true;
        requests[a] = rr;
    }

    function hasReadRequest(address id) public view returns (bool) {
        return hasReq[id];
    }
}