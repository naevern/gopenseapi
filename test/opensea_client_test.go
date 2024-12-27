package opensea_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	opensea "github.com/yourusername/opensea"
)

type testServer struct {
	server *httptest.Server
	client *opensea.Client
}

func newTestServer(t *testing.T, response interface{}, expectedPath string) *testServer {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, expectedPath, r.URL.Path)
		require.Equal(t, "application/json", r.Header.Get("Accept"))
		require.Equal(t, "test-api-key", r.Header.Get("X-API-KEY"))

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		require.NoError(t, err)
	}))

	client := opensea.NewClient(srv.URL, "test-api-key")
	return &testServer{
		server: srv,
		client: client,
	}
}

func TestClient_GetAsset(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		collectionSlug string
		tokenID        string
		mockResponse   *opensea.Asset
		expectedError  bool
	}{
		{
			name:           "successful asset fetch",
			collectionSlug: "test-collection",
			tokenID:        "123",
			mockResponse: &opensea.Asset{
				ID:          "1",
				TokenID:     "123",
				Name:        "Test NFT",
				Description: "Test NFT Description",
				Collection: opensea.Collection{
					Name: "Test Collection",
					Slug: "test-collection",
				},
			},
			expectedError: false,
		},
		{
			name:           "empty collection slug",
			collectionSlug: "",
			tokenID:        "123",
			mockResponse:   nil,
			expectedError:  true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ts := newTestServer(t, tc.mockResponse, "/api/v1/asset/"+tc.collectionSlug+"/"+tc.tokenID)
			defer ts.server.Close()

			ctx := context.Background()
			asset, err := ts.client.GetAsset(ctx, tc.collectionSlug, tc.tokenID)

			if tc.expectedError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.mockResponse.ID, asset.ID)
			assert.Equal(t, tc.mockResponse.Name, asset.Name)
			assert.Equal(t, tc.mockResponse.Collection.Name, asset.Collection.Name)
		})
	}
}

func TestClient_GetAssets(t *testing.T) {
	t.Parallel()

	mockResponse := &opensea.AssetsResponse{
		Assets: []opensea.Asset{
			{
				ID:      "1",
				TokenID: "123",
				Name:    "Test NFT 1",
			},
			{
				ID:      "2",
				TokenID: "124",
				Name:    "Test NFT 2",
			},
		},
	}

	ts := newTestServer(t, mockResponse, "/api/v1/assets")
	defer ts.server.Close()

	ctx := context.Background()
	params := opensea.AssetsParams{
		Owner:          "0x123",
		CollectionSlug: "test-collection",
		Limit:          20,
	}

	assets, err := ts.client.GetAssets(ctx, params)
	require.NoError(t, err)
	assert.Len(t, assets.Assets, 2)
	assert.Equal(t, "Test NFT 1", assets.Assets[0].Name)
	assert.Equal(t, "Test NFT 2", assets.Assets[1].Name)
}

func TestClient_GetCollection(t *testing.T) {
	t.Parallel()

	mockCollection := &opensea.Collection{
		Name:        "Test Collection",
		Slug:        "test-collection",
		Description: "Test Collection Description",
		Stats: opensea.CollectionStats{
			TotalSupply: 1000,
			NumOwners:   500,
			FloorPrice:  0.5,
		},
	}

	ts := newTestServer(t, mockCollection, "/api/v1/collection/test-collection")
	defer ts.server.Close()

	ctx := context.Background()
	collection, err := ts.client.GetCollection(ctx, "test-collection")

	require.NoError(t, err)
	assert.Equal(t, mockCollection.Name, collection.Name)
	assert.Equal(t, mockCollection.Slug, collection.Slug)
	assert.Equal(t, mockCollection.Stats.TotalSupply, collection.Stats.TotalSupply)
}

func TestClient_GetCollectionStats(t *testing.T) {
	t.Parallel()

	mockStats := &opensea.CollectionStats{
		TotalSupply: 1000,
		NumOwners:   500,
		FloorPrice:  0.5,
		TotalVolume: 1000.5,
	}

	ts := newTestServer(t, mockStats, "/api/v1/collection/test-collection/stats")
	defer ts.server.Close()

	ctx := context.Background()
	stats, err := ts.client.GetCollectionStats(ctx, "test-collection")

	require.NoError(t, err)
	assert.Equal(t, mockStats.TotalSupply, stats.TotalSupply)
	assert.Equal(t, mockStats.FloorPrice, stats.FloorPrice)
}

func TestClient_ErrorHandling(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(map[string]string{
			"error": "Not found",
		})
		require.NoError(t, err)
	}))
	defer srv.Close()

	client := opensea.NewClient(srv.URL, "test-api-key")
	ctx := context.Background()

	_, err := client.GetAsset(ctx, "non-existent", "123")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "404")
}
