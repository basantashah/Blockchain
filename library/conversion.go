package library

import (
	"math"
	"math/big"
)

// AmountConversion convert Wei amount to Eth
func AmountConversion(a int) string {
	fbalance := new(big.Float)
	fbalance.SetString(a.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow(18)))
	return ethValue
}
