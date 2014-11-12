## debug

## Install

```bash
go get github.com/favframework/debug
```

## Sample code

```go
package main

import (
	"github.com/favframework/debug"
)

func main() {
	a := make(map[string]int64)

	a["A"] = 1
	a["B"] = 2

	debug.Dump(a)
}
```

Then Print:

```bash
(map[string]int64)
  A(int64) 1
  B(int64) 2
```
