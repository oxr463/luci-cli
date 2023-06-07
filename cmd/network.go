package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var All bool

func init() {
	networkCmd.Flags().BoolVarP(&All, "all", "A", false, "All")
	uciCmd.AddCommand(networkCmd)
}

var networkCmd = &cobra.Command{
	Use: "network",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("network")
	},
}
