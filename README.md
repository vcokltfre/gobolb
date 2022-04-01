# gobolb

Bolb.

## Usage

Add the package: `go get github.com/vcokltfre/gobolb`

```go
package main

import "github.com/vcokltfre/gobolb/bolb"

func main() {
    bolb.Bolb()
}
```

Server:

```go
package main

import "github.com/vcokltfre/gobolb/bolb"

func main() {
    bolb.Serve(8080, "127.0.0.1")
}
```
