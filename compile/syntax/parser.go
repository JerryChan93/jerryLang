package syntax

import (
	"jerryLang/compile/lexical"
)

func Parse(tokens *lexical.Tokens[*lexical.LexicalToken]) {
	tokens.Each(func(item *lexical.LexicalToken) {
		item.Log()
	})
}
