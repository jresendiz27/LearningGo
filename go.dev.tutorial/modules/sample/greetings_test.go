package sample

import (
	"testing"
	"regexp"
)

func TestHelloName(t *testing.T) {
	name := "jresendiz"
	want := regexp.MustCompile(name)
	msg, err := Hello("jresendiz")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("jresendiz") = %q, %v want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg!= "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v want ""`, msg, err)
	}
}