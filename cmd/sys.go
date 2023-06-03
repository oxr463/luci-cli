package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  getCmd.AddCommand(sysCmd)
}

var sysCmd = &cobra.Command{
  Use:   "sys",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("sys")
  },
}
