package tests

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/Miladkhodabandehloo/work_pool_channlege/tests/mock"
	"github.com/Miladkhodabandehloo/work_pool_channlege/tools/workerpool"
)

// getWorkerPool returns a *workerpool.WorkerPool using mock hash retriever
func getWorkerPool() (*workerpool.WorkerPool, error) {
	mockRetriever, err := mock.NewMockHashRetriever()
	if err != nil {
		return nil, err
	}
	return workerpool.NewWorkerPool(workerpool.WithHashRetriever(mockRetriever)), err
}

// TestWorkPool tests whether work pool operates correctly or not
func TestWorkPool(t *testing.T) {
	f, err := os.Open("data/hashes.json")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	var hashes map[string]string
	err = json.NewDecoder(f).Decode(&hashes)
	if err != nil {
		t.Error(err)
	}
	var urls []string
	for url, _ := range hashes {
		urls = append(urls, url)
	}
	urls = append(urls, "http://test.test")

	pool, err := getWorkerPool()
	if err != nil {
		t.Error(f)
	}
	result := pool.Retrieve(context.Background(), urls...)
	if len(result) != len(urls) {
		t.Errorf("expected %d items in result, got %d", len(urls), len(result))
		return
	}
	for _, v := range result {
		hash, ok := hashes[v.Address]
		if ok && hash != v.Hash {
			t.Errorf("expected %s for %s, got %s", hash, v.Address, v.Hash)
		}
		if !ok && v.Err == nil {
			t.Errorf("expected error for %s, go nothing", v.Address)
		}

	}
}
