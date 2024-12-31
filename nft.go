package opensea

import (
	"context"
	"encoding/json"
	"fmt"
)

// NFTFilter represents parameters for filtering NFTs
type NFTFilter struct {
	Collection string   `json:"collection,omitempty"`
	TokenIDs   []string `json:"token_ids,omitempty"`
	Owner      string   `json:"owner,omitempty"`
	Limit      int      `json:"limit,omitempty"`
	Offset     int      `json:"offset,omitempty"`
	OrderBy    string   `json:"order_by,omitempty"`        // created_date, sale_date, etc.
	OrderDir   string   `json:"order_direction,omitempty"` // desc or asc
}

// NFTResponse represents the API response for NFTs
type NFTResponse struct {
	Assets []Asset `json:"assets"`
	Next   string  `json:"next"`
}

// GetNFT retrieves a single NFT by contract address and token ID
func (c *Client) GetNFT(ctx context.Context, contractAddress, tokenID string) (*Asset, error) {
	if contractAddress == "" {
		return nil, fmt.Errorf("contract address cannot be empty")
	}
	if tokenID == "" {
		return nil, fmt.Errorf("token ID cannot be empty")
	}

	path := fmt.Sprintf("%s/%s/%s", assetEP, contractAddress, tokenID)
	resp, err := c.get(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("failed to get NFT: %w", err)
	}

	var asset Asset
	if err := json.Unmarshal(resp, &asset); err != nil {
		return nil, fmt.Errorf("failed to unmarshal NFT: %w", err)
	}

	return &asset, nil
}

// GetNFTs retrieves multiple NFTs based on the provided filters
func (c *Client) GetNFTs(ctx context.Context, filter NFTFilter) (*NFTResponse, error) {
	if filter.Limit == 0 {
		filter.Limit = 20 // Default limit
	}

	// Construct query parameters
	query := fmt.Sprintf("%s?limit=%d", assetEP, filter.Limit)

	if filter.Collection != "" {
		query += fmt.Sprintf("&collection=%s", filter.Collection)
	}
	if filter.Owner != "" {
		query += fmt.Sprintf("&owner=%s", filter.Owner)
	}
	if filter.Offset > 0 {
		query += fmt.Sprintf("&offset=%d", filter.Offset)
	}
	if len(filter.TokenIDs) > 0 {
		for _, id := range filter.TokenIDs {
			query += fmt.Sprintf("&token_ids=%s", id)
		}
	}
	if filter.OrderBy != "" {
		query += fmt.Sprintf("&order_by=%s", filter.OrderBy)
	}
	if filter.OrderDir != "" {
		query += fmt.Sprintf("&order_direction=%s", filter.OrderDir)
	}

	resp, err := c.get(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get NFTs: %w", err)
	}

	var nftResp NFTResponse
	if err := json.Unmarshal(resp, &nftResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal NFTs response: %w", err)
	}

	return &nftResp, nil
}

// GetNFTsByCollection is a convenience method to get NFTs from a specific collection
func (c *Client) GetNFTsByCollection(ctx context.Context, collectionSlug string) (*NFTResponse, error) {
	return c.GetNFTs(ctx, NFTFilter{
		Collection: collectionSlug,
		Limit:      50,
		OrderBy:    "created_date",
		OrderDir:   "desc",
	})
}

// GetNFTsByOwner is a convenience method to get NFTs owned by a specific address
func (c *Client) GetNFTsByOwner(ctx context.Context, ownerAddress string) (*NFTResponse, error) {
	return c.GetNFTs(ctx, NFTFilter{
		Owner:    ownerAddress,
		Limit:    50,
		OrderBy:  "created_date",
		OrderDir: "desc",
	})
}

// GetNFTsByTokenIDs is a convenience method to get NFTs by their token IDs
func (c *Client) GetNFTsByTokenIDs(ctx context.Context, contractAddress string, tokenIDs []string) (*NFTResponse, error) {
	if contractAddress == "" {
		return nil, fmt.Errorf("contract address cannot be empty")
	}

	return c.GetNFTs(ctx, NFTFilter{
		Collection: contractAddress,
		TokenIDs:   tokenIDs,
		Limit:      len(tokenIDs),
	})
}
