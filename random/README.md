# random

A random number generator

## Installation

### Prerequisites

- **[Go](https://go.dev/)** >= 1.20

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```bash
import "github.com/hugiot/pinecone/random"
```

## Demo

```go
package main

import (
	"fmt"
	"github.com/hugiot/pinecone/random"
)

func main() {
	fmt.Println(random.Int(100, 200))
	fmt.Println(random.Int64(100, 200))
}
```