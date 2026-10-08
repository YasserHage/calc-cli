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

func TestCalculateNegativeNumbers(t *testing.T) {
	RegisterTestingT(t)

	r, _ := Calculate("-3 + 5")
	Expect(r).To(Equal(float64(2)))

	r, _ = Calculate("2 * -3")
	Expect(r).To(Equal(float64(-6)))
}

func TestCalculateDivisionByZero(t *testing.T) {
	RegisterTestingT(t)

	_, err := Calculate("1 / 0")

	Expect(err).To(MatchError("division by zero"))
}

func TestCalculateTrailingOperator(t *testing.T) {
	RegisterTestingT(t)

	_, err := Calculate("2 +")

	Expect(err).To(MatchError("expression is not valid: unexpected end of expression"))
}

func TestCalculateInvalidCharacter(t *testing.T) {
	RegisterTestingT(t)

	_, err := Calculate("abc")

	Expect(err).To(MatchError("expression is not valid: invalid character 'a' at position 1"))
}
