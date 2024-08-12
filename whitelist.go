package mkt

import (
	"github.com/gbkr-com/utl"
)

// A WhiteList has one or more T which can be traded.
type WhiteList[T AnyListing] struct {
	listings map[string]T          // Default map
	cache    *utl.Cache[string, T] // Optional cache
}

// WhiteListOption is any option that be applied when the [*WhiteList] is manufactured.
type WhiteListOption[T AnyListing] func(*WhiteList[T])

// WithWhiteListCacheOption is an option to use a [utl.Cache] instead of a simple
// map.
func WithWhiteListCacheOption[T AnyListing](cache *utl.Cache[string, T]) WhiteListOption[T] {
	return func(whitelist *WhiteList[T]) {
		whitelist.cache = cache
	}
}

// NewWhiteList returns a new [*WhiteList] ready to use.
func NewWhiteList[T AnyListing](options ...WhiteListOption[T]) *WhiteList[T] {
	whitelist := &WhiteList[T]{}
	for _, option := range options {
		option(whitelist)
	}
	if whitelist.cache == nil {
		whitelist.listings = make(map[string]T)
	}
	return whitelist
}

// Lookup returns the T for the given symbol, or if not present, the zero value
// and false.
func (x *WhiteList[T]) Lookup(symbol string) (T, bool) {
	var empty T
	if x.cache != nil {
		result, ok := x.cache.Get(symbol)
		if !ok {
			return empty, false
		}
		return result, true
	}
	result, ok := x.listings[symbol]
	if !ok {
		return empty, false
	}
	return result, true
}

// Add the listing to this white list. If the cache option was specified then
// this is a no-op - the cache governs what is in the whitelist.
func (x *WhiteList[T]) Add(listing T) {
	if x.cache != nil {
		return
	}
	x.listings[listing.Definition().Symbol] = listing

}
