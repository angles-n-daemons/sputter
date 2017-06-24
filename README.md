# sputter

_WARNING: I AM NOT CRYPTO FRIENDLY_

Regexp to pseudostring generator

### Usage
```go
package main()

import "github.com/brianasapp/sputter"

func main() {
	s, err := sputter.Gen("[A-Z0-9]")
	if err != nil {
		panic(err)
	}
	// do something with s
}
```

### Supported Operations
 * literal
 * character class
