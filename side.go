package mkt

import (
	"encoding/json"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/shopspring/decimal"
)

// The Side of an order, FIX field 54.
type Side int64

// Recognised Side values.
const (
	Buy  Side = 1
	Sell Side = 2
)

// Equivalent values in QuickFIX.
var (
	fixBuy  field.SideField
	fixSell field.SideField
)

func init() {
	fixBuy = field.NewSide(enum.Side_BUY)
	fixSell = field.NewSide(enum.Side_SELL)
}

// Opposite returns the opposite [Side].
func (x Side) Opposite() Side {
	switch x {
	case Buy:
		return Sell
	case Sell:
		return Buy
	default:
		return 0
	}
}

// Within returns true if the price is within the limit. It also returns true if
// the limit is zero.
func (x Side) Within(price, limit decimal.Decimal) bool {
	if limit.IsZero() {
		return true
	}
	switch x {
	case Buy:
		return price.LessThanOrEqual(limit)
	case Sell:
		return price.GreaterThanOrEqual(limit)
	default:
		return false
	}
}

// Improve returns the price improved by the increment. When buying, the price is
// improved (for the seller) by adding the increment; when selling the price is
// improved (for the buyer) by subtracting the increment.
func (x Side) Improve(price, increment decimal.Decimal) decimal.Decimal {
	if increment.IsZero() {
		return price
	}
	switch x {
	case Buy:
		return price.Add(increment)
	case Sell:
		return price.Sub(increment)
	default:
		return price
	}
}

func (x Side) String() string {
	switch x {
	case Buy:
		return "BUY"
	case Sell:
		return "SELL"
	default:
		return ""
	}
}

// SideFromString returns a recognised [Side] or zero.
func SideFromString(x string) Side {
	switch x {
	case "BUY":
		return Buy
	case "SELL":
		return Sell
	default:
		return 0
	}
}

// MarshalJSON implements [json.Marshaler].
func (x Side) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}

// UnmarshalJSON implements [json.Unmarshaler].
func (x *Side) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*x = SideFromString(s)
	return nil
}

// AsQuickFIX returns this side as a QuickFIX field.
func (x Side) AsQuickFIX() field.SideField {
	switch x {
	case Buy:
		return fixBuy
	case Sell:
		return fixSell
	default:
		return field.NewSide(enum.Side_UNDISCLOSED)
	}
}
