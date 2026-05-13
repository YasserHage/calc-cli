package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/YasserHage/calc-cli/calculator"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calc expression",
	Args:  cobra.ExactArgs(1),
	Short: "Evaluate mathematical expressions from the command line",
	Long: `Calc is a simple command-line calculator written in Go.

It supports:
	- Addition and subtraction
	- Multiplication and division
	- Parentheses
	- Decimal numbers

Examples:

	calc "2 + 3 * 4"
	calc "(10 - 2) / 4"
	calc "3.5 * (2 + 1)"

The expression should be passed as a quoted string
to prevent the shell from interpreting special characters.`,
	Run: func(cmd *cobra.Command, args []string) {
		r, err := calculator.Calculate(args[0])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(r)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
