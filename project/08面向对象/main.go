package main

import "fmt"

////=====================================================================================================

// 面向对象 1.题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。 考察点 ：接口的定义与实现、面向对象编程风格。
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	name string
}

func (r Rectangle) Area() {
	fmt.Println(r.name, "Area")
}

func (r Rectangle) Perimeter() {
	fmt.Println(r.name, "Perimeter")

}

type Circle struct {
	name string
}

func (c Circle) Area() {
	fmt.Println(c.name, "Area")
}

func (c Circle) Perimeter() {
	fmt.Println(c.name, "Perimeter")
}

// 面向对象 题目 ：2.使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。考察点 ：组合的使用、方法接收者。
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() {
	fmt.Println("Name=", e.Name)
	fmt.Println("Age=", e.Age)
	fmt.Println("EmployeeID=", e.EmployeeID)
}

func main() {

	// 面向对象 1.题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。 考察点 ：接口的定义与实现、面向对象编程风格。
	rectangle := Rectangle{
		name: "rectangle",
	}
	rectangle.Area()
	rectangle.Perimeter()

	circle := Circle{
		name: "circle",
	}
	circle.Area()
	circle.Perimeter()

	// 面向对象 2.题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。考察点 ：组合的使用、方法接收者。
	employee := Employee{
		Person: Person{
			Name: "金",
			Age:  29,
		},
		EmployeeID: "0690",
	}
	employee.PrintInfo()

}
