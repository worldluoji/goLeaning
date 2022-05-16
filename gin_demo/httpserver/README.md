# 更新依赖包
go mod tidy 

# 运行程序
go run main.go

# 创建产品
curl -XPOST -H"Content-Type: application/json" -d'{"username":"colin","name":"iphone12","category":"phone","price":8000,"description":"cannot afford"}' http://127.0.0.1:8098/v1/products

# 获取产品信息
curl -XGET http://127.0.0.1:8098/v1/products/iphone12
