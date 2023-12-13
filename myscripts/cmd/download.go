package cmd

import (
	"fmt"
	download "zheng11581/myscripts/pkg"

	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download",
	Long:  "download ingress monitor config file",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var download download.Download
		download.Url = "http://127.0.0.1:5000/ingress_monitor/conf/download"
		if env == "prod" {
			download.Url = "https://ywb.yyuap.com/opsv/ywb-custom/ingress_monitor/conf/download"
		}
		download.Method = "GET"
		download.FileName = fileName
		download.SaveAs = saveas
		download.Download()
		fmt.Printf("FileName: %s, SaveAs: %s, Env: %s\n", download.FileName, download.SaveAs, env)
	},
}

var saveas, fileName, env string

func init() {
	downloadCmd.Flags().StringVar(&saveas, "saveas", "", "the file path to save")
	downloadCmd.Flags().StringVar(&fileName, "filename", "/tmp/ingress_monitor/config.yaml", "the name of remote file")
	downloadCmd.Flags().StringVar(&env, "env", "prod", "the env that the download file from")
	rootCmd.AddCommand(downloadCmd)
}
