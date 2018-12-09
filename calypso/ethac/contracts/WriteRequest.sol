pragma solidity ^0.5.1;
import "./Policy.sol";

contract WriteRequest {
    bytes public data;

    bytes public extraData;

    bytes public LTSID;
    
    bytes public U;

    bytes public Ubar;

    bytes public E;

    bytes public F;

    bytes public X;

    bytes public Cs;

    Policy private policy;

    int64[] public slice;

    constructor(bytes memory d, bytes memory ed, bytes memory l, address p, bytes memory u, bytes memory cs, int64[] memory sl, bytes memory f, bytes memory e, bytes memory ub) public {
        data = d;
        extraData = ed;
        LTSID = l;
        policy = Policy(p);
        U = u;
        Cs = cs;
        slice = sl;
        F = f;
        E = e;
        Ubar = ub;
    }

    function CanRead(address a) public view returns(bool) {
        return policy.CanRead(a);
    }

    function GetSlice() public view returns(int64[] memory) {
        return slice;
    }
}