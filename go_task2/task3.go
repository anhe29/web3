package main

import "fmt"

// func main() {
// 	rect := Rectangle{Width: 5, Height: 3}
// 	circle := Circle{Radius: 2}

// 	fmt.Printf("Rectangle Area: %.2f, Perimeter: %.2f\n", rect.Area(), rect.Perimeter())
// 	fmt.Printf("Circle Area: %.2f, Perimeter: %.2f\n", circle.Area(), circle.Perimeter())

// 	emp := Employee{
// 		Person: Person{
// 			Name: "John",
// 			Age:  30,
// 		},
// 		EmployeeID: 12345,
// 	}
// 	emp.PrintInfo()
// }

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
// 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

//题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
// 组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
//考察点 ：组合的使用、方法接收者。

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person     Person
	EmployeeID int
}

func (emp *Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d\n", emp.Person.Name, emp.Person.Age, emp.EmployeeID)
}
