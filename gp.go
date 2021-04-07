package gp

import "io"

type Parser func(rr io.RuneReader) (*ResultNode, error)

func RuneParser(expectedRune rune) Parser {
	return func(rr io.RuneReader) (rnode *ResultNode, err error) {
		r, _, err := rr.ReadRune()
		if err != nil {
			return
		}
		if r == expectedRune {
		}
		return
	}
}

type Token int

type ParserNode struct {
	meaning  string
	parser   Parser
	children []*ParserNode
}

func New(meaning string, parser Parser) (node *ParserNode) {
	node.meaning = meaning
	node.parser = parser
	return
}
func Maybe(meaning string, pns ...*ParserNode) (node *ParserNode) {
	node.meaning = meaning
	node.children = pns
	return
}
func And(meaning string, pns ...*ParserNode) (node *ParserNode) {
	node.meaning = meaning
	node.children = pns
	return
}
func Xor(meaning string, pns ...*ParserNode) (node *ParserNode) {
	node.meaning = meaning
	node.children = pns
	return
}

type ResultNode struct {
	token    Token
	line     int
	posStart int
	posEnd   int
	literal  string
	children []*ResultNode
}
