package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	fileutils "clitemplate/utils/fileutils"

	_ "clitemplate/config"

	gitworker "clitemplate/worker/gitworker"
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

		woker := &gitworker.GitWorker{
			Dest:   destPath,
			Url:    viper.Get("remoteAddr").(string),
			Branch: viper.Get("branch").(string),
		}

		woker.Do()
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
