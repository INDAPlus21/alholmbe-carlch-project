package misc

type Token struct {
	// Address to token
	address string

	// Value of token in specific value
	balance int
}

type Pair struct {
	// Address to specific pair
	pairAddress string

	// Pair of token
	baseToken     Token
	externalToken Token
}
