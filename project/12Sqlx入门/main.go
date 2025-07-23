package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Sqlx入门 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

// Sqlx入门 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
type Book struct {
	ID     int
	Title  string
	Author string
	Price  float64
}

func main() {
	// 创建db连接
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("===============================", db, err)

	// Sqlx入门 题目1：使用SQL扩展库进行查询
	// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
	// 要求 ：
	// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。

	// 创建表结构
	db.AutoMigrate(&Employee{})

	addEmployee := []Employee{{Name: "Jack", Department: "技术部", Salary: 17000},
		{Name: "Terro", Department: "财务部", Salary: 7000},
		{Name: "Seven", Department: "法务部", Salary: 8000},
		{Name: "Mebius", Department: "技术部", Salary: 27000}}
	db.Debug().Create(&addEmployee)

	// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	rows, _ := db.Debug().Model(&Employee{}).Where("Department", "技术部").Rows()

	defer rows.Close()

	var employee Employee
	employees := []Employee{}
	for rows.Next() {
		// ScanRows 将一行扫描至 employee
		db.ScanRows(rows, &employee)

		employees = append(employees, employee)
	}

	fmt.Println(employees)

	// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	row, _ := db.Debug().Model(&Employee{}).Order("Salary desc").Limit(1).Rows()

	defer row.Close()

	for row.Next() {
		// ScanRows 将一行扫描至 employee
		db.ScanRows(row, &employee)
	}

	fmt.Println(employee)

	// Sqlx入门 题目2：实现类型安全映射
	// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	// 要求 ：
	// 定义一个 Book 结构体，包含与 books 表对应的字段。
	// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

	// 创建表结构
	db.AutoMigrate(&Book{})

	addBook := []Book{{Title: "钢铁是怎样炼成的", Author: "保尔", Price: 69.9},
		{Title: "天方奇谭", Author: "佚名", Price: 39.9},
		{Title: "十万个为什么", Author: "洋洋", Price: 99.9},
		{Title: "爱迪生传奇", Author: "喜多", Price: 109.9}}
	db.Debug().Create(&addBook)

	bookRows, _ := db.Debug().Model(&Book{}).Where("Price > ?", 50).Rows()

	defer bookRows.Close()

	var book Book
	books := []Book{}
	for bookRows.Next() {
		// ScanRows 将一行扫描至 book
		db.ScanRows(bookRows, &book)

		books = append(books, book)
	}

	fmt.Println(books)

}
