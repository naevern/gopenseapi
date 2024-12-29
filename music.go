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

// TrendingMusicFilter represents parameters for filtering trending music NFTs
type TrendingMusicFilter struct {
	TimeWindow string `json:"time_window,omitempty"` // 24h, 7d, 30d
	Limit      int    `json:"limit,omitempty"`
}

// TrendingMusicResponse represents the API response for trending music NFTs
type TrendingMusicResponse struct {
	Assets []Asset `json:"assets"`
	Stats  struct {
		TotalVolume float64 `json:"total_volume"`
		TotalSales  int     `json:"total_sales"`
	} `json:"stats"`
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

// GetTrendingMusic retrieves trending music NFTs based on volume and sales
func (c *Client) GetTrendingMusic(ctx context.Context, filter TrendingMusicFilter) (*TrendingMusicResponse, error) {
	if filter.Limit == 0 {
		filter.Limit = 20 // Default limit
	}
	if filter.TimeWindow == "" {
		filter.TimeWindow = "24h" // Default time window
	}

	// Construct query parameters
	query := fmt.Sprintf("%s/trending?limit=%d&time_window=%s", musicEP, filter.Limit, filter.TimeWindow)

	// Add music-specific filters
	query += "&animation_url_exists=true"
	query += "&order_by=sale_count" // Order by number of sales

	resp, err := c.get(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get trending music NFTs: %w", err)
	}

	var trendingResp TrendingMusicResponse
	if err := json.Unmarshal(resp, &trendingResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal trending music response: %w", err)
	}

	return &trendingResp, nil
}

// GetTrendingMusicLast24Hours is a convenience method to get trending music NFTs in the last 24 hours
func (c *Client) GetTrendingMusicLast24Hours(ctx context.Context) (*TrendingMusicResponse, error) {
	return c.GetTrendingMusic(ctx, TrendingMusicFilter{
		TimeWindow: "24h",
		Limit:      20,
	})
}

// GetTrendingMusicLastWeek is a convenience method to get trending music NFTs in the last week
func (c *Client) GetTrendingMusicLastWeek(ctx context.Context) (*TrendingMusicResponse, error) {
	return c.GetTrendingMusic(ctx, TrendingMusicFilter{
		TimeWindow: "7d",
		Limit:      20,
	})
}

// GetTrendingMusicLastMonth is a convenience method to get trending music NFTs in the last month
func (c *Client) GetTrendingMusicLastMonth(ctx context.Context) (*TrendingMusicResponse, error) {
	return c.GetTrendingMusic(ctx, TrendingMusicFilter{
		TimeWindow: "30d",
		Limit:      20,
	})
}
