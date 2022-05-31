package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg         = pflag.StringP("config", "c", "", "Configuration file.")
	help        = pflag.BoolP("help", "h", false, "Show this help message.")
	CONFIG_PATH = "D:\\go\\src\\github.com\\luoji_demo\\other_demo\\buildtool\\viper"
	token       = pflag.StringP("token", "t", "", "token for test")
)

func init() {
	viper.SetDefault("TestParam", "333")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
	viper.Set("user.username", "colin")
	viper.BindPFlag("token", pflag.Lookup("token"))
}

func main() {
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	// 从配置文件中读取配置
	if *cfg != "" {
		viper.SetConfigFile(*cfg)   // 指定配置文件名，如果调用 SetConfigFile 直接指定了配置文件名，并且配置文件名没有文件扩展名时，需要显式指定配置文件的格式，以使 Viper 能够正确解析配置文件。
		viper.SetConfigType("yaml") // 如果配置文件名中没有文件扩展名，则需要指定配置文件的格式，告诉viper以何种格式解析文件
	} else {
		viper.AddConfigPath(".")         // 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath(CONFIG_PATH) // 配置文件搜索路径，可以设置多个配置文件搜索路径，需要注意添加搜索路径的顺序，Viper 会根据添加的路径顺序搜索配置文件，如果找到则停止搜索。
		viper.SetConfigName("config")    // 配置文件名称（没有文件扩展名）
	}

	// 监听和重新读取配置文件
	// viper.WatchConfig()
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	// 配置文件发生变更之后会调用的回调函数
	// 	fmt.Println("Config file changed:", e.Name)
	// })

	if err := viper.ReadInConfig(); err != nil { // 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	fmt.Printf("Used configuration file is: %s\n", viper.ConfigFileUsed())
	fmt.Println(viper.Get("TestParam"), viper.Get("user.username"))
	fmt.Println(viper.GetInt("test"), viper.GetInt("spec.replicas"), viper.GetString("spec.image"))
	fmt.Println(viper.Get("token") == *token)
}
