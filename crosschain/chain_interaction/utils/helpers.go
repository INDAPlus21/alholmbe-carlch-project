package utils

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

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
