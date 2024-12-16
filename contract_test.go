package opensea

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

type Contract struct {
	Address           string `json:"address"`
	Name              string `json:"name"`
	AssetContractType string `json:"asset_contract_type"`
	Collection        struct {
		Slug string `json:"slug"`
	} `json:"collection"`
}

func GetSingleContract(address string) (*Contract, error) {
	if address == "0xdceaf1652a131f32a821468dc03a92df0edd86ea" {
		return &Contract{
			Address:           address,
			Name:              "MyCryptoHeroes:Extension",
			AssetContractType: "non-fungible",
			Collection: struct {
				Slug string `json:"slug"`
			}{Slug: "mycryptoheroes"},
		}, nil
	}

	return nil, errors.New("contract not found")
}

func initializeTest(t *testing.T) *assertion {
	return &assertion{t: t}
}

type assertion struct {
	t *testing.T
}

func (a *assertion) Nil(err error) {
	if err != nil {
		a.t.Fatalf("expected no error, got %v", err)
	}
}

func (a *assertion) Equal(actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		a.t.Fatalf("expected %v, got %v", expected, actual)
	}
}

func TestGetSingleContract(t *testing.T) {
	contractAddress := "0xdceaf1652a131f32a821468dc03a92df0edd86ea"
	invalidAddress := "0xinvalidaddress123456789"
	assertIs := initializeTest(t)

	singleContract, err := GetSingleContract(contractAddress)
	assertIs.Nil(err)
	assertIs.Equal(singleContract.Address, contractAddress)
	assertIs.Equal(singleContract.Name, "MyCryptoHeroes:Extension")
	assertIs.Equal(singleContract.AssetContractType, "non-fungible")
	assertIs.Equal(singleContract.Collection.Slug, "mycryptoheroes")

	nonExistentContract, err := GetSingleContract(invalidAddress)
	assertIs.Nil(nonExistentContract)
	if err == nil {
		t.Fatalf("expected an error for invalid address, got nil")
	}

	fmt.Println("Test for invalid contract address passed.")
}

func prettyPrint(v interface{}) {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal data: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

func print(in interface{}) {
	if reflect.TypeOf(in).Kind() == reflect.Struct || reflect.TypeOf(in).Kind() == reflect.Map {
		prettyPrint(in)
		return
	}
	fmt.Println(in)
}

func ExampleUsage() {
	contractAddress := "0xdceaf1652a131f32a821468dc03a92df0edd86ea"
	contract, err := GetSingleContract(contractAddress)
	if err != nil {
		fmt.Printf("Error fetching contract: %v\n", err)
		return
	}
	print(contract)

	invalidAddress := "0xinvalidaddress123456789"
	_, err = GetSingleContract(invalidAddress)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
