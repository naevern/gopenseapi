package opensea_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"your_project_path/opensea"
)

func TestParseAddress_Valid(t *testing.T) {
	input := "0x123456789abcdef"
	expected := opensea.Address("0x123456789abcdef")

	address, err := opensea.ParseAddress(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, address)
}

func TestParseAddress_EmptyInput(t *testing.T) {
	input := ""

	address, err := opensea.ParseAddress(input)

	assert.Error(t, err)
	assert.Equal(t, opensea.NullAddress, address)
	assert.Equal(t, "empty address", err.Error())
}

func TestAddressString(t *testing.T) {
	address := opensea.Address("0x123456789abcdef")
	expected := "0x123456789abcdef"

	result := address.String()

	assert.Equal(t, expected, result)
}
