package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	project   string
	mobile    bool
	directory string
	url       string
	branch    string
)

func init() {
	rootCmd.Flags().StringVarP(&project, "project", "p", "", "项目名称")
	rootCmd.Flags().StringVarP(&directory, "directory", "d", "", "拉取代码的本地保存目录，默认为当前目录 + 项目名称")
	// 换成实际的url
	rootCmd.Flags().StringVarP(&url, "url", "u", "https://gitee.com/ant-design/ant-design.git", "github/gitlab仓库地址")
	rootCmd.Flags().StringVarP(&branch, "branch", "b", "", "分支名称")
	rootCmd.Flags().BoolVarP(&mobile, "mobile", "m", false, "是否是移动端, 默认为false, 表示PC端")

	// TODO 根据m参数拉取不同的代码模板
}

var rootCmd = &cobra.Command{
	Use:   "Gittool",
	Short: "Gittool is used for getting code from github",
	Long:  `Gittool is used for getting code from github`,
	Run: func(cmd *cobra.Command, args []string) {
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
