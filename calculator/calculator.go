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
	if expression == "" {
		return 0, errors.New("expression is not valid: expression is empty")
	}

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
			if (curr == '-' || curr == '+') && l == r {
				r = r + 1
				continue
			}
			if !slices.Contains(validOperators, curr) {
				return 0, fmt.Errorf("expression is not valid: invalid character '%c' at position %d", curr, r+1)
			}
			n, err := getNumber(expression, l, r)
			if err != nil {
				return 0, err
			}

			numbers = append(numbers, n)
			operators = append(operators, curr)
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
	numbers, operators, err = multiplyDivide(numbers, operators)
	if err != nil {
		return 0, err
	}
	addSub(numbers, operators)
	return numbers[0], nil
}

func getNumber(expression string, l int, r int) (float64, error) {
	if l >= r {
		if r >= len(expression) {
			return 0, errors.New("expression is not valid: unexpected end of expression")
		}
		return 0, fmt.Errorf("expression is not valid: invalid character '%c' at position %d", expression[r], r+1)
	}
	if expression[l] == '-' || expression[l] == '+' {
		n, err := getNumber(expression, l+1, r)
		if expression[l] == '-' {
			n = -n
		}
		return n, err
	}
	if expression[l] == '(' && expression[r-1] == ')' {
		return Calculate(expression[l+1 : r-1])
	}
	result, err := strconv.ParseFloat(expression[l:r], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s", expression[l:r])
	}
	return result, nil
}

func multiplyDivide(numbers []float64, operators []rune) ([]float64, []rune, error) {
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
			if numbers[i+1] == 0 {
				return nil, nil, errors.New("division by zero")
			}
			numbers[i] = numbers[i] / numbers[i+1]

			numbers = append(numbers[:i+1], numbers[i+2:]...)
			operators = append(operators[:i], operators[i+1:]...)

			i--
			continue
		}
	}
	return numbers, operators, nil
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
