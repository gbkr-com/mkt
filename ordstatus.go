package mkt

import (
	"encoding/json"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
)

// The OrdStatus is the execution status of an order, FIX field 39
type OrdStatus int64

// Recognised OrdStatus values, a subset from FIX 4.4. The zero value is
// reserved for representing 'no value'.
const (
	OrdStatusNew OrdStatus = iota + 1
	OrdStatusPartiallyFilled
	OrdStatusFilled
	OrdStatusCanceled
	OrdStatusPendingCancel
	OrdStatusRejected
	OrdStatusPendingNew
	OrdStatusExpired
	OrdStatusPendingReplace
)

var (
	ordStatusToString map[OrdStatus]string
	stringToOrdStatus map[string]OrdStatus
)

func init() {
	ordStatusToString = map[OrdStatus]string{
		OrdStatusNew:             "NEW",
		OrdStatusPartiallyFilled: "FILL",
		OrdStatusFilled:          "DONE",
		OrdStatusCanceled:        "CXLD",
		OrdStatusPendingCancel:   "PCXL",
		OrdStatusRejected:        "REJD",
		OrdStatusPendingNew:      "PNEW",
		OrdStatusExpired:         "EXPD",
		OrdStatusPendingReplace:  "PRPL",
	}
	stringToOrdStatus = map[string]OrdStatus{
		"NEW":  OrdStatusNew,
		"FILL": OrdStatusPartiallyFilled,
		"DONE": OrdStatusFilled,
		"CXLD": OrdStatusCanceled,
		"PCXL": OrdStatusPendingCancel,
		"REJD": OrdStatusRejected,
		"PNEW": OrdStatusPendingNew,
		"EXPD": OrdStatusExpired,
		"PRPL": OrdStatusPendingReplace,
	}
}

// String returns a mnemonic of the [OrdStatus].
func (x OrdStatus) String() string {
	return ordStatusToString[x]
}

// OrdStatusFromString returns a recognised [OrdStatus] or zero.
func OrdStatusFromString(s string) OrdStatus {
	return stringToOrdStatus[s]
}

// MarshalJSON implements [json.Marshaler].
func (x OrdStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}

// UnmarshalJSON implements [json.Unmarshaler].
func (x *OrdStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*x = OrdStatusFromString(s)
	return nil
}

// AsQuickFIX returns the [OrdStatus]  as a QuickFIX field. If the value is not
// one of those recognised, this function returns a valid value that will
// likely be rejected by the counterparty, rather than panicking.
func (x OrdStatus) AsQuickFIX() field.OrdStatusField {
	switch x {
	case OrdStatusNew:
		return field.NewOrdStatus(enum.OrdStatus_NEW)
	case OrdStatusPartiallyFilled:
		return field.NewOrdStatus(enum.OrdStatus_PARTIALLY_FILLED)
	case OrdStatusFilled:
		return field.NewOrdStatus(enum.OrdStatus_FILLED)
	case OrdStatusCanceled:
		return field.NewOrdStatus(enum.OrdStatus_CANCELED)
	case OrdStatusPendingCancel:
		return field.NewOrdStatus(enum.OrdStatus_PENDING_CANCEL)
	case OrdStatusRejected:
		return field.NewOrdStatus(enum.OrdStatus_REJECTED)
	case OrdStatusPendingNew:
		return field.NewOrdStatus(enum.OrdStatus_PENDING_NEW)
	case OrdStatusExpired:
		return field.NewOrdStatus(enum.OrdStatus_EXPIRED)
	case OrdStatusPendingReplace:
		return field.NewOrdStatus(enum.OrdStatus_PENDING_REPLACE)
	default:
		return field.NewOrdStatus(enum.OrdStatus_SUSPENDED)
	}
}

// OrdStatusFromFIX returns the equivalent [OrdStatus] from the QuickFIX field,
// or zero if there is no equivalence.
func OrdStatusFromFIX(ordStatus field.OrdStatusField) OrdStatus {
	switch ordStatus.Value() {
	case enum.OrdStatus_CANCELED:
		return OrdStatusCanceled
	case enum.OrdStatus_EXPIRED:
		return OrdStatusExpired
	case enum.OrdStatus_FILLED:
		return OrdStatusFilled
	case enum.OrdStatus_NEW:
		return OrdStatusNew
	case enum.OrdStatus_PARTIALLY_FILLED:
		return OrdStatusPartiallyFilled
	case enum.OrdStatus_PENDING_CANCEL:
		return OrdStatusPendingCancel
	case enum.OrdStatus_PENDING_NEW:
		return OrdStatusPendingNew
	case enum.OrdStatus_PENDING_REPLACE:
		return OrdStatusPendingReplace
	case enum.OrdStatus_REJECTED:
		return OrdStatusRejected
		//
		// Unsupported cases.
		//
	case enum.OrdStatus_ACCEPTED_FOR_BIDDING,
		enum.OrdStatus_CALCULATED,
		enum.OrdStatus_DONE_FOR_DAY,
		enum.OrdStatus_REPLACED,
		enum.OrdStatus_STOPPED,
		enum.OrdStatus_SUSPENDED:
		return 0
	}
	return 0
}
