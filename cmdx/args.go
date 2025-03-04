package cmdx

import (
	"fmt"

	"github.com/spf13/cobra"
)

// MinArgs fatals if args does not satisfy min.
func MinArgs(cmd *cobra.Command, args []string, min int) {
	if len(args) < min {
		Fatalf(`%s

Expected at least %d command line arguments but only got %d.`, cmd.UsageString(), min, len(args))
	}
}

// ExactArgs fatals if args does not equal l.
func ExactArgs(cmd *cobra.Command, args []string, l int) {
	if len(args) < l {
		Fatalf(`%s

Expected exactly %d command line arguments but got %d.`, cmd.UsageString(), l, len(args))
	}
}

// RangeArgs fatals if args does not satisfy any of the lengths set in r.
func RangeArgs(cmd *cobra.Command, args []string, r []int) {
	for _, a := range r {
		if len(args) == a {
			return
		}
	}
	Fatalf(`%s

Expected exact %v command line arguments but got %d.`, cmd.UsageString(), r, len(args))
}

// ZeroOrTwoArgs requires either no or 2 arguments.
func ZeroOrTwoArgs(cmd *cobra.Command, args []string) error {
	// zero or exactly two args
	if len(args) != 0 && len(args) != 2 {
		return fmt.Errorf("expected zero or two args, got %d: %+v", len(args), args)
	}
	return nil
}
