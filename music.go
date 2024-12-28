package opensea

import (
	"context"
	"encoding/json"
	"fmt"
)

// MusicFilter represents parameters for filtering music NFTs
type MusicFilter struct {
	Collection string   `json:"collection,omitempty"`
	TokenIDs   []string `json:"token_ids,omitempty"`
	Owner      string   `json:"owner,omitempty"`
	Limit      int      `json:"limit,omitempty"`
	Offset     int      `json:"offset,omitempty"`
}

// MusicResponse represents the API response for music NFTs
type MusicResponse struct {
	Assets []Asset `json:"assets"`
	Next   string  `json:"next"`
}

// GetMusic retrieves music NFTs based on the provided filters
func (c *Client) GetMusic(ctx context.Context, filter MusicFilter) (*MusicResponse, error) {
	if filter.Limit == 0 {
		filter.Limit = 20 // Default limit
	}

	// Construct query parameters
	query := fmt.Sprintf("%s?limit=%d", musicEP, filter.Limit)

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

	// Add filter for assets with animation_url (music NFTs typically have this)
	query += "&animation_url_exists=true"

	resp, err := c.get(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get music NFTs: %w", err)
	}

	var musicResp MusicResponse
	if err := json.Unmarshal(resp, &musicResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal music response: %w", err)
	}

	return &musicResp, nil
}

// GetMusicByCollection is a convenience method to get music NFTs from a specific collection
func (c *Client) GetMusicByCollection(ctx context.Context, collectionSlug string) (*MusicResponse, error) {
	return c.GetMusic(ctx, MusicFilter{
		Collection: collectionSlug,
		Limit:      50,
	})
}

// GetMusicByOwner is a convenience method to get music NFTs owned by a specific address
func (c *Client) GetMusicByOwner(ctx context.Context, ownerAddress string) (*MusicResponse, error) {
	return c.GetMusic(ctx, MusicFilter{
		Owner: ownerAddress,
		Limit: 50,
	})
}
