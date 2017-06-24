package sputter

import (
	"math/rand"
	"regexp/syntax"
	"time"
)

// Gen takes a regular expression and attempts
//  to generate a pseudo-randomized string that
//  matches the input expression.
func Gen(exp string) (string, error) {
	// setup random package
	rand.Seed(time.Now().UTC().UnixNano())

	r, err := syntax.Parse(exp, 0)
	if err != nil {
		return "", err
	}
	return sput(r), nil
}

// simple dfs syntax to string
func sput(r *syntax.Regexp) string {
	switch r.Op {
	case syntax.OpLiteral:
		return literal(r)
	case syntax.OpCharClass:
		return charClass(r)
	}
	return ""
}

func literal(r *syntax.Regexp) string {
	return string(r.Rune)
}

func charClass(r *syntax.Regexp) string {
	switch len(r.Rune) {
	case 1:
		return literal(r)
	default:
		// randomly select from available ranges
		i := random(0, len(r.Rune)/2) * 2
		n1, n2 := int(r.Rune[i]), int(r.Rune[i+1])
		if n2-n1 == 0 {
			return string(r.Rune[i : i+1])
		}
		randRune := rune(random(n1, n2+1))
		s := string(
			[]rune{
				randRune,
			},
		)
		return s
	}
	return ""
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
