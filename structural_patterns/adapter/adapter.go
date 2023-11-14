package main

import "fmt"

/*
适配器模式（Adapter Pattern）是一种结构型设计模式，它允许将一个类的接口转换成客户端希望的另一个接口。适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。

在适配器模式中，有三个主要角色：
目标接口（Target Interface）： 客户端期望的接口，适配器通过实现这个接口来与客户端进行交互。
适配者（Adaptee）： 需要被适配的类或接口。
适配器（Adapter）： 实现了目标接口并持有适配者的引用，在目标接口的方法中调用适配者的方法。
*/

// EuropeanSocket 目标接口
type EuropeanSocket interface {
	PlugIn()
}

// BritishPlug 适配者
type BritishPlug struct{}

func (bp *BritishPlug) Plug() {
	fmt.Println("Plugging in British plug")
}

// BritishToEuropeanAdapter 适配器
type BritishToEuropeanAdapter struct {
	BritishPlug *BritishPlug
}

func (adapter *BritishToEuropeanAdapter) PlugIn() {
	adapter.BritishPlug.Plug()
	fmt.Println("Adapter: Converting British plug to fit European socket")
}

func main() {
	// 使用欧洲插座
	socket := &BritishToEuropeanAdapter{BritishPlug: &BritishPlug{}}
	socket.PlugIn()
}
