package gp_test

import (
	gp "github.com/Contra-Culture/gp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("gp", func() {
	Context("creation", func() {
		Describe("New()", func() {
			It("returns parser node", func() {
				pn := gp.New("test purpose", gp.Optional(false), gp.ExactTokenParser("test", "test"))
				Expect(pn).NotTo(BeNil())
			})
		})
	})
})
