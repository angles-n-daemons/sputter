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
