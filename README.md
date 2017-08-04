# sputter

POSIX basic regular expressions to psuedo random string generator, just for fun :)

### Usage
```go
package main

import (
	"fmt"

	"github.com/brianasapp/sputter"
)

func main() {
	s, err := sputter.Gen("[A-Z0-9]^(Word){1,3}$.+")
	if err != nil {
		panic(err)
	}
	fmt.Printf("generated below: \n%s\n", s)
}
```

```bash
$ go run main.go
generated below:
5
WordWord
Њѯѹկ¢↔≡♲
```

_For cryptographically insecure usage, use the `GenInsecure` function in place of `Gen`_



### Supported Operations
 * literal
 * character class
 * capture
 * any char not newline
 * begin line
 * end line
 * star
 * plus
 * question
 * repeat
 * concat
 * alternation

_note: for randomized repetition, there is a max value of 100. let me know if there's a better way to do it_
