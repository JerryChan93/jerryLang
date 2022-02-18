package lexical

import (
	"fmt"
	"strings"
)

type LexicalTokenType int32

const (
	Token_VAR        LexicalTokenType = 0
	Token_CHAR       LexicalTokenType = 1
	Token_STR        LexicalTokenType = 2
	Token_INT        LexicalTokenType = 3
	Token_FLOAT      LexicalTokenType = 4
	Token_BYTE       LexicalTokenType = 5
	Token_VECTOR     LexicalTokenType = 6
	Token_MAP        LexicalTokenType = 7
	Token_TUPLE      LexicalTokenType = 8
	Token_VAR_ASSIGN LexicalTokenType = 9
)

type LexicalToken struct {
	tokenType  LexicalTokenType
	tokenValue string
}

func (self *LexicalToken) TokenType() LexicalTokenType {
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

func (self *LexicalToken) log() string {
	fmt.Println("tokenType: ", self.tokenType, "tokenValue: ", self.tokenValue)
	return self.tokenValue
}

func NewToken(tokenType LexicalTokenType, tokenValue string) *LexicalToken {
	var p = new(LexicalToken)
	p.tokenType = tokenType
	p.tokenValue = tokenValue

	return p
}
