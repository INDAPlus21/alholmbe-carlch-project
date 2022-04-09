import * as _ from 'lodash';
import { BigNumber, Contract, providers } from 'ethers';
import { UNISWAP_QUERY_ABI } from './abi';
import {
    UNISWAP_LOOKUP_CONTRACT_ADDRESS_BSC as UNISWAP_LOOKUP_CONTRACT_ADDRESS,
    WBNB_ADDRESS as WETH_ADDRESS,
} from './addresses';
import { ETHER } from './utils';
import { readFileSync } from 'fs';

// given a tokenAddress, give me it's markets
type MarketsByToken = { [tokenAddress: string]: Array<UniswapV2EthPair> };

// given a tokenAddress, give me it's balance?
type TokenBalances = { [tokenAddress: string]: BigNumber };

// see what GroupedMarkets look like in groupedMarkets.json
interface GroupedMarkets {
    // hashmap with tokenAddress as key
    marketsByToken: MarketsByToken;
    // array of all markets
    allMarketPairs: Array<UniswapV2EthPair>;
}

export class UniswapV2EthPair {
    // static uniswapInterface = new Contract(WETH_ADDRESS, UNISWAP_PAIR_ABI);
    // every pair has 2 tokens, which both have a balance
    private tokenBalances: TokenBalances;
    private tokens: Array<string>;
    private marketAddress: string;
    private protocol: string;

    constructor(marketAddress: string, tokens: Array<string>, protocol: string) {
        this.tokenBalances = _.zipObject(tokens, [BigNumber.from(0), BigNumber.from(0)]);
        this.tokens = tokens;
        this.marketAddress = marketAddress;
        this.protocol = protocol;
    }

    getBalance(tokenAddress: string): BigNumber {
        const balance = this.tokenBalances[tokenAddress];
        if (balance === undefined) throw new Error('bad token');
        return balance;
    }

    static async getUniswapMarketsByToken(
        provider: providers.JsonRpcProvider,
        factoryAddresses: Array<string>
    ): Promise<GroupedMarkets> {
        // THIS IS FAKING THE GETTING OF PAIRS
        // todo: write a contract that does this on chain
        const rawdata = readFileSync('./randomjson/bsc_allMarketPairs.json');
        const rawAllMarketPairs = JSON.parse(rawdata.toString());
        const allMarketPairs: Array<UniswapV2EthPair> = [];
        for (const pair of rawAllMarketPairs) {
            const p = new UniswapV2EthPair(pair['_marketAddress'], pair['_tokens'], pair['_protocol']);
            allMarketPairs.push(p);
        }

        // we now have all the pairs we care about, and need to update the reserves
        await UniswapV2EthPair.updateReserves(provider, allMarketPairs);

        // only care about pairs with more than 1 ether
        // transform the data into a hashmap with the non-weth token as key
        const marketsByToken = _.chain(allMarketPairs)
            .filter((pair) => pair.getBalance(WETH_ADDRESS).gt(ETHER))
            .groupBy((pair) => (pair.tokens[0] === WETH_ADDRESS ? pair.tokens[1] : pair.tokens[0]))
            .value();

        return {
            marketsByToken,
            allMarketPairs,
        };
    }

    static async updateReserves(
        provider: providers.JsonRpcProvider,
        allMarketPairs: Array<UniswapV2EthPair>
    ): Promise<void> {
        const uniswapQuery = new Contract(UNISWAP_LOOKUP_CONTRACT_ADDRESS, UNISWAP_QUERY_ABI, provider);
        const pairAddresses = allMarketPairs.map((marketPair) => marketPair.marketAddress);
        console.log('Updating markets, count:', pairAddresses.length);

        // get reserves for ALL our pairs, example at reserves.json
        // reserve[i][0] is token0s reserve, reserve[i][1] is token1s reserve
        // reserve[i][2] is not important
        const reserves: Array<Array<BigNumber>> = (await uniswapQuery.functions.getReservesByPairs(pairAddresses))[0];

        // update the balances, one pair at a time
        for (let i = 0; i < allMarketPairs.length; i++) {
            // allMarketPairs[i] is the pair with reserves at reserves[i]
            const marketPair = allMarketPairs[i];
            const reserve = reserves[i];

            // updates this instance of a pair's _tokenBalances
            marketPair.setReservesViaOrderedBalances([reserve[0], reserve[1]]);
        }
    }

    setReservesViaOrderedBalances(balances: Array<BigNumber>): void {
        // balances is an array of the reserves
        // this.tokens is an array with the token addresses for this pair
        this.setReservesViaMatchingArray(this.tokens, balances);
    }

    setReservesViaMatchingArray(tokens: Array<string>, balances: Array<BigNumber>): void {
        // tokenBalances maps address to balance
        const tokenBalances = _.zipObject(tokens, balances);
        if (!_.isEqual(this.tokenBalances, tokenBalances)) {
            // update this.tokenBalances if they are no longer equal
            this.tokenBalances = tokenBalances;
        }
    }
}
