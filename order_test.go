package opensea_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockOpensea struct {
	Response string
	Error    error
}

func (m *MockOpensea) GetPath(ctx context.Context, path string) ([]byte, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	return []byte(m.Response), nil
}

func TestGetOrders(t *testing.T) {
	mockResponse := `{
		"count": 1,
		"orders": [
			{
				"id": 1,
				"asset": {"name": "Test Asset"},
				"created_date": "2024-01-01T00:00:00Z",
				"closing_date": "2024-01-10T00:00:00Z",
				"expiration_time": 1700000000,
				"listing_time": 1690000000,
				"exchange": "0xExchange",
				"maker": {"address": "0xMaker"},
				"taker": {"address": "0xTaker"},
				"current_price": "10.5",
				"maker_relayer_fee": "1.0",
				"taker_relayer_fee": "2.0",
				"maker_protocol_fee": "0.1",
				"taker_protocol_fee": "0.2",
				"fee_recipient": {"address": "0xFeeRecipient"},
				"fee_method": 0,
				"side": 1,
				"sale_kind": 0,
				"target": "0xTarget",
				"how_to_call": 0,
				"calldata": "",
				"replacement_pattern": "",
				"static_target": "",
				"static_extradata": "",
				"payment_token": "",
				"base_price": "10.0",
				"extra": "1.0",
				"quantity": "1",
				"salt": "12345",
				"approved_on_chain": false,
				"cancelled": false,
				"finalized": false,
				"marked_invalid": false
			}
		]
	}`

	opensea := MockOpensea{Response: mockResponse, Error: nil}

	orders, err := opensea.GetOrders("0xAssetContract", 1690000000)

	assert.NoError(t, err, "Expected no error while fetching orders")
	assert.Equal(t, 1, len(orders), "Expected one order in the response")
	assert.Equal(t, int64(1), orders[0].ID, "Order ID should match")
	assert.Equal(t, "0xMaker", orders[0].Maker.Address, "Maker address should match")
}

func TestGetOrdersWithError(t *testing.T) {
	opensea := MockOpensea{Error: errors.New("mock error")}

	_, err := opensea.GetOrders("0xAssetContract", 1690000000)

	assert.Error(t, err, "Expected an error while fetching orders")
	assert.Equal(t, "mock error", err.Error(), "Error message should match")
}
