package utils

import (
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

var factories = map[string]map[string]string{
	"avalanche": {
		"traderjoe": "0x9ad6c38be94206ca50bb0d90783181662f0cfa10",
		"pangolin":  "0xefa94de7a4656d787667c749f7e1223d71e9fd88",
	},
	"bsc": {
		"pancakeswap": "0xca143ce32fe78f1f7019d7d551a6402fc5350c73",
		"sushiswap":   "0x0841BD0B734E4F5853f0dD8d7Ea041c241fb0Da6",
	},
	"polygon": {
		"quickswap": "0x5757371414417b8C6CAad45bAeF941aBc7d3Ab32",
		"sushiswap": "0xc35DADB65012eC5796536bD9864eD8773aBc74C4",
	},
	"fantom": {
		"spookyswap": "0x733A9D1585f2d14c77b49d39BC7d7dd14CdA4aa5",
		"spiritswap": "0xEF45d134b73241eDa7703fa787148D9C9F4950b0",
	},
	"aurora": {
		"trisolaris": "0xc66f594268041db60507f00703b152492fb176e7",
		"wannaswap":  "0x7928d4fea7b2c90c732c10aff59cf403f0c38246",
	},
}

var DeployedUniswapQueryContracts = map[string]string{
	"avalanche": "0xbc37182da7e1f99f5bd75196736bb2ae804cbf6a",
	"bsc":       "0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A",
	"polygon":   "0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A",
	"fantom":    "0xbc37182da7e1f99f5bd75196736bb2ae804cbf6a",
	"aurora":    "0xbc37182da7e1f99f5bd75196736bb2ae804cbf6a",
}

func GetFactories(network string) []string {

	if network == "polygon" {
		return []string{factories[network]["quickswap"], factories[network]["sushiswap"]}
	} else if network == "avalanche" {
		return []string{factories[network]["traderjoe"], factories[network]["pangolin"]}
	} else if network == "bsc" {
		return []string{factories[network]["pancakeswap"], factories[network]["sushiswap"]}
	} else if network == "fantom" {
		return []string{factories[network]["spookyswap"], factories[network]["spiritswap"]}
	} else if network == "aurora" {
		return []string{factories[network]["trisolaris"], factories[network]["wannaswap"]}
	}

	return []string{}
}

func GetTokens(network string) []Token {
	var minLiq, power = big.NewInt(10), big.NewInt(18)
	minLiq.Exp(minLiq, power, nil)
	if network == "polygon" {
		return []Token{{
			Symbol:       "WMATIC",
			Address:      "0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270",
			Protocol:     "polygon",
			MinLiquidity: minLiq, // 1 WMATIC
		},
		// {
		// 	Symbol:       "WETH",
		// 	Address:      "0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619",
		// 	Protocol:     "polygon",
		// 	MinLiquidity: minLiq, // 1 WETH
		// },
		}
	} else if network == "avalanche" {
		return []Token{
			{
				Symbol:       "WAVAX",
				Address:      "0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7",
				Protocol:     "avalanche",
				MinLiquidity: minLiq, // 1 WAVAX
			},
		}
	} else if network == "bsc" {
		return []Token{
			{
				Symbol:       "WBNB",
				Address:      "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
				Protocol:     "bsc",
				MinLiquidity: minLiq, // 1 WBNB
			},
		}
	} else if network == "fantom" {
		return []Token{
			{
				Symbol:       "WFTM",
				Address:      "0x21be370D5312f44cB42ce377BC9b8a0cEF1A4C83",
				Protocol:     "fantom",
				MinLiquidity: minLiq, // 1 WFTM
			},
		}
	} else if network == "aurora" {
		return []Token{
			{
				Symbol:       "WAURORA",
				Address:      "0x8BEc47865aDe3B172A928df8f990Bc7f2A3b9f79",
				Protocol:     "aurora",
				MinLiquidity: minLiq, // 1 AURORA
			},
		}
	}

	return []Token{}
}

// GetClient initializes and returns connection with the chosen network
func GetClient(network string) *ethclient.Client {
	rpc_url := "rpc_" + network
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	rpc := os.Getenv(rpc_url)

	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// helper function for seeing if a token list contains a certain address
func In(address string, tokens []Token) bool {
	for _, token := range tokens {
		if address == token.Address {
			return true
		}
	}
	return false
}

// determines if a markets is liquid enough
func IsLiquidEnough(market *Market, tokenOfInterestAddress common.Address, tokenOfInterest Token) bool {
	if market.Pairs[0].Token0Address == tokenOfInterestAddress {

		if market.Pairs[0].Token0Balance.Cmp(tokenOfInterest.MinLiquidity) == -1 ||
			market.Pairs[1].Token0Balance.Cmp(tokenOfInterest.MinLiquidity) == -1 {
			return false
		}
	} else {
		if market.Pairs[0].Token1Balance.Cmp(tokenOfInterest.MinLiquidity) == -1 ||
			market.Pairs[1].Token1Balance.Cmp(tokenOfInterest.MinLiquidity) == -1 {
			return false
		}
	}

	return true
}
