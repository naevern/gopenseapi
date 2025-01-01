package opensea

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	defaultLimit     = 20
	defaultTimeFrame = "24h"
)

// MusicFilter extends NFTFilter to include music-specific filters
type MusicFilter struct {
	NFTFilter
	AnimationURLExists bool `json:"animation_url_exists,omitempty"`
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
func (c *Client) GetMusic(ctx context.Context, filter MusicFilter) (*NFTResponse, error) {
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}

	// Always set animation_url_exists for music NFTs
	filter.AnimationURLExists = true

	query := buildMusicQuery(musicEP, filter)
	return c.fetchNFTs(ctx, query)
}

// buildMusicQuery constructs the query string for music NFT requests
func buildMusicQuery(endpoint string, filter MusicFilter) string {
	query := fmt.Sprintf("%s?limit=%d&animation_url_exists=true", endpoint, filter.Limit)

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

	return query
}

// GetTrendingMusic retrieves trending music NFTs
func (c *Client) GetTrendingMusic(ctx context.Context, filter TrendingMusicFilter) (*TrendingMusicResponse, error) {
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}
	if filter.TimeWindow == "" {
		filter.TimeWindow = defaultTimeFrame
	}

	query := fmt.Sprintf("%s/trending?limit=%d&time_window=%s&animation_url_exists=true&order_by=sale_count",
		musicEP, filter.Limit, filter.TimeWindow)

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

// Convenience methods for common music NFT queries
func (c *Client) GetMusicByCollection(ctx context.Context, collectionSlug string) (*NFTResponse, error) {
	return c.GetMusic(ctx, MusicFilter{
		NFTFilter: NFTFilter{
			Collection: collectionSlug,
			Limit:      50,
		},
	})
}

func (c *Client) GetMusicByOwner(ctx context.Context, ownerAddress string) (*NFTResponse, error) {
	return c.GetMusic(ctx, MusicFilter{
		NFTFilter: NFTFilter{
			Owner: ownerAddress,
			Limit: 50,
		},
	})
}

// Convenience methods for trending music with different time windows
func (c *Client) GetTrendingMusicLast24Hours(ctx context.Context) (*TrendingMusicResponse, error) {
	return c.GetTrendingMusic(ctx, TrendingMusicFilter{TimeWindow: "24h"})
}

func (c *Client) GetTrendingMusicLastWeek(ctx context.Context) (*TrendingMusicResponse, error) {
	return c.GetTrendingMusic(ctx, TrendingMusicFilter{TimeWindow: "7d"})
}

func (c *Client) GetTrendingMusicLastMonth(ctx context.Context) (*TrendingMusicResponse, error) {
	return c.GetTrendingMusic(ctx, TrendingMusicFilter{TimeWindow: "30d"})
}
