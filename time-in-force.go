package mkt

import (
	"encoding/json"
	"sort"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
)

// The TimeInForce of an order, FIX field 59.
type TimeInForce int64

// Recognised TimeInForce values.
const (
	GTC TimeInForce = 1
	IOC TimeInForce = 3
)

// Equivalent values in QuickFIX.
var (
	fixGTC field.TimeInForceField
	fixIOC field.TimeInForceField
)

func init() {
	fixGTC = field.NewTimeInForce(enum.TimeInForce_GOOD_TILL_CANCEL)
	fixIOC = field.NewTimeInForce(enum.TimeInForce_IMMEDIATE_OR_CANCEL)
}

func (x TimeInForce) String() string {
	switch x {
	case GTC:
		return "GTC"
	case IOC:
		return "IOC"
	default:
		return ""
	}
}

// TimeInForceFromString returns a recognised [TimeInForce] or zero.
func TimeInForceFromString(s string) TimeInForce {
	switch s {
	case "GTC":
		return GTC
	case "IOC":
		return IOC
	default:
		return 0
	}
}

// MarshalJSON implements json.Marshaler.
func (x TimeInForce) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}

// UnmarshalJSON implements json.Unmarshaler.
func (x *TimeInForce) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*x = TimeInForceFromString(s)
	return nil
}

// AsQuickFIX returns this side as a QuickFIX field.
func (x TimeInForce) AsQuickFIX() field.TimeInForceField {
	switch x {
	case GTC:
		return fixGTC
	case IOC:
		return fixIOC
	default:
		return field.NewTimeInForce(enum.TimeInForce_DAY)
	}
}

// HavingTimeInForce is the interface required for [SortImmediateFirst].
type HavingTimeInForce interface {
	TimeInForce() TimeInForce
}

// SortImmediateFirst sorts such that IOC items will be first in the
// slice.
func SortImmediateFirst[T HavingTimeInForce](items []T) {
	sort.Slice(
		items,
		func(i, j int) bool {
			return items[i].TimeInForce() > items[j].TimeInForce()
		},
	)
}
