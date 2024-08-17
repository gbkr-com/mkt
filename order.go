package mkt

import "github.com/google/uuid"

// Order is the prototype for an order sent to a counterparty.
type Order struct {
	MsgType MsgType `json:"msgType"` // FIX field 35
	OrderID string  `json:"orderID"` // FIX field 37
	Side    Side    `json:"side"`    // FIX field 54
	Symbol  string  `json:"symbol"`  // FIX field 55
}

// Definition returns the [*Order].
func (x *Order) Definition() *Order { return x }

// AnyOrder defines all types which can embed, or derive from, [Order].
type AnyOrder interface {
	Definition() *Order
}

// NewOrderID is a convenience function to generate a unique OrderID.
func NewOrderID() string { return uuid.NewString() }
