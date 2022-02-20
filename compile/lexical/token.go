package lexical

import (
	"fmt"
	"jerryLang/compile/commom"
	"strings"
)

type TokenType string

type Tokens[T comparable] struct {
	commom.List[T]
}

const (
	Token_VAR        TokenType = "_VAR_"
	Token_CHAR       TokenType = "char"
	Token_STR        TokenType = "str"
	Token_INT        TokenType = "int"
	Token_FLOAT      TokenType = "float"
	Token_BYTE       TokenType = "byte"
	Token_VECTOR     TokenType = "vector"
	Token_MAP        TokenType = "map"
	Token_TUPLE      TokenType = "tuple"
	Token_VAR_ASSIGN TokenType = "varAssign"
	Token_INT_TYPE   TokenType = "intType"
	Token_FLOAT_TYPE TokenType = "floatType"
	Token_STR_TYPE   TokenType = "strType"
	Token_EQUAL_SIGN TokenType = "="
)

type LexicalToken struct {
	tokenType  TokenType
	tokenValue string
}

func (self *LexicalToken) TokenType() TokenType {
	return self.tokenType
}

func (self *LexicalToken) TokenValue() string {
	return self.tokenValue
}

func genStrToken(data string) *LexicalToken {
	if data == "var" {
		return NewToken(Token_VAR, data)
	} else if data[len(data)-1] == ':' {
		return NewToken(Token_VAR_ASSIGN, data)
	} else if data == "int" {
		return NewToken(Token_INT_TYPE, data)
	} else if data == "float" {
		return NewToken(Token_FLOAT_TYPE, data)
	} else if data == "str" {
		return NewToken(Token_STR_TYPE, data)
	} else if data == "=" {
		return NewToken(Token_EQUAL_SIGN, data)
	} else {
		return NewToken(Token_STR, data)
	}
}

func genDigitalToken(data string) *LexicalToken {
	if strings.Index(data, ".") != -1 {
		return NewToken(Token_FLOAT, data)
	} else {
		return NewToken(Token_INT, data)
	}
}

func (self *LexicalToken) Log() string {
	fmt.Println("tokenType: ", self.tokenType, "tokenValue: ", self.tokenValue)
	return self.tokenValue
}

func NewToken(tokenType TokenType, tokenValue string) *LexicalToken {
	var p = new(LexicalToken)
	p.tokenType = tokenType
	p.tokenValue = tokenValue

	return p
}
