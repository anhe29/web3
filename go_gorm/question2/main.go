package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 使用SQL扩展库进行查询

//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
//要求 ：
//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

type Employees struct {
	ID         int    `grom:"column:id"`
	NAME       string `grom:"column:name"`
	DEPARTMENT string `grom:"column:department"`
	SALARY     int    `grom:"column:salary"`
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}

	err = db.AutoMigrate(&Employees{})
	if err != nil {
		panic("迁移数据库失败" + err.Error())
	}

	var employees []Employees
	err = db.Where("department = ?", "技术部").Find(&employees).Error
	if err != nil {
		panic("查询失败" + err.Error())
	}
	fmt.Println(employees)

	var employee Employees

	err = db.Order("salary desc").First(&employee).Error
	if err != nil {
		panic("查询失败" + err.Error())
	}

	fmt.Println(employee)

	task2()
}

//题目2：实现类型安全映射
//假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
//要求 ：
//定义一个 Book 结构体，包含与 books 表对应的字段。
//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

type Books struct {
	ID     int    `grom:"column:id"`
	TITLE  string `grom:"column:title"`
	AUTHOR string `grom:"column:author"`
	PRICE  int    `grom:"column:price"`
}

func task2() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}
	err = db.AutoMigrate(&Books{})
	if err != nil {
		panic("迁移数据库失败：" + err.Error())
	}
	var books []Books
	err = db.Where("price > ?", 50).Find(&books).Error

	fmt.Println(books)
}
