package mkt

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestTraded(t *testing.T) {

	type statlisting struct {
		Listing
		MedTradeSize decimal.Decimal
	}

	whitelist := NewWhiteList[*statlisting]()
	whitelist.Add(&statlisting{
		Listing:      Listing{Symbol: "A", ContractMultiplier: decimalOne},
		MedTradeSize: decimal.Decimal{},
	})

	decimal0_5 := decimal.New(5, -1)
	decimal1 := decimal.New(1, 0)
	decimal2 := decimal.New(2, 0)
	decimal4 := decimal.New(4, 0)
	decimal4_1 := decimal.New(41, -1)
	decimal4_2 := decimal.New(42, -1)
	decimal5 := decimal.New(5, 0)
	decimal10 := decimal.New(10, 0)
	decimal15 := decimal.New(15, 0)
	decimal20 := decimal.New(20, 0)

	cases := []struct {
		desc             string
		position         *Position[*statlisting]
		side             Side
		lastQty          decimal.Decimal
		lastPx           decimal.Decimal
		expectedAvgPx    decimal.Decimal
		expectedQuantity decimal.Decimal
		expectedRealised decimal.Decimal
	}{
		{
			desc:             "flat, zero trade",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist},
			expectedAvgPx:    decimal.Zero,
			expectedQuantity: decimal.Zero,
		},
		{
			desc:             "flat, buy",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist},
			side:             Buy,
			lastQty:          decimal10,
			lastPx:           decimal4_2,
			expectedAvgPx:    decimal4_2,
			expectedQuantity: decimal10,
		},
		{
			desc:             "flat, sell",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist},
			side:             Sell,
			lastQty:          decimal10,
			lastPx:           decimal4_2,
			expectedAvgPx:    decimal4_2,
			expectedQuantity: decimal10.Neg(),
		},
		{
			desc:             "long, buy",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist, quantity: decimal10, avgPx: decimal4_2},
			side:             Buy,
			lastQty:          decimal10,
			lastPx:           decimal4,
			expectedAvgPx:    decimal4_1,
			expectedQuantity: decimal20,
		},
		{
			desc:             "long, sell part at profit",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist, quantity: decimal20, avgPx: decimal4_1},
			side:             Sell,
			lastQty:          decimal5,
			lastPx:           decimal4_2,
			expectedAvgPx:    decimal4_1,
			expectedQuantity: decimal15,
			expectedRealised: decimal0_5,
		},
		{
			desc:             "long, sell all at profit",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist, quantity: decimal20, avgPx: decimal4_1},
			side:             Sell,
			lastQty:          decimal20,
			lastPx:           decimal4_2,
			expectedAvgPx:    decimal.Zero,
			expectedQuantity: decimal.Zero,
			expectedRealised: decimal2,
		},
		{
			desc:             "long, sell part at loss",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist, quantity: decimal20, avgPx: decimal4_1},
			side:             Sell,
			lastQty:          decimal5,
			lastPx:           decimal4,
			expectedAvgPx:    decimal4_1,
			expectedQuantity: decimal15,
			expectedRealised: decimal0_5.Neg(),
		},
		{
			desc:             "long, sell all at loss",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist, quantity: decimal20, avgPx: decimal4_1},
			side:             Sell,
			lastQty:          decimal20,
			lastPx:           decimal4,
			expectedAvgPx:    decimal.Zero,
			expectedQuantity: decimal.Zero,
			expectedRealised: decimal2.Neg(),
		},
		{
			desc:             "short, sell",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist, quantity: decimal10.Neg(), avgPx: decimal4_2},
			side:             Sell,
			lastQty:          decimal10,
			lastPx:           decimal4,
			expectedAvgPx:    decimal4_1,
			expectedQuantity: decimal20.Neg(),
		},
		{
			desc:             "short, buy part at profit",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist, quantity: decimal20.Neg(), avgPx: decimal4_1},
			side:             Buy,
			lastQty:          decimal5,
			lastPx:           decimal4,
			expectedAvgPx:    decimal4_1,
			expectedQuantity: decimal15.Neg(),
			expectedRealised: decimal0_5,
		},
		{
			desc:             "short, buy all at profit",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist, quantity: decimal20.Neg(), avgPx: decimal4_1},
			side:             Buy,
			lastQty:          decimal20,
			lastPx:           decimal4,
			expectedAvgPx:    decimal.Zero,
			expectedQuantity: decimal.Zero,
			expectedRealised: decimal2,
		},
		{
			desc:             "short, buy part at loss",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist, quantity: decimal20.Neg(), avgPx: decimal4_1},
			side:             Buy,
			lastQty:          decimal5,
			lastPx:           decimal4_2,
			expectedAvgPx:    decimal4_1,
			expectedQuantity: decimal15.Neg(),
			expectedRealised: decimal0_5.Neg(),
		},
		{
			desc:             "short, buy all at loss",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist, quantity: decimal20.Neg(), avgPx: decimal4_1},
			side:             Buy,
			lastQty:          decimal20,
			lastPx:           decimal4_2,
			expectedAvgPx:    decimal.Zero,
			expectedQuantity: decimal.Zero,
			expectedRealised: decimal2.Neg(),
		},
		{
			desc:             "long, reverse at profit",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist, quantity: decimal10, avgPx: decimal4_1},
			side:             Sell,
			lastQty:          decimal20,
			lastPx:           decimal4_2,
			expectedAvgPx:    decimal4_2,
			expectedQuantity: decimal10.Neg(),
			expectedRealised: decimal1,
		},
		{
			desc:             "short, reverse at loss",
			position:         &Position[*statlisting]{symbol: "A", whitelist: whitelist, quantity: decimal10.Neg(), avgPx: decimal4_1},
			side:             Buy,
			lastQty:          decimal20,
			lastPx:           decimal4_2,
			expectedAvgPx:    decimal4_2,
			expectedQuantity: decimal10,
			expectedRealised: decimal1.Neg(),
		},
	}
	for _, c := range cases {
		c.position.Traded(c.side, c.lastQty, c.lastPx)
		assert.True(t, c.expectedQuantity.Equal(c.position.quantity), c.desc)
		assert.True(t, c.expectedAvgPx.Equal(c.position.avgPx), c.desc)
		assert.True(t, c.expectedRealised.Equal(c.position.realised), c.desc)
	}
}

