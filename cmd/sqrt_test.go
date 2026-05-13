package cmd

import (
	"errors"
	"testing"

	. "github.com/onsi/gomega"
)

func TestSqrtWhenZero(t *testing.T) {
	RegisterTestingT(t)

	e := float64(0)
	r, _ := Sqrt("0")

	Expect(r).To(Equal(e))
}

func TestSqrtWhenPositive(t *testing.T) {
	RegisterTestingT(t)

	e := float64(5)
	r, _ := Sqrt("25")

	Expect(r).To(Equal(e))
}

func TestSqrtWhenDecimal(t *testing.T) {
	RegisterTestingT(t)

	e := float64(0.5)
	r, _ := Sqrt("0.25")

	Expect(r).To(Equal(e))
}

func TestSqrtWhenNegative(t *testing.T) {
	RegisterTestingT(t)

	e := errors.New("number must be greater or equal to 0")
	_, err := Sqrt("-9")

	Expect(err).To(Equal(e))
}
