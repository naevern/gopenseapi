package opensea

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"time"
)

const (
	mainnetAPI  = "https://api.opensea.io"
	testnetAPI  = "https://testnets-api.opensea.io"
	rinkebyAPI  = "https://rinkeby-api.opensea.io"
	basePath    = "/api/v1"
	contractEP  = basePath + "/asset_contract"
	assetEP     = basePath + "/asset"
)

type Opensea struct {
	API        string
	APIKey     string
	httpClient *http.Client
}

type errorResponse struct {
	Success bool `json:"success" bson:"success"`
}

func (e errorResponse) Error() string {
	return "Operation unsuccessful"
}

// NewOpensea initializes an Opensea instance for the mainnet.
func NewOpensea(apiKey string) *Opensea {
	return &Opensea{
		API:        mainnetAPI,
		APIKey:     apiKey,
		httpClient: newHttpClient(),
	}
}

// NewOpenseaRinkeby initializes an Opensea instance for the Rinkeby testnet.
func NewOpenseaRinkeby(apiKey string) *Opensea {
	return &Opensea{
		API:        rinkebyAPI,
		APIKey:     apiKey,
		httpClient: newHttpClient(),
	}
}

// newHttpClient creates a default HTTP client with a timeout.
func newHttpClient() *http.Client {
	return &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			IdleConnTimeout:       30 * time.Second,
			TLSHandshakeTimeout:   5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
}

