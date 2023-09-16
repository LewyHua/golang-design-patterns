package main

import "fmt"

/*
抽象工厂模式用的可能不会太多，在不同产品，拥有不同产品族和产品等级，才会发挥作用。
每个工厂生产不同产品族的一系列产品。
例如：生产apple的ipad，iphone，mac，这三个产品为一个产品族
有中国工厂，日本工厂，美国工厂
每个工厂都要生产这三个产品，但是每个工厂的实现不一样。所以每次需要根据抽象工厂实现一个工厂。
*/

// AbstractFactory 是抽象工厂接口，声明一组工厂方法用于创建不同类型的产品
type AbstractFactory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

// ProductA 是产品A接口
type ProductA interface {
	GetName() string
}

// ProductB 是产品B接口
type ProductB interface {
	GetName() string
}

// ConcreteFactory1 是具体工厂1，用于创建产品A1和产品B1
type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProductA() ProductA {
	return &ConcreteProductA1{}
}

func (f *ConcreteFactory1) CreateProductB() ProductB {
	return &ConcreteProductB1{}
}

// ConcreteFactory2 是具体工厂2，用于创建产品A2和产品B2
type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateProductA() ProductA {
	return &ConcreteProductA2{}
}

func (f *ConcreteFactory2) CreateProductB() ProductB {
	return &ConcreteProductB2{}
}

// ConcreteProductA1 是具体产品A1
type ConcreteProductA1 struct{}

func (p *ConcreteProductA1) GetName() string {
	return "Product A1"
}

// ConcreteProductA2 是具体产品A2
type ConcreteProductA2 struct{}

func (p *ConcreteProductA2) GetName() string {
	return "Product A2"
}

// ConcreteProductB1 是具体产品B1
type ConcreteProductB1 struct{}

func (p *ConcreteProductB1) GetName() string {
	return "Product B1"
}

// ConcreteProductB2 是具体产品B2
type ConcreteProductB2 struct{}

func (p *ConcreteProductB2) GetName() string {
	return "Product B2"
}

func main() {
	// 使用具体工厂1创建产品A1和产品B1
	factory1 := &ConcreteFactory1{}
	productA1 := factory1.CreateProductA()
	productB1 := factory1.CreateProductB()

	fmt.Println("Product A1 name:", productA1.GetName())
	fmt.Println("Product B1 name:", productB1.GetName())

	// 使用具体工厂2创建产品A2和产品B2
	factory2 := &ConcreteFactory2{}
	productA2 := factory2.CreateProductA()
	productB2 := factory2.CreateProductB()

	fmt.Println("Product A2 name:", productA2.GetName())
	fmt.Println("Product B2 name:", productB2.GetName())
}
