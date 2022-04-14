import * as _ from 'lodash';
import { BigNumber, Contract, providers } from 'ethers';
import { UNISWAP_QUERY_ABI } from './abi';
import {
    UNISWAP_LOOKUP_CONTRACT_ADDRESS_BSC as UNISWAP_LOOKUP_CONTRACT_ADDRESS,
    WBNB_ADDRESS as WETH_ADDRESS,
} from './addresses';
import { ETHER } from './utils';

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
        // we have all fact addresses, get the token addresses for each one in parallell
        const allPairs = await Promise.all(
            // every DEX (example uniswap) has ONE factory address
            // "for every factoryAddress, call getUniswapMarkets"
            _.map(factoryAddresses, (factoryAddress) => UniswapV2EthPair.getUniswapMarkets(provider, factoryAddress))
        );

        // we now have all pair addresses
        // transform the data into a hashmap with the non-weth token as key
        const marketsByTokenAll = _.chain(allPairs)
            .flatten()
            .groupBy((pair) => (pair.tokens[0] === WETH_ADDRESS ? pair.tokens[1] : pair.tokens[0]))
            .value();

        // here we find crossed markets, e.g we only care about pairs that exists on more than one dex
        // example: a market that exists on both uniswap and sushiswap is something we save
        const allMarketPairs = _.chain(_.pickBy(marketsByTokenAll, (a) => a.length > 1))
            .values()
            .flatten()
            .value();

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

    static async getUniswapMarkets(provider: providers.JsonRpcProvider, factoryAddress: string) {
        const uniswapQuery = new Contract(UNISWAP_LOOKUP_CONTRACT_ADDRESS, UNISWAP_QUERY_ABI, provider);

        const marketPairs = new Array<UniswapV2EthPair>(); // ALL pairs on the chosen dexes
        // const marketPairs = 10000;

        const BATCHES = 1000;
        const BATCH_SIZE = 100;

        // retrieve all pair from uniswap in batches of size 100
        // for (let i = 0; i < 10000; i += BATCH_SIZE) {
        for (let i = 0; i < BATCHES * BATCH_SIZE; i += BATCH_SIZE) {
            // get all pairs in range i to i + BATCH_SIZE

            const pairs: Array<Array<string>> = (
                await uniswapQuery.functions.getPairsByRange(factoryAddress, i, i + BATCH_SIZE)
            )[0];

            for (let i = 0; i < pairs.length; i++) {
                // for every pair, token0Address at 0, token1Address at 1, pairAddress at 2
                const pair = pairs[i];
                const marketAddress = pair[2];
                let tokenAddress: string;

                // determine which pair is not weth
                if (pair[0] === WETH_ADDRESS) {
                    tokenAddress = pair[1];
                } else if (pair[1] === WETH_ADDRESS) {
                    tokenAddress = pair[0];
                } else {
                    continue;
                }

                // initialize a pair, push to array of all pairs
                const uniswapV2EthPair = new UniswapV2EthPair(marketAddress, [pair[0], pair[1]], '');
                marketPairs.push(uniswapV2EthPair);
            }
            if (pairs.length < BATCH_SIZE) {
                break;
            }
        }

        return marketPairs;
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
