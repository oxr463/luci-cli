package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  netCmd.AddCommand(devicesCmd)
}

var devicesCmd = &cobra.Command{
  Use:   "devices",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("devices")
  },
}
