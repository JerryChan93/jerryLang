package lexical

import (
	"jerryLang/compile/commom"
	"unicode"
)

func Parse(filename string) *Tokens {
	var lines = getFileLines(filename)
	tokens := handleParse(lines)
	return tokens
}

func handleParse(lines []string) *Tokens {
	tokens := NewTokens()
	length := len(lines)
	for i := 0; i < length; i++ {
		parseLine(tokens, lines[i])
	}
	return tokens
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

func parseLine(self *Tokens, line string) []*Tokens {
	var specialList = []byte{'=', '>', '<', '!', '+', '-', '*', '\\'}
	for i := 0; i < len(line); i++ {
		if unicode.IsLetter(rune(line[i])) {
			end := stringEnd(line, i)
			var data = genTokenStr(line, i, end)
			self.add(genStrToken(data))
			i = end
			continue
		}
		if unicode.IsDigit(rune(line[i])) {
			end := digitEnd(line, i)
			var data = genTokenStr(line, i, end)
			self.add(genDigitalToken(data))
			i = end
			continue
		}
		code := line[i]
		if commom.InArray(specialList, code) {
			end := stringEnd(line, i)
			var data = genTokenStr(line, i, end)
			self.add(genStrToken(data))
			i = end
			continue
		}
	}
	return nil
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
