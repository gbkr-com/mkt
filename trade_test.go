package mkt

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestTrade(t *testing.T) {

	decimal42 := decimal.New(42, 0)
	decimal43 := decimal.New(43, 0)
	decimal100 := decimal.New(100, 0)

	base := &Trade{Symbol: "A"}

	trade := &Trade{
		Symbol:      "B",
		LastQty:     decimal100,
		LastPx:      decimal42,
		TradeVolume: decimal.Zero,
		AvgPx:       decimal.Zero,
	}
	base.Accumulate(trade, 1)
	assert.True(t, base.LastQty.Equal(decimal.Zero), "different symbol")

	trade.Symbol = "A"
	base.Accumulate(trade, 1)
	assert.True(t, base.LastQty.Equal(decimal100))
	assert.True(t, base.TradeVolume.Equal(decimal100))
	assert.True(t, base.AvgPx.Equal(decimal42))

	trade.LastPx = decimal43
	base.Accumulate(trade, 1)
	assert.True(t, base.LastPx.Equal(decimal43))
	assert.True(t, base.TradeVolume.Equal(decimal.New(200, 0)))
	assert.True(t, base.AvgPx.Equal(decimal.New(425, -1)))

}
