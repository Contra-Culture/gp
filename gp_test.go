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
				pn := gp.Req("test purpose", gp.ExactTokenParser("test", "test"))
				Expect(pn).NotTo(BeNil())
			})
		})
		Describe("Var()", func() {
			It("returns parser node", func() {
				def := gp.Req("method definition", gp.ExactTokenParser("keyword", "def"))
				fnc := gp.Req("function definition", gp.ExactTokenParser("keyword", "func"))
				pn := gp.Var("test purpose", def, fnc)
				Expect(pn).NotTo(BeNil())
			})
		})
		Describe("Seq()", func() {
			It("returns parser node", func() {
				fnc := gp.Req("function definition keyword", gp.ExactTokenParser("keyword", "func"))
				identifierParser, err := gp.PatternTokenParser("indentifier", "/A.?/Z")
				Expect(err).NotTo(HaveOccurred())
				name := gp.Req("function name", identifierParser)
				openBracket := gp.Req("open bracket", gp.ExactTokenParser("bracket", "("))
				closedBracket := gp.Req("closed bracket", gp.ExactTokenParser("bracket", ")"))
				pn := gp.Seq("func definition", fnc, name, openBracket, closedBracket)
				Expect(pn).NotTo(BeNil())
			})
		})
		Describe("Rep()", func() {
			It("returns parser node", func() {
				fnc := gp.Req("function definition keyword", gp.ExactTokenParser("keyword", "func"))
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
						funcKwd := gp.Req("function definition", gp.ExactTokenParser("keyword", "func"))
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
						funcKwd := gp.Req("function definition", gp.ExactTokenParser("keyword", "func"))
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
						funcToken := gp.Req("function definition keyword", gp.ExactTokenParser("keyword", "func"))
						identifierParser, err := gp.PatternTokenParser("identifier", "^([\\w]+)")
						Expect(err).NotTo(HaveOccurred())
						identifierToken := gp.Req("function name", identifierParser)
						spaceToken := gp.Req("space", gp.ExactTokenParser("space", " "))
						openingBracketToken := gp.Req("opening bracket", gp.ExactTokenParser("opening bracket", "("))
						closingBracketToken := gp.Req("closing bracket", gp.ExactTokenParser("closing bracket", ")"))
						openingCurlyBracketToken := gp.Req("opening curly bracket", gp.ExactTokenParser("opening curly bracket", "{"))
						closingCurlyBracketToken := gp.Req("closing curly bracket", gp.ExactTokenParser("closing curly bracket", "}"))
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
						Expect(result.Lines).To(Equal([]int{1}))
						Expect(result.PosStart).To(Equal(1))
						Expect(result.PosEnd).To(Equal(14))
						Expect(result.Literal).To(Equal(""))
						Expect(result.Token).To(Equal(""))
						Expect(result.Children).To(BeEmpty())
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
