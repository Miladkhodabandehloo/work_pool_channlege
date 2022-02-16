// Package mock provides functionality for having a mock implementation of httphash.HashRetriever using mock data from
//the file hashes.json
package mock

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/Miladkhodabandehloo/work_pool_channlege/tools/httphash"
)

// mockHashRetriever implement httphash.HashRetriever using mock data
type mockHashRetriever struct {
	hashes map[string]string
}

// Retrieve calculates digest regarding an url
func (r *mockHashRetriever) Retrieve(_ context.Context, url string) (string, error) {
	hash, ok := r.hashes[url]
	if !ok {
		return "", errors.New("not found")
	}
	return hash, nil
}

// NewMockHashRetriever returns an implementation of httphash.HashRetriever using mockHashRetriever
func NewMockHashRetriever() (httphash.HashRetriever, error) {
	retriever := new(mockHashRetriever)
	f, err := os.Open("data/hashes.json")
	defer f.Close()
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(f).Decode(&retriever.hashes)
	return retriever, err

}
