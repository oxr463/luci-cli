package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  sysCmd.AddCommand(uptimeCmd)
}

var uptimeCmd = &cobra.Command{
  Use:   "uptime",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("uptime")
  },
}
