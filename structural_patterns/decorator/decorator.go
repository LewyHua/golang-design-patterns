package main

import "fmt"

/*

装饰器模式（Decorator Pattern）是一种结构型设计模式，它允许在不改变原始对象接口的情况下，动态地扩展对象的功能。
这种模式是通过创建一个包装（wrapper）对象来实现的，这个包装对象包含了原始对象，并且具有相同的接口，使得客户端无法区分原始对象和装饰对象。

在装饰器模式中，有几个关键角色：

Component（组件）： 定义了一个抽象接口，可以是抽象类或接口，是具体组件和装饰器的共同接口。
ConcreteComponent（具体组件）： 实现了组件接口的具体类，是我们要最终装饰的对象。
Decorator（装饰器）： 也是组件的具体实现类，但它同时包含了一个指向组件对象的引用，通过该引用可以调用被装饰对象的方法。装饰器通常是抽象类，它的子类可以添加额外的行为。
ConcreteDecorator（具体装饰器）： 是装饰器的具体实现类，负责向组件添加新的行为。
*/

// Component 定义组件接口
type CoffeeComponent interface {
	Cost() float64
}

// ConcreteComponent 具体组件
type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() float64 {
	return 2.0
}

// Decorator 装饰器抽象类
type CoffeeDecorator struct {
	component CoffeeComponent
}

func (d *CoffeeDecorator) Cost() float64 {
	return d.component.Cost()
}

// ConcreteDecorator 具体装饰器
type SugarDecorator struct {
	CoffeeDecorator
}

func (s *SugarDecorator) Cost() float64 {
	return s.CoffeeDecorator.Cost() + 0.5
}

// ConcreteDecorator 具体装饰器
type MilkDecorator struct {
	CoffeeDecorator
}

func (m *MilkDecorator) Cost() float64 {
	return m.CoffeeDecorator.Cost() + 1.0
}

func main() {
	// 创建一个简单咖啡对象
	coffee := &SimpleCoffee{}
	fmt.Printf("Cost of simple coffee: $%.2f\n", coffee.Cost())

	// 使用装饰器给咖啡加糖
	sugarCoffee := &SugarDecorator{CoffeeDecorator: CoffeeDecorator{component: coffee}}
	fmt.Printf("Cost of coffee with sugar: $%.2f\n", sugarCoffee.Cost())

	// 使用装饰器给咖啡加糖和牛奶
	milkSugarCoffee := &MilkDecorator{CoffeeDecorator: CoffeeDecorator{component: sugarCoffee}}
	fmt.Printf("Cost of coffee with sugar and milk: $%.2f\n", milkSugarCoffee.Cost())
}
