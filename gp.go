package gp

import (
	"errors"

	"github.com/Contra-Culture/gp/reader"
	"github.com/Contra-Culture/gp/store"
	"github.com/Contra-Culture/gp/tester"
)

type (
	sequence struct {
		title   string
		symbols store.Symbols
	}
	Frame struct {
		sequences []sequence
		title     string
	}
	Token       string
	ParserMaker func() Parser
	Parser      interface {
		Parse(rdr *reader.Reader) (frame *Frame)
	}
	Breaker      func() bool
	BreakerMaker func() Breaker
	ParserNode   interface {
		Parse(rdr *reader.Reader, result Parsed) (err error)
	}
	basicNode struct {
		meaning    string
		makeParser ParserMaker
	}
	seqNode struct {
		meaning  string
		children []ParserNode
	}
	orNode struct {
		meaning  string
		variants []ParserNode
	}
	repNode struct {
		meaning     string
		repeatable  ParserNode
		makeBreaker BreakerMaker
	}
	noneNode struct {
		meaning string
	}
	exactParser struct {
		token    Token
		expected []rune
	}
	fixedParser struct {
		token Token
		len   int
	}
	continuousParser struct {
		token        Token
		testerMakers []tester.TesterMaker
	}
	Parsed struct {
		Frame    Frame
		Children []Parsed
	}
	dumbTesterT struct{}
)

const (
	CANT_PARSE = "can't parse"
	BEGIN      = "BEGIN:"
)

var (
	dumbTester *dumbTesterT
)

func (d *dumbTesterT) Test(_ rune) bool {
	return true
}
func FixedParserMaker(token Token, len int) (maker ParserMaker) {
	maker = func() Parser {
		return &fixedParser{
			token,
			len,
		}
	}
	return
}
func (p *fixedParser) Parse(rdr *reader.Reader) *Frame {
	var err error
	for i := 0; i < p.len; i++ {
		_, err = rdr.ReadSymbol()
		if err != nil {
			panic(err)
		}
	}
	return &Frame{
		title: string(p.token),
		sequences: []sequence{
			{
				title:   string(p.token),
				symbols: rdr.UncommittedRead(),
			},
		},
	}
}
func ExactParserMaker(token Token, expected string) (maker ParserMaker) {
	runes := []rune(expected)
	maker = func() (parser Parser) {
		parser = &exactParser{
			token:    token,
			expected: runes,
		}
		return
	}
	return
}
func (p *exactParser) Parse(rdr *reader.Reader) *Frame {
	var err error
	var s store.Symbol
	var ok = false
	for _, r := range p.expected {
		s, err = rdr.ReadSymbol()
		if err != nil {
			panic(err)
		}
		ok = s.Rune == r
		if !ok {
			return nil
		}
	}
	return &Frame{
		title: string(p.token),
		sequences: []sequence{
			{
				title:   string(p.token),
				symbols: rdr.UncommittedRead(),
			},
		},
	}
}
func ContinuousParserMaker(token Token, testerMakers ...tester.TesterMaker) (maker ParserMaker) {
	maker = func() (parser Parser) {
		parser = &continuousParser{
			token,
			testerMakers,
		}
		return
	}
	return
}
func (p *continuousParser) Parse(rdr *reader.Reader) *Frame {
	var tester tester.Tester
	var s store.Symbol
	var err error
	var smap []sequence
	var symbols store.Symbols
	var title string
testers:
	for _, testerMaker := range p.testerMakers {
		title = testerMaker.Title()
		tester = testerMaker.Tester()
		symbols = store.Symbols{}
		for {
			s, err = rdr.ReadSymbol()
			if err != nil {
				panic(err)
			}
			ok, cntn := tester.Test(s.Rune)
			if ok {
				symbols = append(symbols, s)
				if !cntn {
					smap = append(smap, sequence{
						title,
						symbols,
					})
					continue testers
				}
				continue
			}
		}
	}
	return &Frame{
		sequences: smap,
		title:     string(p.token),
	}
}
func BasicNode(meaning string, makeParser ParserMaker) ParserNode {
	return &basicNode{
		meaning,
		makeParser,
	}
}
func (n *basicNode) Parse(rdr *reader.Reader, result Parsed) (err error) {
	p := n.makeParser()
	frame := p.Parse(rdr)
	if frame != nil {
		result.Children = append(result.Children, Parsed{
			Frame:    *frame,
			Children: make([]Parsed, 8),
		})
		return
	}
	err = errors.New(CANT_PARSE)
	return
}
func SeqNode(meaning string, children ...ParserNode) ParserNode {
	return &seqNode{
		meaning,
		children,
	}
}
func BeginNode(children ...ParserNode) ParserNode {
	return &seqNode{
		meaning:  BEGIN,
		children: children,
	}
}
func (n *seqNode) Parse(rdr *reader.Reader, result Parsed) (err error) {
	parsed := Parsed{
		Children: make([]Parsed, len(n.children)),
	}
	var newRdr *reader.Reader
	for _, n := range n.children {
		newRdr = rdr.Continuation(reader.NOLIMIT)
		err = n.Parse(newRdr, parsed)
		if err != nil {
			return
		}
	}
	parsed.Frame = Frame{
		sequences: []sequence{
			{
				title:   n.meaning,
				symbols: rdr.UncommittedRead(),
			},
		},
	}
	result.Children = append(result.Children, parsed)
	return
}
func OrNode(meaning string, variants ...ParserNode) ParserNode {
	return &orNode{
		meaning,
		variants,
	}
}
func (nd *orNode) Parse(rdr *reader.Reader, result Parsed) (err error) {
	var newRdr *reader.Reader
	var parsed = Parsed{
		Children: make([]Parsed, 1),
	}
	for _, n := range nd.variants {
		newRdr = rdr.Continuation(reader.NOLIMIT)
		err = n.Parse(newRdr, parsed)
		if err != nil {
			continue
		}
		newRdr.CommitToParent()
		break
	}
	symbols := rdr.UncommittedRead()
	if len(symbols) > 0 {
		parsed.Frame = Frame{
			title: nd.meaning,
			sequences: []sequence{
				{
					title:   nd.meaning,
					symbols: symbols,
				},
			},
		}
		result.Children = append(result.Children, parsed)
		return
	}
	err = errors.New(CANT_PARSE)
	return
}
func RepNode(meaning string, makeBreaker BreakerMaker, repeatable ParserNode) ParserNode {
	return &repNode{
		meaning,
		repeatable,
		makeBreaker,
	}
}
func (rn *repNode) Parse(rdr *reader.Reader, result Parsed) (err error) {
	var newRdr *reader.Reader
	brk := rn.makeBreaker()
	parsed := Parsed{
		Children: make([]Parsed, 8),
	}
	for !brk() {
		newRdr = rdr.Continuation(reader.NOLIMIT)
		err = rn.repeatable.Parse(newRdr, parsed)
		if err != nil {
			return
		}
		newRdr.CommitToParent()
	}
	parsed.Frame = Frame{
		title: rn.meaning,
		sequences: []sequence{
			{
				title:   rn.meaning,
				symbols: rdr.UncommittedRead(),
			},
		},
	}
	result.Children = append(result.Children, parsed)
	return
}
func NoneNode(meaning string) ParserNode {
	return &noneNode{meaning}
}
func (nn *noneNode) Parse(rdr *reader.Reader, result Parsed) (err error) {
	result.Children = append(result.Children, Parsed{
		Frame: Frame{
			title:     nn.meaning,
			sequences: []sequence{},
		},
	})
	return
}