func TestPositionContractMultiplier(t *testing.T) {

	whitelist := NewWhiteList[*Listing]()
	whitelist.Add(&Listing{Symbol: "WTI", ContractMultiplier: decimal.New(50, 0)})

	position := NewPosition("WTI", whitelist)

	position.Traded(Buy, decimalOne, decimal.New(4900, 0))
	position.Traded(Sell, decimalOne, decimal.New(4901, 0))

	memo := position.Memo()
	assert.NotNil(t, memo)
	assert.True(t, memo.Realised.Equal(decimal.New(50, 0)))

}

func TestMark(t *testing.T) {

	whitelist := NewWhiteList[*Listing]()
	whitelist.Add(&Listing{Symbol: "A", ContractMultiplier: decimalOne})

	decimal10 := decimal.New(10, 0)
	decimal20 := decimal.New(20, 0)
	decimal40 := decimal.New(40, 0)
	decimal42 := decimal.New(42, 0)
	decimal44 := decimal.New(44, 0)

	cases := []struct {
		desc               string
		position           *Position[*Listing]
		mark               decimal.Decimal
		expectedUnrealised decimal.Decimal
	}{
		{
			desc:               "flat",
			position:           &Position[*Listing]{symbol: "A", whitelist: whitelist},
			mark:               decimal42,
			expectedUnrealised: decimal.Zero,
		},
		{
			desc:               "profitable long",
			position:           &Position[*Listing]{symbol: "A", whitelist: whitelist, quantity: decimal10, avgPx: decimal40},
			mark:               decimal42,
			expectedUnrealised: decimal20,
		},
		{
			desc:               "unprofitable long",
			position:           &Position[*Listing]{symbol: "A", whitelist: whitelist, quantity: decimal10, avgPx: decimal44},
			mark:               decimal42,
			expectedUnrealised: decimal20.Neg(),
		},
		{
			desc:               "profitable short",
			position:           &Position[*Listing]{symbol: "A", whitelist: whitelist, quantity: decimal10.Neg(), avgPx: decimal44},
			mark:               decimal42,
			expectedUnrealised: decimal20,
		},
		{
			desc:               "unprofitable short",
			position:           &Position[*Listing]{symbol: "A", whitelist: whitelist, quantity: decimal10.Neg(), avgPx: decimal40},
			mark:               decimal42,
			expectedUnrealised: decimal20.Neg(),
		},
	}
	for _, c := range cases {
		_, pnl := c.position.Mark(c.mark)
		assert.True(t, c.expectedUnrealised.Equal(pnl), fmt.Sprintf("%s %s", c.desc, pnl.String()))
	}
}

func TestPositionChannel(t *testing.T) {

	decimal42 := decimal.New(42, 0)
	decimal100 := decimal.New(100, 0)

	whitelist := NewWhiteList[*Listing]()
	whitelist.Add(&Listing{Symbol: "A", ContractMultiplier: decimalOne})

	c := make(chan *PositionMemo, 2)

	position := NewPosition("A", whitelist, WithPositionChannel[*Listing](c))

	position.Traded(Buy, decimal100, decimal42)
	assert.Equal(t, 1, len(c))
	m := <-c
	assert.True(t, m.Quantity.Equal(decimal100))
	assert.True(t, m.AvgPx.Equal(decimal42))

}
