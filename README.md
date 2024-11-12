# A Brief Overview of Go

![gopher](gopher.png)

## Goals

* introduction to the go programming language (golang) for developers familiar with other languages
* from practitioners for practitioners
* a little bit of the history of go, comparison with other languages (chart)
* many tiny fully self-contained coding examples to exemplify common coding patterns and pitfalls in golang
* distilled version: save you time reading a thick textbook or watch hours of tutorials (maybe)
* focus on relevant features and examples

## Resources

### Golang Core

* Go Playground https://go.dev/play/
* A Tour Of Go https://go.dev/tour/list
* Go By Example https://gobyexample.com/
* How to write Go code https://go.dev/doc/code
* Effective Go https://go.dev/doc/effective_go
* Go Blog https://go.dev/blog/strings
* Downloads https://go.dev/dl/

### Talks

* Rob Pike, Lexical Scanning in Go (channels) - https://www.youtube.com/watch?v=HxaD_trXwRE
* Liz Rice, Building a container from scratch in Go (unix system commands) - https://www.youtube.com/watch?v=Utf-A4rODH8

## The Language - Syntax and Concepts

* [simplicity](simplicity/main.go) - simple, concise and C-inspired syntax
* [shadowing](shadowing/main.go) - shadowing variables based on scope
* [interface](interface/main.go) - interface, type casts and type conversions, no generics
* [pointers](pointers/main.go) - pointers work very much like they do in C
* [functional](functional/main.go) - functional programming
* [variadic](variadic/main.go) - introduction to variadic functions
* [defer](defer/main.go) - use defer when you don't want to forget to clean up
* [panic](panic/main.go) - recover from panics
* [mapsandslices](mapsandslices/main.go) - basic operations with maps and slices
* [safemap](safemap/main.go) - a thread safe map implementation
* [jsoninjsonout](jsoninjsonout/main.go) - parse and serialize json payloads
* [duck](duck/main.go) - duck typing, no need to explicitly declare which interfaces are implemented
* [composition](composition/main.go) - composition over inheritance
* [oop](oop/main.go) - limitations regarding overriding functions, workaround using functional programming
* [concurrency](concurrency/main.go) - use go routines for concurrent execution
* [channel](channel/main.go) - creative use of goroutines and channels
* [workerpool](workerpool/main.go) - combine go routines with go channels to send jobs to workers in a pool 
* [fanoutfanin](workerpool/main.go) - use worker pool to fan-out work and main thread to fan-in results
* [producerconsumer](producerconsumer/main.go) / producer sends work to consumer via channel
* [packages](packages/main.go) - importing and using packages, packages live inside a directory, everything uppercase is visible from outside the package
* [module](module/main.go) - manage package versions and dependencies

## Patterns

* [aws](aws/main.go) - use aws sdk to interact with amazon web services
* [options](options/main.go) - functional options
* [context](context/main.go) - use go context to manage lifecycle of long-running tasks with cancel and / or timeout
* [errors](errors/main.go) - errors.As(), errors.Is()
* [tinywebservice](tinywebservice/main.go) - a tiny web service using the http package from the standard library
* [controlc](controlc/main.go) - capture ctrl-c OS signal, then wait for all goroutines to finish
* [shell](shell/main.go) execute shell command
* [shellong](shellong/main.go) execute long running shell command
* File I/O and HTTP
* init function
* Commandline parameters

## Frameworks for Building RESTful Web Services

*  [webservice](webservice/main.go) - building REST APIs using viper, cobra, gorilla mux, zerolog and other frameworks

## Useful Libraries

* https://github.com/aws/aws-sdk-go-v2 - AWS SDK
* https://github.com/IBM/sarama - Kafka library
* https://github.com/rs/zerolog - structured logging
* https://github.com/gorilla/mux - http web services
* https://github.com/spf13/cobra - cli commands
* https://github.com/spf13/viper - configuration

## Special Features

* [embed](embed/main.go) - embed files in go binary with //go:embed
* [cgo](cgo/main.go) - use cgo to call C library functions

## The Tool Chain

### Packages, Modules and go mod

```
go mod init
```

```
go mod tidy
```

```
go get somedomain.com/somepackage
```

### Go Build

```
go build
```

### Cross-Compiling

```
GOOS=linux GOARCH=arm go build -o myprogram myprogram.go
```

### Go Test

https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package

#### Unit Tests

[unittest](unittest/factorial_test.go)

```
go test -v
```

https://pkg.go.dev/testing

#### Benchmarks

```
go test -bench=.
```

Note: Runtime of benchmarked function must converge!

https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

#### Race Conditions

```
go test -race
```

#### Golden Files

[goldie](goldie/main_test.go) - easy testing with golden files

https://github.com/sebdah/goldie

#### Mock Services

#### Table Driven Testing

### Linter

[linter](linter/main.go)

```
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

```
golangci-lint run
```

https://golangci-lint.run/usage/configuration/#config-file

### Profiling

[pprof](pprof/main.go) - use pprof for heap profiling etc

Profile interactively for example with 

```
go tool pprof http://localhost:8080/debug/pprof/heap
```

then use `web` command.

https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/

### Go Releaser

### Go Generate

## Misc

* Core features
  * compiled
  * statically typed
  * garbage collected
  * statically linked
  * C-style pointers
  * no objects (structs and composition instead)
  * no generics
  * no method overloading
  * conventions (tests, exported functions, init function etc.)
* Strong support for concurrency built in (channels and concurrent execution as first class object, suited for multi-core architectures)
* History of golang and resources, Rob Pike et al.
* "Young" language
* Easy to transition, especially if you know a little about C (structs, pointers)
* Intellij Goland and debugging with https://github.com/go-delve/delve
* Semantic versioning



