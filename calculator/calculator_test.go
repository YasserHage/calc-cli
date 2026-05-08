package calculator

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestCalculateAddition(t *testing.T) {
	RegisterTestingT(t)

	e := float64(2)
	r, _ := Calculate("1 + 1")

	Expect(r).To(Equal(e))
}

func TestCalculateSubtraction(t *testing.T) {
	RegisterTestingT(t)

	e := float64(1)
	r, _ := Calculate("3 - 2")

	Expect(r).To(Equal(e))
}

func TestCalculateMultiplication(t *testing.T) {
	RegisterTestingT(t)

	e := float64(12)
	r, _ := Calculate("6 * 2")

	Expect(r).To(Equal(e))
}

func TestCalculateDivision(t *testing.T) {
	RegisterTestingT(t)

	e := float64(15)
	r, _ := Calculate("45 / 3")

	Expect(r).To(Equal(e))
}

func TestCalculateParentheses(t *testing.T) {
	RegisterTestingT(t)

	e := float64(55)
	r, _ := Calculate("5 * (2 * (1 + 2) + 4) + 5")

	Expect(r).To(Equal(e))
}
