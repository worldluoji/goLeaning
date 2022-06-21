package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory" // 在实例化 warden 时指定了用内存方式做数据持久化，因此需要导入 memory
)

/*
* ID：策略的标识。
* Description：策略的描述。
* Subjects：策略的主题。<> 内为正则表达式。
* Resources：策略的资源。
* Actions：策略的操作类型。
* Effect：AllowAccess 表示允许。DenyAccess 表示拒绝。
* Conditions：，描述策略生效的约束条件。

如下策略的意思是：对于主题users:manager和users:admin下的资源resources:users, username为luoji的有
对该资源的delete、create和update权限
*/
var pol = &ladon.DefaultPolicy{
	ID:          "1",
	Description: "demo for ladon test",
	Subjects:    []string{"users:<manager|admin>"},
	Resources: []string{
		"resources:users",
	},
	Actions: []string{"delete", "<create|update>"},
	Effect:  ladon.AllowAccess,
	Conditions: ladon.Conditions{
		"username": &ladon.StringEqualCondition{
			Equals: "luoji",
		},
	},
}

// // Metric is used to expose metrics about authz
// type Metric interface {
// 	// RequestDeniedBy is called when we get explicit deny by policy
// 	RequestDeniedBy(Request, Policy)
// 	// RequestAllowedBy is called when a matching policy has been found.
// 	RequestAllowedBy(Request, Policies)
// 	// RequestNoMatch is called when no policy has matched our request
// 	RequestNoMatch(Request)
// 	// RequestProcessingError is called when unexpected error occured
// 	RequestProcessingError(Request, Policy, error)
// }

type prometheusMetrics struct{}

// prometheusMetrics实现了下面4个函数，那么它就是一个ladon.Metric
func (mtr *prometheusMetrics) RequestDeniedBy(r ladon.Request, p ladon.Policy) {
	log.Println("Request deny ", r.Subject)
}

func (mtr *prometheusMetrics) RequestAllowedBy(r ladon.Request, policies ladon.Policies) {
	log.Println("Request allowed ", r.Subject)
}

func (mtr *prometheusMetrics) RequestNoMatch(r ladon.Request) {
	log.Println("Request Not Match ", r.Subject)
}

func (mtr *prometheusMetrics) RequestProcessingError(r ladon.Request, p ladon.Policy, err error) {
	log.Println("Request error ", r.Subject, err)
}

// 实例化 warden
var warden = &ladon.Ladon{
	// 数据持久化
	Manager: memory.NewMemoryManager(),
	// ladon.AuditLoggerInfo，该 AuditLogger 会在授权时打印调用的策略到标准错误
	// 要实现一个新的 AuditLogger，你只需要实现 AuditLogger 接口就可以了。比如，我们可以实现一个 AuditLogger，将授权日志保存到 Redis 或者 MySQL 中。
	AuditLogger: &ladon.AuditLoggerInfo{},

	Metric: &prometheusMetrics{},
}

func init() {
	// 添加策略
	warden.Manager.Create(pol)
}

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello Ladon!!!")
	})

	r.POST("/check", func(c *gin.Context) {
		accessRequest := &ladon.Request{}

		if err := c.BindJSON(accessRequest); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
			return
		}

		// 判断是否拥有权限
		var message string
		if err := warden.IsAllowed(accessRequest); err != nil {
			message = "Not allowed"
		} else {
			message = "Allowed"
		}

		c.JSON(200, gin.H{
			"message": message,
		})

	})

	srv := &http.Server{
		Addr:    ":8083",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

/*
$ curl -X POST -H 'Content-Type: application/json' -d '{"context":{"username":"luoji"}, "subject": "users:manager","action" : "delete", "resource": "resources:users"}' http://localhost:8083/check
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   132  100    21  100   111   2333  12333 --:--:-- --:--:-- --:--:-- 14666{"message":"Allowed"}


$ curl -X POST -H 'Content-Type: application/json' -d '{"context":{"username":"zhangmiaomiao"}, "subject": "users:manager","action" : "delete", "resource": "resources:users"}' http://localhost:8083/check
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   144  100    25  100   119   2777  13222 --:--:-- --:--:-- --:--:-- 18000{"message":"Not allowed"}

*/
