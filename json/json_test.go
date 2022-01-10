package json_test

import (
	"github.com/Contra-Culture/gp"
	. "github.com/Contra-Culture/gp/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("json parser", func() {
	var (
		parser gp.Parser
		err    error
		rs     *gp.GPRuneScanner
		n      *gp.Node
	)
	BeforeSuite(func() {
		parser, err = New()
		Expect(err).NotTo(HaveOccurred())
		Expect(parser).NotTo(BeNil())
		// file, err := os.Open("./test/test.json")
		// Expect(err).NotTo(HaveOccurred())
		// defer file.Close()
		// readBytes, err := ioutil.ReadAll(file)
		// rs = gp.NewRuneScanner(string(readBytes))
		// Expect(err).NotTo(HaveOccurred())
	})
	Describe("parsing", func() {
		It("returns AST", func() {
			rs = gp.NewRuneScanner("null")
			n, err = parser.Parse(rs)
			Expect(err).NotTo(HaveOccurred())
			Expect(n.Tags()).To(Equal([]string{}))
			Expect(n.Parsed()).To(Equal([]rune{'n', 'u', 'l', 'l'}))
		})
	})
})
