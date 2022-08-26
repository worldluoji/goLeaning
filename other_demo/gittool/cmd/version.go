package cmd

// 添加一个 version 命令

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	viper.Set("author", "luoji")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of GitTool",
	Long:  `All software has versions. This is GitTool's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GitTool v1.0 -- HEAD", "Author:", viper.Get("author"))
	},
}
