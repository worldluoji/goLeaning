package cmd

// 添加一个 version 命令

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "clitemplate/config"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of GitTool",
	Long:  `All software has versions. This is GitTool's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.Get("cliName"), " Vertion:", viper.Get("version"), " Author:", viper.Get("author"))
	},
}
