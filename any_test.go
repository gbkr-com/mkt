package mkt

import "time"

type order struct {
	orderID      string
	timeInForce  TimeInForce
	transactTime time.Time
}

func (x *order) TimeInForce() TimeInForce { return x.timeInForce }
func (x *order) TransactTime() time.Time  { return x.transactTime }
