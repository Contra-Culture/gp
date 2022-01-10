package gp_test

import (
	. "github.com/Contra-Culture/gp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("gp", func() {
	Describe("parsers", func() {
		Describe("symbol parser", func() {
			It("parses symbol", func() {
				s1 := Symbol('{')
				rs := NewRuneScanner("{}")
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
				rs := NewRuneScanner("end")
				n, err := str.Parse(rs)
				Expect(err).NotTo(HaveOccurred())
				Expect(n).NotTo(BeNil())
				Expect(n.Parsed()).To(Equal([]rune{'e', 'n', 'd'}))
			})
		})
		Describe("sequence parser", func() {
			It("parses sequence", func() {
				seq := Seq(Symbol('<'), Symbol('='))
				rs := NewRuneScanner("<=")
				n, err := seq.Parse(rs)
				Expect(err).NotTo(HaveOccurred())
				Expect(n).NotTo(BeNil())
				Expect(n.Parsed()).To(Equal([]rune{'<', '='}))
			})
		})
		Describe("repeat parser", func() {
			It("parses repeatable stuff", func() {
				rep := Repeat(Symbol('='))
				rs := NewRuneScanner("=====")
				n, err := rep.Parse(rs)
				Expect(err).NotTo(HaveOccurred())
				Expect(n.Parsed()).To(Equal([]rune{'=', '=', '=', '=', '='}))
			})
		})
		Describe("variant parser", func() {
			It("parses variants", func() {
				vars := Variant(
					String("!="),
					String("=="),
					String("<="),
					String(">="),
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
					rs := NewRuneScanner(op)
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
				}
				for t, runes := range tests {
					rs := NewRuneScanner(t)
					n, err := optional.Parse(rs)
					Expect(err).NotTo(HaveOccurred())
					Expect(n.Parsed()).To(Equal(runes))
				}
				n, err := optional.Parse(NewRuneScanner("x"))
				Expect(err).NotTo(HaveOccurred())
				Expect(n).To(BeNil())
			})
		})
		Describe("range parser", func() {
			Context("when Digits range", func() {
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
					sp := Range(Digits())
					for _, t := range tests {
						rs := NewRuneScanner(t)
						n, err := sp.Parse(rs)
						Expect(err).NotTo(HaveOccurred())
						Expect(n).NotTo(BeNil())
						Expect(n.Parsed()).To(Equal([]rune(t)))
					}
				})
			})
			Context("when low alphabet symbols range", func() {
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
					sp := Range(LowASCIIAlphabet())
					for _, t := range tests {
						rs := NewRuneScanner(t)
						n, err := sp.Parse(rs)
						Expect(err).NotTo(HaveOccurred())
						Expect(n).NotTo(BeNil())
						Expect(n.Parsed()).To(Equal([]rune(t)))
					}
				})
			})
			Context("when high alphabet symbols range", func() {
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
					sp := Range(HighASCIIAlphabet())
					for _, t := range tests {
						rs := NewRuneScanner(t)
						n, err := sp.Parse(rs)
						Expect(err).NotTo(HaveOccurred())
						Expect(n).NotTo(BeNil())
						Expect(n.Parsed()).To(Equal([]rune(t)))
					}
				})
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
					rs := NewRuneScanner(t)
					n, err := sp.Parse(rs)
					Expect(err).NotTo(HaveOccurred())
					Expect(n).NotTo(BeNil())
					Expect(n.Parsed()).To(Equal([]rune(t)))
				}
			})
		})
		Describe("proxy parser", func() {
			It("parses using predefined parser", func() {
				p, err := New(func(u *UnivCfgr) {
					u.Define("alpha", Range(LowASCIIAlphabet()))
					u.Top(u.Get("alpha"))
				})
				Expect(err).NotTo(HaveOccurred())
				Expect(p).NotTo(BeNil())
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
				for _, t := range tests {
					rs := NewRuneScanner(t)
					n, err := p.Parse(rs)
					Expect(err).NotTo(HaveOccurred())
					Expect(n).NotTo(BeNil())
					Expect(n.Parsed()).To(Equal([]rune(t)))
				}
			})
		})
	})
})
