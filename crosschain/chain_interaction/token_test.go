package main

import (
	"testing"
)

func TestAddSet(t *testing.T) {
	set := initSet()
	set.add(WETH_ETHEREUM)
	if !set.exists(WETH_ETHEREUM) {
		t.Errorf("Expected element not found in set.")
	}
}

func TestGetSet(t *testing.T) {
	set := initSet()
	set.add(WETH_ETHEREUM)
	set.add(WETH_FANTOM)
	set.add(WETH_BINANCE)

	expectedSet := []string{
		WETH_ETHEREUM,
		WETH_FANTOM,
		WETH_BINANCE,
	}
	for i := 0; i < 3; i++ {
		if !set.exists(expectedSet[i]) {
			t.Errorf("Expected output %s was not found.", expectedSet[i])
		}
	}
}

func TestInitSets(t *testing.T) {
	initSets()
	expectedSet := []string{
		AVAX_AVALANCE,
		AVAX_BINANCE,
		AVAX_FANTOM,
	}
	for i := 0; i < 3; i++ {
		if !AVAX.exists(expectedSet[i]) {
			t.Errorf("Expected output %s was not found.", expectedSet[i])
		}
	}
}
