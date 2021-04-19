package gp_test

import (
	"strings"

	gp "github.com/Contra-Culture/gp"
	"github.com/Contra-Culture/gp/reader"
	"github.com/Contra-Culture/gp/store"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("gp", func() {
	Context("creation", func() {
		Describe("New()", func() {
			It("returns parser node", func() {
				pn := gp.New("test purpose", gp.ExactTokenParser("test", "test"))
				Expect(pn).NotTo(BeNil())
			})
		})
		Describe("Opt()", func() {
			It("returns parser node", func() {

			})
		})
		Describe("Var()", func() {
			It("returns parser node", func() {
				def := gp.New("method definition", gp.ExactTokenParser("keyword", "def"))
				fnc := gp.New("function definition", gp.ExactTokenParser("keyword", "func"))
				pn := gp.Var("test purpose", def, fnc)
				Expect(pn).NotTo(BeNil())
			})
		})
		Describe("Seq()", func() {
			It("returns parser node", func() {
				fnc := gp.New("function definition keyword", gp.ExactTokenParser("keyword", "func"))
				identifierParser, err := gp.PatternTokenParser("indentifier", "/A.?/Z")
				Expect(err).NotTo(HaveOccurred())
				name := gp.New("function name", identifierParser)
				openBracket := gp.New("open bracket", gp.ExactTokenParser("bracket", "("))
				closedBracket := gp.New("closed bracket", gp.ExactTokenParser("bracket", ")"))
				pn := gp.Seq("func definition", fnc, name, openBracket, closedBracket)
				Expect(pn).NotTo(BeNil())
			})
		})
		Describe("Rep()", func() {
			It("returns parser node", func() {
				fnc := gp.New("function definition keyword", gp.ExactTokenParser("keyword", "func"))
				pn := gp.Rep("rep", fnc)
				Expect(pn).NotTo(BeNil())
			})
		})
	})
	Context("parsing", func() {
		Describe(".Parse()", func() {
			Context("when custom node", func() {
				Context("when matches", func() {
					It("returns result node and ok", func() {
						sr := strings.NewReader("func myFunc() {}")
						s, err := store.New(sr)
						Expect(err).NotTo(HaveOccurred())
						r := reader.New(s, 0)
						funcKwd := gp.New("function definition", gp.ExactTokenParser("keyword", "func"))
						result, ok, err := funcKwd.Parse(r)
						Expect(err).NotTo(HaveOccurred())
						Expect(ok).To(BeTrue())
						Expect(result.Lines).To(Equal([]int{1, 1}))
						Expect(result.PosStart).To(Equal(1))
						Expect(result.PosEnd).To(Equal(4))
						Expect(result.Token).To(Equal("keyword"))
						Expect(result.Literal).To(Equal("func"))
						Expect(result.Children).To(BeEmpty())
					})
				})
				Context("when not matches", func() {
					It("returns not ok", func() {
						sr := strings.NewReader("myFunc() {}")
						s, err := store.New(sr)
						Expect(err).NotTo(HaveOccurred())
						r := reader.New(s, 0)
						funcKwd := gp.New("function definition", gp.ExactTokenParser("keyword", "func"))
						result, ok, err := funcKwd.Parse(r)
						Expect(err).NotTo(HaveOccurred())
						Expect(ok).To(BeFalse())
						Expect(result).To(BeNil())
					})
				})
			})
			Context("when sequence node", func() {
				Context("when matches", func() {
					It("returns result node and ok", func() {
						sr := strings.NewReader("func myFunc() { ignore }")
						s, err := store.New(sr)
						Expect(err).NotTo(HaveOccurred())
						r := reader.New(s, 0)
						funcToken := gp.New("function definition keyword", gp.ExactTokenParser("keyword", "func"))
						identifierParser, err := gp.PatternTokenParser("identifier", "^([a-zA-Z]+)")
						Expect(err).NotTo(HaveOccurred())
						identifierToken := gp.New("function name", identifierParser)
						spaceToken := gp.New("space", gp.ExactTokenParser("space", " "))
						openingBracketToken := gp.New("opening bracket", gp.ExactTokenParser("opening bracket", "("))
						closingBracketToken := gp.New("closing bracket", gp.ExactTokenParser("closing bracket", ")"))
						openingCurlyBracketToken := gp.New("opening curly bracket", gp.ExactTokenParser("opening curly bracket", "{"))
						closingCurlyBracketToken := gp.New("closing curly bracket", gp.ExactTokenParser("closing curly bracket", "}"))
						funcDef := gp.Seq("function definition",
							funcToken,
							spaceToken,
							identifierToken,
							openingBracketToken,
							closingBracketToken,
							spaceToken,
							openingCurlyBracketToken,
							spaceToken,
							identifierToken,
							spaceToken,
							closingCurlyBracketToken,
						)
						result, ok, err := funcDef.Parse(r)
						Expect(err).NotTo(HaveOccurred())
						Expect(ok).To(BeTrue())
						Expect(result.Lines).To(Equal([]int{1, 1}))
						Expect(result.PosStart).To(Equal(1))
						Expect(result.PosEnd).To(Equal(24))
						Expect(result.Literal).To(Equal("func myFunc() { ignore }"))
						Expect(result.Token).To(Equal("expression"))
						Expect(result.Children).To(HaveLen(11))
						Expect(result.Children[0]).To(Equal(&gp.ResultNode{
							Lines:    []int{1, 1},
							PosStart: 1,
							PosEnd:   4,
							Token:    "keyword",
							Literal:  "func",
							Children: nil,
						}))
						Expect(result.Children[1]).To(Equal(&gp.ResultNode{
							Lines:    []int{1, 1},
							PosStart: 5,
							PosEnd:   5,
							Token:    "space",
							Literal:  " ",
							Children: nil,
						}))
						Expect(result.Children[2]).To(Equal(&gp.ResultNode{
							Lines:    []int{1, 1},
							PosStart: 6,
							PosEnd:   11,
							Token:    "identifier",
							Literal:  "myFunc",
							Children: nil,
						}))
						Expect(result.Children[3]).To(Equal(&gp.ResultNode{
							Lines:    []int{1, 1},
							PosStart: 12,
							PosEnd:   12,
							Token:    "opening bracket",
							Literal:  "(",
							Children: nil,
						}))
						Expect(result.Children[4]).To(Equal(&gp.ResultNode{
							Lines:    []int{1, 1},
							PosStart: 13,
							PosEnd:   13,
							Token:    "closing bracket",
							Literal:  ")",
							Children: nil,
						}))
						Expect(result.Children[5]).To(Equal(&gp.ResultNode{
							Lines:    []int{1, 1},
							PosStart: 14,
							PosEnd:   14,
							Token:    "space",
							Literal:  " ",
							Children: nil,
						}))
						Expect(result.Children[6]).To(Equal(&gp.ResultNode{
							Lines:    []int{1, 1},
							PosStart: 15,
							PosEnd:   15,
							Token:    "opening curly bracket",
							Literal:  "{",
							Children: nil,
						}))
						Expect(result.Children[7]).To(Equal(&gp.ResultNode{
							Lines:    []int{1, 1},
							PosStart: 16,
							PosEnd:   16,
							Token:    "space",
							Literal:  " ",
							Children: nil,
						}))
						Expect(result.Children[8]).To(Equal(&gp.ResultNode{
							Lines:    []int{1, 1},
							PosStart: 17,
							PosEnd:   22,
							Token:    "identifier",
							Literal:  "ignore",
							Children: nil,
						}))
						Expect(result.Children[9]).To(Equal(&gp.ResultNode{
							Lines:    []int{1, 1},
							PosStart: 23,
							PosEnd:   23,
							Token:    "space",
							Literal:  " ",
							Children: nil,
						}))
						Expect(result.Children[10]).To(Equal(&gp.ResultNode{
							Lines:    []int{1, 1},
							PosStart: 24,
							PosEnd:   24,
							Token:    "closing curly bracket",
							Literal:  "}",
							Children: nil,
						}))
					})
				})
				Context("when not matches", func() {
					It("returns not ok", func() {

					})
				})
			})
			Context("when repeatable node", func() {
				Context("when matches", func() {
					It("returns result node and ok", func() {
						sr := strings.NewReader("func func func garbage")
						s, err := store.New(sr)
						Expect(err).NotTo(HaveOccurred())
						r := reader.New(s, 0)
						funcTokenParser := gp.New("rep func keywords", gp.ExactTokenParser("keyword", "func"))
						spaceTokenParser := gp.New("space", gp.ExactTokenParser("space", " "))
						rep := gp.Rep("repeatable", gp.Seq("func def", funcTokenParser, spaceTokenParser))
						result, ok, err := rep.Parse(r)
						Expect(err).NotTo(HaveOccurred())
						Expect(ok).To(BeTrue())
						Expect(result.Lines).To(Equal([]int{1, 1}))
						Expect(result.PosStart).To(Equal(1))
						Expect(result.PosEnd).To(Equal(15))
						Expect(result.Token).To(Equal(""))
						Expect(result.Literal).To(Equal("func func func "))
						Expect(result.Children).To(HaveLen(3))
						Expect(result.Children[0]).To(Equal(&gp.ResultNode{
							Lines:    []int{1, 1},
							PosStart: 1,
							PosEnd:   5,
							Token:    "expression",
							Literal:  "func ",
							Children: []*gp.ResultNode{
								{
									Lines:    []int{1, 1},
									PosStart: 1,
									PosEnd:   4,
									Token:    "keyword",
									Literal:  "func",
									Children: nil,
								},
								{
									Lines:    []int{1, 1},
									PosStart: 5,
									PosEnd:   5,
									Token:    "space",
									Literal:  " ",
									Children: nil,
								},
							},
						}))
						Expect(result.Children[1]).To(Equal(&gp.ResultNode{
							Lines:    []int{1, 1},
							PosStart: 6,
							PosEnd:   10,
							Token:    "expression",
							Literal:  "func ",
							Children: []*gp.ResultNode{
								{
									Lines:    []int{1, 1},
									PosStart: 6,
									PosEnd:   9,
									Token:    "keyword",
									Literal:  "func",
									Children: nil,
								},
								{
									Lines:    []int{1, 1},
									PosStart: 10,
									PosEnd:   10,
									Token:    "space",
									Literal:  " ",
									Children: nil,
								},
							},
						}))
						Expect(result.Children[2]).To(Equal(&gp.ResultNode{
							Lines:    []int{1, 1},
							PosStart: 11,
							PosEnd:   15,
							Token:    "expression",
							Literal:  "func ",
							Children: []*gp.ResultNode{
								{
									Lines:    []int{1, 1},
									PosStart: 11,
									PosEnd:   14,
									Token:    "keyword",
									Literal:  "func",
									Children: nil,
								},
								{
									Lines:    []int{1, 1},
									PosStart: 15,
									PosEnd:   15,
									Token:    "space",
									Literal:  " ",
									Children: nil,
								},
							},
						}))
					})
				})
				Context("when not matches", func() {
					It("returns not ok", func() {

					})
				})
			})
			Context("when variant node", func() {
				Context("when matches", func() {
					It("returns result node and ok", func() {

					})
				})
				Context("when not matches", func() {
					It("returns not ok", func() {

					})
				})
			})
		})
	})
	Describe("helpers", func() {
		Describe("ExactTokenParser()", func() {
			It("returns parser", func() {

			})
		})
		Describe("PatternTokenParser()", func() {
			It("returns parser", func() {

			})
		})
		Describe("DictTokenParser()", func() {
			It("returns parser", func() {

			})
		})
	})
})
