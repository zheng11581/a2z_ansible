package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "download",
	Short: "download",
	Long:  "Download config YAML file",
}

func Excute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("destination", "D", "/tmp/ingress_monitor/config.yaml", "the name of download file")
}
