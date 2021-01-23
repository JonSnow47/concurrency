// Copyright 2021 JonSnow47, Inc. All rights reserved.

// concurrency function example.

package main

import (
	"context"
	"time"

	"github.com/JonSnow47/concurrency"
)

func main() {
	newWorkerFunc := func() concurrency.Worker {
		return &sleeper{}
	}

	engine := concurrency.New(100, newWorkerFunc)
	engine.Run()

	// this part will sleep 100 * time.Second without concurrency engine,
	// but just sleep about 1 second with 100 workers concurrency engine.
	for i := 0; i < 100; i++ {
		ctx := context.Background()
		engine.Submit(ctx)
	}

	time.Sleep(time.Second)
	engine.Close()
}

// sleep is the function which we want concurrent process.
func sleep() {
	time.Sleep(time.Second)
}

type sleeper struct{}

func (s *sleeper) Do(ctx context.Context) {
	sleep()
}
