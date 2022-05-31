# Gorm
GORM 是 Go 语言的 ORM 包。GORM 有两个版本，V1和V2。遵循用新不用旧的原则。

## Gorm的优势
- 功能全。使用 ORM 操作数据库的接口，GORM 都有，可以满足我们开发中对数据库调用的各类需求。
- 支持钩子方法。这些钩子方法可以应用在 Create、Save、Update、Delete、Find 方法中。
- 开发者友好，调用方便。
- 支持 Auto Migration。
- 支持关联查询。
- 支持多种关系数据库，例如 MySQL、Postgres、SQLite、SQLServer 等

## demo
-》main.go

## 默认映射字段名称
```
type Animal struct {
  AnimalID int64        // 列名 `animal_id`
  Birthday time.Time    // 列名 `birthday`
  Age      int64        // 列名 `age`
}
```
上述模型对应的表名为 animals ，列名分别为 animal_id 、 birthday 和 age,
但是建议像demo中一样，通过注解明确指定映射字段。


## 增删查改说明

### Create
#### db.Create 函数会返回如下 3 个值：
- user.ID：返回插入数据的主键，这个是直接赋值给 user 变量。
- result.Error：返回 error。
- result.RowsAffected：返回插入记录的条数。

#### 数据量大时可以批量插入，提高效率
```
type User struct {  
    gorm.Model  Name         
    string  Age          
    uint8  Birthday     
    *time.Time
}

var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
DB.Create(&users)
for _, user := range users {
  user.ID // 1,2,3
}
```

### Delete
#### Where + Delete删除
```
// DELETE from users where id = 10 AND name = "jinzhu";
db.Where("name = ?", "jinzhu").Delete(&user)
```

#### 通过主键删除：
```
// DELETE FROM users WHERE id = 10;
db.Delete(&User{}, 10)
```

#### 软删除
如果模型包含了一个 gorm.DeletedAt 字段，GORM 在执行删除操作时，会软删除该记录。
```
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;
db.Where("age = ?", 20).Delete(&User{})

// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;
db.Where("age = 20").Find(&user)
```
对于软删除，GORM 并没有真正把记录从数据库删除掉，而是只更新了 deleted_at 字段；
在查询时，GORM 查询条件中新增了AND deleted_at IS NULL条件，所以这些被设置过 deleted_at 字段的记录不会被查询到。

#### 硬删除
```
// DELETE FROM orders WHERE id=10;
db.Unscoped().Delete(&order)
```
或者，你也可以在模型中去掉 gorm.DeletedAt。


### Update
#### 最常用的更新方法
```
db.First(&user)

user.Name = "jinzhu 2"
user.Age = 100
// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
db.Save(&user)
```

#### 更新单列
```
// UPDATE users SET age=200, updated_at='2013-11-17 21:34:10' WHERE name='colin';
db.Model(&User{}).Where("name = ?", "colin").Update("age", 200)
```

#### 更新多列

```
// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE name = 'colin';
db.Model(&user).Where("name", "colin").Updates(User{Name: "hello", Age: 18, Active: false})
```

### Query
#### 检索单个记录的示例代码
```
// 获取第一条记录（主键升序）
// SELECT * FROM users ORDER BY id LIMIT 1;
db.First(&user)

// 获取最后一条记录（主键降序）
// SELECT * FROM users ORDER BY id DESC LIMIT 1;
db.Last(&user)
result := db.First(&user)
result.RowsAffected // 返回找到的记录数
result.Error        // returns error

// 检查 ErrRecordNotFound 错误
errors.Is(result.Error, gorm.ErrRecordNotFound)
```
如果 model 类型没有定义主键，则按第一个字段排序。


#### 查询所有符合条件的
```
users := make([]*User, 0)
// SELECT * FROM users WHERE name <> 'jinzhu';
db.Where("name <> ?", "jinzhu").Find(&users)
```

#### 智能选择字段
```
type APIUser struct {
  ID   uint
  Name string
}

// SELECT `id`, `name` FROM `users` LIMIT 10;
db.Model(&User{}).Limit(10).Find(&APIUser{})
```

### 高级查询
#### 排序
```
// SELECT * FROM users ORDER BY age desc, name;
db.Order("age desc, name").Find(&users)
```

#### limit & offset
```
// SELECT * FROM users OFFSET 5 LIMIT 10; Offset 指定从第几条记录开始查询，Limit 指定返回的最大记录数
db.Limit(10).Offset(5).Find(&users)
```

#### Distinct
```
db.Distinct("name", "age").Order("name, age desc").Find(&results)
```

#### Count
```
var count int64
// SELECT count(1) FROM users WHERE name = 'jinzhu'; (count)
db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
```
GORM 还支持很多高级查询功能，比如内联条件、Not 条件、Or 条件、Group & Having、Joins、Group、FirstOrInit、FirstOrCreate、迭代、FindInBatches 等。


### 原生SQL
#### 原生查询
```
type Result struct {
  ID   int
  Name string
  Age  int
}

var result Result
db.Raw("SELECT id, name, age FROM users WHERE name = ?", 3).Scan(&result)
```

#### 原生执行
```
db.Exec("DROP TABLE users")
db.Exec("UPDATE orders SET shipped_at=? WHERE id IN ?", time.Now(), []int64{1,2,3})
```

### Gorm钩子
#### 例如下面这个在插入记录前执行的钩子
```
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    u.UUID = uuid.New()

    if u.Name == "admin" {
        return errors.New("invalid name")
    }
    return
}
```

### 参考文档
https://gorm.io/zh_CN/docs/index.html