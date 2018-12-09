pragma solidity ^0.5.1;

import "./Owners.sol";
import "./WriteRequestHolder.sol";
import "./WriteRequest.sol";
import "./ReadRequest.sol";
import "./ReadRequestHolder.sol";


contract Calypso {
    Owners private owners;
    WriteRequestHolder private wrHolder;
    mapping(bytes32 => bool) hasReq;
    ReadRequestHolder private rrHolder;
    mapping(address => bool) public  ownersMap;


    //The addresses of the owners of the contract
    //This is a mapping that maps the hash of a request to the
    //associated WriteRequest of that hash
    
    constructor(address holder, address rrholder, address o) public {
        owners = Owners(o);
        wrHolder = WriteRequestHolder(holder);
        rrHolder = ReadRequestHolder(rrholder);
    }

    function canWrite(address a) public view returns (bool) {
        return owners.canWrite(a);
    }

    function addWriteRequest(address a) public {
        require(canWrite(msg.sender), "You are not a registered owner");
        wrHolder.addWriteRequest(a);
    }

    function AddReadRequest(address a) public {
        ReadRequest rr = ReadRequest(a);
        address write = rr.writeRequest();
        require(wrHolder.canRead(write), "There is no corresponding write request");
        rrHolder.addReadRequest(a);
    }

    function AddOwner(address a) public {
        //require(canWrite(msg.sender), "You are not one of the original owners");
        ownersMap[a] = true;
    }

    function checkIfRequestIsValid(address a) public returns (bool) {
        return wrHolder.canRead(a);
    }

    function isValidReadRequest(address a)  public view returns (bool) {
        /*require(rrHolder.hasReadRequest(a), "It must at first be a valid read request.");
        */
        ReadRequest rr = ReadRequest(a);
        address wrA = rr.writeRequest();
        WriteRequest wr = WriteRequest(wrA);
        require(wr.CanRead(rr.owner()), "Is this a member of the policy of the pointed WriteRequest");
        return true;
    }

}



