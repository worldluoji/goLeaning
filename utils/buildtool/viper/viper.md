# viper读取配置文件
大型项目配置很多，通过命令行参数传递就变得很麻烦，不好维护。标准的解决方案是将这些配置信息保存在配置文件中，由程序启动时加载和解析。
Go 生态中有很多包可以加载并解析配置文件，目前最受欢迎的是 Viper 包。

## 优先级
Viper 可以从不同的位置读取配置，不同位置的配置具有不同的优先级，高优先级的配置会覆盖低优先级相同的配置，按优先级从高到低排列如下：

- 通过 viper.Set 函数显示设置的配置
- 命令行参数
- 环境变量
- 配置文件
- Key/Value 存储
- 默认值
<br>
这里需要注意，Viper 配置键不区分大小写。

<br>

## 读入配置
读入配置，就是将配置读入到 Viper 中，有如下读入方式：
- 设置默认的配置文件名
```
viper.SetDefault("ContentDir", "content")
viper.SetDefault("LayoutDir", "layouts")
viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
```
- 读取配置文件
```
Viper 提供了如下方法来读取配置：
Get(key string) interface{}
Get<Type>(key string) <Type>
AllSettings() map[string]interface{}
IsSet(key string) : bool

每一个 Get 方法在找不到值的时候都会返回零值。为了检查给定的键是否存在，可以使用 IsSet() 方法。
<Type>可以是 Viper 支持的类型，首字母大写：Bool、Float64、Int、IntSlice、String、StringMap、StringMapString、StringSlice、Time、Duration。例如：GetInt()。

demo见main/main.go

```
- 监听和重新读取配置文件
```
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
   // 配置文件发生变更之后会调用的回调函数
  fmt.Println("Config file changed:", e.Name)
})
```
- 从 io.Reader 读取配置
- 从环境变量读取
- 从命令行标志读取
- 从远程 Key/Value 存储读取
<br>


<br>

## 环境变量
Viper 还支持环境变量，通过如下 5 个函数来支持环境变量：
- AutomaticEnv()
- BindEnv(input …string) error
```
BindEnv 需要一个或两个参数。第一个参数是键名，第二个是环境变量的名称，环境变量的名称区分大小写。如果未提供 Env 变量名，则 Viper 将假定 Env 变量名为：环境变量前缀_键名全大写。例如：前缀为 VIPER，key 为 username，则 Env 变量名为VIPER_USERNAME。当显示提供 Env 变量名（第二个参数）时，它不会自动添加前缀。例如，如果第二个参数是 ID，Viper 将查找环境变量 ID。
```
- SetEnvPrefix(in string)
```
通过使用 SetEnvPrefix，可以告诉 Viper 在读取环境变量时使用前缀。BindEnv 和 AutomaticEnv 都将使用此前缀。比如，我们设置了 viper.SetEnvPrefix(“VIPER”)，当使用 viper.Get(“apiversion”) 时，实际读取的环境变量是VIPER_APIVERSION。
```
- SetEnvKeyReplacer(r *strings.Replacer)
- AllowEmptyEnv(allowEmptyEnv bool)

<br>
这里要注意：Viper 读取环境变量是区分大小写的
<br>

```
// 使用环境变量
os.Setenv("VIPER_USER_SECRET_ID", "QLdywI2MrmDVjSSv6e95weNRvmteRjfKAuNV")
os.Setenv("VIPER_USER_SECRET_KEY", "bVix2WBv0VPfrDrvlLWrhEdzjLpPCNYb")

viper.AutomaticEnv()                                             // 读取环境变量
viper.SetEnvPrefix("VIPER")                                      // 设置环境变量前缀：VIPER_，如果是viper，将自动转变为大写。
viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_")) // 将viper.Get(key) key字符串中'.'和'-'替换为'_'
viper.BindEnv("user.secret-key")
viper.BindEnv("user.secret-id", "USER_SECRET_ID") // 绑定环境变量名到key
```
这里要注意：Viper 读取环境变量是区分大小写的。

<br>

## Viper 结合 Pflag 
Viper 支持 Pflag 包，能够绑定 key 到 Flag。与 BindEnv 类似，在调用绑定方法时，不会设置该值，但在访问它时会设置。
- 对于单个标志，可以调用 BindPFlag() 进行绑定：
```
viper.BindPFlag("token", pflag.Lookup("token")) // 绑定单个标志
```
- 还可以绑定一组现有的 pflags（pflag.FlagSet）
```
viper.BindPFlags(pflag.CommandLine)             //绑定标志集
```

<br>

## 序列化和反序列化
### 反序列化
Viper 可以支持将所有或特定的值解析到结构体、map 等。可以通过两个函数来实现：
- Unmarshal(rawVal interface{}) error
- UnmarshalKey(key string, rawVal interface{}) error
```
type config struct {
  Port int
  Name string
  PathMap string `mapstructure:"path_map"`
}

var C config

err := viper.Unmarshal(&C)
if err != nil {
  t.Fatalf("unable to decode into struct, %v", err)
}
```

如果想要解析那些键本身就包含.(默认的键分隔符）的配置，则需要修改分隔符:
```
v := viper.NewWithOptions(viper.KeyDelimiter("::"))

v.SetDefault("chart::values", map[string]interface{}{
    "ingress": map[string]interface{}{
        "annotations": map[string]interface{}{
            "traefik.frontend.rule.type":                 "PathPrefix",
            "traefik.ingress.kubernetes.io/ssl-redirect": "true",
        },
    },
})

type config struct {
    Chart struct{
        Values map[string]interface{}
    }
}

var C config

v.Unmarshal(&C)
```

### 序列化成字符串
```
import (
    yaml "gopkg.in/yaml.v2"
    // ...
)

func yamlStringSettings() string {
    c := viper.AllSettings()
    bs, err := yaml.Marshal(c)
    if err != nil {
        log.Fatalf("unable to marshal config to YAML: %v", err)
    }
    return string(bs)
}
```