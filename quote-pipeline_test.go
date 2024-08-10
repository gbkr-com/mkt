package mkt

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestQuotePipeline(t *testing.T) {

	pipeline := NewQuotePipeline(8)

	quote := pipeline.Get()
	assert.NotNil(t, quote)

	quote.Symbol = "A"
	quote.BidPx = decimal.New(42, 0)
	quote.BidSize = decimal.New(100, 0)
	quote.AskPx = decimal.New(43, 0)
	quote.AskSize = decimal.New(100, 0)

	pipeline.Publish(quote)

	quote = pipeline.Get()

	quote.Symbol = "A"
	quote.BidPx = decimal.New(43, 0)
	quote.BidSize = decimal.New(100, 0)
	quote.AskPx = decimal.New(44, 0)
	quote.AskSize = decimal.New(100, 0)

	pipeline.Publish(quote)

	// -------------------------------------------------------------------------

	rcv := pipeline.Receive()
	assert.NotNil(t, rcv)
	assert.True(t, rcv.BidPx.Equal(decimal.New(43, 0)))

	pipeline.Recycle(rcv)

}
