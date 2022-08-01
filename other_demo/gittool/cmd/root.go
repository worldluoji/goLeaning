package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	mobile     bool
	typescript bool
	directory  string
	url        string
	branch     string
)

func init() {
	rootCmd.Flags().StringVarP(&directory, "directory", "d", "", "拉取代码的本地保存目录，默认为当前目录")
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "github/gitlab仓库地址")
	rootCmd.Flags().StringVarP(&branch, "branch", "b", "master", "分支名称")

	rootCmd.Flags().BoolVarP(&mobile, "mobile", "m", false, "是否是移动端模板，默认为false, 表示PC端模板")
	rootCmd.Flags().BoolVarP(&typescript, "typescript", "t", false, "是否使用typescript模板，默认为false, 表示使用JavaScript")
	if url == "" {
		// 换成实际的url
		if mobile {
			if typescript {
				url = "https://gitee.com/ant-design/ant-design.git"
			} else {
				url = "https://gitee.com/ant-design/ant-design.git"
			}
		} else {
			if typescript {
				url = "https://gitee.com/ant-design/ant-design.git"
			} else {
				url = "https://gitee.com/ant-design/ant-design.git"
			}
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "Gittool",
	Short: "Gittool is used for getting code from github",
	Long:  `Gittool is used for getting code from github`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("begin to get code...")
		if _, err := GitClone(directory, url, branch); err != nil {
			fmt.Println(err)
		}
		fmt.Println("get code finished...")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
