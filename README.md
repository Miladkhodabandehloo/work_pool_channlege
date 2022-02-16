# Code challenge

The aim of this project is to provide an **url hash retriever worker pool** for getting hash digest for a collection of
urls. At the moment, it makes a Get HTTP call for each url, then retrieves [MD5](https://en.wikipedia.org/wiki/MD5)
digest of the body of the response.

- The project is developed in a way that it is flexible for changing the hashing algorithm. The only thing to do is to
  change the default hashing algorithm using an option.
- It is flexible to change how we should calculate the digest regarding an url. The only thing to do is to implement
  **httphash.HashRetriever**
- It is possible to specify the number of workers that work in parallel to accomplish the task.

### Getting started

In order to use the worker pool follow the below instructions.

1- Install Go programming language ([installation manual](https://go.dev/doc/install))

2- Change the working directory to the project directory

3- use following command to build an executable for using the tool

```shell
go build -o build/tool cmd/url_hash_retriever/main.go
```

Then you can use the tool by following command

```shell
build/tool 5  http://google.com http://time.ir
```

The first argument is the number of workers that should accomplish the task, and then you can pass as many urls as you
want by other argument.

if you do not provide number of workers its default number is 10.

#### result

```
incorrecturl: Get "incorrecturl": unsupported protocol scheme ""
http://google.com: 4232d22639faa16ec9218307919e371b
http://time.ir: 2542908a294bf906f28cc7ee01c24c9e

```

in the case of occurring an error, instead of hash digest, the error description will be printed after the url.

### Structure

The following directory tree is the representation of the structure in which project files and directories are located.

```
project
│   README.md       project description and guide
│   .gitignore      for specifying files that git must ignore   
│   go.mod          for dependency management
│   go.sum          for dependency management
│      
└─── cmd            Entry point for using the worker pool
│   
└─── tests          unit tests 
│   
└─── tools          packages for having the url hash retriever worker pool
```

### Tests

To run the tests in your development environment that has [Go](https://go.dev/) installed, run following command:

```shell
go test ./tests/...
```