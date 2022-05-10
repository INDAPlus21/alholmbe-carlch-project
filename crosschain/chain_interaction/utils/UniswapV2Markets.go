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
const BATCH_COUNT_LIMIT int = 50
const UNISWAP_BATCH_SIZE int = 50

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
	Symbol   string
	Address  string
	Protocol string
}

// helper function for seeing if a list contains something
func in(address string, tokens []Token) bool {
	for _, token := range tokens {
		if address == token.Address {
			return true
		}
	}
	return false
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

			if !in(pair[0].String(), tokensOfInterest) && !in(pair[1].String(), tokensOfInterest) {
				// we don't care if none of the tokens in the pair is weth or wbnb
				continue
			}

			uniswapV2EthPair := &UniswapV2EthPair{pairAddress, pair[0], pair[1], big.NewInt(0), big.NewInt(0)}

			marketPairs = append(marketPairs, uniswapV2EthPair)
		}
	}

	return marketPairs

}

func (uniswapMarkets *UniswapV2Markets) UpdateMarkets(
	client *ethclient.Client,
	addresses []string,
	queryContractAddress string,
	tokensOfInterest []Token) {
	// WETH_ADDRESS := "0x2170Ed0880ac9A755fd29B2688956BD959F933F8"
	WBNB_ADDRESS := "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"
	WBNB := common.HexToAddress(WBNB_ADDRESS)

	// for every address, get markets
	// markets is a list with all pairs on this network
	allMarkets := [][]*UniswapV2EthPair{}
	for i := 0; i < len(addresses); i++ {
		marketPairs := uniswapMarkets.uniswapV2MarketByFactory(client, addresses[i],
			queryContractAddress, tokensOfInterest)

		allMarkets = append(allMarkets, marketPairs)
	}

	// group markets by non weth token address
	// mapped from the non weth token address
	allMarketsByToken := map[string][]*UniswapV2EthPair{}
	// a flat list of all markets
	allMarketsFlat := []*UniswapV2EthPair{}

	// groups all pairs into a dictionary with the non weth token as the key
	// O(n^2) doesn't matter, this function will only run on startup
	for i := 0; i < len(allMarkets); i++ {
		for j := 0; j < len(allMarkets[i]); j++ {
			allMarketsFlat = append(allMarketsFlat, allMarkets[i][j])
			if allMarkets[i][j].Token0Address == WBNB {
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

	for _, token := range tokensOfInterest {
		uniswapMarkets.Asset[token.Symbol][token.Protocol].AllMarkets = allMarketsFlat

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
			// remove depending on the reserve balance
			// reserve[0] is token0s reserve, reserve[1] is token1s reserve, reserve[2] is last interaction
			(uniswapMarkets.Asset[token.Symbol][token.Protocol].CrossMarkets)[i].Token0Balance = reserves[i][0]
			(uniswapMarkets.Asset[token.Symbol][token.Protocol].CrossMarkets)[i].Token1Balance = reserves[i][1]

			uniswapMarkets.Asset[token.Symbol][token.Protocol].CrossMarketsFiltered =
				append(uniswapMarkets.Asset[token.Symbol][token.Protocol].CrossMarketsFiltered,
					(uniswapMarkets.Asset[token.Symbol][token.Protocol].CrossMarkets)[i])

		}
	}

}

func (uniswapMarkets *UniswapV2Markets) EvaluateCrossMarkets() {
	WBNB_ADDRESS := "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"
	WBNB := common.HexToAddress(WBNB_ADDRESS)
	_ = WBNB

	for tokenAddress, market := range uniswapMarkets.Asset["WBNB"]["bsc"].CrossMarketsByToken {
		// fmt.Println(tokenAddress)
		_ = tokenAddress
		bnbPrices := []*big.Float{}
		otherTokenPrices := []*big.Float{}
		for _, market := range market.Pairs {
			// fmt.Println(market)
			if market.Token0Address == WBNB {
				bnbPrice := new(big.Float).Quo(new(big.Float).SetInt(market.Token0Balance), new(big.Float).SetInt(market.Token1Balance))
				otherTokenPrice := new(big.Float).Quo(new(big.Float).SetInt(market.Token1Balance), new(big.Float).SetInt(market.Token0Balance))
				bnbPrices = append(bnbPrices, bnbPrice)
				otherTokenPrices = append(otherTokenPrices, otherTokenPrice)
			} else {
				otherTokenPrice := new(big.Float).Quo(new(big.Float).SetInt(market.Token0Balance), new(big.Float).SetInt(market.Token1Balance))
				bnbPrice := new(big.Float).Quo(new(big.Float).SetInt(market.Token1Balance), new(big.Float).SetInt(market.Token0Balance))
				bnbPrices = append(bnbPrices, bnbPrice)
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

func getTokensIn(reserveIn *big.Int, reserveOut *big.Int, amountOut *big.Int) *big.Int {
	fmt.Println("reserveIn:", reserveIn)
	fmt.Println("amountOut:", amountOut)
	numerator := reserveIn.Mul(reserveIn, amountOut)
	fmt.Println("numerator:", numerator)
	numerator = numerator.Mul(numerator, big.NewInt(1000))
	fmt.Println("numerator:", numerator)
	denominator := reserveOut.Sub(reserveOut, amountOut).Mul(reserveOut, big.NewInt(997))
	fmt.Println("denominator:", denominator)
	res := numerator.Div(numerator, denominator)
	// res = res.Add(res, big.NewInt(1))

	return res
}

func getTokensOut(reserveIn *big.Int, reserveOut *big.Int, amountIn *big.Int) *big.Int {
	amountInWithFee := amountIn.Mul(amountIn, big.NewInt(997))
	numerator := amountInWithFee.Mul(amountInWithFee, reserveOut)
	denominator := reserveIn.Mul(reserveIn, big.NewInt(1000))
	denominator = denominator.Add(denominator, amountInWithFee)
	res := numerator.Div(numerator, denominator)

	return res
}

// abigen --bin=./builds/UniswapQuery.bin --abi=./builds/UniswapQuery.abi --pkg=generatedContracts --out=./generatedContracts/UniswapQuery.go
