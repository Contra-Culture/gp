package gp

import (
	"fmt"
	"regexp"
	"strings"

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
			Token:    token,
			Lines:    []int{sr.Frame()[0].Line, sr.Frame()[0].Line},
			PosStart: sr.Frame()[0].Position, //fix
			PosEnd:   sr.Frame()[0].Position,
			Literal:  "dumb",
			Children: []*ResultNode{},
		}
		return
	}
}
func PatternTokenParser(token string, pattern string) (parser Parser, err error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return
	}
	parser = func(sr *reader.BaseSymbolReader) (rnode *ResultNode, ok bool, err error) {
		forSearch := sr.Continuation()
		loc := re.FindReaderIndex(forSearch)
		ok = loc != nil
		if !ok {
			return
		}
		var literal strings.Builder
		var r rune
		for i := loc[0]; i < loc[1]; i++ {
			r, _, err = sr.ReadRune()
			if err != nil {
				return
			}
			literal.WriteRune(r)
		}
		rnode = &ResultNode{
			Token:    token,
			Lines:    []int{sr.Frame()[0].Line, sr.Frame()[len(sr.Frame())-1].Line},
			PosStart: sr.Frame()[0].Position,
			PosEnd:   sr.Frame()[len(sr.Frame())-1].Position,
			Literal:  literal.String(),
		}
		return
	}
	return
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
				fmt.Printf("\n\tExactTokenParser: error %s\n", err.Error())
				return
			}
			fmt.Printf("\n\tExactTokenParser: expected rune %s , got: %s, %#v\n", string(expectedRune), string(s.Rune), s.Rune)
			if posStart == 0 {
				posStart = s.Position
			}
			ok = s.Rune == expectedRune
			if !ok {
				return
			}
		}
		rnode = &ResultNode{
			Token:    token,
			Lines:    []int{s.Line, s.Line},
			PosStart: posStart,
			PosEnd:   s.Position,
			Literal:  exactTokenValue,
			Children: nil,
		}
		return
	}
}

type ParserNode struct {
	meaning  string
	parser   Parser
	children []*ParserNode
}

func New(meaning string, parser Parser) *ParserNode {
	return &ParserNode{
		meaning:  meaning,
		parser:   parser,
		children: []*ParserNode{},
	}
}
func Opt(meaning string, pn *ParserNode) *ParserNode {
	parser := func(reader *reader.BaseSymbolReader) (result *ResultNode, ok bool, err error) {
		result, ok, err = pn.Parse(reader)
		if err != nil {
			return
		}
		if !ok {
			return
		}
		return
	}
	return &ParserNode{
		meaning:  meaning,
		parser:   parser,
		children: []*ParserNode{},
	}
}
func Seq(meaning string, pns ...*ParserNode) (node *ParserNode) {
	parser := func(reader *reader.BaseSymbolReader) (rn *ResultNode, ok bool, err error) {
		rn = &ResultNode{}
		var childNode *ResultNode
		var literal strings.Builder
		for _, pn := range pns {
			fmt.Printf("\n\tparses node: %s\n", pn.meaning)
			childNode, ok, err = pn.Parse(reader)
			if err != nil {
				return
			}
			if !ok {
				rn = nil
				return
			}
			literal.WriteString(childNode.Literal)
			rn.Children = append(rn.Children, childNode)
			reader = reader.Continuation()
		}
		fmt.Printf("\nresult node: %#v\n", rn.Children)
		rn.Lines = rn.Children[0].Lines
		rn.PosStart = rn.Children[0].PosStart
		last := len(rn.Children) - 1
		rn.PosEnd = rn.Children[last].PosEnd
		rn.Literal = literal.String()
		rn.Token = "expression"
		return
	}
	node = &ParserNode{
		meaning:  meaning,
		parser:   parser,
		children: pns,
	}
	return
}
func Var(meaning string, pns ...*ParserNode) (node *ParserNode) {
	parser := func(reader *reader.BaseSymbolReader) (rn *ResultNode, ok bool, err error) {
		rn = &ResultNode{}
		var childNode *ResultNode
		for _, pn := range pns {
			childNode, ok, err = pn.Parse(reader)
			if err != nil {
				return
			}
			if ok {
				rn.Children = append(rn.Children, childNode)
				break
			}
		}
		return
	}
	node = &ParserNode{
		meaning:  meaning,
		parser:   parser,
		children: pns,
	}
	return
}
func Rep(meaning string, pn *ParserNode) (node *ParserNode) {
	parser := func(reader *reader.BaseSymbolReader) (rn *ResultNode, ok bool, err error) {
		rn = &ResultNode{}
		var childNode *ResultNode
		var literal strings.Builder
		for {
			childNode, ok, err = pn.Parse(reader)
			fmt.Printf("\n\tRep(): %#v, ok: %#v, err: %#v\n", childNode, ok, err)
			if err != nil {
				return
			}
			if !ok {
				break
			}
			rn.Children = append(rn.Children, childNode)
			literal.WriteString(childNode.Literal)
			_, err = reader.ReadSymbol()
			if err != nil {
				return
			}
		}
		ok = len(rn.Children) > 0
		rn.Lines = []int{rn.Children[0].Lines[0], rn.Children[len(rn.Children)-1].Lines[1]}
		rn.Literal = literal.String()
		rn.PosStart = rn.Children[0].PosStart
		rn.PosEnd = rn.Children[len(rn.Children)-1].PosEnd
		return
	}
	node = &ParserNode{
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
	Lines    []int // pair: first line number and last line number
	PosStart int
	PosEnd   int
	Token    string
	Literal  string
	Children []*ResultNode
}
