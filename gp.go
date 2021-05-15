package gp

import (
	"strings"

	"github.com/Contra-Culture/gp/store"
)

type (
	ParserMaker func() Parser
	Parser      interface {
		Test(s *store.Symbol) (rnode *RecSeq, ok bool, done bool, err error)
	}
	ParserNode struct {
		meaning  string
		parser   Parser
		children []*ParserNode
	}
	LinesRange struct {
		First, Last int
	}
	exactTokenParser struct {
		expected string
		acc      []*store.Symbol
	}
	untilTokenParser struct {
		main  Parser
		final Parser
	}
	Context struct {
		dict Dict
	}
	Dict          []string
	

	RecSeq struct { // regognised sequence
		LinesRange LinesRange
		PosStart   int
		PosEnd     int
		Token      string
		Literal    string
		Children   []*RecSeq
	}
)

const (
	RUNE_TOKEN   = "rune"
	PSEUDO_TOKEN = "<pseudo>"
	NO_TOKEN     = ""
)

var (
	NO_MEANING = []string{""}
)

func New(parser Parser, meaning ...string) *ParserNode {
	if meaning == nil {
		meaning = NO_MEANING
	}
	return &ParserNode{
		parser:   parser,
		meaning:  strings.Join(meaning, "\n"),
		children: []*ParserNode{},
	}
}
func ExactTokenParserMaker(expected string) (maker ParserMaker) {
	return func() (parser Parser) {
		parser = &exactTokenParser{
			expected: expected,
			acc:      []*store.Symbol{},
		}
		return
	}
}
func (p *exactTokenParser) Test(s *store.Symbol) (rnode *RecSeq, ok bool, done bool, err error) {

	return
}

func UntilTokenParserMaker(main ParserMaker, final ParserMaker) (maker ParserMaker) {
	maker = func() (parser Parser) {
		m := main()
		f := final()
		parser = &untilTokenParser{
			main:  m,
			final: f,
		}
		return
	}
	return
}
func (p *untilTokenParser) Test(s *store.Symbol) (rnode *RecSeq, ok bool, done bool, err error) {

	return
}
