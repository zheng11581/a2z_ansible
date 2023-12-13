package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myscript",
	Short: "myscript",
	Long:  "Plenty of script for ops",
	Args:  cobra.NoArgs,
}

func Excute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
