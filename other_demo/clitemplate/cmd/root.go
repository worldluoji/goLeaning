package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "clitemplate/config"
)

var cliName = viper.Get("cliName").(string)

var rootCmd = &cobra.Command{
	Use:   cliName,
	Short: cliName,
	Long:  cliName,
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
