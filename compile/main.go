package main

import (
	"jerryLang/compile/lexical"
	"os"
)

// 词法分析
func main() {
	path, _ := os.Getwd()
	filename := path + "/compile/resource/simple1.jl"

	lexical.Parse(filename)
}
