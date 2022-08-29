package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	fileUtils "clitemplate/fileutils"
	gitOperation "clitemplate/gitoperation"

	_ "clitemplate/config"
)

var (
	project string
	mobile  bool
	remote  bool
)

// var cliName = viper.Get("cliName").(string)

var createCmd = &cobra.Command{
	Use: "create",
	Short: `create a new project with create command,
	      for example: ` + cliName + ` create -p demo
	      see more, using command: ` + cliName + ` create -h`,
	Long: `create a new project with create command,for example:
	       ` + cliName + " create -p demo",
	Run: func(cmd *cobra.Command, args []string) {
		var destPath string
		if fileUtils.IsDir(project) {
			destPath = project
			fileUtils.MkDir(destPath)
		} else {
			destPath = fileUtils.GetCurrentJointDir(project)
			fileUtils.MkDir(destPath)
			if !fileUtils.IsDir(destPath) {
				fmt.Println("dest path error....", destPath)
				return
			}
		}

		if !remote {
			var sourcePath string
			if mobile {
				sourcePath = filepath.Join("template", "mobile")
			} else {
				sourcePath = filepath.Join("template", "pc")
			}
			// fmt.Println(sourcePath, destPath)
			if err := fileUtils.CopyDir(sourcePath, destPath); err != nil {
				fmt.Println("error", sourcePath, destPath, err)
			}
			return
		}

		remoteAddr := viper.Get("remoteAddr").(string)
		var branch string
		if mobile {
			branch = viper.Get("branchPC").(string)
		} else {
			branch = viper.Get("branchMobile").(string)
		}
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
	// 换成实际的url
	createCmd.Flags().BoolVarP(&remote, "remote", "r", false, "是否从远程GitLab仓库拉取, 需要Gitlab仓库权限")
	createCmd.Flags().BoolVarP(&mobile, "mobile", "m", false, "是否是移动端, 默认为false, 表示PC端")
	createCmd.MarkFlagRequired("project")

}
