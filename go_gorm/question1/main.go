package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Students struct {
	ID    int    `gorm:"column:id"`
	NAME  string `gorm:"column:name"`
	AGE   int    `gorm:"column:age"`
	GRADE string `gorm:"column:grade"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}

	// 自动迁移：确保表存在
	err = db.AutoMigrate(&Students{})
	if err != nil {
		panic("迁移数据库失败：" + err.Error())
	}
	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	db.Create(&Students{ID: 1, NAME: "张三", AGE: 20, GRADE: "三年级"})
	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	var students []Students
	err = db.Where("age > ?", 18).Find(&students).Error
	if err != nil {
		panic("查询年龄大于18失败" + err.Error())
	}
	fmt.Println("年龄大于18:", students)

	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	err = db.Model(&Students{}).Where("name = ?", "张三").Update("grade", "四年级").Error
	if err != nil {
		panic("更新失败： " + err.Error())
	}
	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	if err := db.Where("age < ?", 15).Delete(&Students{}).Error; err != nil {
		panic("删除失败: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("获取底层数据库连接失败：" + err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}
}
