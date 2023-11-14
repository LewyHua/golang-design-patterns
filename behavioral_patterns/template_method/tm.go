package main

import "fmt"

/*
模板方法模式（Template Method Pattern）是一种行为设计模式，定义了一个算法的骨架，将算法中的某些步骤延迟到子类中实现。
这种模式使得子类可以在不改变算法结构的情况下重新定义算法中的某些步骤。
*/

// Beverage 抽象接口定义了算法骨架
type Beverage interface {
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
}

// Template 具体结构体实现算法的具体步骤
type Template struct {
	b Beverage
}

// PrepareRecipe 模板方法，定义了算法骨架
func (t *Template) PrepareRecipe() {
	if t.b == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()
	t.b.AddCondiments()
}

//func (t *Template) Brew() {
//	fmt.Println("Brewing")
//}
//
//func (t *Template) AddCondiments() {
//	fmt.Println("Adding condiments")
//}
//
//func (t *Template) BoilWater() {
//	fmt.Println("Boiling water")
//}
//
//func (t *Template) PourInCup() {
//	fmt.Println("Pouring into cup")
//}

// Coffee 具体结构体，实现了算法中的某些步骤
type Coffee struct {
	t Template
}

func (c *Coffee) BoilWater() {
	//c.BoilWater() // 调用父类的方法
	fmt.Println("Boiling water")
}

func (c *Coffee) PourInCup() {
	//c.PourInCup() // 调用父类的方法
	fmt.Println("Pouring into cup")
}

func (c *Coffee) Brew() {
	//c.Brew() // 调用父类的方法
	fmt.Println("Brewing coffee grounds")
}

func (c *Coffee) AddCondiments() {
	//c.AddCondiments() // 调用父类的方法
	fmt.Println("Adding sugar and milk")
}

// Tea 具体结构体，实现了算法中的某些步骤
type Tea struct {
	t Template
}

func (t *Tea) Brew() {
	//t.Brew() // 调用父类的方法
	fmt.Println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
	//t.AddCondiments() // 调用父类的方法
	fmt.Println("Adding lemon")
}

func (t *Tea) BoilWater() {
	//t.BoilWater() // 调用父类的方法
	fmt.Println("Boiling water")
}

func (t *Tea) PourInCup() {
	//t.PourInCup() // 调用父类的方法
	fmt.Println("Pouring into cup")
}

func main() {
	coffeeTemplate := &Template{b: &Coffee{}}
	coffeeTemplate.PrepareRecipe()

	fmt.Println()

	teaTemplate := &Template{b: &Tea{}}
	teaTemplate.PrepareRecipe()
}
