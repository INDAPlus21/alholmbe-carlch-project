# Cross-Chain Arbitrage Finder

To compile and execute the program:

- `go run main.go tui` to run with a user interface.
- `go run main go print` to only display new oppertunites.

**Note**: A `.env` file needs to exist to be able to run the porgram. An example `.env` can be found in `.env.example`.

## Specifications

The core source code is found in `UniswapV2Market.go`. It retrieves, processes and evaluate all markets found between all chosen chains. Each time the program interacts with a chain for non-view operations it pays a fixed price to the chain called _gas_ which create a incentive to perform as much off-chain computation as possible. It it also more optimized to do the computations off-chain as in most cases it has better performance than on-chain computation.

The retrieval of markets is made on the chain however, as it is faster to get all markets in one call than to make thousand of calls. This interaction is also free to make (no gas is required to view the chain).

### Chains

We have deployed `UniswapQuery.sol` (binded to Go by `UniswapQuery.go`) on 6 different chains right now, and the contracts can be found at:

- Ethereum: https://etherscan.io/address/0x4180d411c7fdaf77c2a8056ca712550ecca07fcd#code
- Polygon: https://polygonscan.com/address/0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A#code
- Binance: https://bscscan.com/address/0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A#code
- Avalanche: https://snowtrace.io/address/0xbc37182da7e1f99f5bd75196736bb2ae804cbf6a#code
- Aurora: https://aurorascan.dev/address/0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A
- Fantom: https://ftmscan.com/address/0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A

### Go-Contract Binding

We can gather information from an _Automated Market Maker_ (AMM) on any specified blockchain by implementing a binding between Go and our `UniswapV2Factory` contract. The binding acts as an API by exposing methods to us from the chain.

NOTE: A "_factory contract_" is something almost every AMM has, and it's not written by us, just retrieved from the chain. For example, here is the original UniswapFactory contract on ethereum: https://etherscan.io/address/0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f#code . You can also view it here at `UniswapV2Factory.sol`.

<br>

The functions used to communicate with a chain is found in `network.go`. For each chain, the main routine creates a separate goroutine that runs until termination of the program. The code is specifically written to make the process of adding a new blockchain as effortless as possible.

<br>
