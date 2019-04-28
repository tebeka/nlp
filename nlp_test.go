package nlp

import (
	//	"reflect"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	testCases := []struct {
		text     string
		expected []string
	}{
		{"Who's on first?", []string{"who", "on", "first"}},
		{"", []string(nil)},
	}

	for _, tc := range testCases {
		name := tc.text
		if name == "" {
			name = "<empty>"
		}
		t.Run(name, func(t *testing.T) {
			require := require.New(t)
			out := Tokenize(tc.text)
			require.Equal(tc.expected, out, "tokenize")
			// if expected != out { // can't compare slices
			/*
				if !reflect.DeepEqual(tc.expected, out) {
					t.Fatalf("%#v != %#v", tc.expected, out)
				}
			*/
		})
	}
}

func TestQuick(t *testing.T) {
	require := require.New(t)
	fn := func(text string) bool {
		tokens := Tokenize(text)
		if len(wordRe.FindAllString(text, -1)) != len(tokens) {
			t.Log(text)
			return false
		}
		return true
	}

	require.NoError(quick.Check(fn, nil))
}
