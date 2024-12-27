package opensea

import "errors"

// ErrEmptyContractAddress is returned when attempting to get a contract with an empty address
var ErrEmptyContractAddress = errors.New("contract address cannot be empty")
