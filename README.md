# Arbitrage Bots on ethereum virtual machine compatible blockchains

Written by Carl Chemnitz and Alexander Holmberg

## Summary

Crypto arbitrage trading is a form of _arbitrage_ trading that utilizes the vast amount of _decentralized exchanges_ to make a profit with low to no risk. At the moment there are two parallel projects, `Atomic arbitrage` written in TypeScript and Solidity and `Cross-chain arbitrage` which is the main project and is written in GoLang and Solidity.

## Definitions

<dl>
    <dt>Arbitrage</dt>
    <dd>An investment strategy in which an investor simultaneously buys and sells an asset in different markets to take advantage of a price difference and generate a profit.</dd>
    <dt>Atomic Arbitrage</dt>
    <dd>A form of arbitrage where if you find an opportunity and manage to get your transaction through, you are guaranteed to make a profit without risk. This is because the you are both selling AND buying in the same transaction, and because a transaction is atomic, you are guaranteed to have both your sell and buy order go through, or neither. This atomicity is not possible to get across chains, so this is only possible if you sell and buy on the same chain.</dd>
    <dt>Cross-chain Arbitrage</dt>
    <dd>This is more like the kind of arbitrage that exists in the "traditional world", where you sell on one place and buy on another. Unlike atomic arbitrages, this comes with risks since it's possible that your sell order goes through but not your buy order, and vice versa.</dd>
    <dt>Decentralized exchanges (DEXes)</dt>
    <dd>A type of cryptocurrency exchange that allows for peer-to-peer swaps between tokens without an intermediary. These DEXes are almost always "automated markets makers", AMMs, instead of order books which most regular exchanges are. To enable trading between two tokens on an AMM, they are both placed in a "pair contract". Every pair has two tokens, and every token has a balance. <br>
    price of token x (denominated in terms of y) = balance token y / balance token x <br>
    price of token y (denominated in terms of x) = balance token x / balance token y <br>
     Two examples of AMMs are app.uniswap.org and app.sushi.com.</dd>

</dl>

## How it works

- **Connect**: Connect to multiple Ethereum related networks. (EVM-compatible networks)
- **Read data**: Get the current state of the chain (especially the state of the chosen DEXes), along with transaction data from the most recent block. The state of a DEX is basically the balances that a "pair" has. This is done on-chain with "smart contracts". You can get pairs in batches this way, which isn't possible to do off-chain.
- **Transform data**: Transform the raw blockchain data into more appropriate data structures so we later can search through everything easier and faster. This is done off-chain because it's more efficient.
- **Evaluate data**: With data of token prices from multiple chains, what's left is to evaluate everything and see if an opportunity is found.
- **Execute transaction(s)**: If an opportunity is found, send a transaction(s) to the relevant network(s).

## Roadmap

### Must-have features

### QoL features
