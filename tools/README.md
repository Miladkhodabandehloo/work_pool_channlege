# tools

This directory comprises types, interfaces, and functionalities for having url hash retriever worker pool

### Structure

```
tools
│      
└─── httphash                       the package that defines and implements url hash retriever 
│   │
│   └─── httphash.go                defenition of HashRetriever interface
│   │
│   └─── bodyhash/bodyhash.go       implementation of HashRetiever with calculation of hash digest for body of HTTP Get response
│
│
└─── workerpool                     the package for having a url hash retriever worker pool
    │
    └─── workerpool.go              implementation of url hash retriever worker pool
```