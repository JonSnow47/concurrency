// Copyright 2021 JonSnow47, Inc. All rights reserved.

package concurrency

import (
	"context"
)

// NewWorkerFunc represents the method to create a practical Worker.
type NewWorkerFunc func() Worker

// Engine implement a concurrent engine.
// You can use Engine do many tasks concurrent.
type Engine struct {
	// newWorker function
	newWorker NewWorkerFunc

	// contextChan receive task from outside.
	contextChan chan context.Context

	// workerNum represent maximum worker number for concurrent working.
	workerNum int
}

func New(
	workerNum int,
	newWorkerFunc NewWorkerFunc,
) *Engine {
	engine := &Engine{
		newWorker:   newWorkerFunc,
		contextChan: make(chan context.Context),
		workerNum:   workerNum,
	}
	return engine
}

// Run is represent launch engine and listening context.
func (e *Engine) Run() {
	// TODO: make contextChan at here to make engine reusable
	for i := 0; i < e.workerNum; i++ {
		createWorker(e.newWorker(), e.contextChan)
	}
}

// Submit a context to Engine, it will be received while any Worker ready.
func (e *Engine) Submit(ctx context.Context) {
	e.contextChan <- ctx
}

// Close close Engine, and you cant reuse or relaunch Engine.
func (e *Engine) Close() {
	close(e.contextChan)
}

// createWorker create a worker to listening context.
// you may take a task in context.
func createWorker(worker Worker, contextChan <-chan context.Context) {
	go func() {
		for {
			ctx, ok := <-contextChan
			if !ok {
				break
			}

			// handle context.
			worker.Do(ctx)
		}
	}()
}
