# Gin
## Gin 是什么？ 
Gin是用 Go 语言编写的 Web 框架，功能完善，使用简单，性能很高。Gin 核心的路由功能是通过一个定制版的HttpRouter来实现的，具有很高的路由性能。

## Gin 具有如下特性：
- 轻量级，代码质量高，性能比较高；
- 项目目前很活跃，并有很多可用的 Middleware；
- 作为一个 Web 框架，功能齐全，使用起来简单。

## Gin 的一些核心功能
- 支持 HTTP 方法：GET、POST、PUT、PATCH、DELETE、OPTIONS。
- 支持不同位置的 HTTP 参数：路径参数（path）、查询字符串参数（query）、表单参数（form）、HTTP 头参数（header）、消息体参数（body）。
- 支持 HTTP 路由和路由分组。
- 支持 middleware 和自定义 middleware。
- 支持自定义 Log。
- 支持 binding 和 validation，支持自定义 validator。
- 可以 bind 如下参数：query、path、body、header、form。
- 支持重定向。
- 支持 basic auth middleware。支持自定义 HTTP 配置。
- 支持优雅关闭。
- 支持 HTTP2。
- 支持设置和获取 cookie。

<br>

## 中间件

### 1. Gin的中间件可以做什么？
Gin 支持中间件，HTTP 请求在转发到实际的处理函数之前，会被一系列加载的中间件进行处理。在中间件中，可以解析 HTTP 请求做一些逻辑处理，例如：跨域处理或者生成 X-Request-ID 并保存在 context 中，以便追踪某个请求。处理完之后，可以选择中断并返回这次请求，也可以选择将请求继续转交给下一个中间件处理。当所有的中间件都处理完之后，请求才会转给路由函数进行处理。

<br>

### 2. Gin中间件的优缺点
通过中间件，可以实现对所有请求都做统一的处理，提高开发效率，并使我们的代码更简洁。但是，因为所有的请求都需要经过中间件的处理，可能会增加请求延时。一些建议如下：
- 中间件做成可加载的，通过配置文件指定程序启动时加载哪些中间件。
- 只将一些通用的、必要的功能做成中间件。
- 在编写中间件时，一定要保证中间件的代码质量和性能。

<br>

### 3. Gin中使用中间件
#### 加载中间件
在 Gin 中，可以通过 gin.Engine 的 Use 方法来加载中间件。中间件可以加载到不同的位置上，而且不同的位置作用范围也不同。例子如下：
```
router := gin.New()
router.Use(gin.Logger(), gin.Recovery()) // 中间件作用于所有的HTTP请求
v1 := router.Group("/v1").Use(gin.BasicAuth(gin.Accounts{"foo": "bar", "colin": "colin404"})) // 中间件作用于v1 group
v1.POST("/login", Login).Use(gin.BasicAuth(gin.Accounts{"foo": "bar", "colin": "colin404"})) //中间件只作用于/v1/login API接口
```

#### 内置中间件
- gin.Logger()：Logger 中间件会将日志写到 gin.DefaultWriter，gin.DefaultWriter 默认为 os.Stdout。
- gin.Recovery()：Recovery 中间件可以从任何 panic 恢复，并且写入一个 500 状态码。
- gin.CustomRecovery(handle gin.RecoveryFunc)：类似 Recovery 中间件，但是在恢复时还会调用传入的 handle 方法进行处理。
- gin.BasicAuth()：HTTP 请求基本认证（使用用户名和密码进行认证）。

#### 自定义中间件
- Gin 还支持自定义中间件。中间件其实是一个函数，函数类型为 gin.HandlerFunc，HandlerFunc 底层类型为 func(*Context)
- Logger中间件例子见-> ./middleware/main.go