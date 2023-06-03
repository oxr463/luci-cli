package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  sysCmd.AddCommand(netCmd)
}

var netCmd = &cobra.Command{
  Use:   "net",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("net")
  },
}
