# openssl

A openssl tool

## Installation

### Prerequisites

- **[Go](https://go.dev/)** >= 1.20

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```bash
import "github.com/hugiot/pinecone/openssl"
```

## Demo

```go
package main

import (
	"fmt"
	"github.com/hugiot/pinecone/openssl"
)

func main() {
	fmt.Println(openssl.MD5([]byte("hello world")))
	fmt.Println(openssl.SHA1ToString([]byte("hello world")))
	fmt.Println(openssl.SHA256ToString([]byte("hello world")))
}
```