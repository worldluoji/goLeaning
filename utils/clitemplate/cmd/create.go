package cmd

import (
	"github.com/spf13/cobra"

	fileutils "clitemplate/utils/fileutils"

	_ "clitemplate/config"

	worker "clitemplate/worker"
)

var (
	project string
	w       worker.Worker
)

var createCmd = &cobra.Command{
	Use: "create",
	Short: `create a new project with create command,
	      for example: ` + cliName + ` create -p demo
	      see more, using command: ` + cliName + ` create -h`,
	Long: `create a new project with create command,for example:
	       ` + cliName + " create -p demo",
	Run: func(cmd *cobra.Command, args []string) {
		var destPath = fileutils.GetDestPath(project)
		if destPath == "" {
			return
		}

		w.Do(destPath)
	},
}

func init() {
	processFlags()
	initWorker()
	rootCmd.AddCommand(createCmd)
}

func processFlags() {
	createCmd.Flags().StringVarP(&project, "project", "p", "", "项目名称，可以是目录，为必选项")
	createCmd.MarkFlagRequired("project")
}

func initWorker() {
	w = worker.GetWorker("git")
}
