package mkt

import "github.com/shopspring/decimal"

// Precision returns the number of decimals in the number.
func Precision(number decimal.Decimal) int32 {
	e := number.Exponent()
	if e >= 0 {
		return 0
	}
	return -e
}

// Units returns x as an integer number of units greater than or equal to the
// minimum. This function is used for both prices (to be round ticks) and
// quantities (to be round lots).
func Units(x, unit, min decimal.Decimal) decimal.Decimal {
	if x.LessThanOrEqual(decimal.Zero) {
		return decimal.Zero
	}
	if unit.IsZero() {
		return x
	}
	u := x.Sub(x.Mod(unit))
	if u.LessThan(min) {
		return decimal.Zero
	}
	return u
}

// CumQtyAvgPx returns the cumulative quantity and average price, the latter
// rounded to 'n' places.
func CumQtyAvgPx(cumQty, avgPx, lastQty, lastPx decimal.Decimal, n int32) (decimal.Decimal, decimal.Decimal) {

	v := cumQty.Mul(avgPx)
	v = v.Add(lastQty.Mul(lastPx))

	cumQty = cumQty.Add(lastQty)

	avgPx = v.DivRound(cumQty, n)
	return cumQty, avgPx

}
