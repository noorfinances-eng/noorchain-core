pragma solidity ^0.8.20;

contract MultiEmit {
    event E1(uint256 indexed a);
    event E2(uint256 indexed b);

    function emitTwo(uint256 a, uint256 b) external {
        emit E1(a);
        emit E2(b);
    }
}
