package opensea

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// OrderStatus represents the current state of an order
type OrderStatus struct {
	ApprovedOnChain bool `json:"approved_on_chain" bson:"approved_on_chain"`
	Cancelled       bool `json:"cancelled" bson:"cancelled"`
	Finalized       bool `json:"finalized" bson:"finalized"`
	MarkedInvalid   bool `json:"marked_invalid" bson:"marked_invalid"`
}

// OrderFees contains all fee-related fields
type OrderFees struct {
	MakerRelayerFee  Number  `json:"maker_relayer_fee" bson:"maker_relayer_fee"`
	TakerRelayerFee  Number  `json:"taker_relayer_fee" bson:"taker_relayer_fee"`
	MakerProtocolFee Number  `json:"maker_protocol_fee" bson:"maker_protocol_fee"`
	TakerProtocolFee Number  `json:"taker_protocol_fee" bson:"taker_protocol_fee"`
	MakerReferrerFee Number  `json:"maker_referrer_fee" bson:"maker_referrer_fee"`
	FeeRecipient    Account `json:"fee_recipient" bson:"fee_recipient"`
	FeeMethod       FeeMethod `json:"fee_method" bson:"fee_method"`
}

// OrderSignature contains signature-related fields
type OrderSignature struct {
	V    *uint8 `json:"v" bson:"v"`
	R    *Bytes `json:"r" bson:"r"`
	S    *Bytes `json:"s" bson:"s"`
	Salt Number `json:"salt" bson:"salt"`
}

// Order represents an OpenSea order
type Order struct {
	ID            int64      `json:"id" bson:"id"`
	Asset         Asset      `json:"asset" bson:"asset"`
	CreatedDate   *TimeNano  `json:"created_date" bson:"created_date"`
	ClosingDate   *TimeNano  `json:"closing_date" bson:"closing_date"`
	ExpirationTime int64     `json:"expiration_time" bson:"expiration_time"`
	ListingTime    int64     `json:"listing_time" bson:"listing_time"`
	
	// Participants
	Exchange      Address  `json:"exchange" bson:"exchange"`
	Maker         Account  `json:"maker" bson:"maker"`
	Taker         Account  `json:"taker" bson:"taker"`
	
	// Price information
	CurrentPrice  Number   `json:"current_price" bson:"current_price"`
	BasePrice     Number   `json:"base_price" bson:"base_price"`
	Extra         Number   `json:"extra" bson:"extra"`
	Quantity      string   `json:"quantity" bson:"quantity"`
	PaymentToken  Address  `json:"payment_token" bson:"payment_token"`
	
	// Fee information
	Fees OrderFees `json:"fees" bson:"fees"`
	
	// Order parameters
	Side               Side      `json:"side" bson:"side"`
	SaleKind           SaleKind  `json:"sale_kind" bson:"sale_kind"`
	Target             Address   `json:"target" bson:"target"`
	HowToCall          HowToCall `json:"how_to_call" bson:"how_to_call"`
	Calldata           Bytes     `json:"calldata" bson:"calldata"`
	ReplacementPattern Bytes     `json:"replacement_pattern" bson:"replacement_pattern"`
	StaticTarget       Address   `json:"static_target" bson:"static_target"`
	StaticExtradata    Bytes     `json:"static_extradata" bson:"static_extradata"`
	
	// Status and signature
	Status    OrderStatus    `json:"status" bson:"status"`
	Signature OrderSignature `json:"signature" bson:"signature"`
}

// IsPrivate returns true if the order has a specific taker address
func (o Order) IsPrivate() bool {
	return o.Taker.Address != NullAddress
}

// IsExpired returns true if the order has expired
func (o Order) IsExpired() bool {
	return time.Unix(o.ExpirationTime, 0).Before(time.Now())
}

// IsValid returns true if the order is currently valid
func (o Order) IsValid() bool {
	return !o.Status.MarkedInvalid && 
	       !o.Status.Cancelled && 
	       !o.IsExpired() &&
	       o.Status.ApprovedOnChain
}

// IsSellOrder returns true if this is a sell order
func (o Order) IsSellOrder() bool {
	return o.Side == SideSell
}
