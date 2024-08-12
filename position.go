package mkt

import "github.com/shopspring/decimal"

// A Position as a result of one or more trades in a [Listing].
type Position[T AnyListing] struct {
	symbol    string
	whitelist *WhiteList[T]
	quantity  decimal.Decimal    // Long is positive, short is negative.
	avgPx     decimal.Decimal    // Average price of building the position.
	realised  decimal.Decimal    // Realised profit/loss.
	c         chan *PositionMemo // Optional channel.
}

// PositionMemo is the key information for each position.
type PositionMemo struct {
	Symbol   string          `json:"symbol"`
	Quantity decimal.Decimal `json:"quantity"`
	AvgPx    decimal.Decimal `json:"avgPx"`
	Realised decimal.Decimal `json:"realised"`
}

// PositionOption is any option which can be applied when constructing the position.
type PositionOption[T AnyListing] func(*Position[T])

// WithPositionChannel writes a [*PositionMemo] to the channel at every trade.
func WithPositionChannel[T AnyListing](c chan *PositionMemo) PositionOption[T] {
	return func(x *Position[T]) {
		x.c = c
	}
}

// NewPosition returns a flat position for the given symbol.
func NewPosition[T AnyListing](symbol string, whitelist *WhiteList[T], options ...PositionOption[T]) *Position[T] {
	position := &Position[T]{symbol: symbol, whitelist: whitelist}
	for _, option := range options {
		option(position)
	}
	return position
}

// Memo returns a [*PositionMemo] for the current position.
func (x *Position[T]) Memo() *PositionMemo {
	return &PositionMemo{
		Symbol:   x.symbol,
		Quantity: x.quantity,
		AvgPx:    x.avgPx,
		Realised: x.realised,
	}
}

// Traded adjusts this position for the given trade.
func (x *Position[T]) Traded(side Side, lastQty decimal.Decimal, lastPx decimal.Decimal) {

	if lastQty.IsZero() {
		return
	}

	defer func() {
		if x.c == nil {
			return
		}
		x.c <- x.Memo()
	}()

	//
	// Adjust the sign for sales.
	//
	if side == Sell {
		lastQty = lastQty.Neg()
	}
	//
	// Flat?
	//
	if x.quantity.IsZero() {
		x.quantity = lastQty
		x.avgPx = lastPx
		return
	}

	precision, multiplier := x.fromListing()

	//
	// Increasing the position?
	//
	if x.quantity.Sign() == lastQty.Sign() {
		x.quantity, x.avgPx = CumQtyAvgPx(x.quantity, x.avgPx, lastQty, lastPx, precision)
		return
	}
	//
	// Reducing the position creates realised profit/loss ...
	//
	profit := x.unitProfit(lastPx, multiplier)
	//
	// ... from either flattening the position ...
	//
	if lastQty.Abs().Equal(x.quantity.Abs()) {
		x.realised = x.realised.Add(x.quantity.Abs().Mul(profit))
		x.quantity = decimal.Zero
		x.avgPx = decimal.Zero
		return
	}
	//
	// ... reversing the position ...
	//
	if lastQty.Abs().GreaterThan(x.quantity.Abs()) {
		x.realised = x.realised.Add(x.quantity.Abs().Mul(profit))
		x.quantity = x.quantity.Add(lastQty)
		x.avgPx = lastPx
		return
	}
	//
	// ... or decreasing the position.
	//
	x.realised = x.realised.Add(lastQty.Abs().Mul(profit))
	x.quantity = x.quantity.Add(lastQty)

}

// Cash adds cash to the realised profit/loss, representing a cash only movement
// such as a dividend. The cash may be negative.
func (x *Position[T]) Cash(cash decimal.Decimal) {
	x.realised = x.realised.Add(cash)
}

// Reset the realised profit/loss. Typically this would be done at the end of
// an accounting period.
func (x *Position[T]) Reset() {
	x.realised = decimal.Zero
}

// Mark returns the valuation of the position at the given price and the
// unrealised profit/loss at that price.
func (x *Position[T]) Mark(price decimal.Decimal) (valuation, unrealised decimal.Decimal) {

	_, multiplier := x.fromListing()

	valuation = x.quantity.Mul(price).Mul(multiplier)
	unrealised = x.quantity.Abs().Mul(x.unitProfit(price, multiplier))

	return
}

func (x *Position[T]) fromListing() (precision int32, contractMultiplier decimal.Decimal) {

	precision = 8
	contractMultiplier = decimalOne

	listing, ok := x.whitelist.Lookup(x.symbol)
	if ok {
		def := listing.Definition()
		precision = Precision(def.TickIncrement) + 1
		contractMultiplier = def.ContractMultiplier
	}

	return
}

func (x *Position[T]) unitProfit(price, multiplier decimal.Decimal) decimal.Decimal {

	profit := decimal.Zero

	if x.quantity.IsZero() {
		return profit
	}

	if x.quantity.IsPositive() {
		profit = price.Sub(x.avgPx)
	} else {
		profit = x.avgPx.Sub(price)
	}

	return profit.Mul(multiplier)

}
