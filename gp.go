package gp

import (
	"regexp"

	"github.com/Contra-Culture/gp/reader"
	"github.com/Contra-Culture/gp/store"
)

type Parser func(*reader.BaseSymbolReader) (*ResultNode, bool, error)

func DictTokenParser(token string, dict []string) Parser {
	return func(sr *reader.BaseSymbolReader) (rnode *ResultNode, ok bool, err error) {
		var s store.Symbol
		for _, word := range dict {
			wordRunes := []rune(word)
			for _, expectedRune := range wordRunes {
				s, err = sr.ReadSymbol()
				if err != nil {
					return
				}
				ok = s.Rune == expectedRune
				if !ok {
					return
				}
			}
		}
		rnode = &ResultNode{
			token:    token,
			line:     sr.Frame()[0].Line,
			posStart: sr.Frame()[0].Position, //fix
			posEnd:   sr.Frame()[0].Position,
			literal:  "dumb",
			children: []*ResultNode{},
		}
		return
	}
}
func PatternTokenParser(token string, pattern string) Parser {
	return func(sr *reader.BaseSymbolReader) (rnode *ResultNode, ok bool, err error) {
		ok, err = regexp.MatchReader(pattern, sr)
		if err != nil {
			return
		}
		rnode = &ResultNode{
			token:    token,
			line:     sr.Frame()[0].Line,
			posStart: sr.Frame()[0].Position, //fix
			posEnd:   sr.Frame()[0].Position,
			literal:  "dumb",
			children: []*ResultNode{},
		}
		return
	}
}
func ExactTokenParser(token string, exactTokenValue string) Parser {
	return func(sr *reader.BaseSymbolReader) (rnode *ResultNode, ok bool, err error) {
		expectedRunes := []rune(exactTokenValue)
		var (
			s        store.Symbol
			posStart int
		)
		for _, expectedRune := range expectedRunes {
			s, err = sr.ReadSymbol()
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
			token:    token,
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
	return func(sr *reader.BaseSymbolReader) (rnode *ResultNode, ok bool, err error) {
		return
	}
}

type ParserNode struct {
	optional Optional
	meaning  string
	parser   Parser
	children []*ParserNode
}
type Optional bool

func New(meaning string, optional Optional, parser Parser) *ParserNode {
	return &ParserNode{
		optional: optional,
		meaning:  meaning,
		parser:   parser,
		children: []*ParserNode{},
	}
}
func Seq(meaning string, optional Optional, pns ...*ParserNode) (node *ParserNode) {
	parser := func(reader *reader.BaseSymbolReader) (rn *ResultNode, ok bool, err error) {
		rn = &ResultNode{}
		var childNode *ResultNode
		for _, pn := range pns {
			childNode, ok, err = pn.Parse(reader)
			if err != nil {
				return
			}
			if !ok {
				rn = nil
				return
			}
			rn.children = append(rn.children, childNode)
		}
		return
	}
	node = &ParserNode{
		optional: optional,
		meaning:  meaning,
		parser:   parser,
		children: pns,
	}
	return
}
func Var(meaning string, optional Optional, pns ...*ParserNode) (node *ParserNode) {
	parser := func(reader *reader.BaseSymbolReader) (rn *ResultNode, ok bool, err error) {
		rn = &ResultNode{}
		var childNode *ResultNode
		for _, pn := range pns {
			childNode, ok, err = pn.Parse(reader)
			if err != nil {
				return
			}
			if ok {
				rn.children = append(rn.children, childNode)
				break
			}
		}
		return
	}
	node = &ParserNode{
		optional: optional,
		meaning:  meaning,
		parser:   parser,
		children: pns,
	}
	return
}
func Rep(meaning string, optional Optional, pn *ParserNode) (node *ParserNode) {
	parser := func(reader *reader.BaseSymbolReader) (rn *ResultNode, ok bool, err error) {
		rn = &ResultNode{}
		var childNode *ResultNode
		for {
			childNode, ok, err = pn.Parse(reader)
			if err != nil {
				return
			}
			if !ok {
				return
			}
			rn.children = append(rn.children, childNode)
		}
	}
	node = &ParserNode{
		optional: optional,
		meaning:  meaning,
		parser:   parser,
		children: []*ParserNode{pn},
	}
	return
}
func (pn *ParserNode) Parse(reader *reader.BaseSymbolReader) (result *ResultNode, ok bool, err error) {
	result, ok, err = pn.parser(reader)
	return
}

type ResultNode struct {
	line     int
	posStart int
	posEnd   int
	token    string
	literal  string
	children []*ResultNode
}
