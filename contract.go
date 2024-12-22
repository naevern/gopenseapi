package opensea

import (
	"context"
	"encoding/json"
	"fmt"
)

// AssetContract represents an NFT contract on OpenSea
type AssetContract struct {
	Collection     Collection `json:"collection" bson:"collection"`
	Address        Address    `json:"address" bson:"address"`
	ContractType   string     `json:"asset_contract_type" bson:"asset_contract_type"`
	CreatedDate    string     `json:"created_date" bson:"created_date"`
	Name           string     `json:"name" bson:"name"`
	NFTVersion     string     `json:"nft_version" bson:"nft_version"`
	OpenseaVersion any        `json:"opensea_version" bson:"opensea_version"`
	Owner          int64      `json:"owner" bson:"owner"`
	SchemaName     string     `json:"schema_name" bson:"schema_name"`
	Symbol         string     `json:"symbol" bson:"symbol"`
	TotalSupply    any        `json:"total_supply" bson:"total_supply"`
	Description    string     `json:"description" bson:"description"`
	ExternalLink   string     `json:"external_link" bson:"external_link"`
	ImageURL       string     `json:"image_url" bson:"image_url"`

	// Fee configuration
	DefaultToFiat              bool    `json:"default_to_fiat" bson:"default_to_fiat"`
	DevBuyerFeeBasisPoints     int64   `json:"dev_buyer_fee_basis_points" bson:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints    int64   `json:"dev_seller_fee_basis_points" bson:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers       bool    `json:"only_proxied_transfers" bson:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints int64   `json:"opensea_buyer_fee_basis_points" bson:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoint int64   `json:"opensea_seller_fee_basis_points" bson:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints        int64   `json:"buyer_fee_basis_points" bson:"buyer_fee_basis_points"`
	SellerFeeBasisPoints       int64   `json:"seller_fee_basis_points" bson:"seller_fee_basis_points"`
	PayoutAddress              Address `json:"payout_address" bson:"payout_address"`
}

// GetContract retrieves a single contract by its address
func (c *Client) GetContract(ctx context.Context, contractAddress string) (*AssetContract, error) {
	if contractAddress == "" {
		return nil, ErrEmptyContractAddress
	}

	path := fmt.Sprintf("%s/%s", singleContractEndpoint, contractAddress)
	resp, err := c.get(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("failed to get contract: %w", err)
	}

	var contract AssetContract
	if err := json.Unmarshal(resp, &contract); err != nil {
		return nil, fmt.Errorf("failed to unmarshal contract: %w", err)
	}

	return &contract, nil
}

// GetContractWithoutContext is a convenience wrapper around GetContract
func (c *Client) GetContractWithoutContext(contractAddress string) (*AssetContract, error) {
	return c.GetContract(context.Background(), contractAddress)
}
