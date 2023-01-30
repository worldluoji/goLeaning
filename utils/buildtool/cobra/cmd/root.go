package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Verbose bool
	Source  string
	author  string
	Region  string
)

// 可以在 init() 函数中定义标志和处理配置
func init() {
	// 标志可以是“持久的”，这意味着该标志可用于它所分配的命令以及该命令下的每个子命令。可以在 rootCmd 上定义持久标志.
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// 也可以分配一个本地标志，本地标志只能在它所绑定的命令上使用,--source标志只能在 rootCmd 上引用，而不能在 rootCmd 的子命令上引用。
	rootCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")

}

// 将标志绑定到 Viper,这样就可以使用 viper.Get() 获取标志的值
func init() {
	rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))

	rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
	// 设置为必选
	rootCmd.MarkFlagRequired("region")
	viper.BindPFlag("region", rootCmd.Flags().Lookup("region"))
}

var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "cobra is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println(viper.Get("author"), viper.Get("region"))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
