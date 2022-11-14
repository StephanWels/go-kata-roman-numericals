package roman

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRomanNumeralDecoder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Roman Numeral Suite")
}

var _ = Describe("test roman to decimal converter", func() {
	It("should give decimal number from roman", func() {
// 		Expect(Decode(...)).To(Equal(...))
	})
})
