package mkt

import "github.com/shopspring/decimal"

// A Trade in a market.
type Trade struct {
	Symbol  string
	LastQty decimal.Decimal
	LastPx  decimal.Decimal
	Volume  decimal.Decimal
	AvgPx   decimal.Decimal
}

// Accumulate the given trade with this. The LastQty and LastPx are copied
// from the given trade and used to update total volume and average price.
func (x *Trade) Accumulate(trade *Trade, precision int32) {
	if x == nil {
		return
	}
	if trade.Symbol != x.Symbol {
		return
	}
	x.LastQty, x.LastPx = trade.LastQty, trade.LastPx
	x.Volume, x.AvgPx = CumQtyAvgPx(x.Volume, x.AvgPx, trade.LastQty, trade.LastPx, precision)
}
