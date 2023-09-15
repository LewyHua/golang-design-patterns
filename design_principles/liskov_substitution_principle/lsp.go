package main

import "fmt"

/*
里氏替换原则（Liskov Substitution Principle，简称LSP）是面向对象编程中的一个重要原则，
提出了子类（派生类）应该能够替换其基类（父类）而不引起意外行为的要求。具体来说，LSP强调：
子类对象必须能够替换掉父类对象，并且程序的行为不应该受到影响。
*/

/*
在上面的示例中，我们有一个基类 Bird，它有一个 Fly 方法表示鸟可以飞行。
然后，我们有两个子类 Ostrich 和 Duck，它们继承了 Bird 类，并且分别在 Fly 方法中覆盖了基类的行为。
最后，我们创建了基类和子类的对象，演示了子类对象能够替换父类对象并且行为正常。
*/

// Bird 是鸟类的基类
type Bird struct{}

// Fly 方法表示鸟可以飞行
func (b *Bird) Fly() {
	fmt.Println("鸟可以飞行")
}

// Ostrich 是鸵鸟的子类
type Ostrich struct {
	Bird
}

// Fly 方法在鸵鸟中覆盖了基类的行为
func (o *Ostrich) Fly() {
	fmt.Println("鸵鸟不能飞行")
}

// Duck 是鸭子的子类
type Duck struct {
	Bird
}

func main() {
	ostrich := &Ostrich{}
	duck := &Duck{}

	// 使用基类 Bird 来表示鸟，可以替换成子类对象
	bird1 := &Bird{}
	bird2 := ostrich
	bird3 := duck

	// 调用 Fly 方法，行为应该符合预期
	bird1.Fly() // 输出："鸟可以飞行"
	bird2.Fly() // 输出："鸵鸟不能飞行"
	bird3.Fly() // 输出："鸟可以飞行"
}
