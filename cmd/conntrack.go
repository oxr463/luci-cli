package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  netCmd.AddCommand(conntrackCmd)
}

var conntrackCmd = &cobra.Command{
  Use:   "conntrack",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("conntrack")
  },
}
