package fizzbuzz

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{-5, ""},
		{0, ""},
		{1, "1"},
		{5, "1, 2, Fizz, 4, Buzz"},
		{15, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz"},
	}

	for _, c := range cases {
		got := FizzBuzz(c.in)
		if got != c.want {
			t.Errorf("FizzBuzz(%v) == %q, wanted %q", c.in, got, c.want)
		}
	}
}
