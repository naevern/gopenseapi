package opensea

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// TestGetSingleContract verifies the contract retrieval functionality
func TestGetSingleContract(t *testing.T) {
	// Test contract address
	contractAddress := "0xdceaf1652a131f32a821468dc03a92df0edd86ea"
	
	// Initialize test assertions
	assertIs := initializeTest(t)

	// Fetch contract details
	singleContract, err := o.GetSingleContract(contractAddress)

	// Verify no errors occurred
	assertIs.Nil(err)

	// Verify contract details match expected values
	assertIs.Equal(singleContract.Address, contractAddress)
	assertIs.Equal(singleContract.Name, "MyCryptoHeroes:Extension")
	assertIs.Equal(singleContract.AssetContractType, "non-fungible")
	assertIs.Equal(singleContract.Collection.Slug, "mycryptoheroes")
}

// print is a helper function to print structs and other types
func print(in interface{}) {
	if reflect.TypeOf(in).Kind() == reflect.Struct {
		in, _ = json.Marshal(in)
		in = string(in.([]byte))
	}
	fmt.Println(in)
}
