package gp

import (
	"regexp"

	iterator "github.com/Contra-Culture/gp/iterator"
	"github.com/Contra-Culture/gp/store"
)

type Parser func(*iterator.SymbolsIterator) (*ResultNode, bool, error)

func PatternTokenParser(pattern string) Parser {
	return func(iter *iterator.SymbolsIterator) (rnode *ResultNode, ok bool, err error) {
		matched, err := regexp.MatchReader(pattern, iter)
		
	}
}

func ExactTokenParser(exactTokenValue string) Parser {
	return func(iter *iterator.SymbolsIterator) (rnode *ResultNode, ok bool, err error) {
		expectedRunes := []rune(exactTokenValue)
		var (
			s        store.Symbol
			posStart int
		)
		for _, expectedRune := range expectedRunes {
			s, err = iter.Next()
			if err != nil {
				return
			}
			if posStart == 0 {
				posStart = s.Position
			}
			ok = s.Rune == expectedRune
			if !ok {
				return
			}
		}
		rnode = &ResultNode{
			token:    exactTokenValue,
			line:     s.Line,
			posStart: posStart,
			posEnd:   s.Position,
			literal:  exactTokenValue,
			children: nil,
		}
		return
	}
}
func SubStringParser(str string) Parser {
	return func(iter *iterator.SymbolsIterator) (rnode *ResultNode, ok bool, err error) {
		return
	}
}

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
func Many(meaning string, pns ...*ParserNode) (node *ParserNode) {
	return
}

type ResultNode struct {
	token    string
	line     int
	posStart int
	posEnd   int
	literal  string
	children []*ResultNode
}
