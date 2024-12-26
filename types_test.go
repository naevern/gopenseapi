package opensea

import (
	"testing"
)

func TestParseAddress(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Address
		wantErr bool
	}{
		{"Valid address", "0x1234567890abcdef", Address("0x1234567890abcdef"), false},
		{"Empty address", "", NullAddress, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAddress(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddressString(t *testing.T) {
	addr := Address("0x1234567890abcdef")
	if got := addr.String(); got != "0x1234567890abcdef" {
		t.Errorf("Address.String() = %v, want %v", got, "0x1234567890abcdef")
	}
}

func TestTimeNanoString(t *testing.T) {
	var time TimeNano = 1234567890
	if got := time.String(); got != "1234567890" {
		t.Errorf("TimeNano.String() = %v, want %v", got, "1234567890")
	}
}

func TestNumberString(t *testing.T) {
	number := Number("123.45")
	if got := number.String(); got != "123.45" {
		t.Errorf("Number.String() = %v, want %v", got, "123.45")
	}
}

func TestBytesString(t *testing.T) {
	bytes := Bytes([]byte("hello"))
	if got := bytes.String(); got != "hello" {
		t.Errorf("Bytes.String() = %v, want %v", got, "hello")
	}
}
