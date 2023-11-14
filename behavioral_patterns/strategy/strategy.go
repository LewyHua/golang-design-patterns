package main

import "fmt"

/*
策略模式（Strategy Pattern）是一种行为设计模式，它定义了算法家族，分别封装起来，让它们之间可以互相替换，使得算法的变化独立于使用算法的客户。
在策略模式中，定义一组算法，将每个算法都封装起来，并且使它们可以互相替换。客户端代码通过一个抽象接口调用不同的具体算法实现。
*/

// PriceCalculator 定义了价格计算的策略接口
type PriceCalculator interface {
	CalculatePrice(originalPrice float64) float64
}

// NormalPriceCalculator 实现了普通价格计算策略
type NormalPriceCalculator struct{}

func (c *NormalPriceCalculator) CalculatePrice(originalPrice float64) float64 {
	return originalPrice
}

// DiscountPriceCalculator 实现了折扣价格计算策略
type DiscountPriceCalculator struct {
	DiscountRate float64
}

func (c *DiscountPriceCalculator) CalculatePrice(originalPrice float64) float64 {
	return originalPrice * (1 - c.DiscountRate)
}

// Context 定义了上下文，包含一个价格计算策略
type Context struct {
	PriceCalculator PriceCalculator
}

// SetCalculator 设置价格计算策略
func (ctx *Context) SetCalculator(calculator PriceCalculator) {
	ctx.PriceCalculator = calculator
}

// Calculate 使用当前策略计算最终价格
func (ctx *Context) Calculate(originalPrice float64) float64 {
	return ctx.PriceCalculator.CalculatePrice(originalPrice)
}

func main() {
	// 创建上下文
	context := &Context{}

	// 默认使用普通价格计算策略
	context.SetCalculator(&NormalPriceCalculator{})

	originalPrice := 100.0
	finalPrice := context.Calculate(originalPrice)
	fmt.Printf("Original Price: %.2f, Final Price: %.2f\n", originalPrice, finalPrice)
	// Output: Original Price: 100.00, Final Price: 100.00

	// 切换到折扣价格计算策略
	context.SetCalculator(&DiscountPriceCalculator{DiscountRate: 0.2})

	finalPrice = context.Calculate(originalPrice)
	fmt.Printf("Original Price: %.2f, Final Price with Discount: %.2f\n", originalPrice, finalPrice)
	// Output: Original Price: 100.00, Final Price with Discount: 80.00
}
