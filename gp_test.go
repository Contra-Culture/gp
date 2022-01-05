package gp_test

import (
	"fmt"
	"io"

	. "github.com/Contra-Culture/gp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("gp", func() {
	Describe("parsers", func() {
		Describe("symbol parser", func() {
			It("parses symbol", func() {
				s1 := Symbol('{')
				rs := RuneScanner("{}")
				n, err := s1.Parse(rs)
				Expect(err).NotTo(HaveOccurred())
				Expect(n).NotTo(BeNil())
				Expect(n.Parsed()).To(Equal([]rune{'{'}))
				s2 := Symbol('}')
				n, err = s2.Parse(rs)
				Expect(err).NotTo(HaveOccurred())
				Expect(n).NotTo(BeNil())
				Expect(n.Parsed()).To(Equal([]rune{'}'}))
			})
		})
		Describe("string parser", func() {
			It("parses string", func() {
				str := String("end")
				rs := RuneScanner("end")
				n, err := str.Parse(rs)
				Expect(err).NotTo(HaveOccurred())
				Expect(n).NotTo(BeNil())
				Expect(n.Parsed()).To(Equal([]rune{'e', 'n', 'd'}))
			})
		})
		Describe("sequence parser", func() {
			It("parses sequence", func() {
				seq := Seq(Symbol('<'), Symbol('='))
				rs := RuneScanner("<=")
				n, err := seq.Parse(rs)
				Expect(err).NotTo(HaveOccurred())
				Expect(n).NotTo(BeNil())
				Expect(n.Parsed()).To(Equal([]rune{'<', '='}))
			})
		})
		Describe("repeat parser", func() {
			It("parses repeatable stuff", func() {
				rep := Repeat(Symbol('='))
				rs := RuneScanner("=====")
				n, err := rep.Parse(rs)
				Expect(err).NotTo(HaveOccurred())
				Expect(n.Parsed()).To(Equal([]rune{'=', '=', '=', '=', '='}))
			})
		})
		Describe("variant parser", func() {
			It("parses variants", func() {
				vars := Variant(
					Seq(Symbol('='), Symbol('=')),
					Seq(Symbol('<'), Symbol('=')),
					Seq(Symbol('>'), Symbol('=')),
					Seq(Symbol('!'), Symbol('=')),
					Symbol('<'),
					Symbol('>'),
				)
				operators := []string{
					">",
					"<",
					">=",
					"<=",
					"!=",
					"==",
				}
				for _, op := range operators {
					rs := RuneScanner(op)
					n, err := vars.Parse(rs)
					Expect(err).NotTo(HaveOccurred())
					Expect(n).NotTo(BeNil())
					Expect(n.Parsed()).To(Equal([]rune(op)))
				}
			})
		})
		Describe("optional parser", func() {
			It("parses optional stuff", func() {
				optional := Optional(Variant(Symbol('-'), Symbol('+')))
				tests := map[string][]rune{
					"-": {'-'},
					"+": {'+'},
					"x": nil,
				}
				for t, runes := range tests {
					rs := RuneScanner(t)
					n, err := optional.Parse(rs)
					Expect(err).NotTo(HaveOccurred())
					Expect(n.Parsed()).To(Equal(runes))
				}
			})
		})
		Describe("digit symbol parser", func() {
			It("parses special symbol", func() {
				tests := []string{
					"0",
					"1",
					"2",
					"3",
					"4",
					"5",
					"6",
					"7",
					"8",
					"9",
				}
				sp := Digit()
				for _, t := range tests {
					rs := RuneScanner(t)
					n, err := sp.Parse(rs)
					Expect(err).NotTo(HaveOccurred())
					Expect(n).NotTo(BeNil())
					Expect(n.Parsed()).To(Equal([]rune(t)))
				}
			})
		})
		Describe("any of listed runes parser", func() {
			It("parses special symbol", func() {
				tests := []string{
					"0",
					"1",
					"2",
					"3",
					"4",
					"5",
					"6",
					"7",
					"8",
					"9",
				}
				sp := AnyOneOfRunes('0', '1', '2', '3', '4', '5', '6', '7', '8', '9')
				for _, t := range tests {
					rs := RuneScanner(t)
					n, err := sp.Parse(rs)
					Expect(err).NotTo(HaveOccurred())
					Expect(n).NotTo(BeNil())
					Expect(n.Parsed()).To(Equal([]rune(t)))
				}
			})
		})
		Describe("low alpha symbol parser", func() {
			It("parses low alpha symbol", func() {
				tests := []string{
					"a",
					"b",
					"c",
					"d",
					"e",
					"f",
					"g",
					"h",
					"i",
					"j",
					"k",
					"l",
					"m",
					"n",
					"o",
					"p",
					"q",
					"r",
					"s",
					"t",
					"u",
					"v",
					"w",
					"x",
					"y",
					"z",
				}
				sp := LowAlpha()
				for _, t := range tests {
					rs := RuneScanner(t)
					n, err := sp.Parse(rs)
					Expect(err).NotTo(HaveOccurred())
					Expect(n).NotTo(BeNil())
					Expect(n.Parsed()).To(Equal([]rune(t)))
				}
			})
		})
		Describe("high alpha symbol parser", func() {
			It("parses low alpha symbol", func() {
				tests := []string{
					"A",
					"B",
					"C",
					"D",
					"E",
					"F",
					"G",
					"H",
					"I",
					"J",
					"K",
					"L",
					"M",
					"N",
					"O",
					"P",
					"Q",
					"R",
					"S",
					"T",
					"U",
					"V",
					"W",
					"X",
					"Y",
					"Z",
				}
				sp := HighAlpha()
				for _, t := range tests {
					rs := RuneScanner(t)
					n, err := sp.Parse(rs)
					Expect(err).NotTo(HaveOccurred())
					Expect(n).NotTo(BeNil())
					Expect(n.Parsed()).To(Equal([]rune(t)))
				}
			})
		})
		Describe("alpha symbol parser", func() {
			It("parses low alpha symbol", func() {
				tests := []string{
					"a",
					"b",
					"c",
					"d",
					"e",
					"f",
					"g",
					"h",
					"i",
					"j",
					"k",
					"l",
					"m",
					"n",
					"o",
					"p",
					"q",
					"r",
					"s",
					"t",
					"u",
					"v",
					"w",
					"x",
					"y",
					"z",
					"A",
					"B",
					"C",
					"D",
					"E",
					"F",
					"G",
					"H",
					"I",
					"J",
					"K",
					"L",
					"M",
					"N",
					"O",
					"P",
					"Q",
					"R",
					"S",
					"T",
					"U",
					"V",
					"W",
					"X",
					"Y",
					"Z",
				}
				sp := Alpha()
				for _, t := range tests {
					rs := RuneScanner(t)
					n, err := sp.Parse(rs)
					Expect(err).NotTo(HaveOccurred())
					Expect(n).NotTo(BeNil())
					Expect(n.Parsed()).To(Equal([]rune(t)))
				}
			})
		})
	})
})

type runeScanner struct {
	runes  []rune
	cursor int
}

func RuneScanner(s string) io.RuneScanner {
	return &runeScanner{
		runes:  []rune(s),
		cursor: -1,
	}
}
func (s *runeScanner) ReadRune() (r rune, l int, err error) {
	if s.cursor > len(s.runes)-2 {
		err = io.ErrUnexpectedEOF
		return
	}
	s.cursor++
	r = s.runes[s.cursor]
	l = len([]byte(string(r)))
	return
}
func (s *runeScanner) UnreadRune() (err error) {
	if s.cursor >= 0 {
		s.cursor--
		return
	}
	return fmt.Errorf("can't unread rune")
}
