package emailvalidator

import (
	"errors"
	"testing"
)

var e error = errors.New("Invalid email")

var samples = []struct{
	email string
	expected error
}{
	{email: "email123@email.com", expected: nil},
	{email: "test@two.com", expected: nil},
	{email: "test @blablabla.gg", expected: e},
	{email: "someone@emailcom", expected: e},
	{email: "aaa.com", expected: e},
	{email: " test@test.com ", expected: e},
	{email: "test@test.com " , expected: e},
	{email: "a@3ema il.com", expected: e},
	{email: "test@.com", expected: e},
	{email: "blablabla@blablabla.lol", expected: nil},
}

func TestValid(t *testing.T) {
	for _, e := range samples {
		err := Valid(e.email)
		if (err != nil && e.expected == nil) || (err == nil && e.expected != nil) {
			t.Fatalf("🔴 %s\n", e.email)
		}
	}
}
