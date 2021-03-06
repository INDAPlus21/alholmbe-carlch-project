import { providers } from 'ethers';
import { FACTORY_ADDRESSES } from './addresses';
import { UniswapV2EthPair } from './UniswapV2EthPair';
import 'dotenv/config';

const ETHEREUM_RPC_URL = 'https://bsc-dataseed.binance.org/';

// a provider to interact with the chain
// about providers: https://docs.ethers.io/v5/api/providers/
const provider = new providers.StaticJsonRpcProvider(ETHEREUM_RPC_URL);

const main = async () => {
    console.log('main');

    // a market is two different tokens that are tradeable, for example ETH/USDC
    // get ALL markets that exists on the factory addresses (which are exchanges)
    const markets = await UniswapV2EthPair.getUniswapMarketsByToken(provider, FACTORY_ADDRESSES);
    console.log('Number of pairs:', markets.allMarketPairs.length);

    provider.on('block', async (blockNumber) => {
        console.log('block', blockNumber);

        // on every new block, update reserves
        await UniswapV2EthPair.updateReserves(provider, markets.allMarketPairs);
    });
};

main();
