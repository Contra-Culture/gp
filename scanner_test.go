package gp_test

import (
	. "github.com/Contra-Culture/gp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("gp", func() {
	Describe("GPRuneScanner", func() {
		Describe("creation", func() {
			Describe("New()", func() {
				It("returns rune scanner", func() {
					rs := NewRuneScanner("abcdefg")
					Expect(rs).NotTo(BeNil())
				})
			})
		})
		Describe("reading", func() {
			Describe(".ReadOne()", func() {
				Context("when there is a rune to read", func() {
					It("returns rune", func() {
						rs := NewRuneScanner("abc")
						r, n, err := rs.ReadOne()
						Expect(err).NotTo(HaveOccurred())
						Expect(n).To(Equal(0))
						Expect(r).To(Equal('a'))
						r, n, err = rs.ReadOne()
						Expect(err).NotTo(HaveOccurred())
						Expect(n).To(Equal(1))
						Expect(r).To(Equal('b'))
						r, n, err = rs.ReadOne()
						Expect(err).NotTo(HaveOccurred())
						Expect(n).To(Equal(2))
						Expect(r).To(Equal('c'))
					})
				})
				Context("when there is no rune to read", func() {
					It("returns error", func() {
						rs := NewRuneScanner("")
						r, n, err := rs.ReadOne()
						Expect(err).To(HaveOccurred())
						Expect(err.Error()).To(Equal("no rune to read"))
						Expect(n).To(Equal(0))
						Expect(r).To(Equal(int32(0)))
					})
				})
			})
		})
		Describe("unreading", func() {
			Context("when there are runes to unread", func() {
				It("unreads rune", func() {
					rs := NewRuneScanner("abc")
					r, n, err := rs.ReadOne()
					Expect(err).NotTo(HaveOccurred())
					Expect(n).To(Equal(0))
					Expect(r).To(Equal('a'))
					n, err = rs.Unread(1)
					Expect(err).NotTo(HaveOccurred())
					Expect(n).To(Equal(-1))
					r, n, err = rs.ReadOne()
					Expect(err).NotTo(HaveOccurred())
					Expect(n).To(Equal(0))
					Expect(r).To(Equal('a'))
					r, n, err = rs.ReadOne()
					Expect(err).NotTo(HaveOccurred())
					Expect(n).To(Equal(1))
					Expect(r).To(Equal('b'))
					r, n, err = rs.ReadOne()
					Expect(err).NotTo(HaveOccurred())
					Expect(n).To(Equal(2))
					Expect(r).To(Equal('c'))
					n, err = rs.Unread(1)
					Expect(err).NotTo(HaveOccurred())
					Expect(n).To(Equal(1))
					n, err = rs.Unread(1)
					Expect(err).NotTo(HaveOccurred())
					Expect(n).To(Equal(0))
					n, err = rs.Unread(1)
					Expect(err).NotTo(HaveOccurred())
					Expect(n).To(Equal(-1))
					r, n, err = rs.ReadOne()
					Expect(err).NotTo(HaveOccurred())
					Expect(n).To(Equal(0))
					Expect(r).To(Equal('a'))
					r, n, err = rs.ReadOne()
					Expect(err).NotTo(HaveOccurred())
					Expect(n).To(Equal(1))
					Expect(r).To(Equal('b'))
					r, n, err = rs.ReadOne()
					Expect(err).NotTo(HaveOccurred())
					Expect(n).To(Equal(2))
					Expect(r).To(Equal('c'))
				})
			})
		})
	})
})
