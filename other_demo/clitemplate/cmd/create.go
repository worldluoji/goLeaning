package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	gitOperation "clitemplate/gitOperation"
	fileUtils "clitemplate/fileUtils"

	_ "clitemplate/config"
)


var (
	project   string
	mobile    bool
    remote    bool
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new project with create command",
	Long:  `create a new project with create command,for example:
	` + viper.Get("cliName").(string) + " create demo",
	Run: func(cmd *cobra.Command, args []string) {
		var destPath string
		if fileUtils.IsDir(project) {
			destPath = project
		} else {
			destPath = fileUtils.GetCurrentJointDir(project)
			if !fileUtils.IsDir(destPath) {
				fmt.Println("dest path error....", destPath)
				return
			}
		}
		fileUtils.MkDir(destPath)

		if (!remote) {
			var sourcePath string
			if (mobile) {
				sourcePath = viper.Get("SourceTemplateMobile").(string)
			} else {
				sourcePath = viper.Get("SourceTemplatePC").(string)
			}
			fmt.Println(sourcePath, destPath)
			fileUtils.CopyDir(sourcePath, destPath)
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
	createCmd.Flags().StringVarP(&project, "project", "p", "", "项目名称，可以是目录，默认为当前目录")
	// 换成实际的url
	createCmd.Flags().BoolVarP(&remote, "remote", "r", false, "是否从远程GitLab仓库拉取, 需要Gitlab仓库权限")
	createCmd.Flags().BoolVarP(&mobile, "mobile", "m", false, "是否是移动端, 默认为false, 表示PC端")
	rootCmd.MarkFlagRequired("project")
	
}