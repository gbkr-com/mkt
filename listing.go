package mkt

import (
	"github.com/shopspring/decimal"
)

// A Listing which can be traded with counterparties. These are the essential
// details required for trading.
//
// Listing is principally a data object and stateless, so its fields are
// exported for convenience.
type Listing struct {
	Symbol             string          // FIX field 55
	TickIncrement      decimal.Decimal // FIX field 1208
	RoundLot           decimal.Decimal // FIX field 561
	MinTradeVol        decimal.Decimal // FIX field 562
	ContractMultiplier decimal.Decimal // FIX field 231
}

// Definition returns the [*Listing].
func (x *Listing) Definition() *Listing { return x }

// AnyListing defines all types which can embed, or derive from, [Listing].
type AnyListing interface {
	Definition() *Listing
}
