//Package workerpool provides functionalities to have a worker pool with specified number of workers that use an
//instance of httphash.HashRetriever to retrieve hashes for a slice of urls in parallel
package workerpool

import (
	"context"
	"fmt"

	"github.com/Miladkhodabandehloo/work_pool_channlege/tools/httphash"
	"github.com/Miladkhodabandehloo/work_pool_channlege/tools/httphash/bodyhash"
)

var (
	_defaultWorkersNumber = 10
	_defaultHashRetriever = bodyhash.NewRetriever()
)

// WorkerPool provides functionality for retrieving hashes for a slice of urls in parallel with specified number of
// workers using and specified HashRetriever
type WorkerPool struct {
	hashRetriever httphash.HashRetriever
	workersNumber int
}

// HashResult is the type that comprises result of digest calculation for an url
type HashResult struct {
	Address string
	Hash    string
	Err     error
}

// String returns string representation of a HashResult
func (hr HashResult) String() string {
	if hr.Err != nil {
		return fmt.Sprintf("%s: %s", hr.Address, hr.Err.Error())
	} else {
		return fmt.Sprintf("%s: %s", hr.Address, hr.Hash)
	}
}

// work is representation of a worker that listens to a channel or urls and publishes result of retrieving its hash on
// a channel for results
func (wp *WorkerPool) work(ctx context.Context, urlCh <-chan string, resultCh chan<- HashResult) {
	for url := range urlCh {
		hash, err := wp.hashRetriever.Retrieve(ctx, url)
		resultCh <- HashResult{
			Address: url,
			Hash:    hash,
			Err:     err}
	}
}

// Retrieve runs workers for retrieving hash digest of urls
func (wp *WorkerPool) Retrieve(ctx context.Context, urls ...string) []HashResult {
	urlCh := publishURLs(urls...)
	resultCh := make(chan HashResult)
	defer close(resultCh)
	for i := 0; i < wp.workersNumber; i++ {
		go wp.work(ctx, urlCh, resultCh)
	}
	var hashes []HashResult
	for i := 0; i < len(urls); i++ {
		hashes = append(hashes, <-resultCh)
	}
	return hashes
}

// publishURLs publishes a slice of urls on a channel
func publishURLs(urls ...string) <-chan string {
	urlCh := make(chan string)
	go func() {
		for _, url := range urls {
			urlCh <- url
		}
		close(urlCh)
	}()
	return urlCh
}

// Option is a function that takes a worker pool and changes its attributes
type Option func(pool *WorkerPool)

// WithHashRetriever returns an Option for setting WorkerPool.hashRetriever
func WithHashRetriever(retriever httphash.HashRetriever) Option {
	return func(pool *WorkerPool) {
		pool.hashRetriever = retriever
	}
}

// WithWorkers returns an option for setting WorkerPool.workersNumber
func WithWorkers(workersNumber int) Option {
	return func(pool *WorkerPool) {
		pool.workersNumber = workersNumber
	}
}

// NewWorkerPool returns a *WorkerPool specified by options
func NewWorkerPool(options ...Option) *WorkerPool {
	pool := &WorkerPool{
		hashRetriever: _defaultHashRetriever,
		workersNumber: _defaultWorkersNumber,
	}
	for _, o := range options {
		o(pool)
	}
	return pool
}
