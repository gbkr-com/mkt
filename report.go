package mkt

import (
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// Report is an augmented execution report from the counterparty.
//
// [Report.OrderID] is that known by the originator of the order, whereas
// [Report.SecondaryOrderID] is that assigned by the counterparty.
//
// [Report.ExecInst] values are ignored except for 'e', meaning 'work to target
// strategy'. If that is returned to the originator then it signals that the
// originator may resume sending requests to the counterparty.
type Report struct {
	OrderID          string          `json:"orderID"`                    // FIX field 37
	Symbol           string          `json:"symbol,omitempty"`           // FIX field 55
	Side             Side            `json:"side,omitempty"`             // FIX field 54
	SecondaryOrderID string          `json:"secondaryOrderID,omitempty"` // FIX field 198
	ClOrdID          string          `json:"clOrdID,omitempty"`          // FIX field 11
	OrdStatus        OrdStatus       `json:"ordStatus,omitempty"`        // FIX field 39
	Account          string          `json:"account,omitempty"`          // FIX field 1
	TimeInForce      TimeInForce     `json:"timeInForce,omitempty"`      // FIX field 59
	LastQty          decimal.Decimal `json:"lastQty"`                    // FIX field 32
	LastPx           decimal.Decimal `json:"lastPx"`                     // FIX field 31
	TransactTime     time.Time       `json:"transactTime"`               // FIX field 60
	ExecInst         string          `json:"execInst,omitempty"`         // FIX field 18
}

// WorkToTarget returns true if the report indicates the originator may
// continue sending requests.
func (x *Report) WorkToTarget() bool {
	return strings.Contains(x.ExecInst, "e")
}
