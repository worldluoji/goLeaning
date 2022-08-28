package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "clitemplate/config"
)

var rootCmd = &cobra.Command{
	Use:   viper.Get("cliName").(string),
	Short: "cliTemplate",
	Long:  "cliTemplate -h",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
