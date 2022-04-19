//SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.13;

interface IUniswapV2Pair {
    function token0() external view returns (address);
    function token1() external view returns (address);
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
}

abstract contract UniswapV2Factory  {
    mapping(address => mapping(address => address)) public getPair;
    address[] public allPairs;
    function allPairsLength() external view virtual returns (uint);
}

contract UniswapQuery {


    /// @notice Thrown when the stop index is greater than the start index
    error BadIndex();

    /// @param _pairs the array of pairs
	/// @return reserves a 2D array where the inner arrays each contain the token0 reserves, 
    /// token1 reserves and the timestamp of the last interaction
    function getReservesByPairs(IUniswapV2Pair[] calldata _pairs) external view returns (uint256[3][] memory) {
        uint256[3][] memory reserves = new uint256[3][](_pairs.length);
        for (uint i = 0; i < _pairs.length; i++) {
            (reserves[i][0], reserves[i][1], reserves[i][2]) = _pairs[i].getReserves();
        }
        return reserves;
    }

	/// @param _uniswapFactory The uniswap factory we are querying
	/// @param _startIndex the index we start querying from 
    /// @param _stopIndex the index we stop querying at
	/// @return pairs a 2D array where the inner arrays each contain the token0 address, token1 address and pair address
	/// @notice we use _start and _stop for efficiency
    function getPairsByRange(UniswapV2Factory _uniswapFactory, uint256 _startIndex, uint256 _stopIndex) external view returns (address[3][] memory)  {
        // how many pairs exists on this DEX?
        uint256 allPairsLength = _uniswapFactory.allPairsLength();
        if (_stopIndex > allPairsLength) {
            _stopIndex = allPairsLength;
        }

        if (_stopIndex < _startIndex) revert BadIndex();
        uint256 quantity = _stopIndex - _startIndex;

        // outer array is of size quantity, inner array is of size 3
        address[3][] memory pairs = new address[3][](quantity);
        for (uint i = 0; i < quantity; i++) {
            // get every pair in this range and save them
            IUniswapV2Pair _uniswapPair = IUniswapV2Pair(_uniswapFactory.allPairs(_startIndex + i));
            pairs[i][0] = _uniswapPair.token0();
            pairs[i][1] = _uniswapPair.token1();
            pairs[i][2] = address(_uniswapPair);
        }
        return pairs;
    }
}