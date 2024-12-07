package opensea

import (
	"encoding/hex"
	"errors"
	"math/big"
	"strconv"
	"strings"
)

// Number represents a numeric value that can be converted to big.Int
type Number string

// Big converts the Number to a big.Int, ignoring decimal places
func (n Number) Big() *big.Int {
	s := strings.Split(string(n), ".")
	result, _ := new(big.Int).SetString(s[0], 10)
	return result
}

// Address represents an Ethereum address
type Address string

// NullAddress represents the Ethereum zero address
const NullAddress Address = "0x0000000000000000000000000000000000000000"

// IsHexAddress validates if a string is a valid Ethereum address
func IsHexAddress(s string) bool {
	if s == "0x0" {
		return true
	}
	if len(s) < 2 || s[:2] != "0x" {
		return false
	}
	
	const addressLength = 42 // 2 (0x) + 40 (hex chars)
	if len(s) != addressLength {
		return false
	}
	
	_, err := hex.DecodeString(s[2:])
	return err == nil
}

// ParseAddress converts a string to an Address type with validation
func ParseAddress(address string) (Address, error) {
	if !IsHexAddress(address) {
		return "", errors.New("invalid address: " + address)
	}
	return Address(strings.ToLower(address)), nil
}

// String returns the string representation of the address
func (a Address) String() string {
	return string(a)
}

// IsNullAddress checks if the address is the null address
func (a Address) IsNullAddress() bool {
	return a.String() == NullAddress.String()
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (a *Address) UnmarshalJSON(b []byte) error {
	var s string
	if string(b) == "null" {
		s = NullAddress.String()
	} else {
		var err error
		s, err = strconv.Unquote(string(b))
		if err != nil {
			return err
		}
	}
	
	parsed, err := ParseAddress(s)
	if err != nil {
		return err
	}
	*a = parsed
	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (a Address) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(a.String())), nil
}

// Bytes represents a byte slice that can be marshaled/unmarshaled to/from hex strings
type Bytes []byte

// Bytes32 converts the byte slice to a fixed-size [32]byte array
func (b Bytes) Bytes32() [32]byte {
    var result [32]byte
    // Only copy up to 32 bytes to prevent potential panic on shorter slices
    copyLen := min(len(b), 32)
    copy(result[:copyLen], b)
    return result
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (b *Bytes) UnmarshalJSON(data []byte) error {
    // Unquote the JSON string
    hexStr, err := strconv.Unquote(string(data))
    if err != nil {
        return fmt.Errorf("failed to unquote JSON string: %w", err)
    }

    // Handle empty string case
    if hexStr == "" {
        *b = Bytes{}
        return nil
    }

    // Validate hex string format
    if !strings.HasPrefix(hexStr, "0x") {
        return fmt.Errorf("hex string must start with 0x")
    }

    // Decode hex string (skip "0x" prefix)
    decoded, err := hex.DecodeString(hexStr[2:])
    if err != nil {
        return fmt.Errorf("failed to decode hex string: %w", err)
    }

    *b = decoded
    return nil
}

// MarshalJSON implements the json.Marshaler interface
func (b Bytes) MarshalJSON() ([]byte, error) {
    // Convert to hex string with "0x" prefix
    hexStr := "0x" + hex.EncodeToString(b)
    
    // Quote the string for JSON
    return []byte(strconv.Quote(hexStr)), nil
}

// min returns the minimum of two integers
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
