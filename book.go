package mkt

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// Book is a set of [*Position]. A book may only hold a position in a [Listing]
// that is included in the associated [WhiteList].
type Book[T AnyListing] struct {
	name      string
	whitelist *WhiteList[T]
	positions map[string]*Position[T]
	c         chan *PositionMemo
}

// BookOption is any option that can be applied when constructing the book.
type BookOption[T AnyListing] func(*Book[T])

// WithBookChannel writes a [*PositionMemo] to the channel after any trade in
// the book.
func WithBookChannel[T AnyListing](c chan *PositionMemo) BookOption[T] {
	return func(book *Book[T]) {
		book.c = c
	}
}

// NewBook returns a [*Book] ready to use.
func NewBook[T AnyListing](name string, whitelist *WhiteList[T], options ...BookOption[T]) *Book[T] {
	book := &Book[T]{name: name, whitelist: whitelist, positions: map[string]*Position[T]{}}
	for _, option := range options {
		option(book)
	}
	return book
}

// Name returns the name of the book.
func (x *Book[T]) Name() string { return x.name }

// ForEachPosition visits every position in the book.
func (x *Book[T]) ForEachPosition(visitor func(*Position[T])) {
	for _, position := range x.positions {
		visitor(position)
	}
}

// Traded applies the trade to the book. If the symbol is not recognised this
// function will return an error.
func (x *Book[T]) Traded(symbol string, side Side, lastQty decimal.Decimal, lastPx decimal.Decimal) error {

	if lastQty.IsZero() {
		return nil
	}

	position := x.positions[symbol]
	if position == nil {
		var err error
		position, err = x.makePosition(symbol)
		if err != nil {
			return err
		}
	}

	position.Traded(side, lastQty, lastPx)
	return nil

}

func (x *Book[T]) makePosition(symbol string) (*Position[T], error) {

	_, ok := x.whitelist.Lookup(symbol)
	if !ok {
		return nil, fmt.Errorf("inv.Book: %s is not whitelisted", symbol)
	}

	options := []PositionOption[T]{}
	if x.c != nil {
		options = append(options, WithPositionChannel[T](x.c))
	}

	position := NewPosition(symbol, x.whitelist, options...)
	x.positions[symbol] = position

	return position, nil
}
