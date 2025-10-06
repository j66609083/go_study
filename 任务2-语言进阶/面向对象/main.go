package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length, width float64
}

type Circle struct {
	radius float64
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID string
	Person
}

func (e *Employee) PrintInfo() {
	fmt.Printf("EmployeeID: %s, Name: %s, Age: %d\n", e.EmployeeID, e.Name, e.Age)
}

func main() {

	/*
		题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
		然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
		在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
	*/
	rect := Rectangle{length: 5, width: 3}
	circ := Circle{radius: 4}

	fmt.Println("Rectangle Area:", rect.Area(), "  Perimeter:", rect.Perimeter())
	fmt.Println("Circle Area:", circ.Area(), "  Perimeter:", circ.Perimeter())

	/*
		题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
		再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
		为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	*/
	e := Employee{EmployeeID: "E123", Person: Person{Name: "Alice", Age: 30}}
	e.PrintInfo()
}
