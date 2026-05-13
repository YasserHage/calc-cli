package cmd

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

var sqrtCmd = &cobra.Command{
	Use:   "sqrt number",
	Short: "Calculate the square root of a number",
	Long: `Calculate the square root of a positive number.

The sqrt command returns the non-negative square root
of the provided value.

Examples:

  calc sqrt 9
  calc sqrt 16
  calc sqrt 2

Results are returned as floating point numbers when necessary.

The input number must be greater than or equal to zero.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := Sqrt(args[0])
		if err != nil {
			return err
		}
		fmt.Println(s)
		return nil
	},
}

func Sqrt(arg string) (float64, error) {
	n, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return 0, err
	}
	if n < 0 {
		return 0, errors.New("number must be greater or equal to 0")
	}

	return math.Sqrt(n), nil
}

func init() {
	rootCmd.AddCommand(sqrtCmd)
}
