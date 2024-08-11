package mkt

import "github.com/gbkr-com/utl"

// QuotePipeline carries quotes from the publisher to a consumer.
type QuotePipeline struct {
	pool  *utl.Pool[*Quote]
	queue *utl.ConflatingQueue[string, *Quote]
}

// NewQuotePipeline returns a pipeline ready to use. The argument specifies the
// size of the pool for [*Quote].
func NewQuotePipeline(n int) *QuotePipeline {
	pool := utl.NewPool(n, ZeroQuote)
	queue := utl.NewConflatingQueue(
		func(quote *Quote) string { return quote.Symbol },
		utl.WithPoolOption[string, *Quote](pool),
	)
	return &QuotePipeline{
		pool:  pool,
		queue: queue,
	}
}

// Get a zeroed, reusable [*Quote].
func (x *QuotePipeline) Get() *Quote {
	return x.pool.Get()
}

// Push a quote to the pipeline.
func (x *QuotePipeline) Push(quote *Quote) {
	x.queue.Push(quote)
}

// C returns the notification channel for the consumer.
func (x *QuotePipeline) C() chan struct{} {
	return x.queue.C()
}

// Pop a quote from the pipeline.
func (x *QuotePipeline) Pop() *Quote {
	return x.queue.Pop()
}

// Recycle the [*Quote].
func (x *QuotePipeline) Recycle(quote *Quote) {
	x.pool.Recycle(quote)
}
