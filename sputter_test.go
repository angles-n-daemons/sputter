package sputter

import (
	"fmt"
	"regexp"
	"testing"
)

func TestLiteral(t *testing.T) {
	testSputHundredEmoji(t, "abc")
}

func TestCharClassSingle(t *testing.T) {
	testSputHundredEmoji(t, "[A]")
}

func TestCharClassRange(t *testing.T) {
	testSputHundredEmoji(t, "[A0-9]")
}

func testSputHundredEmoji(t *testing.T, exp string) {
	for i := 0; i < 100; i++ {
		s, err := Gen(exp)
		fmt.Println(s)
		if err != nil {
			t.Error("error from Gen:", err)
		}

		match, err := regexp.Match(exp, []byte(s))
		if !match {
			t.Errorf(`output string "%s" does not match expression "%s"`, s, exp)
		}

		if err != nil {
			t.Error(err)
		}
	}
}
