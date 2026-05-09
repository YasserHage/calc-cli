package cmd

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log number",
	Short: "Calculate the logarithm of a number",
	Long: `Calculate logarithms using base 10, the natural logarithm, or a custom base.

By default, the log command uses base 10.

Use the -b flag to specify a different base.

Examples:

  calc log 100
  calc log 1000
  calc log 8 -b 2

- The input number must be greater than zero.
- The logarithm base must also be greater than zero and cannot be equal to 1`,
	RunE: func(cmd *cobra.Command, args []string) error {
		n, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return err
		}

		b, err := cmd.Flags().GetFloat64("base")
		if err != nil {
			return err
		}

		r, err := Calculate(n, b)
		if err != nil {
			return err
		}
		fmt.Println(r)
		return nil
	},
}

func Calculate(n float64, b float64) (float64, error) {
	if n <= 0 {
		return 0, errors.New("number must be greater than 0")
	}
	if b <= 0 || b == 1 {
		return 0, errors.New("base should be greater than 0 and different than 1")
	}

	return (math.Log(n) / math.Log(b)), nil
}

func init() {
	rootCmd.AddCommand(logCmd)

	logCmd.Flags().Float64P("base", "b", 10, "Define a custom base (b != 1 && b > 0)")
}
