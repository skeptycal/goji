package goji

import (
	"context"
	"net/http"

	// "github.com/skeptycal/goji/internal"
	"github.com/skeptycal/goji/internal"
)

type match struct {
	context.Context
	p Pattern
	h http.Handler
}

func (m match) Value(key interface{}) interface{} {
	switch key {
	case internal.Pattern:
		return m.p
	case internal.Handler:
		return m.h
	default:
		return m.Context.Value(key)
	}
}

var _ context.Context = match{}
