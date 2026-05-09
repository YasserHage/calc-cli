package cmd

import (
	"errors"
	"testing"

	. "github.com/onsi/gomega"
)

func TestLog(t *testing.T) {
	RegisterTestingT(t)

	e := float64(2)
	r, _ := Calculate(100, 10)

	Expect(r).To(Equal(e))
}

func TestLogDiffBase(t *testing.T) {
	RegisterTestingT(t)

	e := float64(-3)
	r, _ := Calculate(8, 0.5)

	Expect(r).To(Equal(e))
}

func TestLogBaseOne(t *testing.T) {
	RegisterTestingT(t)

	e := errors.New("base should be greater than 0 and different than 1")
	_, err := Calculate(10, 1)

	Expect(err).To(Equal(e))
}

func TestLogNegativeBase(t *testing.T) {
	RegisterTestingT(t)

	e := errors.New("base should be greater than 0 and different than 1")
	_, err := Calculate(10, -1)

	Expect(err).To(Equal(e))
}

func TestLogNegativeNumber(t *testing.T) {
	RegisterTestingT(t)

	e := errors.New("number must be greater than 0")
	_, err := Calculate(-10, 1)

	Expect(err).To(Equal(e))
}
