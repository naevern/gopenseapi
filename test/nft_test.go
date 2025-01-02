package opensea_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockClient struct {
	Responses map[string][]byte
	Errors    map[string]error
}

func (m *MockClient) get(ctx context.Context, path string) ([]byte, error) {
	if err, ok := m.Errors[path]; ok {
		return nil, err
	}
	if resp, ok := m.Responses[path]; ok {
		return resp, nil
	}
	return nil, errors.New("unexpected path")
}

func TestGetNFT(t *testing.T) {
	mockClient := &MockClient{
		Responses: map[string][]byte{
			"/assetEP/0x123/1": []byte(`{"name":"SampleNFT","token_id":"1"}`),
		},
		Errors: map[string]error{
			"/assetEP/0x123/2": errors.New("NFT not found"),
		},
	}

	client := &Client{mockClient}
	ctx := context.Background()

	t.Run("Successful NFT retrieval", func(t *testing.T) {
		asset, err := client.GetNFT(ctx, "0x123", "1")
		require.NoError(t, err)
		assert.Equal(t, "SampleNFT", asset.Name)
		assert.Equal(t, "1", asset.TokenID)
	})

	t.Run("Empty contract address", func(t *testing.T) {
		asset, err := client.GetNFT(ctx, "", "1")
		assert.Error(t, err)
		assert.Nil(t, asset)
		assert.EqualError(t, err, "contract address cannot be empty")
	})

	t.Run("Empty token ID", func(t *testing.T) {
		asset, err := client.GetNFT(ctx, "0x123", "")
		assert.Error(t, err)
		assert.Nil(t, asset)
		assert.EqualError(t, err, "token ID cannot be empty")
	})

	t.Run("NFT not found", func(t *testing.T) {
		asset, err := client.GetNFT(ctx, "0x123", "2")
		assert.Error(t, err)
		assert.Nil(t, asset)
		assert.EqualError(t, err, "failed to get NFT: NFT not found")
	})

	t.Run("Invalid JSON response", func(t *testing.T) {
		mockClient.Responses["/assetEP/0x123/3"] = []byte(`invalid json`)
		asset, err := client.GetNFT(ctx, "0x123", "3")
		assert.Error(t, err)
		assert.Nil(t, asset)
		assert.Contains(t, err.Error(), "failed to unmarshal NFT")
	})
}
