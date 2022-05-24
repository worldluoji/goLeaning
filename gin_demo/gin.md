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