package stringutil

import "testing"
import "reflect"

func TestWordCount(t *testing.T) {
	cases := []struct {
		in   string
		want map[string]int
	}{
		{"eu sei quantas palavras tem sei eu", map[string]int{"eu": 2, "sei": 2, "quantas": 1, "palavras": 1, "tem": 1}},
		{"teste teste", map[string]int{"teste":2}},
	}
	for _, c := range cases {
		got := WordCount(c.in)
		if ! reflect.DeepEqual(got,c.want) {
			t.Errorf("Wordcount(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
