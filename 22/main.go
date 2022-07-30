package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type student struct {
	id           int32
	name         string
	age          int32
	grade        int32
	phone_number string
}

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

func add(db *sql.DB, stu *student) int64 {
	prepare, _ := db.Prepare("INSERT INTO gx_students SET name=?,age=?,grade=?,phone_number=?")
	res, _ := prepare.Exec(stu.name, stu.age, stu.grade, stu.phone_number)

	idVal, _ := res.LastInsertId()

	fmt.Println("插入数据id为", idVal)

	return idVal
}

func del(db *sql.DB, id int64) {
	prepare, _ := db.Prepare("DELETE FROM gx_students where id=?")
	res, _ := prepare.Exec(id)

	affectLine, _ := res.RowsAffected()

	fmt.Println("删除数据", string(affectLine), '条')
}

func update(db *sql.DB, stu *student, id int32) {
	prepare, _ := db.Prepare(`Update gx_students SET name=?,age=?,grade=? where id=?`)
	prepare.Exec(stu.name, stu.age, stu.grade, id)

	fmt.Println("更新成功")
}

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

func main() {
	stu := student{
		name:         "李四",
		age:          20,
		grade:        7,
		phone_number: "13111111111",
	}

	db := connectDB()

	id := add(db, &stu)

	stu.name = "王五"

	update(db, &stu, int32(id))
	del(db, id)

	stus := query(db)
	fmt.Println(stus)

}
