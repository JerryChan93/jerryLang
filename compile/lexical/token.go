package lexical

import (
	"fmt"
	"jerryLang/compile/commom"
	"strings"
)

type TokenType int32

type Tokens[T comparable] struct {
	commom.List[T]
}

const (
	Token_VAR        TokenType = 0
	Token_CHAR       TokenType = 1
	Token_STR        TokenType = 2
	Token_INT        TokenType = 3
	Token_FLOAT      TokenType = 4
	Token_BYTE       TokenType = 5
	Token_VECTOR     TokenType = 6
	Token_MAP        TokenType = 7
	Token_TUPLE      TokenType = 8
	Token_VAR_ASSIGN TokenType = 9
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
