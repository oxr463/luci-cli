package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(uciCmd)
}

var uciCmd = &cobra.Command{
	Use: "uci",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uci")
	},
}
