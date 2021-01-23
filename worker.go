// Copyright 2021 JonSnow47, Inc. All rights reserved.

package concurrency

import "context"

// Worker represents who can handle context.
type Worker interface {
	Do(ctx context.Context)
}
