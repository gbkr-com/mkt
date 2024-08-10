package mkt

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestListingExtension(t *testing.T) {

	listingA := &Listing{
		Symbol:        "A",
		TickIncrement: decimal.New(1, 0),
		RoundLot:      decimal.New(1, 0),
		MinTradeVol:   decimal.New(1, 0),
	}

	def := listingA.Definition()
	assert.NotNil(t, def)

	type statlisting struct {
		Listing
		MedTradeSize decimal.Decimal
	}

	listingB := &statlisting{
		Listing: Listing{
			Symbol:        "B",
			TickIncrement: decimal.New(2, 0),
			RoundLot:      decimal.New(2, 0),
			MinTradeVol:   decimal.New(2, 0),
		},
		MedTradeSize: decimal.New(100, 0),
	}

	def = listingB.Definition()
	assert.NotNil(t, def)
	assert.Equal(t, "B", listingB.Symbol)
	assert.True(t, listingB.MedTradeSize.IsPositive())

}

func TestWhiteList(t *testing.T) {

	whitelist := NewWhiteListWithReload(
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
	assert.NotNil(t, whitelist)

	listing := whitelist.Lookup("A")
	assert.NotNil(t, listing)
	assert.Equal(t, "A", listing.Symbol)
	assert.True(t, listing.TickIncrement.Equal(decimal.New(1, 0)))

	listing = whitelist.Lookup("B")
	assert.Nil(t, listing)

}
