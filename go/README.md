# Golang - Features That Make It Great

## A Few Good Things About Go

### Simple Language: Easy to Read, Easy to Understand
Go is easy to read, the syntax is quite easy to follow. Go has usually one way of doing things on the community, which reduces noise.
To be honest, all the code looks pretty much the same, following the same conventions and style. Go team has a [style guide](https://golang.org/doc/effective_go.html) that is widely followed by the community.

I like the way how go follows a [line of sight](https://medium.com/@matryer/line-of-sight-in-code-186dd7cdea88) idea, and from my experience, most of the code you will see in the wild follows this principle, dive in any major open-source project and you will see that the code is easy to read and understand: no magic, no tricks, just plain code. Even something as complex as [`docker-compose build`](https://github.com/docker/compose/blob/main/pkg/compose/build.go#L66) can be read easily.



### Great Standard Package
Go has a great standard library. The standard library is well-designed and easy to use. It has packages for working with files (i/o handle), web servers, data manipulation, JSON serializarion/deserialization, testing, concurrency, and much more.

The standard library is well-documented, and the documentation is easy to read. The standard library is also well-tested, which means you can rely on it to work correctly.

### Built for Concurrency
Go has built-in support for concurrent programming. Go's goroutines make it easy to write concurrent code. Goroutines are lightweight threads that are managed by the Go runtime. You can create thousands of goroutines in a Go program without running out of memory.

Go's channels make it easy to communicate between goroutines. Using goroutine(s) and channel(s) it’s almost trivial.

### Interfaces everywhere
Go uses interfaces to define behavior. Interfaces are a powerful feature of Go that allows you to define behavior without specifying the implementation. This makes it easy to write code that is decoupled and easy to test.

### Type Safe
Go is a statically typed language. This means that the compiler checks the types of variables at compile time. This helps catch errors early in the development process.

### Garbage Collection
Go has a garbage collector that automatically manages memory. The garbage collector runs in the background and reclaims memory that is no longer in use. This makes it easy to write code without worrying about memory management.

## Some more important concepts

## Go Memory Management

Heap and Stack are two important concepts in memory management. In Go, memory is managed by the garbage collector. The garbage collector is responsible for allocating and freeing memory.
Heap is usually used for dynamic memory allocation, while the stack is used for static memory allocation. Practically, the stack is used for storing local variables, function parameters, and return addresses, while the heap is used for storing objects that are created at runtime.

### Memory Leaks
Memory leaks are a common problem in programming. A memory leak occurs when a program allocates memory but does not free it. This can lead to memory exhaustion and cause the program to crash. Example:
1. go routine that is not closed
2. db connection that is not closed
