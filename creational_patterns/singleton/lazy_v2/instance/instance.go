package instance

import (
	"fmt"
	"sync"
)

/*
通过Mutex实现的简单版并发安全懒汉式
*/

var lock sync.Mutex

type singleton struct {
	name string
}

func (s *singleton) GetName() string {
	return s.name
}

var instance *singleton

func NewInstance() *singleton {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func TestConcurrentGetInstance() {
	const testTimes = 50
	var wg sync.WaitGroup
	instances := make([]*singleton, testTimes)
	for i := 0; i < testTimes; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			instances[index] = NewInstance()
		}(i)
	}
	wg.Wait()

	baseInstance := instances[0]
	for i := 1; i < len(instances); i++ {
		if instances[i] != baseInstance {
			fmt.Printf("not the same\n")
		}
	}

	fmt.Println("ok")

}
