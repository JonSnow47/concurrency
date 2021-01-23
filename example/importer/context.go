// Copyright 2021 JonSnow47, Inc. All rights reserved.

package main

import "context"

type importerContext struct{}

type Context struct {
	key string
	val interface{}
}

func WithContext(parent context.Context, key string, val interface{}) context.Context {
	c := &Context{
		key: key,
		val: val,
	}
	return context.WithValue(parent, importerContext{}, c)
}

func GetContext(ctx context.Context) *Context {
	return ctx.Value(importerContext{}).(*Context)
}
