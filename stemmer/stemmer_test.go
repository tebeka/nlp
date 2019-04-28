package stemmer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testCases = []struct {
	word     string
	expected string
}{
	{"runs", "run"},
	{"sleeping", "sleep"},
	{"fish", "fish"},
	{"", ""},
}

func TestStem(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.word, func(t *testing.T) {
			out := Stem(tc.word)
			require.Equal(t, tc.expected, out)
		})
	}
}
