// Package bodyhash implements httphash.HashRetriever by calculating digest of the body of GET HTTP call response for
//an url
package bodyhash

import (
	"context"
	"crypto/md5"
	"fmt"
	"hash"
	"io"
	"net/http"

	"github.com/Miladkhodabandehloo/work_pool_channlege/tools/httphash"
)

var (
	_defaultHashFunc = md5.New
	_defaultClient   = http.DefaultClient
)

// retriever implements httphash.HashRetriever by calling HTTP GET for an url and hashing body of response
type retriever struct {
	client   *http.Client
	hashFunc func() hash.Hash
}

// Retrieve calculates digest regarding an url
func (r *retriever) Retrieve(ctx context.Context, url string) (string, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	response, err := r.client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	h := r.hashFunc()
	_, err = io.Copy(h, response.Body)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// Option is a function that takes a retriever and changes its attributes
type Option func(r *retriever)

// WithHashFunc returns an Option for setting retriever.hashFunc
func WithHashFunc(hashFunc func() hash.Hash) Option {
	return func(r *retriever) {
		r.hashFunc = hashFunc
	}
}

// WithClient returns an Option for setting retriever.client
func WithClient(client *http.Client) Option {
	return func(r *retriever) {
		r.client = client
	}
}

// NewRetriever returns an implementation of httphash.HashRetriever using retriever specified by options
func NewRetriever(options ...Option) httphash.HashRetriever {
	r := &retriever{
		client:   _defaultClient,
		hashFunc: _defaultHashFunc,
	}
	for _, o := range options {
		o(r)
	}
	return r
}
