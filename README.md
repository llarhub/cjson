# cJSON

LLGo bindings for cJSON 1.7.19.

## Usage

Install cJSON and the Go binding:

```bash
brew install cjson
 go get github.com/llarhub/cjson
```

Parse JSON and read a string field:

```go
package main

import (
	"github.com/goplus/lib/c"
	"github.com/llarhub/cjson"
)

func main() {
	root := cjson.Parse(c.AllocaCStr(`{"name":"LLGo"}`))
	if root == nil {
		panic("invalid JSON")
	}
	defer root.Delete()

	name := root.ObjectItemCaseSensitive(c.AllocaCStr("name"))
	if name != nil && name.IsString() != 0 {
		println(c.GoString(name.StringValue()))
	}
}
```

Save the example as `main.go` and run it with `llgo run .`.
