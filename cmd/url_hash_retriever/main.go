package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/Miladkhodabandehloo/work_pool_channlege/tools/workerpool"
)

func main() {
	// first element is ignored because it is related to the executable file
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	ctx := context.Background()
	workersNumber, err := strconv.Atoi(args[0])
	var pool *workerpool.WorkerPool

	if err != nil {
		pool = workerpool.NewWorkerPool()
	} else {
		pool = workerpool.NewWorkerPool(workerpool.WithWorkers(workersNumber))
		args = args[1:]
	}

	hashes := pool.Retrieve(ctx, args...)
	for _, hash := range hashes {
		fmt.Println(hash)
	}
}
