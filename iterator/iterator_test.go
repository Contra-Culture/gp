package iterator_test

import (
	"bufio"
	"os"

	store "github.com/Contra-Culture/gp/store"

	. "github.com/Contra-Culture/gp/iterator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("iterator", func() {
	var (
		file         *os.File
		err          error
		content      *bufio.Reader
		symbolsStore *store.SymbolsStore
		iterator     *SymbolsIterator
	)
	BeforeEach(func() {
		file, err = os.Open("../test/test")
		if err != nil {
			panic(err)
		}
		content = bufio.NewReader(file)
		symbolsStore, err = store.New(content)
		Expect(err).NotTo(HaveOccurred())
		iterator = New(symbolsStore, 0)
		Expect(iterator).NotTo(BeNil())
	})
	AfterEach(func() {
		err = file.Close()
		if err != nil {
			panic(err)
		}
	})
	Describe(".Next()", func() {
		It("iterates over runes", func() {
			s, err := iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('s'))
			Expect(s.Line).To(Equal(1))
			Expect(s.Position).To(Equal(1))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('o'))
			Expect(s.Line).To(Equal(1))
			Expect(s.Position).To(Equal(2))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('m'))
			Expect(s.Line).To(Equal(1))
			Expect(s.Position).To(Equal(3))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('e'))
			Expect(s.Line).To(Equal(1))
			Expect(s.Position).To(Equal(4))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal(' '))
			Expect(s.Line).To(Equal(1))
			Expect(s.Position).To(Equal(5))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('c'))
			Expect(s.Line).To(Equal(1))
			Expect(s.Position).To(Equal(6))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('o'))
			Expect(s.Line).To(Equal(1))
			Expect(s.Position).To(Equal(7))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('n'))
			Expect(s.Line).To(Equal(1))
			Expect(s.Position).To(Equal(8))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('t'))
			Expect(s.Line).To(Equal(1))
			Expect(s.Position).To(Equal(9))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('e'))
			Expect(s.Line).To(Equal(1))
			Expect(s.Position).To(Equal(10))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('n'))
			Expect(s.Line).To(Equal(1))
			Expect(s.Position).To(Equal(11))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('t'))
			Expect(s.Line).To(Equal(1))
			Expect(s.Position).To(Equal(12))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('\n'))
			Expect(s.Line).To(Equal(1))
			Expect(s.Position).To(Equal(13))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('h'))
			Expect(s.Line).To(Equal(2))
			Expect(s.Position).To(Equal(1))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('e'))
			Expect(s.Line).To(Equal(2))
			Expect(s.Position).To(Equal(2))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('r'))
			Expect(s.Line).To(Equal(2))
			Expect(s.Position).To(Equal(3))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('e'))
			Expect(s.Line).To(Equal(2))
			Expect(s.Position).To(Equal(4))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('\n'))
			Expect(s.Line).To(Equal(2))
			Expect(s.Position).To(Equal(5))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('f'))
			Expect(s.Line).To(Equal(3))
			Expect(s.Position).To(Equal(1))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('o'))
			Expect(s.Line).To(Equal(3))
			Expect(s.Position).To(Equal(2))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('r'))
			Expect(s.Line).To(Equal(3))
			Expect(s.Position).To(Equal(3))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal(' '))
			Expect(s.Line).To(Equal(3))
			Expect(s.Position).To(Equal(4))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('t'))
			Expect(s.Line).To(Equal(3))
			Expect(s.Position).To(Equal(5))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('e'))
			Expect(s.Line).To(Equal(3))
			Expect(s.Position).To(Equal(6))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('s'))
			Expect(s.Line).To(Equal(3))
			Expect(s.Position).To(Equal(7))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('t'))
			Expect(s.Line).To(Equal(3))
			Expect(s.Position).To(Equal(8))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('\n'))
			Expect(s.Line).To(Equal(3))
			Expect(s.Position).To(Equal(9))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('p'))
			Expect(s.Line).To(Equal(4))
			Expect(s.Position).To(Equal(1))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('u'))
			Expect(s.Line).To(Equal(4))
			Expect(s.Position).To(Equal(2))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('r'))
			Expect(s.Line).To(Equal(4))
			Expect(s.Position).To(Equal(3))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('p'))
			Expect(s.Line).To(Equal(4))
			Expect(s.Position).To(Equal(4))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('o'))
			Expect(s.Line).To(Equal(4))
			Expect(s.Position).To(Equal(5))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('s'))
			Expect(s.Line).To(Equal(4))
			Expect(s.Position).To(Equal(6))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('e'))
			Expect(s.Line).To(Equal(4))
			Expect(s.Position).To(Equal(7))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('s'))
			Expect(s.Line).To(Equal(4))
			Expect(s.Position).To(Equal(8))
			s, err = iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('\n'))
			Expect(s.Line).To(Equal(4))
			Expect(s.Position).To(Equal(9))
			s, err = iterator.Next()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("36 index is out of symbols range (lenght: 36)"))
			Expect(s.Rune).To(Equal(rune(0)))
			Expect(s.Line).To(Equal(0))
			Expect(s.Position).To(Equal(0))
		})
	})
	Describe(".Fork()", func() {
		It("returns fork", func() {
			forked := iterator.Fork()
			s, err := iterator.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(s.Rune).To(Equal('s'))
			Expect(s.Line).To(Equal(1))
			Expect(s.Position).To(Equal(1))
			sf, err := forked.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(sf).To(Equal(s))
			sf, err = forked.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(sf).To(Equal(store.Symbol{
				Rune:     'o',
				Line:     1,
				Position: 2,
			}))
			forked2 := forked.Fork()
			sff, err := forked2.Next()
			Expect(err).NotTo(HaveOccurred())
			Expect(sff).To(Equal(store.Symbol{
				Rune:     'm',
				Line:     1,
				Position: 3,
			}))
		})
	})
})
