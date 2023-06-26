# hmac

A hmac tool

## Installation

### Prerequisites

- **[Go](https://go.dev/)** >= 1.20

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```bash
import "github.com/hugiot/pinecone/hmac"
```

## Demo

```go
package main

import (
	"fmt"
	"github.com/hugiot/pinecone/hmac"
)

func main() {
	fmt.Println(hmac.SHA1ToString([]byte("123456"), []byte("hello world")))
	fmt.Println(hmac.SHA224ToString([]byte("123456"), []byte("hello world")))
	fmt.Println(hmac.SHA256ToString([]byte("123456"), []byte("hello world")))
	fmt.Println(hmac.SHA512ToString([]byte("123456"), []byte("hello world")))
}
```