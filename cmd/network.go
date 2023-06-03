package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  networkCmd.Flags().BoolVarP("all", "A", false, "All")
  uciCmd.AddCommand(networkCmd)
}

var networkCmd = &cobra.Command{
  Use:   "network",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("network")
  },
}
