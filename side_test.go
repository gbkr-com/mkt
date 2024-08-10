package mkt

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestSideOpposite(t *testing.T) {
	assert.Equal(t, Sell, Buy.Opposite())
	assert.Equal(t, Buy, Sell.Opposite())
}

func TestSideImprove(t *testing.T) {

	PRICE := decimal.New(42, 0)
	IMPROVEMENT := decimal.New(5, -1)

	assert.True(t, Buy.Improve(PRICE, decimal.Zero).Equal(PRICE))
	assert.True(t, Buy.Improve(PRICE, IMPROVEMENT).Equal(decimal.New(425, -1)))
	assert.True(t, Sell.Improve(PRICE, IMPROVEMENT).Equal(decimal.New(415, -1)))

}

func TestSideWithin(t *testing.T) {

	PLUS := decimal.New(425, -1)
	PRICE := decimal.New(42, 0)
	MINUS := decimal.New(415, -1)

	assert.True(t, Buy.Within(PRICE, PLUS))
	assert.True(t, Buy.Within(PRICE, PRICE))
	assert.False(t, Buy.Within(PRICE, MINUS))

	assert.True(t, Sell.Within(PRICE, MINUS))
	assert.True(t, Sell.Within(PRICE, PRICE))
	assert.False(t, Sell.Within(PRICE, PLUS))

}

func TestSideJSON(t *testing.T) {

	b, err := Buy.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, `"BUY"`, string(b))

	var side Side
	err = side.UnmarshalJSON(b)
	assert.Nil(t, err)
	assert.Equal(t, Buy, side)

	side = 0
	b, err = side.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, `""`, string(b))

	side = Buy
	err = side.UnmarshalJSON(b)
	assert.Nil(t, err)
	assert.Equal(t, Side(0), side)

}
