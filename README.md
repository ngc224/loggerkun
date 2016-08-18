# loggerkun

```
package main

import (
	"fmt"
	"os"

	"github.com/ngc224/loggerkun"
)

type ABC struct {
	A string
	B string
	C string
}

func main() {
	l, err := loggerkun.New("{{.A}} - {{.B}} - {{.C}}")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	l.PUT(&ABC{
		A: "aaa",
		B: "bbb",
		C: "ccc",
	})
}
```