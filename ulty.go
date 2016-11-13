// Package ulty has utility functions for liby library
package ulty

import (
	"context"
	"net/http"
)

// GetContext will get the context from the spesefic request
func GetContext(r *http.Request, key string) interface{} {
	ctx := r.Context()
	val := ctx.Value(key)
	return val
}

// NewContext creates and return the request with context
func NewContext(r *http.Request, key string, value interface{}) *http.Request {
	ctx := r.Context()
	ctx = context.WithValue(ctx, key, value)
	r = r.WithContext(ctx)
	return r
}
