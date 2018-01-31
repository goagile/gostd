package string

import "testing"

type Case struct {
	s, want string
}
	
func Test(t *testing.T) {
	var cases = []Case{
		{"Back", "kcaB"},
		{"", ""},
	}

	for _, c := range cases {
		got := Reverse(c.s)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}
