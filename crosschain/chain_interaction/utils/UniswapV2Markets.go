package utils

import (
	"fmt"
	"log"
	"math/big"

	UniswapQuery "chain_interaction/UniswapQuery"
	UniswapV2Factory "chain_interaction/generatedContracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// useful for testing purposes, so we don't have to load all markets (takes time)
const BATCH_COUNT_LIMIT int = 100
const UNISWAP_BATCH_SIZE int = 100

const MY_ADDRESS string = "0x30429A2FfAE3bE74032B6ADD7ac4A971AbAd4d02"

// maps asset (e.g WETH) to networks (e.g ethereum and bsc)
type UniswapV2Markets struct {
	Asset map[string]map[string]*Network
}

// maps network (e.g ethereum) to pairs
type Network struct {
	Asset                string
	Protocol             string
	Pairs                []*UniswapV2EthPair
	AllMarkets           []*UniswapV2EthPair
	CrossMarkets         []*UniswapV2EthPair
	CrossMarketsFiltered []*UniswapV2EthPair
	// CrossMarketsByToken map[string][]*UniswapV2EthPair
	CrossMarketsByTokenFiltered map[string]*Market
	CrossMarketsByToken         map[string]*Market
}

type Market struct {
	Pairs               []*UniswapV2EthPair
	CurrentArbitrageOpp *big.Float
}

type UniswapV2EthPair struct {
	PairAddress   common.Address
	Token0Address common.Address
	Token1Address common.Address
	Token0Balance *big.Int
	Token1Balance *big.Int
}

type Token struct {
	Symbol       string
	Address      string
	Protocol     string
	MinLiquidity *big.Int
}

// Setup is an initialization function that populates the UniswapV2Markets struct
func (uniswapMarkets *UniswapV2Markets) Setup() {

	ethereum_WETH := Network{Asset: "WETH", Protocol: "ethereum"}
	ethereum_WBNB := Network{Asset: "WBNB", Protocol: "ethereum"}
	ethereum_WMATIC := Network{Asset: "WMATIC", Protocol: "ethereum"}

	bsc_WETH := Network{Asset: "WETH", Protocol: "bsc"}
	bsc_WBNB := Network{Asset: "WBNB", Protocol: "bsc"}
	bsc_WMATIC := Network{Asset: "WMATIC", Protocol: "bsc"}

	polygon_WETH := Network{Asset: "WETH", Protocol: "polygon"}
	polygon_WBNB := Network{Asset: "WBNB", Protocol: "polygon"}
	polygon_WMATIC := Network{Asset: "WMATIC", Protocol: "polygon"}

	uniswapMarkets.Asset = make(map[string]map[string]*Network)
	uniswapMarkets.Asset["WETH"] = make(map[string]*Network)
	uniswapMarkets.Asset["WETH"]["ethereum"] = &ethereum_WETH
	uniswapMarkets.Asset["WETH"]["bsc"] = &bsc_WETH
	uniswapMarkets.Asset["WETH"]["polygon"] = &polygon_WETH

	uniswapMarkets.Asset["WBNB"] = make(map[string]*Network)
	uniswapMarkets.Asset["WBNB"]["ethereum"] = &ethereum_WBNB
	uniswapMarkets.Asset["WBNB"]["bsc"] = &bsc_WBNB
	uniswapMarkets.Asset["WBNB"]["polygon"] = &polygon_WBNB

	uniswapMarkets.Asset["WMATIC"] = make(map[string]*Network)
	uniswapMarkets.Asset["WMATIC"]["ethereum"] = &ethereum_WMATIC
	uniswapMarkets.Asset["WMATIC"]["bsc"] = &bsc_WMATIC
	uniswapMarkets.Asset["WMATIC"]["polygon"] = &polygon_WMATIC
}

func (uniswapMarkets *UniswapV2Markets) uniswapV2MarketByFactory(client *ethclient.Client,
	address string, queryContractAddress string, tokensOfInterest []Token) []*UniswapV2EthPair {
	uniswapQueryAddress := common.HexToAddress(queryContractAddress)
	factoryAddress := common.HexToAddress(address)

	// list with all marketPairs
	marketPairs := []*UniswapV2EthPair{}
	uniswapV2Factory, err := UniswapV2Factory.NewGeneratedContracts(factoryAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	uniswapQuery, err := UniswapQuery.NewUniswapQuery(uniswapQueryAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	bigNum, err := uniswapV2Factory.AllPairsLength(nil)
	if err != nil {
		log.Fatal(err)
	}
	numberOfPairs := int(bigNum.Int64())
	fmt.Printf("number of pairs: %d\n", numberOfPairs)

	var x int
	if numberOfPairs > UNISWAP_BATCH_SIZE*BATCH_COUNT_LIMIT {
		x = UNISWAP_BATCH_SIZE * BATCH_COUNT_LIMIT
	} else {
		x = numberOfPairs
	}

	for i := 0; i < x; i += UNISWAP_BATCH_SIZE {
		// get pairs from the network
		pairs, err := uniswapQuery.GetPairsByRange(nil, factoryAddress, big.NewInt(int64(i)), big.NewInt(int64(i+UNISWAP_BATCH_SIZE)))
		if err != nil {
			// this happens because something is weird in the sushiswap factory contract
			fmt.Printf("revert at i = %d, factoryAddress = %s\n", i, factoryAddress)
			log.Fatal(err)
		}

		for j := 0; j < len(pairs); j++ {
			pair := pairs[j]
			pairAddress := pair[2]

			if !In(pair[0].String(), tokensOfInterest) && !In(pair[1].String(), tokensOfInterest) {
				// we don't care if none of the tokens in the pair is weth or wbnb
				continue
			}

			uniswapV2EthPair := &UniswapV2EthPair{pairAddress, pair[0], pair[1], big.NewInt(0), big.NewInt(0)}

			marketPairs = append(marketPairs, uniswapV2EthPair)
		}
	}

	return marketPairs

}

// UpdateMarkets retrieves and structure market data from the chosen chain/network
// client is connected to chain x and queryContractAddress is an instance of UniswapQuery.sol deployed on chain x
func (uniswapMarkets *UniswapV2Markets) UpdateMarkets(
	client *ethclient.Client,
	factoryAddresses []string,
	queryContractAddress string,
	tokensOfInterest []Token) {
	// WETH_ADDRESS := "0x2170Ed0880ac9A755fd29B2688956BD959F933F8"
	NATIVE_CURRENCY_ADDRESS := tokensOfInterest[0].Address
	NATIVE_CURRENCY := common.HexToAddress(NATIVE_CURRENCY_ADDRESS)

	// for every address, get markets
	// markets is a list with all pairs on this network
	allMarkets := [][]*UniswapV2EthPair{}
	for i := 0; i < len(factoryAddresses); i++ {
		marketPairs := uniswapMarkets.uniswapV2MarketByFactory(client, factoryAddresses[i],
			queryContractAddress, tokensOfInterest)

		allMarkets = append(allMarkets, marketPairs)
	}

	allMarketsByToken := map[string][]*UniswapV2EthPair{}
	allMarketsFlat := []*UniswapV2EthPair{}

	// groups all pairs into a dictionary with the non weth token as the key
	// O(n^2) doesn't matter, this function will only run on startup
	for i := 0; i < len(allMarkets); i++ {
		for j := 0; j < len(allMarkets[i]); j++ {
			allMarketsFlat = append(allMarketsFlat, allMarkets[i][j])
			if allMarkets[i][j].Token0Address == NATIVE_CURRENCY {
				if _, ok := allMarketsByToken[allMarkets[i][j].Token1Address.String()]; ok {
					allMarketsByToken[allMarkets[i][j].Token1Address.String()] =
						append(allMarketsByToken[allMarkets[i][j].Token1Address.String()], allMarkets[i][j])
				} else {
					allMarketsByToken[allMarkets[i][j].Token1Address.String()] = []*UniswapV2EthPair{allMarkets[i][j]}
				}
			} else {
				if _, ok := allMarketsByToken[allMarkets[i][j].Token0Address.String()]; ok {
					allMarketsByToken[allMarkets[i][j].Token0Address.String()] =
						append(allMarketsByToken[allMarkets[i][j].Token0Address.String()], allMarkets[i][j])
				} else {
					allMarketsByToken[allMarkets[i][j].Token0Address.String()] = []*UniswapV2EthPair{allMarkets[i][j]}
				}
			}
		}
	}

	// uniswapMarkets.UpdateReserves(client, queryContractAddress, tokensOfInterest)

	// only keep pairs with decent reserves

	// a cross markets exists if the same market exists on 2+ places on 1 network
	// crossMarketsByToken := make(map[string][]*UniswapV2EthPair)
	crossMarketsByToken := make(map[string]*Market)
	crossMarketsFlat := []*UniswapV2EthPair{}
	for tokenAddress, markets := range allMarketsByToken {
		if len(markets) > 1 {
			crossMarketsByToken[tokenAddress] = &Market{markets, big.NewFloat(0)}
			crossMarketsFlat = append(crossMarketsFlat, markets...)
		}
	}

	for _, token := range tokensOfInterest {
		uniswapMarkets.Asset[token.Symbol][token.Protocol].AllMarkets = allMarketsFlat
		uniswapMarkets.Asset[token.Symbol][token.Protocol].CrossMarkets = crossMarketsFlat
		uniswapMarkets.Asset[token.Symbol][token.Protocol].CrossMarketsByTokenFiltered = make(map[string]*Market)
		uniswapMarkets.Asset[token.Symbol][token.Protocol].CrossMarketsByToken = crossMarketsByToken
	}

}

func (uniswapMarkets *UniswapV2Markets) UpdateReserves(
	client *ethclient.Client,
	queryContractAddress string,
	tokensOfInterest []Token) {
	for _, token := range tokensOfInterest {

		uniswapQueryAddress := common.HexToAddress(queryContractAddress)
		uniswapQuery, err := UniswapQuery.NewUniswapQuery(uniswapQueryAddress, client)
		if err != nil {
			log.Fatal(err)
		}
		pairAddresses := []common.Address{}
		for _, market := range uniswapMarkets.Asset[token.Symbol][token.Protocol].CrossMarkets {
			pairAddresses = append(pairAddresses, (*market).PairAddress)
		}

		reserves, err := uniswapQuery.GetReservesByPairs(nil, pairAddresses)
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(uniswapMarkets.Asset[token.Symbol][token.Protocol].CrossMarkets); i++ {
			// reserve[0] is token0s reserve, reserve[1] is token1s reserve, reserve[2] is last interaction
			(uniswapMarkets.Asset[token.Symbol][token.Protocol].CrossMarkets)[i].Token0Balance = reserves[i][0]
			(uniswapMarkets.Asset[token.Symbol][token.Protocol].CrossMarkets)[i].Token1Balance = reserves[i][1]

		}
	}

}

func (uniswapMarkets *UniswapV2Markets) EvaluateCrossMarkets(tokensOfInterest []Token) {
	WBNB_ADDRESS := "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"
	WBNB := common.HexToAddress(WBNB_ADDRESS)
	_ = WBNB

	for _, token := range tokensOfInterest {
		var tokenOfInterestAddress common.Address = common.HexToAddress(token.Address)

		for tokenAddress, market := range uniswapMarkets.Asset[token.Symbol][token.Protocol].CrossMarketsByToken {
			if !IsLiquidEnough(market, tokenOfInterestAddress, token) {
				continue
			}

			_ = tokenAddress

			otherTokenPrices := []*big.Float{}
			for _, market := range market.Pairs {
				if market.Token0Address == tokenOfInterestAddress {
					otherTokenPrice := new(big.Float).Quo(new(big.Float).SetInt(market.Token1Balance), new(big.Float).SetInt(market.Token0Balance))
					otherTokenPrices = append(otherTokenPrices, otherTokenPrice)
				} else {
					otherTokenPrice := new(big.Float).Quo(new(big.Float).SetInt(market.Token0Balance), new(big.Float).SetInt(market.Token1Balance))
					otherTokenPrices = append(otherTokenPrices, otherTokenPrice)
				}
			}

			priceDiff0 := new(big.Float).Quo(otherTokenPrices[0], otherTokenPrices[1])
			priceDiff1 := new(big.Float).Quo(otherTokenPrices[1], otherTokenPrices[0])
			if priceDiff0.Cmp(priceDiff1) == 1 {
				market.CurrentArbitrageOpp = priceDiff0
			} else {
				market.CurrentArbitrageOpp = priceDiff1
			}
		}
	}

}

var cache map[string]*big.Float = make(map[string]*big.Float)

func (uniswapMarkets *UniswapV2Markets) UpdateScreen(asset string, protocol string) {
	// iterate and see if latest state has changed
	reserveChanges := 0
	for tokenAddress, market := range uniswapMarkets.Asset[asset][protocol].CrossMarketsByToken {
		if market.CurrentArbitrageOpp.Cmp(big.NewFloat(0)) == 1 {
			_, ok := cache[tokenAddress]
			if ok && market.CurrentArbitrageOpp.Cmp(cache[tokenAddress]) != 0 {
				fmt.Printf("%s at %s: %f\n", asset, tokenAddress, market.CurrentArbitrageOpp)
				reserveChanges++
			} else if !ok {
				fmt.Printf("%s at %s: %f\n", asset, tokenAddress, market.CurrentArbitrageOpp)
				reserveChanges++
			}
		}
		cache[tokenAddress] = market.CurrentArbitrageOpp
	}

	fmt.Printf("reserve changes for %s on %s: %d\n", asset, protocol, reserveChanges)
}

// abigen --bin=./builds/UniswapQuery.bin --abi=./builds/UniswapQuery.abi --pkg=generatedContracts --out=./generatedContracts/UniswapQuery.go
