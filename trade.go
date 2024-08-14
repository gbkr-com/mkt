package mkt

import "github.com/shopspring/decimal"

// A Trade in a market.
//
// Trade is principally a data object and very very little state, so its fields
// are exported for convenience.
type Trade struct {
	Symbol      string          // FIX field 55
	LastQty     decimal.Decimal // FIX field 32
	LastPx      decimal.Decimal // FIX field 31
	TradeVolume decimal.Decimal // FIX field 1020
	AvgPx       decimal.Decimal // FIX field 6
}

// Aggregate the given trade with this. The LastQty and LastPx are copied
// from the given trade and used to update total volume and average price.
func (x *Trade) Aggregate(trade *Trade, precision int32) {

	if x == nil || trade == nil {
		return
	}
	if trade.Symbol != x.Symbol {
		return
	}

	switch {
	case x.TradeVolume.IsZero() && trade.TradeVolume.IsZero():
		x.TradeVolume, x.AvgPx = CumQtyAvgPx(x.LastQty, x.LastPx, trade.LastQty, trade.LastPx, precision)
	case trade.TradeVolume.IsZero():
		x.TradeVolume, x.AvgPx = CumQtyAvgPx(x.TradeVolume, x.AvgPx, trade.LastQty, trade.LastPx, precision)
	default:
		x.TradeVolume, x.AvgPx = CumQtyAvgPx(x.TradeVolume, x.AvgPx, trade.TradeVolume, trade.AvgPx, precision)
	}

	x.LastQty, x.LastPx = trade.LastQty, trade.LastPx

}

// TradeKey is a convenience function to use when constructing a
// [utl.ConflatingQueue].
func TradeKey(trade *Trade) string { return trade.Symbol }
