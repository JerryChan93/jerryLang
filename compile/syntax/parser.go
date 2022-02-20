package syntax

import (
	"jerryLang/compile/lexical"
)

func Parse(tokens *lexical.Tokens[*lexical.LexicalToken]) {
	//tokens.Each(func(item *lexical.LexicalToken) {
	//	item.Log()
	//})
	ast := genAST(tokens)
	ast.Log()
}

func genAST(tokens *lexical.Tokens[*lexical.LexicalToken]) *node {
	program := genNodeWithType(NodeType_PROGRAM)

	for p := tokens.Pop(); p != nil; p = tokens.Pop() {
		if p.TokenType() == lexical.Token_INT {
			node := genNodeWithType(NodeType_INT)
			node.value = p.TokenValue()
			return node
		}
		if p.TokenType() == lexical.Token_VAR_ASSIGN {
			node := genNodeWithType(NodeType_VAR_ASSIGN)
			node.value = p.TokenValue()
			return node
		}
		if p.TokenType() == lexical.Token_INT_TYPE {
			nodeType := ASTNodeType(p.TokenType())
			node := genNodeWithType(nodeType)
			node.value = p.TokenValue()
			return node
		}
		if p.TokenType() == lexical.Token_EQUAL_SIGN {
			node := genNodeWithType(NodeType_EQUAL_SIGN)
			node.value = p.TokenValue()
			return node
		}
		if p.TokenType() == lexical.Token_INT {
			nodeType := ASTNodeType(p.TokenType())
			node := genNodeWithType(nodeType)
			node.value = p.TokenValue()
			return node
		}
		if p.TokenType() == lexical.Token_VAR {
			n1 := genNodeWithType(NodeType_VAR)
			varAssign := genAST(tokens)
			n1.body.Push(varAssign)
			varType := genAST(tokens)
			n1.body.Push(varType)
			equal := genAST(tokens)
			n1.body.Push(equal)
			val := genAST(tokens)
			n1.body.Push(val)

			program.body.Push(n1)
		}
		p.Log()
	}

	return program
}
