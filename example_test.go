package nlp_test

import (
	"fmt"

	"github.com/tebeka/nlp"
)

func ExampleTokenize() {
	text := "What's up doc?"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)

	// Output:
	// [what s up doc]
}
