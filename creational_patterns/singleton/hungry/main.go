package main

import (
	"design_patterns/creational_patterns/singleton/hungry/instance"
	"fmt"
)

func main() {
	instanceA := instance.NewInstance()
	instanceB := instance.NewInstance()
	fmt.Println(instanceA == instanceB)
}
