package main

import "fmt"

/*
工厂方法模式（Factory Method Pattern）是一种创建型设计模式，它定义了一个用于创建对象的接口，但将对象的实际创建延迟到子类中。
这意味着工厂方法模式允许一个类在其子类中选择创建对象的类型。
工厂方法模式的核心思想是：
1. 定义一个抽象的工厂接口，该接口声明一个或多个工厂方法，用于创建对象。
2. 创建具体的工厂类，实现工厂接口，并提供对象的实际创建逻辑。
3. 客户端代码通过调用工厂方法来创建对象，而不需要知道具体对象的类。
*/

// Product 是产品接口
type Product interface {
	GetName() string
}

// ConcreteProductA 是具体产品A
type ConcreteProductA struct{}

func (p *ConcreteProductA) GetName() string {
	return "Product A"
}

// ConcreteProductB 是具体产品B
type ConcreteProductB struct{}

func (p *ConcreteProductB) GetName() string {
	return "Product B"
}

// Factory 是工厂接口，声明工厂方法 CreateProduct
type Factory interface {
	CreateProduct() Product
}

// ConcreteFactoryA 是具体工厂A，用于创建产品A
type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteFactoryB 是具体工厂B，用于创建产品B
type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	// 使用具体工厂A创建产品A
	factoryA := &ConcreteFactoryA{}
	productA := factoryA.CreateProduct()
	fmt.Println("Product A name:", productA.GetName())

	// 使用具体工厂B创建产品B
	factoryB := &ConcreteFactoryB{}
	productB := factoryB.CreateProduct()
	fmt.Println("Product B name:", productB.GetName())
}
