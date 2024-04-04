package leetcode0383

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name        string
		ransomeNote string
		magazine    string
		want        bool
	}{
		{
			name:        "test 1",
			ransomeNote: "a",
			magazine:    "b",
			want:        false,
		},
		{
			name:        "test 2",
			ransomeNote: "aa",
			magazine:    "ab",
			want:        false,
		},
		{
			name:        "test 3",
			ransomeNote: "aa",
			magazine:    "aab",
			want:        true,
		},
	}

	for _, test := range tests {
		have := canConstruct(test.ransomeNote, test.magazine)
		assert.Equalf(t, test.want, have, "%s: mismatch in canConstruct", test.name)
	}
}
