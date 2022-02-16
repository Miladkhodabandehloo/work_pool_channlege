# tests

This directory comprises unit tests, mocked functionalities, and test data.

### Structure

```
tests
│      
└─── data                           test data
│   │
│   └─── hashes.json                test data that includes hashes for some urls
│
│
└─── mock                           mocked functionalities
│   │
│   └─── hash_retriever.go          mock implementation of a url hash retriever that reads hash digests from hashes.json
│
│
└─── worker_pool_test               unit tests for testing the url hash retriever worker pool
```