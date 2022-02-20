package syntax

import (
	"jerryLang/compile/commom"
)

type ASTNodeType string

type params[T comparable] struct {
	commom.List[T]
}
type body[T comparable] struct {
	commom.List[T]
}

type node struct {
	nodeType ASTNodeType
	name     string
	params   params[*node]
	body     body[*node]
	value    string
}

const (
	NodeType_VAR        ASTNodeType = "_VAR_"
	NodeType_CHAR       ASTNodeType = "char"
	NodeType_STR        ASTNodeType = "str"
	NodeType_INT        ASTNodeType = "int"
	NodeType_FLOAT      ASTNodeType = "float"
	NodeType_BYTE       ASTNodeType = "byte"
	NodeType_VECTOR     ASTNodeType = "vector"
	NodeType_MAP        ASTNodeType = "map"
	NodeType_TUPLE      ASTNodeType = "tuple"
	NodeType_VAR_ASSIGN ASTNodeType = "varAssign"
	NodeType_BODY       ASTNodeType = "body"
	NodeType_PROGRAM    ASTNodeType = "program"
	NodeType_INT_TYPE   ASTNodeType = "intType"
	NodeType_FLOAT_TYPE ASTNodeType = "floatType"
	NodeType_STR_TYPE   ASTNodeType = "strType"
	NodeType_EQUAL_SIGN ASTNodeType = "="
)

func genNodeWithType(nodeType ASTNodeType) *node {
	n := node{}
	n.nodeType = nodeType
	n.params = params[*node]{}
	n.body = body[*node]{}

	return &n
}

func (self *node) Log() {
	self.logIter(1)
}

func (self *node) logIter(level int) {
	space := ""
	for i := 0; i < level; i++ {
		space += "--"
	}
	//println("nodeName:", self.name)
	println(space, "nodeType:", self.nodeType)
	println(space, "nodeValue:", self.value)
	println(space, "nodeBody")
	self.body.Each(func(self *node) {
		self.logIter(level + 1)
	})
	println(space, "nodeParams")
	self.params.Each(func(self *node) {
		self.logIter(level + 1)
	})
	println("")
}
