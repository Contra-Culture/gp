package gp_test

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/Contra-Culture/gp"
	. "github.com/Contra-Culture/gp/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("*Node", func() {
	Describe(".Traverse()", func() {
		It("traverses over all the nodes", func() {
			parser, err := New()
			Expect(err).NotTo(HaveOccurred())
			Expect(parser).NotTo(BeNil())
			file, err := os.Open("./json/test/test.json")
			Expect(err).NotTo(HaveOccurred())
			defer file.Close()
			readBytes, err := ioutil.ReadAll(file)
			rs := gp.NewRuneScanner(string(readBytes))
			Expect(err).NotTo(HaveOccurred())
			ast, err := parser.Parse(rs)
			Expect(err).NotTo(HaveOccurred())
			var sb strings.Builder
			err = ast.Traverse(func(d int, idx int, node *gp.Node) (err error) {
				sb.WriteRune('\n')
				for i := 0; i <= d; i++ {
					sb.WriteRune('\t')
				}
				sb.WriteString("parsed: ")
				sb.WriteString(string(node.Parsed()))
				return
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(sb.String()).To(Equal(""))

		})
	})
})
