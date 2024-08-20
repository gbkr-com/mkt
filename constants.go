package mkt

import "github.com/shopspring/decimal"

// FIX timestamp formats. Timestamps are always UTC.
const (
	FIXUTCMillis = "20060102-15:04:05.000"    // FIXUTCMillis is the conventional timestamp format.
	FIXUTCMicros = "20060102-15:04:05.000000" // FIXUTCMicros is used by Binance SPOT.
)

// DecimalOne is the number 1 as a decimal.
var DecimalOne = decimal.New(1, 0)
