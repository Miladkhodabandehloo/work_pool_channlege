// Package httphash defines what behavior a HashRetriever must show
package httphash

import "context"

// HashRetriever is the interface that defines behavior of an url hash retriever
type HashRetriever interface {
	Retrieve(ctx context.Context, url string) (string, error)
}
