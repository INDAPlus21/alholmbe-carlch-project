### cross-chain arbitrage searcher

For the solidity and smart contracts, go to contracts/src/UniswapQuery.sol/ <br>
We have deployed UniswapQuery.sol on 5 different chains right now, and the contracts can be found at:

- on polygon: https://polygonscan.com/address/0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A#code
- on bsc: https://bscscan.com/address/0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A#code
- on avalanche: https://snowtrace.io/address/0xbc37182da7e1f99f5bd75196736bb2ae804cbf6a#code
- on aurora: https://aurorascan.dev/address/0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A
- on fantom: https://ftmscan.com/address/0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A

<br>

generatedContracts/UniswapV2Factory.go is a go binding of the UniswapV2Factory contract. It allows us the get information from whatever AMM (automated market maker) on whatever blockchain we want. It's possible because the contract essentially acts as an API, exposing methods for us to call from go via RPC.

NOTE: A "factory contract" is something almost every AMM has, and it's not written by us, just retrieved from the chain. For example, here is the original UniswapFactory contract on ethereum: https://etherscan.io/address/0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f#code . You can also view it here at UniswapV2Factory.sol.

<br>

UniswapQuery/UniswapQuery.go is a the go binding for our UniswapQuery.sol contract.

<br>

/interface contains the code for generating a terminal user interface that show opportunities

<br>

/networks/network.go contains the function that gets called once in each goroutine for each chain, and it goes until someone terminates the program. It's written in a way so it's easy to add new chains.

<br>
 /utils/helper.go contains a bunch of helper functions that is used in all the other files.<br>
 /utils/UniswapV2Market.go can be seen as the main file of the program, where most of the stuff are being done. Here is where all the markets from every blockchain gets retrieved, processed and evaluated. We try do as much of the computation we can off-chain (so that we don't have to pay gas and because it's often faster), but something that is better to on chain is the retrieval of markets, first of all it's only a "view operation" so it's free but also it allows us to get tens of thousands of markets in one call.

<br>

`go run main.go tui` to run the program with the terminal user interface <br>
`go run main.go print` to run the program with print statements on every new opportunity<br>

a .env file need to exist in the root of the folder, a .env.example file is provided with some RPC url that's enough for now. (We don't send any transactions right now, so it's not necessary to have a private key)
