export const UNISWAP_QUERY_ABI = [
    {
        inputs: [],
        name: 'BadIndex',
        type: 'error',
    },
    {
        inputs: [
            {
                internalType: 'contract UniswapV2Factory',
                name: '_uniswapFactory',
                type: 'address',
            },
            {
                internalType: 'uint256',
                name: '_startIndex',
                type: 'uint256',
            },
            {
                internalType: 'uint256',
                name: '_stopIndex',
                type: 'uint256',
            },
        ],
        name: 'getPairsByRange',
        outputs: [
            {
                internalType: 'address[3][]',
                name: '',
                type: 'address[3][]',
            },
        ],
        stateMutability: 'view',
        type: 'function',
    },
    {
        inputs: [
            {
                internalType: 'contract IUniswapV2Pair[]',
                name: '_pairs',
                type: 'address[]',
            },
        ],
        name: 'getReservesByPairs',
        outputs: [
            {
                internalType: 'uint256[3][]',
                name: '',
                type: 'uint256[3][]',
            },
        ],
        stateMutability: 'view',
        type: 'function',
    },
];
