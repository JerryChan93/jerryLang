package lexical

import (
	"io/ioutil"
	"jerryLang/compile/commom"
	"os"
	"strings"
	"unicode"
)

type Tokens struct {
	head   *TokenNode
	tail   *TokenNode
	length int32
}

type TokenNode struct {
	token *LexicalToken
	next  *TokenNode
}

func Parse(filename string) *Tokens {
	getFileLines(filename)
	return nil
}

// 读取文件内容
// 过滤注释

func getFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(file)
	lines := strings.Split(string(content), "\n")
	lines = validateLines(lines)

	commom.LogList(lines)
	handleParse(lines)
	return lines
}

func validateLines(lines []string) []string {
	lines = commom.Map(lines, func(line string) string {
		line = strings.Trim(line, " ")
		// 去除空注释
		index := strings.Index(line, "//")
		if index != -1 {
			line = line[0:index]
		}

		return line
	})

	lines = commom.FilterIn(lines, func(line string) bool {
		return len(line) > 0
	})

	return lines
}

func handleParse(lines []string) *Tokens {
	length := len(lines)
	for i := 0; i < length; i++ {
		parseLine(lines[i])
	}
	return nil
}

func genTokenStr(line string, start int, end int) string {
	var data = ""
	if end == start {
		data = string(line[end])
	} else {
		data = line[start:end]
	}

	return data
}

func parseLine(line string) []*Tokens {
	var specialList = []byte{'=', '>', '<', '!', '+', '-', '*', '\\'}
	for i := 0; i < len(line); i++ {
		if unicode.IsLetter(rune(line[i])) {
			end := stringEnd(line, i)
			var data = genTokenStr(line, i, end)
			genStrToken(data)
			i = end
			continue
		}
		if unicode.IsDigit(rune(line[i])) {
			end := digitEnd(line, i)
			var data = genTokenStr(line, i, end)
			genStrToken(data)
			i = end
			continue
		}
		code := line[i]
		if commom.InArray(specialList, code) {
			end := stringEnd(line, i)
			var data = genTokenStr(line, i, end)
			genStrToken(data)
			i = end
			continue
		}

	}
	return nil
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

func stringEnd(line string, startIndex int) int {
	length := len(line)
	for i := startIndex + 1; i < length; i++ {
		if line[i] == ' ' {
			return i
		}

	}

	return length - 1
}

func digitEnd(line string, startIndex int) int {
	length := len(line)
	for i := startIndex + 1; i < length; i++ {
		if line[i] == ' ' {
			return i
		}

	}

	return length - 1
}

func (self *Tokens) add(token *LexicalToken) {
	var node = new(TokenNode)
	node.token = token
	node.next = nil

	self.tail.next = node
	self.tail = node
	self.length += 1

}
