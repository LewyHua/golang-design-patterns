package main

import (
	"fmt"
	"sync"
)

/*
观察者模式（Observer Pattern）是一种行为设计模式，它定义了一种一对多的依赖关系，让多个观察者对象同时监听某一个主题对象，当主题对象状态发生变化时，所有依赖它的观察者都会得到通知并自动更新。

观察者模式有三个主要角色：

Subject（主题）： 维护一份观察者对象的列表，提供方法来添加和删除观察者，以及通知观察者的方法。
Observer（观察者）： 定义一个更新接口，使得在收到主题通知时能够执行相应的操作。
ConcreteSubject（具体主题）： 继承自Subject，实现具体的业务逻辑，当状态发生变化时通知所有观察者。
ConcreteObserver（具体观察者）： 继承自Observer，实现更新接口，定义具体的业务逻辑，以便在状态发生变化时做出相应的动作。
*/

// Observer 定义了观察者接口
type Observer interface {
	Update(temperature float64)
}

// Subject 定义了主题接口
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

// ConcreteSubject 实现了主题接口
type ConcreteSubject struct {
	observers    []Observer
	temperature  float64
	observerLock sync.Mutex
}

func (s *ConcreteSubject) RegisterObserver(observer Observer) {
	s.observerLock.Lock()
	defer s.observerLock.Unlock()
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	s.observerLock.Lock()
	defer s.observerLock.Unlock()
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyObservers() {
	s.observerLock.Lock()
	defer s.observerLock.Unlock()
	for _, observer := range s.observers {
		observer.Update(s.temperature)
	}
}

func (s *ConcreteSubject) SetTemperature(temperature float64) {
	s.temperature = temperature
	s.NotifyObservers()
}

// ConcreteObserver 实现了观察者接口
type ConcreteObserver struct {
	name string
}

func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{name: name}
}

func (o *ConcreteObserver) Update(temperature float64) {
	fmt.Printf("Observer %s received update: Temperature is %.2f\n", o.name, temperature)
}

func main() {
	subject := &ConcreteSubject{}
	observer1 := NewConcreteObserver("Device 1")
	observer2 := NewConcreteObserver("Device 2")

	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)

	subject.SetTemperature(25.5)
	// Output:
	// Observer Device 1 received update: Temperature is 25.50
	// Observer Device 2 received update: Temperature is 25.50

	subject.RemoveObserver(observer1)

	subject.SetTemperature(30.0)
	// Output:
	// Observer Device 2 received update: Temperature is 30.00
}
