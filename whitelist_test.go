package mkt

import (
	"testing"
	"time"

	"github.com/gbkr-com/utl"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestWhiteListWithCacheOption(t *testing.T) {

	cache := utl.NewCache(
		time.Hour,
		func(symbol string) (*Listing, bool) {
			if symbol == "B" {
				return nil, false
			}
			listing := &Listing{
				Symbol:        "A",
				TickIncrement: decimal.New(1, 0),
				RoundLot:      decimal.New(1, 0),
				MinTradeVol:   decimal.New(1, 0),
			}
			return listing, true
		},
	)

	whitelist := NewWhiteList(WithWhiteListCacheOption(cache))
	assert.NotNil(t, whitelist)

	listing, ok := whitelist.Lookup("A")
	assert.True(t, ok)
	assert.NotNil(t, listing)
	assert.Equal(t, "A", listing.Symbol)
	assert.True(t, listing.TickIncrement.Equal(decimal.New(1, 0)))

	listing, ok = whitelist.Lookup("B")
	assert.False(t, ok)
	assert.Nil(t, listing)

}
