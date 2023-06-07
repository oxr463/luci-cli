package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(uciCmd)
}

var uciCmd = &cobra.Command{
	Use: "uci",
        Short: "Unified Configuration Interface",
}
