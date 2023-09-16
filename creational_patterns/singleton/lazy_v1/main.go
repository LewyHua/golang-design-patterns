package main

import (
	"design_patterns/creational_patterns/singleton/lazy_v1/instance"
)

func main() {
	//instanceA := instance.NewInstance()
	//instanceB := instance.NewInstance()
	//fmt.Println(instanceA == instanceB)

	// 会有线程安全问题
	instance.TestConcurrentGetInstance()
}
