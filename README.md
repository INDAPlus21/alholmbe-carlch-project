# Ethereum Arbitrage Bots
Written by Carl Chemnitz and Alexander Holmberg

## Summary
Crypto arbitrage trading is a form of *arbitrage* trading that utilizes the vast amount of *decentralized exchanges* to make a profit with low to no risk. At the moment there are two parallel projects, `Atomic arbitrage` written in TypeScript and Solidity and `Cross-chain arbitrage` which is the main project and is written in GoLang and Solidity.

## Definitions
Arbitrage
: an investment strategy in which an investor simultaneously buys and sells an asset in different markets to take advantage of a price difference and generate a profit.

Decentralized exchanges
: A type of cryptocurrency exhange which allows for direct peer-to-peer cryptocurrent transactions to take palce securly without the need for an intermediary.

## How it works
- **Connect**: Connect to multiple Ethereum related networks.
- **Scrape**: Fetch transaction data from the most recent block of transactions.
- **Read data**: Translate transaction data with `Smart Contracts` and load pairs of tokens and their respective values.
- **Evaluate**: Compare a pair's value in all relevant chains and search for an oppertunity.
- **Execute transaction**: If an oppertunity is found, send a transaction to the Ethereum network.

## Roadmap

### Must-have features

### QoL features
