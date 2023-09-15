package main

import "fmt"

/*
合成复用原则（Composite Reuse Principle，简称CRP）是SOLID原则的一部分，
它强调应该优先使用对象组合（composition）而不是继承（inheritance）来实现代码的复用。
合成复用原则的核心思想是：应该优先选择将已有的类组合起来，而不是通过继承现有的类来实现代码的复用。
使用合成复用原则有以下优点：
1. 降低耦合度：使用组合可以减少类之间的耦合，因为对象之间的关系更加松散。
2. 提高灵活性：通过组合，可以在运行时动态改变对象的行为，而无需修改代码。
3. 支持单一职责原则：使用组合可以更容易地实现单一职责原则，因为每个对象只需要关注自己的职责。
*/

type Engine interface {
	Start()
	Stop()
}

type Wing interface {
	Lift()
}

type JetEngine struct{}

func (j *JetEngine) Start() {
	fmt.Println("喷气引擎启动")
}

func (j *JetEngine) Stop() {
	fmt.Println("喷气引擎停止")
}

type PropellerEngine struct{}

func (p *PropellerEngine) Start() {
	fmt.Println("螺旋桨引擎启动")
}

func (p *PropellerEngine) Stop() {
	fmt.Println("螺旋桨引擎停止")
}

type FixedWing struct{}

func (f *FixedWing) Lift() {
	fmt.Println("固定机翼起飞")
}

type RotatingWing struct{}

func (r *RotatingWing) Lift() {
	fmt.Println("旋翼机翼起飞")
}

/*
在上面的代码中，我们定义了两个接口 Engine 和 Wing，并实现了两种不同类型的引擎和机翼。
然后，我们可以通过组合这些组件来构建不同类型的飞行器，而不需要使用继承。
例如，构建一架使用喷气引擎和固定机翼的飞行器：
*/

type Aircraft struct {
	engine Engine
	wing   Wing
}

func NewAircraft(engine Engine, wing Wing) *Aircraft {
	return &Aircraft{
		engine: engine,
		wing:   wing,
	}
}

func (a *Aircraft) TakeOff() {
	a.engine.Start()
	a.wing.Lift()
	fmt.Println("飞行器起飞")
}

func main() {
	jetEngine := &JetEngine{}
	fixedWing := &FixedWing{}

	aircraft1 := NewAircraft(jetEngine, fixedWing)
	aircraft1.TakeOff()

}
