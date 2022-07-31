# 持久化存储之数据库

本章主要通过数据库的增删改查来实现数据的持久化，一个简单的demo

## 安装第三方包

1. 初始化一个 go 的项目`go mod init`
2. 使用`go install/get`来安装对应的第三方依赖

```bash
# 初始化依赖
go mod init sql-exp

# 安装第三方包
go get -u github.com/go-sql-driver/mysql
```

## 使用 mysql 包进行数据库操作

具体的使用方法可以参考开源包的官方文档: [go-sql-driver](https://github.com/go-sql-driver/mysql#usage)

### 数据库的连接

引用包：`database/sql` 和 `github.com/go-sql-driver/mysql`

创建步骤：
1. 使用`sql.Open`打开数据库
2. 对`db`进行参数的设置

```go
func connectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@/students")
	if err != nil {
		fmt.Println("数据库链接错误", err)
		panic(err)
	}

	// 重用连接的最大时间
	db.SetConnMaxLifetime(time.Hour * 1)
	// 最大连接数量
	db.SetMaxOpenConns(5)
	// 最大空闲数量
	db.SetMaxIdleConns(5)

	fmt.Println("链接成功")

	return db
}
```

### 数据库数据库的增删改查

#### 插入数据

数据修改操作：
1. 使用`Prepare`输入需要预执行的sql语句
2. 通过`prepare.Exec`方法对占位的参数进行填充
3. 完成相关操作

```go
func add(db *sql.DB, stu *student) int64 {
	prepare, _ := db.Prepare("INSERT INTO gx_students SET name=?,age=?,grade=?,phone_number=?")
	res, _ := prepare.Exec(stu.name, stu.age, stu.grade, stu.phone_number)

	idVal, _ := res.LastInsertId()

	fmt.Println("插入数据id为", idVal)

	return idVal
}
```

#### 修改数据

```go
func update(db *sql.DB, stu *student, id int32) {
	prepare, _ := db.Prepare(`Update gx_students SET name=?,age=?,grade=? where id=?`)
	prepare.Exec(stu.name, stu.age, stu.grade, id)

	fmt.Println("更新成功")
}
```

#### 删除数据

```go
func del(db *sql.DB, id int64) {
	prepare, _ := db.Prepare("DELETE FROM gx_students where id=?")
	res, _ := prepare.Exec(id)

	affectLine, _ := res.RowsAffected()

	fmt.Println("删除数据", string(affectLine), '条')
}
```

#### 查询数据

查询方法：
1. 使用`db.Query`输入sql语句进行执行
2. 对查询出来的`rows`使用 `for rows.Next()`查询是否还有吓一条数据
3. 使用`rows.Scan`逐条获取下条数据内容，，**在使用Scan**的时候需要传入和表中数量一样字段的引用，进行数据的接收处理后，整理成切片输出

```go
func query(db *sql.DB) []student {
	var stus []student

	rows, err := db.Query(`Select * from gx_students`)

	if err != nil {
		fmt.Println("err ->", err)
		panic(err)
	}

	var stu student

	var (
		createAt interface{}
		updateAt interface{}
		deleteAt interface{}
	)

	for rows.Next() {
		fmt.Println()
		if err := rows.Scan(&stu.id, &stu.name, &stu.age, &createAt, &stu.phone_number, &deleteAt, &stu.grade, &updateAt); err != nil {
			fmt.Println("数据库查询错误", err)
			panic(err)
		}

		stus = append(stus, stu)
	}

	return stus
}
```