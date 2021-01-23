// Copyright 2021 JonSnow47, Inc. All rights reserved.

package main

import (
	"context"
	"fmt"

	"github.com/JonSnow47/concurrency"
)

func main() {
	taskNum := 1000

	importer := NewImporter()
	importer.Add(int64(taskNum))

	newWorkerFunc := func() concurrency.Worker {
		return importer
	}

	engine := concurrency.New(100, newWorkerFunc)
	engine.Run()

	for i := 0; i < taskNum; i++ {
		ctx := context.Background()
		key := fmt.Sprintf("key_%d", i)
		val := fmt.Sprintf("val_%d", i)
		engine.Submit(WithContext(ctx, key, val))
	}

	for !importer.AllDone() {
	}
	engine.Close()
}
