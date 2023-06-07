package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(sysCmd)
}

var sysCmd = &cobra.Command{
	Use: "sys",
        Short: "System Information",
}
