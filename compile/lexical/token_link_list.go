package lexical

type Tokens struct {
	head   *TokenNode
	tail   *TokenNode
	length int32
}

type TokenNode struct {
	token *LexicalToken
	next  *TokenNode
}

func NewTokens() *Tokens {
	var p = new(Tokens)
	return p
}

func (self *Tokens) add(token *LexicalToken) {
	var node = new(TokenNode)
	node.token = token
	node.next = nil

	if self.tail != nil {
		self.tail.next = node
	} else {
		self.head = node
	}
	self.tail = node
	self.length += 1
}

func (self *Tokens) Log() {
	var node = self.head
	for node != nil {
		node.token.log()
		node = node.next
	}
}
