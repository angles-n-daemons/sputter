package sputter

import (
	"bytes"
	"fmt"
	"math/rand"
	"regexp/syntax"
	"time"
	"unicode/utf8"
)

const (
	repetitionMax = 100
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
	return sput(r)
}

// simple dfs syntax to string
func sput(r *syntax.Regexp) (string, error) {
	switch r.Op {
	case syntax.OpLiteral:
		return literal(r), nil
	case syntax.OpCharClass:
		return charClass(r), nil
	case syntax.OpAnyCharNotNL:
		return any(r), nil
	case syntax.OpBeginLine:
		return "\n", nil
	case syntax.OpEndLine:
		return "\n", nil
	case syntax.OpCapture:
		return sput(r.Sub[0])
	case syntax.OpRepeat:
		return repeat(r)
	case syntax.OpStar:
		return star(r)
	case syntax.OpPlus:
		return plus(r)
	case syntax.OpQuest:
		return quest(r)
	case syntax.OpConcat:
		return concat(r)
	case syntax.OpAlternate:
		return alternate(r)
	default:
		return "", fmt.Errorf("unsupported syntax operation %d", r.Op)
	}
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

func any(r *syntax.Regexp) string {
	c := '\n'
	for utf8.ValidRune(c) && c == '\n' {
		c = rune(random(1, int(utf8.MaxRune)))
	}
	return string([]rune{c})
}

func begin(r *syntax.Regexp) string {
	c := '\n'
	for utf8.ValidRune(c) && c == '\n' {
		c = rune(random(1, int(utf8.MaxRune)))
	}
	return string([]rune{c})
}

func end(r *syntax.Regexp) string {
	c := '\n'
	for utf8.ValidRune(c) && c == '\n' {
		c = rune(random(1, int(utf8.MaxRune)))
	}
	return string([]rune{c})
}

func repeat(r *syntax.Regexp) (string, error) {
	var buffer bytes.Buffer
	n := random(r.Min, r.Max)
	if r.Max == 0 {
		return "", nil
	}

	for i := 0; i < n; i++ {
		s, err := sput(r.Sub[0])
		if err != nil {
			return "", err
		}

		_, err = buffer.WriteString(s)
		if err != nil {
			return "", err
		}
	}

	return buffer.String(), nil
}

func star(r *syntax.Regexp) (string, error) {
	var buffer bytes.Buffer
	n := random(0, repetitionMax)
	for i := 0; i < n; i++ {
		s, err := sput(r.Sub[0])
		if err != nil {
			return "", err
		}

		_, err = buffer.WriteString(s)
		if err != nil {
			return "", err
		}
	}
	return buffer.String(), nil
}

func plus(r *syntax.Regexp) (string, error) {
	var buffer bytes.Buffer
	n := random(1, repetitionMax)
	for i := 0; i < n; i++ {
		s, err := sput(r.Sub[0])
		if err != nil {
			return "", err
		}

		_, err = buffer.WriteString(s)
		if err != nil {
			return "", err
		}
	}
	return buffer.String(), nil
}

func quest(r *syntax.Regexp) (string, error) {
	if rand.Int()%2 == 0 {
		return sput(r.Sub[0])
	}
	return "", nil
}

func concat(r *syntax.Regexp) (string, error) {
	var buffer bytes.Buffer
	for _, sub := range r.Sub {
		s, err := sput(sub)
		if err != nil {
			return "", err
		}

		_, err = buffer.WriteString(s)
		if err != nil {
			return "", err
		}
	}

	return buffer.String(), nil
}

func alternate(r *syntax.Regexp) (string, error) {
	i := random(0, len(r.Sub))
	return sput(r.Sub[i])
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
