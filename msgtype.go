package mkt

import (
	"encoding/json"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
)

// The MsgType for an order, FIX field 35.
type MsgType int64

// Recognised MsgType values.
const (
	OrderNew MsgType = iota + 13
	_
	OrderCancel
	OrderReplace
)

// Equivalent values in QuickFIX.
var (
	fixOrderNew     field.MsgTypeField
	fixOrderCancel  field.MsgTypeField
	fixOrderReplace field.MsgTypeField
)

func init() {
	fixOrderNew = field.NewMsgType(enum.MsgType_ORDER_SINGLE)
	fixOrderCancel = field.NewMsgType(enum.MsgType_ORDER_CANCEL_REQUEST)
	fixOrderReplace = field.NewMsgType(enum.MsgType_ORDER_CANCEL_REPLACE_REQUEST)
}

// String returns a short mnemonic for the [MsgType], avoiding the FIX values.
func (x MsgType) String() string {
	switch x {
	case OrderNew:
		return "NEW"
	case OrderCancel:
		return "CXL"
	case OrderReplace:
		return "RPL"
	default:
		return ""
	}
}

// MsgTypeFromString returns a recognised [MsgType] or zero.
func MsgTypeFromString(s string) MsgType {
	switch s {
	case "NEW":
		return OrderNew
	case "CXL":
		return OrderCancel
	case "RPL":
		return OrderReplace
	default:
		return 0
	}
}

// MarshalJSON implements json.Marshaler.
func (x MsgType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}

// UnmarshalJSON implements json.Unmarshaler.
func (x *MsgType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	*x = MsgTypeFromString(s)
	return nil
}

// AsQuickFIX returns this [MsgType] as a QuickFIX field.
func (x MsgType) AsQuickFIX() field.MsgTypeField {
	switch x {
	case OrderNew:
		return fixOrderNew
	case OrderCancel:
		return fixOrderCancel
	case OrderReplace:
		return fixOrderReplace
	default:
		return field.NewMsgType(enum.MsgType_USERNOTIFICATION)
	}
}
