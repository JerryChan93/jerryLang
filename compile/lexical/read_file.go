package lexical

import (
	"io/ioutil"
	"jerryLang/compile/commom"
	"os"
	"strings"
)

func getFileLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	content, _ := ioutil.ReadAll(file)
	lines := strings.Split(string(content), "\n")
	lines = validateLines(lines)

	commom.LogList(lines)
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
