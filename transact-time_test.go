package mkt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSortTransactTime(t *testing.T) {

	now := time.Now()

	orders := []*order{
		{orderID: "A", transactTime: now.Add(-time.Second)},
		{orderID: "B", transactTime: now.Add(-time.Minute)},
		{orderID: "C", transactTime: now},
	}

	SortRecentFirst(orders)
	assert.Equal(t, "C", orders[0].orderID)

}
