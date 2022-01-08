package gen

import (
	"testing"

	"github.com/spudtrooper/goutil/or"
)

func TestIncTag(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{
			input: "v0.2.1",
			want:  "v0.2.2",
		},
		{
			input: "v0.2.10",
			want:  "v0.2.11",
		},
	}
	for _, test := range tests {
		t.Run(or.String(test.name, test.input), func(t *testing.T) {
			newTag, err := incTag(test.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if want, got := test.want, newTag; want != got {
				t.Errorf("incTag(%q) want(%q) != got(%q)", test.input, want, got)
			}
		})
	}
}
