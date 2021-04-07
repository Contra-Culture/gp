package store_test

import (
	"bufio"
	"os"

	. "github.com/Contra-Culture/gp/store"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("store", func() {
	var (
		store   *SymbolsStore
		file    *os.File
		err     error
		content *bufio.Reader
	)
	BeforeEach(func() {
		file, err = os.Open("../test/test")
		if err != nil {
			panic(err)
		}
		content = bufio.NewReader(file)
		store, err = New(content)
		Expect(err).NotTo(HaveOccurred())
		Expect(store).NotTo(BeNil())
	})
	AfterEach(func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	})
	Context("creation", func() {
		Describe("New()", func() {
			It("returns store", func() {
				Expect(store).To(BeAssignableToTypeOf(&SymbolsStore{}))
			})
		})
	})
	Context("use cases", func() {
		Describe(".GetSymbol()", func() {
			Context("when symbol index is in symbols range", func() {
				It("returns symbol", func() {
					// line1: some content
					symbol, err := store.GetSymbol(0)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(1))
					Expect(symbol.Position).To(Equal(1))
					Expect(symbol.Rune).To(Equal('s'))
					symbol, err = store.GetSymbol(1)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(1))
					Expect(symbol.Position).To(Equal(2))
					Expect(symbol.Rune).To(Equal('o'))
					symbol, err = store.GetSymbol(2)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(1))
					Expect(symbol.Position).To(Equal(3))
					Expect(symbol.Rune).To(Equal('m'))
					symbol, err = store.GetSymbol(3)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(1))
					Expect(symbol.Position).To(Equal(4))
					Expect(symbol.Rune).To(Equal('e'))
					symbol, err = store.GetSymbol(4)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(1))
					Expect(symbol.Position).To(Equal(5))
					Expect(symbol.Rune).To(Equal(' '))
					symbol, err = store.GetSymbol(5)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(1))
					Expect(symbol.Position).To(Equal(6))
					Expect(symbol.Rune).To(Equal('c'))
					symbol, err = store.GetSymbol(6)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(1))
					Expect(symbol.Position).To(Equal(7))
					Expect(symbol.Rune).To(Equal('o'))
					symbol, err = store.GetSymbol(7)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(1))
					Expect(symbol.Position).To(Equal(8))
					Expect(symbol.Rune).To(Equal('n'))
					symbol, err = store.GetSymbol(8)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(1))
					Expect(symbol.Position).To(Equal(9))
					Expect(symbol.Rune).To(Equal('t'))
					symbol, err = store.GetSymbol(9)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(1))
					Expect(symbol.Position).To(Equal(10))
					Expect(symbol.Rune).To(Equal('e'))
					symbol, err = store.GetSymbol(10)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(1))
					Expect(symbol.Position).To(Equal(11))
					Expect(symbol.Rune).To(Equal('n'))
					symbol, err = store.GetSymbol(11)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(1))
					Expect(symbol.Position).To(Equal(12))
					Expect(symbol.Rune).To(Equal('t'))
					symbol, err = store.GetSymbol(12)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(1))
					Expect(symbol.Position).To(Equal(13))
					Expect(symbol.Rune).To(Equal('\n'))
					// line 2: here
					symbol, err = store.GetSymbol(13)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(2))
					Expect(symbol.Position).To(Equal(1))
					Expect(symbol.Rune).To(Equal('h'))
					symbol, err = store.GetSymbol(14)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(2))
					Expect(symbol.Position).To(Equal(2))
					Expect(symbol.Rune).To(Equal('e'))
					symbol, err = store.GetSymbol(15)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(2))
					Expect(symbol.Position).To(Equal(3))
					Expect(symbol.Rune).To(Equal('r'))
					symbol, err = store.GetSymbol(16)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(2))
					Expect(symbol.Position).To(Equal(4))
					Expect(symbol.Rune).To(Equal('e'))
					symbol, err = store.GetSymbol(17)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(2))
					Expect(symbol.Position).To(Equal(5))
					Expect(symbol.Rune).To(Equal('\n'))
					// line 3: for test
					symbol, err = store.GetSymbol(18)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(3))
					Expect(symbol.Position).To(Equal(1))
					Expect(symbol.Rune).To(Equal('f'))
					symbol, err = store.GetSymbol(19)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(3))
					Expect(symbol.Position).To(Equal(2))
					Expect(symbol.Rune).To(Equal('o'))
					symbol, err = store.GetSymbol(20)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(3))
					Expect(symbol.Position).To(Equal(3))
					Expect(symbol.Rune).To(Equal('r'))
					symbol, err = store.GetSymbol(21)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(3))
					Expect(symbol.Position).To(Equal(4))
					Expect(symbol.Rune).To(Equal(' '))
					symbol, err = store.GetSymbol(22)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(3))
					Expect(symbol.Position).To(Equal(5))
					Expect(symbol.Rune).To(Equal('t'))
					symbol, err = store.GetSymbol(23)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(3))
					Expect(symbol.Position).To(Equal(6))
					Expect(symbol.Rune).To(Equal('e'))
					symbol, err = store.GetSymbol(24)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(3))
					Expect(symbol.Position).To(Equal(7))
					Expect(symbol.Rune).To(Equal('s'))
					symbol, err = store.GetSymbol(25)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(3))
					Expect(symbol.Position).To(Equal(8))
					Expect(symbol.Rune).To(Equal('t'))
					symbol, err = store.GetSymbol(26)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(3))
					Expect(symbol.Position).To(Equal(9))
					Expect(symbol.Rune).To(Equal('\n'))
					// line 4: purposes
					symbol, err = store.GetSymbol(27)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(4))
					Expect(symbol.Position).To(Equal(1))
					Expect(symbol.Rune).To(Equal('p'))
					symbol, err = store.GetSymbol(28)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(4))
					Expect(symbol.Position).To(Equal(2))
					Expect(symbol.Rune).To(Equal('u'))
					symbol, err = store.GetSymbol(29)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(4))
					Expect(symbol.Position).To(Equal(3))
					Expect(symbol.Rune).To(Equal('r'))
					symbol, err = store.GetSymbol(30)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(4))
					Expect(symbol.Position).To(Equal(4))
					Expect(symbol.Rune).To(Equal('p'))
					symbol, err = store.GetSymbol(31)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(4))
					Expect(symbol.Position).To(Equal(5))
					Expect(symbol.Rune).To(Equal('o'))
					symbol, err = store.GetSymbol(32)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(4))
					Expect(symbol.Position).To(Equal(6))
					Expect(symbol.Rune).To(Equal('s'))
					symbol, err = store.GetSymbol(33)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(4))
					Expect(symbol.Position).To(Equal(7))
					Expect(symbol.Rune).To(Equal('e'))
					symbol, err = store.GetSymbol(34)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(4))
					Expect(symbol.Position).To(Equal(8))
					Expect(symbol.Rune).To(Equal('s'))
					symbol, err = store.GetSymbol(35)
					Expect(err).NotTo(HaveOccurred())
					Expect(symbol.Line).To(Equal(4))
					Expect(symbol.Position).To(Equal(9))
					Expect(symbol.Rune).To(Equal('\n'))
				})
			})
			Context("when symbol index is out of range", func() {
				It("returns error", func() {
					symbol, err := store.GetSymbol(36)
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("36 index is out of symbols range (lenght: 36)"))
					Expect(symbol.Line).To(Equal(0))
					Expect(symbol.Position).To(Equal(0))
					Expect(symbol.Rune).To(Equal(rune(0)))
				})
			})
		})
		Describe(".GetLine()", func() {
			Context("when line index is in lines range", func() {
				It("returns line", func() {
					line, err := store.GetLine(0)
					Expect(err).NotTo(HaveOccurred())
					Expect(line).To(HaveLen(13))
					Expect(line[0]).To(Equal(Symbol{
						Rune:     's',
						Line:     1,
						Position: 1,
					}))
					Expect(line[1]).To(Equal(Symbol{
						Rune:     'o',
						Line:     1,
						Position: 2,
					}))
					Expect(line[2]).To(Equal(Symbol{
						Rune:     'm',
						Line:     1,
						Position: 3,
					}))
					Expect(line[3]).To(Equal(Symbol{
						Rune:     'e',
						Line:     1,
						Position: 4,
					}))
					Expect(line[4]).To(Equal(Symbol{
						Rune:     ' ',
						Line:     1,
						Position: 5,
					}))
					Expect(line[5]).To(Equal(Symbol{
						Rune:     'c',
						Line:     1,
						Position: 6,
					}))
					Expect(line[6]).To(Equal(Symbol{
						Rune:     'o',
						Line:     1,
						Position: 7,
					}))
					Expect(line[7]).To(Equal(Symbol{
						Rune:     'n',
						Line:     1,
						Position: 8,
					}))
					Expect(line[8]).To(Equal(Symbol{
						Rune:     't',
						Line:     1,
						Position: 9,
					}))
					Expect(line[9]).To(Equal(Symbol{
						Rune:     'e',
						Line:     1,
						Position: 10,
					}))
					Expect(line[10]).To(Equal(Symbol{
						Rune:     'n',
						Line:     1,
						Position: 11,
					}))
					Expect(line[11]).To(Equal(Symbol{
						Rune:     't',
						Line:     1,
						Position: 12,
					}))
					Expect(line[12]).To(Equal(Symbol{
						Rune:     '\n',
						Line:     1,
						Position: 13,
					}))
					line, err = store.GetLine(1)
					Expect(err).NotTo(HaveOccurred())
					Expect(line).To(HaveLen(5))
					Expect(line[0]).To(Equal(Symbol{
						Rune:     'h',
						Line:     2,
						Position: 1,
					}))
					Expect(line[1]).To(Equal(Symbol{
						Rune:     'e',
						Line:     2,
						Position: 2,
					}))
					Expect(line[2]).To(Equal(Symbol{
						Rune:     'r',
						Line:     2,
						Position: 3,
					}))
					Expect(line[3]).To(Equal(Symbol{
						Rune:     'e',
						Line:     2,
						Position: 4,
					}))
					Expect(line[4]).To(Equal(Symbol{
						Rune:     '\n',
						Line:     2,
						Position: 5,
					}))
					line, err = store.GetLine(2)
					Expect(err).NotTo(HaveOccurred())
					Expect(line).To(HaveLen(9))
					Expect(line[0]).To(Equal(Symbol{
						Rune:     'f',
						Line:     3,
						Position: 1,
					}))
					Expect(line[1]).To(Equal(Symbol{
						Rune:     'o',
						Line:     3,
						Position: 2,
					}))
					Expect(line[2]).To(Equal(Symbol{
						Rune:     'r',
						Line:     3,
						Position: 3,
					}))
					Expect(line[3]).To(Equal(Symbol{
						Rune:     ' ',
						Line:     3,
						Position: 4,
					}))
					Expect(line[4]).To(Equal(Symbol{
						Rune:     't',
						Line:     3,
						Position: 5,
					}))
					Expect(line[5]).To(Equal(Symbol{
						Rune:     'e',
						Line:     3,
						Position: 6,
					}))
					Expect(line[6]).To(Equal(Symbol{
						Rune:     's',
						Line:     3,
						Position: 7,
					}))
					Expect(line[7]).To(Equal(Symbol{
						Rune:     't',
						Line:     3,
						Position: 8,
					}))
					Expect(line[8]).To(Equal(Symbol{
						Rune:     '\n',
						Line:     3,
						Position: 9,
					}))
					line, err = store.GetLine(3)
					Expect(err).NotTo(HaveOccurred())
					Expect(line).To(HaveLen(9))
					Expect(line[0]).To(Equal(Symbol{
						Rune:     'p',
						Line:     4,
						Position: 1,
					}))
					Expect(line[1]).To(Equal(Symbol{
						Rune:     'u',
						Line:     4,
						Position: 2,
					}))
					Expect(line[2]).To(Equal(Symbol{
						Rune:     'r',
						Line:     4,
						Position: 3,
					}))
					Expect(line[3]).To(Equal(Symbol{
						Rune:     'p',
						Line:     4,
						Position: 4,
					}))
					Expect(line[4]).To(Equal(Symbol{
						Rune:     'o',
						Line:     4,
						Position: 5,
					}))
					Expect(line[5]).To(Equal(Symbol{
						Rune:     's',
						Line:     4,
						Position: 6,
					}))
					Expect(line[6]).To(Equal(Symbol{
						Rune:     'e',
						Line:     4,
						Position: 7,
					}))
					Expect(line[7]).To(Equal(Symbol{
						Rune:     's',
						Line:     4,
						Position: 8,
					}))
					Expect(line[8]).To(Equal(Symbol{
						Rune:     '\n',
						Line:     4,
						Position: 9,
					}))
					line, err = store.GetLine(4)
					Expect(err).NotTo(HaveOccurred())
					Expect(line).To(BeEmpty())
				})
			})
			Context("when line index is out of lines range", func() {
				It("returns error", func() {
					line, err := store.GetLine(5)
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("line index 5 is out of lines range (length: 5)"))
					Expect(line).To(BeNil())
				})
			})
		})
		Describe(".GetLineBySymbolIndex()", func() {
			Context("when symbol index is in symbols range", func() {
				It("returns line", func() {
					ln, line, err := store.GetLineBySymbolIndex(0)
					Expect(err).NotTo(HaveOccurred())
					Expect(ln).To(Equal(1))
					Expect(line).To(HaveLen(13))
					Expect(line[0]).To(Equal(Symbol{
						Rune:     's',
						Line:     1,
						Position: 1,
					}))
					ln, line, err = store.GetLineBySymbolIndex(13)
					Expect(err).NotTo(HaveOccurred())
					Expect(ln).To(Equal(2))
					Expect(line).To(HaveLen(5))
					Expect(line[0]).To(Equal(Symbol{
						Rune:     'h',
						Line:     2,
						Position: 1,
					}))
					ln, line, err = store.GetLineBySymbolIndex(15)
					Expect(err).NotTo(HaveOccurred())
					Expect(ln).To(Equal(2))
					Expect(line).To(HaveLen(5))
					Expect(line[0]).To(Equal(Symbol{
						Rune:     'h',
						Line:     2,
						Position: 1,
					}))
					ln, line, err = store.GetLineBySymbolIndex(12)
					Expect(err).NotTo(HaveOccurred())
					Expect(ln).To(Equal(1))
					Expect(line).To(HaveLen(13))
					Expect(line[0]).To(Equal(Symbol{
						Rune:     's',
						Line:     1,
						Position: 1,
					}))
					ln, line, err = store.GetLineBySymbolIndex(27)
					Expect(err).NotTo(HaveOccurred())
					Expect(ln).To(Equal(4))
					Expect(line).To(HaveLen(9))
					Expect(line[0]).To(Equal(Symbol{
						Rune:     'p',
						Line:     4,
						Position: 1,
					}))
				})
			})
			Context("when symbol index is out of symbols range", func() {
				It("returns error", func() {
					ln, line, err := store.GetLineBySymbolIndex(50)
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("50 index is out of symbols range (lenght: 36)"))
					Expect(ln).To(Equal(0))
					Expect(line).To(BeEmpty())
				})
			})
		})
		Describe(".LinesIndex()", func() {
			It("returns lines index", func() {
				li := store.LineIndex()
				Expect(li).To(HaveLen(5))
				Expect(li[0]).To(Equal(0))
				Expect(li[1]).To(Equal(13))
				Expect(li[2]).To(Equal(18))
				Expect(li[4]).To(Equal(36))
			})
		})
	})
})
