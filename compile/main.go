package main

import (
	"jerryLang/compile/lexical"
	"os"
)

// 词法分析
func main() {
	path, _ := os.Getwd()
	filename := path + "/compile/resource/simple1.jl"

	// 词法解析 -> tokens
	tokens := lexical.Parse(filename)
	// 语法解析 -> 语法树
	tokens.Log()
}
