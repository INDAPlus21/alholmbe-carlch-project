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
}

var DeployedUniswapQueryContracts = map[string]string{
	"avalanche": "0xbc37182da7e1f99f5bd75196736bb2ae804cbf6a",
	"bsc":       "0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A",
	"polygon":   "0xBc37182dA7E1f99f5Bd75196736BB2ae804Cbf6A",
}

func GetFactories(network string) []string {

	if network == "polygon" {
		return []string{factories[network]["quickswap"], factories[network]["sushiswap"]}
	} else if network == "avalanche" {
		return []string{factories[network]["traderjoe"], factories[network]["pangolin"]}
	} else if network == "bsc" {
		return []string{factories[network]["pancakeswap"], factories[network]["sushiswap"]}
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
		}}
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
