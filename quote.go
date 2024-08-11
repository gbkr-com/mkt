package mkt

import (
	"github.com/shopspring/decimal"
)

// Quote is the best bid and ask for a symbol.
//
// Quote is principally a data object and stateless, so its fields are exported
// for convenience.
type Quote struct {
	Symbol  string          // FIX field 55
	BidPx   decimal.Decimal // FIX field 132
	BidSize decimal.Decimal // FIX field 134
	AskPx   decimal.Decimal // FIX field 133, renamed from OfferPx
	AskSize decimal.Decimal // FIX field 135, renamed from OfferSize
}

// Near returns the passive price and size for the given [Side].
func (x *Quote) Near(side Side) (decimal.Decimal, decimal.Decimal) {
	if x == nil {
		return decimal.Zero, decimal.Zero
	}
	switch side {
	case Buy:
		return x.BidPx, x.BidSize
	case Sell:
		return x.AskPx, x.AskSize
	default:
		return decimal.Zero, decimal.Zero
	}
}

// Far returns the aggressive price and size for the given [Side].
func (x *Quote) Far(side Side) (decimal.Decimal, decimal.Decimal) {
	return x.Near(side.Opposite())
}

// Spread returns the difference between the ask and the bid. This will be an
// integral number of ticks.
func (x *Quote) Spread() decimal.Decimal {
	if x == nil {
		return decimal.Zero
	}
	if x.BidPx.IsZero() {
		return decimal.Zero
	}
	if x.AskPx.IsZero() {
		return decimal.Zero
	}
	return x.AskPx.Sub(x.BidPx)
}

// MidPrice returns the mean of the bid and ask. The mean is not expected to
// be tick aligned.
func (x *Quote) MidPrice() decimal.Decimal {
	if x == nil {
		return decimal.Zero
	}
	if x.BidPx.IsZero() {
		return decimal.Zero
	}
	if x.AskPx.IsZero() {
		return decimal.Zero
	}
	return decimal.Avg(x.BidPx, x.AskPx)
}

// ZeroQuote will 'zero' all the fields in the [*Quote], never returning nil.
// This is a convenience for using quotes in a [utl.Pool].
func ZeroQuote(quote *Quote) *Quote {
	if quote == nil {
		return &Quote{}
	}
	quote.Symbol = ""
	quote.BidPx, quote.BidSize = decimal.Zero, decimal.Zero
	quote.AskPx, quote.AskSize = decimal.Zero, decimal.Zero
	return quote
}
