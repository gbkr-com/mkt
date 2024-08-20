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

func TestTransactTimeFormat(t *testing.T) {

	now := time.Date(2024, 8, 20, 8, 4, 32, 397561000, time.UTC)

	assert.Equal(t, "20240820-08:04:32.397", now.Format(FIXUTCMillis))
	assert.Equal(t, "20240820-08:04:32.397561", now.Format(FIXUTCMicros))

}
