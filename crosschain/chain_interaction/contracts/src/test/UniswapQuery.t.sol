// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.13;

import "ds-test/test.sol";
import "../UniswapQuery.sol";

contract UniswapQueryTest is DSTest {
    function setUp() public {}

    function testExampleUni() public {
        assertTrue(true);
    }

    function testDeploy() public {
        new UniswapQuery();
    }
}
