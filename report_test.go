package mkt

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestReport(t *testing.T) {

	minimal := &Report{
		OrderID:          NewOrderID(),
		Symbol:           "",
		Side:             0,
		SecondaryOrderID: "",
		ClOrdID:          "",
		OrdStatus:        0,
		Account:          "",
		TimeInForce:      0,
		LastQty:          decimal.Decimal{},
		LastPx:           decimal.Decimal{},
		TransactTime:     time.Now(),
		ExecInst:         "e",
	}

	b, err := json.Marshal(minimal)
	assert.Nil(t, err)

	var report Report
	err = json.Unmarshal(b, &report)
	assert.Nil(t, err)

	assert.True(t, report.WorkToTarget())
}
