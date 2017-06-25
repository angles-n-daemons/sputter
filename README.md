# sputter

_WARNING: I AM NOT CRYPTO FRIENDLY_

POSIX basic regular expressions to psuedo random string generator

### Usage
```go
package main()

import "github.com/brianasapp/sputter"

func main() {
	s, err := sputter.Gen("[A-Z0-9](Word){1,3}.")
	if err != nil {
		panic(err)
	}
	fmt.Printf("generated: ",s)
}
```

```bash
$ go run main.go
generated: 5WordWordꄈ
```



### Supported Operations
 * literal
 * character class
 * capture
 * any char not newline
 * star
 * plus
 * question
 * repeat
 * concat
 * alternation

_note: for randomized repetition, there is a max value of 100. let me know if there's a better way to do it_
