package main

import "fmt"

/*
原型模式（Prototype Pattern）是一种创建型设计模式，它通过复制现有对象来创建新对象，而不是通过实例化新的对象。
在原型模式中，我们创建一个原型对象，然后通过复制这个原型对象来创建新的对象。
这种方式通常比通过构造函数直接创建对象更加灵活和高效。
*/

// Cloneable 定义了克隆方法
type Cloneable interface {
	Clone() Cloneable
}

// ConcretePrototype 实现了Cloneable接口的具体原型
type ConcretePrototype struct {
	data int
}

// Clone 实现了克隆方法
func (c *ConcretePrototype) Clone() Cloneable {
	// 创建一个新对象并复制原始对象的数据
	return &ConcretePrototype{data: c.data}
}

// NewConcretePrototype 创建ConcretePrototype的实例
func NewConcretePrototype(data int) *ConcretePrototype {
	return &ConcretePrototype{data: data}
}

func main() {
	// 创建原型对象
	prototype := NewConcretePrototype(42)

	// 克隆原型对象
	clone1 := prototype.Clone().(*ConcretePrototype)
	clone2 := prototype.Clone().(*ConcretePrototype)

	// 验证克隆对象是否成功
	fmt.Println("Original Object:", prototype.data)
	fmt.Println("Clone 1 Object:", clone1.data)
	fmt.Println("Clone 2 Object:", clone2.data)
}
