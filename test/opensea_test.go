package opensea

import (
	"github.com/cheekybits/is"
	"math/big"
	"os"
	"testing"
)

var (
	openseaClient *Opensea
)

func TestMain(m *testing.M) {
	// Initialize the Opensea client before running the tests
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		panic("API_KEY environment variable is not set")
	}

	var err error
	openseaClient, err = NewOpensea(apiKey)
	if err != nil {
		panic("Failed to initialize Opensea client: " + err.Error())
	}

	os.Exit(m.Run())
}

func TestGetSingleAsset(t *testing.T) {
	assert := is.New(t)

	contractAddress := "0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb"
	tokenID := big.NewInt(1)

	singleAsset, err := openseaClient.GetSingleAsset(contractAddress, tokenID)
	assert.Nil(err)

	assert.Equal(singleAsset.Name, "CryptoPunk #1")
	assert.Equal(singleAsset.ExternalLink, "https://www.larvalabs.com/cryptopunks/details/1")
	assert.Equal(singleAsset.AssetContract.AssetContractType, "non-fungible")
}

func TestGetSingleAssetContract(t *testing.T) {
	assert := is.New(t)

	contractAddress := "0x06012c8cf97bead5deae237070f9587f8e7a266d"

	cryptoKittiesContract, err := openseaClient.GetSingleContract(contractAddress)
	assert.Nil(err)

	assert.Equal(cryptoKittiesContract.Name, "CryptoKitties")
	assert.Equal(cryptoKittiesContract.SchemaName, "ERC721")
	assert.Equal(cryptoKittiesContract.Collection.ExternalUrl, "https://www.cryptokitties.co/")
	assert.Equal(cryptoKittiesContract.Collection.DiscordUrl, "https://discord.gg/cryptokitties")
}

