import { BigNumber } from 'ethers';

export interface TokenBalances {
    [tokenAddress: string]: BigNumber;
}

export interface MultipleCallData {
    targets: Array<string>;
    data: Array<string>;
}

export interface CallDetails {
    target: string;
    data: string;
    value?: BigNumber;
}

export abstract class EthMarket {}
