# sputter

_WARNING: I AM NOT CRYPTO FRIENDLY_

Regexp to pseudostring generator

### Usage
```go
package main()

import "github.com/brianasapp/sputter"

func main() {
	s, err := sputter.Gen("[A-Z0-9](Word){1,3}")
	// returns something like "CWordWord"
	if err != nil {
		panic(err)
	}
	// do something with s
}
```

### Supported Operations
 * literal
 * character class
 * capture
 * star
 * plus
 * question
 * repeat
 * concat
 * alternation

_note: for randomized repetition, there is a max value of 100. let me know if there's a better way to do it_
