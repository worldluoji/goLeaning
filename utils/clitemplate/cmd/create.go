package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	fileutils "clitemplate/fileutils"
	gitOperation "clitemplate/gitoperation"

	_ "clitemplate/config"
)

var (
	project string
)

var createCmd = &cobra.Command{
	Use: "create",
	Short: `create a new project with create command,
	      for example: ` + cliName + ` create -p demo
	      see more, using command: ` + cliName + ` create -h`,
	Long: `create a new project with create command,for example:
	       ` + cliName + " create -p demo",
	Run: func(cmd *cobra.Command, args []string) {
		var destPath string
		if fileutils.IsDir(project) {
			destPath = project
			fileutils.MkDir(destPath)
		} else {
			destPath = fileutils.GetCurrentJointDir(project)
			fileutils.MkDir(destPath)
			if !fileutils.IsDir(destPath) {
				fmt.Println("dest path error....", destPath)
				return
			}
		}

		// if !remote {
		// 	if err := template.CopyEmbededFiles(mobile, destPath); err != nil {
		// 		fmt.Println("copy error....", err)
		// 		return
		// 	}

		// 	return
		// }

		remoteAddr := viper.Get("remoteAddr").(string)
		var branch = viper.Get("branch").(string)

		fmt.Println("begin to get code from gitlab...")
		if _, err := gitOperation.GitClone(destPath, remoteAddr, branch); err != nil {
			fmt.Println(err)
		}
		fmt.Println("get code from gitlab finished...")
	},
}

func init() {
	processFlags()
	rootCmd.AddCommand(createCmd)
}

func processFlags() {
	createCmd.Flags().StringVarP(&project, "project", "p", "", "项目名称，可以是目录，为必选项")
	createCmd.MarkFlagRequired("project")
}
