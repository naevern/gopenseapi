package opensea

import (
	"math/big"
	"strings"
)

// Big converts the Number to a big.Int, ignoring decimal places
func (n Number) Big() *big.Int {
	s := strings.Split(string(n), ".")
	result, _ := new(big.Int).SetString(s[0], 10)
	return result
}
