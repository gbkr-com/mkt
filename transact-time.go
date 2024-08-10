package mkt

import (
	"sort"
	"time"
)

// HavingTransactTime is the interface required for [SortRecentFirst].
type HavingTransactTime interface {
	TransactTime() time.Time
}

// SortRecentFirst sorts the items in ascending time priority, with the most
// recent TransactTime first.
func SortRecentFirst[T HavingTransactTime](items []T) {
	sort.Slice(
		items,
		func(i, j int) bool {
			return items[i].TransactTime().After(items[j].TransactTime())
		},
	)
}
