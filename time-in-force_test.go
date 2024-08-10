package mkt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortTimeInForce(t *testing.T) {

	orders := []*order{
		{orderID: "A", timeInForce: GTC},
		{orderID: "B", timeInForce: GTC},
		{orderID: "C", timeInForce: IOC},
	}

	SortImmediateFirst(orders)
	assert.Equal(t, "C", orders[0].orderID)

}
