package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of luci",
	Long:  `All software has versions. This is luci's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("luci v0.0.1")
	},
}
