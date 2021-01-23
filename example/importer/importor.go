// Copyright 2021 JonSnow47, Inc. All rights reserved.

package main

import (
	"context"
	"fmt"
	"sync/atomic"
)

type Importer struct {
	// what ever in need

	// success
	total   int64
	success int64
	failed  int64
}

func NewImporter() *Importer {
	return &Importer{}
}

func (i *Importer) Do(ctx context.Context) {
	importerCtx := GetContext(ctx)

	// handle the ctx
	fmt.Printf("importer context: %v\n", importerCtx)

	// maybe your process will generate error(s), deal it by your self.
	var err error
	_ = err // handle this error by your self
	if err != nil {
		i.Fail()
		return
	}

	i.Done()
}

func (i *Importer) Add(n int64) {
	atomic.AddInt64(&i.total, n)
}

func (i *Importer) Done() {
	atomic.AddInt64(&i.success, 1)
}

func (i *Importer) Fail() {
	atomic.AddInt64(&i.failed, 1)
}

func (i *Importer) AllDone() bool {
	return i.failed+i.success == i.total
}
