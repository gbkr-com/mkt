package mkt

import (
	"github.com/shopspring/decimal"
)

// Quote is the best bid and ask for a symbol.
type Quote struct {
	Symbol  string
	Bid     decimal.Decimal
	BidSize decimal.Decimal
	Ask     decimal.Decimal
	AskSize decimal.Decimal
}

// Near returns the passive price and size for the given [Side].
func (x *Quote) Near(side Side) (decimal.Decimal, decimal.Decimal) {
	if x == nil {
		return decimal.Zero, decimal.Zero
	}
	switch side {
	case Buy:
		return x.Bid, x.BidSize
	case Sell:
		return x.Ask, x.AskSize
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
	if x.Bid.IsZero() {
		return decimal.Zero
	}
	if x.Ask.IsZero() {
		return decimal.Zero
	}
	return x.Ask.Sub(x.Bid)
}

// MidPrice returns the mean of the bid and ask. The mean is not expected to
// be tick aligned.
func (x *Quote) MidPrice() decimal.Decimal {
	if x == nil {
		return decimal.Zero
	}
	if x.Bid.IsZero() {
		return decimal.Zero
	}
	if x.Ask.IsZero() {
		return decimal.Zero
	}
	return decimal.Avg(x.Bid, x.Ask)
}

// ZeroQuote will 'zero' all the fields in the [*Quote], never returning nil.
// This is a convenience for using quotes in a [utl.Pool].
func ZeroQuote(quote *Quote) *Quote {
	if quote == nil {
		return &Quote{}
	}
	quote.Symbol = ""
	quote.Bid, quote.BidSize = decimal.Zero, decimal.Zero
	quote.Ask, quote.AskSize = decimal.Zero, decimal.Zero
	return quote
}
