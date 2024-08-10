package mkt

import (
	"time"

	"github.com/gbkr-com/utl"
	"github.com/shopspring/decimal"
)

// A Listing which can be traded with counterparties. These are the essential
// details required for trading.
//
// Listing is principally a data object and stateless, so its fields are
// exported for convenience.
type Listing struct {
	Symbol        string          // FIX field 55
	TickIncrement decimal.Decimal // FIX field 1208
	RoundLot      decimal.Decimal // FIX field 561
	MinTradeVol   decimal.Decimal // FIX field 562
}

// Definition returns the [*Listing].
func (x *Listing) Definition() *Listing { return x }

// Listable defines all types which can embed ('extend') [Listing].
type Listable interface {
	Definition() *Listing
}

// A WhiteList has one or more [Listable] which can be traded.
type WhiteList[T Listable] interface {
	Lookup(symbol string) T
}

// WhiteListWithReload is [WhiteList] which will reload T from another
// source.
type WhiteListWithReload[T Listable] struct {
	cache *utl.Cache[string, T]
}

// NewWhiteListWithReload returns a new [WhiteList] with the given time-to-live
// for each entry and the given replace function.
func NewWhiteListWithReload[T Listable](ttl time.Duration, replace func(string) (T, bool)) WhiteList[T] {
	return &WhiteListWithReload[T]{
		cache: utl.NewCache(ttl, replace),
	}
}

// Lookup a symbol.
func (x *WhiteListWithReload[T]) Lookup(symbol string) T {
	result, ok := x.cache.Get(symbol)
	if !ok {
		var empty T
		return empty
	}
	return result
}
