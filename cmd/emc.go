package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bvarberg/wwork/internal/emc"
	"github.com/spf13/cobra"
)

var emcCmd = &cobra.Command{
	Use:   "emc temperature relative_humidity",
	Args:  cobra.ExactArgs(2),
	Short: "Calculate the EMC for the given climate",
	Long: `Calculate the equilibrium moisture content (EMC) given a temperature and relative humidity.
The EMC is the moisture content at which wood neither gains nor loses moisture, and is representative
of whether the wood is stable enough to work.
`,
	Example: "emc 70.0 0.35",
	Run: func(cmd *cobra.Command, args []string) {
		t, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid temperature value")
			os.Exit(1)
		}

		rh, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid relative humidity value")
			os.Exit(1)
		}

		formattedEmc := fmt.Sprintf("%.1f", emc.Simpson(t, rh))

		fmt.Printf("Equilibrium moisture content is %v%%\n", formattedEmc)
	},
}
