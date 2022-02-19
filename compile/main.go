package main

import (
	"jerryLang/compile/lexical"
	"jerryLang/compile/syntax"
	"os"
)

// 词法分析
func main() {
	path, _ := os.Getwd()
	filename := path + "/compile/resource/simple1.jl"

	// 词法解析 -> tokens
	tokens := lexical.Parse(filename)
	// 语法解析 -> 抽象语法树
	syntax.Parse(tokens)
}
