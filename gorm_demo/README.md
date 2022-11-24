# README

可以使用docker启动一个mariadb进行测试：
```
docker run -d -p 3309:3306 -e MARIADB_ROOT_PASSWORD=xxxxxx mariadb
```

启动
```
go run main.go
```