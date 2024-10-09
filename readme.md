# TypedSyncMap

TypedSyncMap is a thread-safe map with typed keys and values. It is a wrapper around `sync.Map` that provides type safety in Go.

## Installation

To install the package, use the following command:

```sh
go get github.com/yourusername/typedSyncMap
```

## Usage

### Creating a TypedSyncMap

To create a new `TypedSyncMap`, use the `NewTypedSyncMap` function:

```go
package main

import (
    "fmt"
    "typedSyncMap/common"
)

func main() {
    m := common.NewTypedSyncMap[string, string]()
    m.Store("key1", "value1")
    value, ok := m.Load("key1")
    if ok {
        fmt.Println("Loaded value:", value)
    }
}
```

### Methods

#### `Load`

Returns the value stored in the map for a key, or `nil` if no value is present. The `ok` result indicates whether the value was found in the map.

```go
value, ok := m.Load("key1")
```

#### `Store`

Sets the value for a key.

```go
m.Store("key1", "value1")
```

#### `Delete`

Deletes the value for a key.

```go
m.Delete("key1")
```

#### `Range`

Calls a function sequentially for each key and value present in the map.

```go
m.Range(func(key, value string) bool {
    fmt.Println(key, value)
    return true
})
```

#### `CopyFrom`

Copies all key-value pairs from the source map to the destination map.

```go
src := common.NewTypedSyncMap[string, string]()
src.Store("key1", "value1")
m.CopyFrom(&src)
```

## Testing

To run the tests, use the following command:

```sh
go test ./...
```

## License

This project is licensed under the MIT License.
```

This `README.md` provides an overview of the `TypedSyncMap` package, including installation instructions, usage examples, method descriptions, and testing instructions.