package mkt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookWhitelist(t *testing.T) {

	whitelist := NewWhiteList[*Listing]()

	book := NewBook("HEDGE", whitelist)

	err := book.Traded("A", Buy, DecimalOne, DecimalOne)
	assert.NotNil(t, err)

}
