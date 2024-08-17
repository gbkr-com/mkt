package mkt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderJSON(t *testing.T) {

	order := &Order{
		MsgType: OrderNew,
		OrderID: "abc",
		Side:    Buy,
		Symbol:  "XRP-USD",
	}

	b, err := json.Marshal(order)
	assert.Nil(t, err)
	assert.Equal(t, `{"MsgType":"NEW","OrderID":"abc","Side":"BUY","Symbol":"XRP-USD"}`, string(b))

	var o Order
	err = json.Unmarshal(b, &o)
	assert.Nil(t, err)
	assert.Equal(t, OrderNew, o.MsgType)
	assert.Equal(t, Buy, o.Side)

}
