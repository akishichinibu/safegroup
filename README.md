# safegroup

`safegroup` is a wrapper around [golang.org/x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup), which automatically catches panics inside goroutines and converts them into errors, preventing the entire program from crashing.

ref: https://github.com/golang/go/issues/53757

## Features

- ✅ Compatible API with `errgroup`
- ✅ Captures stack trace information on panic

## Installation

```bash
go get github.com/akishichinibu/safegroup
```

## Usage

```go
package main

import (
 "fmt"
 "github.com/akishichinibu/safegroup"
)

func main() {
 sg := safegroup.New()

 sg.Go(func() error {
  fmt.Println("this task runs normally")
  return nil
 })

 sg.Go(func() error {
  panic("something went wrong")
 })

 err := sg.Wait()
 if err != nil {
  fmt.Printf("SafeGroup caught an error: %v\n", err)
 }
}
```

## Panic Handling

When a panic occurs inside a goroutine started by `Go` or `TryGo`, SafeGroup automatically recovers it and returns a `*PanicError` containing:

- The panic value (`Expt`)

- The captured stack trace (`Stack`)

You can safely log, report, or analyze panics without crashing your program.

### Example

```go
if pErr, ok := err.(*safegroup.PanicError); ok {
 fmt.Printf("panic occurred: %v\nstack trace:\n%s\n", pErr.Expt, string(pErr.Stack))
}
```

## License

MIT License.
Parts of the test code are adapted from the Go project (BSD license).
