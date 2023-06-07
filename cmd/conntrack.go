package cmd

import (
	"github.com/spf13/cobra"

        luci "github.com/oxr463/luci-cli/pkg/luci"
)

func init() {
	netCmd.AddCommand(conntrackCmd)
}

var conntrackCmd = &cobra.Command{
	Use: "conntrack",
	Run: func(cmd *cobra.Command, args []string) {
		luci.GetSysNetConntrack()
	},
}
