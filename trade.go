package mkt

import "github.com/shopspring/decimal"

// A Trade in a market.
//
// [Trade] is principally a data object and has little state, so its fields
// are exported for convenience.
//
// Trades may be aggregated, for example in a [utl.ConflatingQueue]. When
// inspecting a Trade the [Trade.LastQty] and [Trade.LastPx] fields will always
// refer to the last trade that was included. However, if non zero, [Trade.TradeVolume]
// and [Trade.AvgPx] reflect all the trades aggregated in the struct: if simply
// counting volume and VWAP use those fields.
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
// [utl.ConflatingQueue] for [Trade].
func TradeKey(trade *Trade) string { return trade.Symbol }
