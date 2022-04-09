import { Contract, providers, Wallet } from 'ethers';
import { FACTORY_ADDRESSES } from './addresses';
import 'dotenv/config';

const ETHEREUM_RPC_URL = process.env.BSC_RPC_URL; //process.env.ETHEREUM_RPC_URL || 'http://127.0.0.1:8545';
const PRIVATE_KEY = process.env.PRIVATE_KEY || '';

if (PRIVATE_KEY == '') {
    console.error('no private key provided');
    process.exit();
}

// a provider to interact with the chain
const provider = new providers.StaticJsonRpcProvider(ETHEREUM_RPC_URL);

const main = async () => {
    const markets = await UniV2EthPair.getUniswapMarkets(provider, FACTORY_ADDRESSES);
};
