package mkt

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestQuote(t *testing.T) {

	decimal1 := decimal.New(1, 0)
	decimal42 := decimal.New(42, 0)
	decimal42_5 := decimal.New(425, -1)
	decimal43 := decimal.New(43, 0)
	decimal100 := decimal.New(100, 0)
	decimal200 := decimal.New(200, 0)

	quote := &Quote{
		Bid:     decimal42,
		BidSize: decimal100,
		Ask:     decimal43,
		AskSize: decimal200,
	}

	px, qty := quote.Near(Buy)
	assert.True(t, px.Equal(decimal42))
	assert.True(t, qty.Equal(decimal100))

	px, qty = quote.Far(Sell)
	assert.True(t, px.Equal(decimal42))
	assert.True(t, qty.Equal(decimal100))

	px, qty = quote.Far(Buy)
	assert.True(t, px.Equal(decimal43))
	assert.True(t, qty.Equal(decimal200))

	px, qty = quote.Near(Sell)
	assert.True(t, px.Equal(decimal43))
	assert.True(t, qty.Equal(decimal200))

	assert.True(t, quote.Spread().Equal(decimal1))
	assert.True(t, quote.MidPrice().Equal(decimal42_5))

	quote.Bid = decimal.Zero

	assert.True(t, quote.Spread().Equal(decimal.Zero))
	assert.True(t, quote.MidPrice().Equal(decimal.Zero))

}
