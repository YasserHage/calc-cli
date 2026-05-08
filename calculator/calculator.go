package calculator

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func Calculate(expression string) (float64, error) {
	l, r := 0, 0
	numbers := []float64{}
	operators := []rune{}
	validOperators := []rune{'*', '/', '+', '-'}
	validDigits := []rune{'.', '(', ')'}
	parentheses := 0
	expression = strings.ReplaceAll(expression, " ", "")

	for r < len(expression) {
		curr := rune(expression[r])

		if curr == ')' {
			if parentheses == 0 {
				return 0, errors.New("expression is not valid: unmatched parentheses")
			}
			parentheses = parentheses - 1
		}

		if curr == '(' {
			parentheses = parentheses + 1
		}

		if !unicode.IsDigit(curr) && !slices.Contains(validDigits, curr) {
			if parentheses != 0 {
				r = r + 1
				continue
			}
			n, err := getNumber(expression, l, r)
			if err != nil {
				return 0, err
			}

			numbers = append(numbers, n)
			if slices.Contains(validOperators, curr) {
				operators = append(operators, curr)
			} else {
				return 0, fmt.Errorf("expression is not valid: invalid character '%c' at position '%d'", curr, r)
			}
			l = r + 1
		}
		r = r + 1
	}

	if parentheses != 0 {
		return 0, errors.New("expression is not valid: unmatched parentheses")
	}

	n, err := getNumber(expression, l, r)
	if err != nil {
		return 0, err
	}
	numbers = append(numbers, n)
	numbers, operators = multiplyDivide(numbers, operators)
	addSub(numbers, operators)
	return numbers[0], nil
}

func getNumber(expression string, l int, r int) (float64, error) {
	if expression[l] == '(' && expression[r-1] == ')' {
		return Calculate(expression[l+1 : r-1])
	}
	result, err := strconv.ParseFloat(expression[l:r], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s", expression[l:r])
	}
	return result, nil
}

func multiplyDivide(numbers []float64, operators []rune) ([]float64, []rune) {
	for i := 0; i < len(operators); i++ {
		o := operators[i]

		if o == '*' {
			numbers[i] = numbers[i] * numbers[i+1]

			numbers = append(numbers[:i+1], numbers[i+2:]...)
			operators = append(operators[:i], operators[i+1:]...)

			i--
			continue
		}

		if o == '/' {
			numbers[i] = numbers[i] / numbers[i+1]

			numbers = append(numbers[:i+1], numbers[i+2:]...)
			operators = append(operators[:i], operators[i+1:]...)

			i--
			continue
		}
	}
	return numbers, operators
}

func addSub(numbers []float64, operators []rune) {
	for i := 0; i < len(operators); i++ {
		o := operators[i]

		if o == '+' {
			numbers[i] = numbers[i] + numbers[i+1]

			numbers = append(numbers[:i+1], numbers[i+2:]...)
			operators = append(operators[:i], operators[i+1:]...)

			i--
			continue
		}

		if o == '-' {
			numbers[i] = numbers[i] - numbers[i+1]

			numbers = append(numbers[:i+1], numbers[i+2:]...)
			operators = append(operators[:i], operators[i+1:]...)

			i--
			continue
		}
	}
}
