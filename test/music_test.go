package opensea

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMusic(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		if r.URL.Path != "/api/v1/assets" {
			t.Errorf("Expected path '/api/v1/assets', got %s", r.URL.Path)
		}

		// Check query parameters
		q := r.URL.Query()
		if !q.Has("animation_url_exists") {
			t.Error("Missing animation_url_exists parameter")
		}

		// Return mock response
		response := MusicResponse{
			Assets: []Asset{
				{
					ID:           1,
					TokenID:      "123",
					Name:         "Test Music NFT",
					AnimationURL: "https://example.com/music.mp3",
				},
			},
			Next: "next_page_token",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create client with mock server
	client := &Client{
		baseURL:    server.URL,
		httpClient: server.Client(),
	}

	// Test GetMusic
	filter := MusicFilter{
		Collection: "test-collection",
		Limit:      20,
	}

	resp, err := client.GetMusic(context.Background(), filter)
	if err != nil {
		t.Fatalf("GetMusic failed: %v", err)
	}

	// Verify response
	if len(resp.Assets) != 1 {
		t.Errorf("Expected 1 asset, got %d", len(resp.Assets))
	}
	if resp.Assets[0].Name != "Test Music NFT" {
		t.Errorf("Expected name 'Test Music NFT', got %s", resp.Assets[0].Name)
	}
}

func TestGetTrendingMusic(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		if r.URL.Path != "/api/v1/assets/trending" {
			t.Errorf("Expected path '/api/v1/assets/trending', got %s", r.URL.Path)
		}

		// Check query parameters
		q := r.URL.Query()
		if !q.Has("time_window") {
			t.Error("Missing time_window parameter")
		}

		// Return mock response
		response := TrendingMusicResponse{
			Assets: []Asset{
				{
					ID:           1,
					TokenID:      "123",
					Name:         "Trending Music NFT",
					AnimationURL: "https://example.com/trending.mp3",
					NumSales:     100,
				},
			},
			Stats: struct {
				TotalVolume float64 `json:"total_volume"`
				TotalSales  int     `json:"total_sales"`
			}{
				TotalVolume: 1000.0,
				TotalSales:  500,
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create client with mock server
	client := &Client{
		baseURL:    server.URL,
		httpClient: server.Client(),
	}

	// Test convenience methods
	t.Run("Last 24 Hours", func(t *testing.T) {
		resp, err := client.GetTrendingMusicLast24Hours(context.Background())
		if err != nil {
			t.Fatalf("GetTrendingMusicLast24Hours failed: %v", err)
		}

		// Verify response
		if len(resp.Assets) != 1 {
			t.Errorf("Expected 1 asset, got %d", len(resp.Assets))
		}
		if resp.Stats.TotalSales != 500 {
			t.Errorf("Expected 500 total sales, got %d", resp.Stats.TotalSales)
		}
	})

	t.Run("Custom Time Window", func(t *testing.T) {
		filter := TrendingMusicFilter{
			TimeWindow: "7d",
			Limit:      10,
		}

		resp, err := client.GetTrendingMusic(context.Background(), filter)
		if err != nil {
			t.Fatalf("GetTrendingMusic failed: %v", err)
		}

		// Verify response
		if len(resp.Assets) != 1 {
			t.Errorf("Expected 1 asset, got %d", len(resp.Assets))
		}
		if resp.Assets[0].Name != "Trending Music NFT" {
			t.Errorf("Expected name 'Trending Music NFT', got %s", resp.Assets[0].Name)
		}
	})
}

func TestMusicFilters(t *testing.T) {
	tests := []struct {
		name   string
		filter MusicFilter
		want   string
	}{
		{
			name: "Collection Filter",
			filter: MusicFilter{
				Collection: "test-collection",
				Limit:      20,
			},
			want: "collection=test-collection",
		},
		{
			name: "Owner Filter",
			filter: MusicFilter{
				Owner: "0x123",
				Limit: 20,
			},
			want: "owner=0x123",
		},
		{
			name: "Token IDs Filter",
			filter: MusicFilter{
				TokenIDs: []string{"1", "2"},
				Limit:    20,
			},
			want: "token_ids=1&token_ids=2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if !contains(r.URL.RawQuery, tt.want) {
					t.Errorf("Expected query to contain %s, got %s", tt.want, r.URL.RawQuery)
				}
				json.NewEncoder(w).Encode(MusicResponse{})
			}))
			defer server.Close()

			client := &Client{
				baseURL:    server.URL,
				httpClient: server.Client(),
			}

			_, err := client.GetMusic(context.Background(), tt.filter)
			if err != nil {
				t.Fatalf("GetMusic failed: %v", err)
			}
		})
	}
}

// Helper function to check if a string contains another string
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[len(s)-len(substr):] == substr
}
