package mkt

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestPrecision(t *testing.T) {

	cases := []struct {
		number   decimal.Decimal
		expected int32
	}{
		{
			number:   decimal.New(10, 0),
			expected: 0,
		},
		{
			number:   decimal.New(1, 0),
			expected: 0,
		},
		{
			number:   decimal.New(1, -1),
			expected: 1,
		},
		{
			number:   decimal.New(1, -2),
			expected: 2,
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, Precision(c.number))
	}

}

func TestUnits(t *testing.T) {

	cases := []struct {
		desc     string
		original decimal.Decimal
		incr     decimal.Decimal
		min      decimal.Decimal
		expected decimal.Decimal
	}{
		{
			desc:     "round down to 44",
			original: decimal.New(444, -1),
			incr:     decimal.New(5, -1),
			min:      decimal.New(5, -1),
			expected: decimal.New(44, 0),
		},
		{
			desc:     "round down to 44.5",
			original: decimal.New(446, -1),
			incr:     decimal.New(5, -1),
			min:      decimal.New(5, -1),
			expected: decimal.New(445, -1),
		},
		{
			desc:     "below min",
			original: decimal.New(4, -1),
			incr:     decimal.New(5, -1),
			min:      decimal.New(5, -1),
			expected: decimal.Zero,
		},
		{
			desc:     "to min",
			original: decimal.New(6, -1),
			incr:     decimal.New(5, -1),
			min:      decimal.New(5, -1),
			expected: decimal.New(5, -1),
		},
	}

	for _, c := range cases {
		result := Units(c.original, c.incr, c.min)
		assert.True(t, result.Equal(c.expected), c.desc)
	}

}

func TestCumQtyAvgPx(t *testing.T) {

	cumQty := decimal.New(100, 0)
	avgPx := decimal.New(42, 0)

	cumQty, avgPx = CumQtyAvgPx(cumQty, avgPx, decimal.New(10, 0), decimal.New(45, 0), 3)
	assert.True(t, cumQty.Equal(decimal.New(110, 0)))
	assert.True(t, avgPx.Equal(decimal.New(42273, -3)))

}
