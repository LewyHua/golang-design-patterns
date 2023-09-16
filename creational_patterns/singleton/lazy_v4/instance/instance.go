package instance

import (
	"fmt"
	"sync"
)

/*
通过atomic + DCL实现的简单版并发安全懒汉式
*/

var once sync.Once

type singleton struct {
	name string
}

func (s *singleton) GetName() string {
	return s.name
}

var instance *singleton

func NewInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})

	return instance
}

func TestConcurrentGetInstance() {
	const numRoutines = 50

	var wg sync.WaitGroup
	instances := make([]*singleton, numRoutines)

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			instances[index] = NewInstance()
		}(i)
	}

	wg.Wait()

	// 检查所有实例是否相同
	baseInstance := instances[0]
	for i := 1; i < numRoutines; i++ {
		if instances[i] != baseInstance {
			fmt.Printf("Instances at index %d and index 0 are not the same\n", i)
		}
	}
}
