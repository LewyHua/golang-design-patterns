package main

import "fmt"

/*
接口隔离原则（Interface Segregation Principle，简称ISP）是SOLID原则中的一项，
它强调客户端不应该依赖于它不需要的接口。
具体来说，接口隔离原则要求将大型接口拆分为多个更小的、特定于客户端的接口，以减少不必要的依赖和复杂性。
*/

type Shape interface {
	Render()
	Move(x, y int)
}

type Rectangle struct {
	// 矩形的属性
}

func (r *Rectangle) Render() {
	fmt.Println("渲染矩形")
}

func (r *Rectangle) Move(x, y int) {
	fmt.Printf("移动矩形到坐标 (%d, %d)\n", x, y)
}

/*
在上面的代码中，我们定义了一个通用的 Shape 接口，它包含了 Render 和 Move 两个方法。
然后，我们实现了一个矩形类型 Rectangle，它实现了 Shape 接口的方法。
然而，问题出现了，如果我们添加了一个新的形状类型，比如圆形，但它不需要移动操作，那么按照目前的设计，
我们还是需要实现 Move 方法，这违反了接口隔离原则。
为了改进这个设计，我们可以拆分 Shape 接口为两个更小的接口：Renderer 和 Mover，以满足不同形状的需求
*/

type Renderer interface {
	Render()
}

type Mover interface {
	Move(x, y int)
}

type Circle struct {
	// 圆形的属性
}

func (c *Circle) Render() {
	fmt.Println("渲染圆形")
}
