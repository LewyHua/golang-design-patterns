package main

import "fmt"

/*
简单工厂模式虽然是一种简单的创建型设计模式，但它仍然在某些情况下有用，特别是在项目初期或某些特定的场景下。
然而，它也存在一些限制和缺点，导致在某些现代应用中被认为不够灵活或不推荐使用。
以下是一些关于简单工厂模式的观点和限制：
1. 创造单一产品族：简单工厂适用于创建单一产品族的情况，即产品之间有相似的结构或行为。如果产品族变得复杂或多样化，简单工厂模式可能不够灵活。
2. 不支持开放-封闭原则：每次添加新产品类型时，都需要修改工厂类的代码，这违反了开放-封闭原则。在面对经常变化的产品类型时，这可能会导致维护问题。
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

// Factory 是简单工厂
type Factory struct{}

func (f *Factory) CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return &ConcreteProductA{}
	case "B":
		return &ConcreteProductB{}
	default:
		return nil
	}
}

func main() {
	factory := &Factory{}

	// 创建具体产品A
	productA := factory.CreateProduct("A")
	fmt.Println("Product A name:", productA.GetName())

	// 创建具体产品B
	productB := factory.CreateProduct("B")
	fmt.Println("Product B name:", productB.GetName())
}
