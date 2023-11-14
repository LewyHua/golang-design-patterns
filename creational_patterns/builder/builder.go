package main

import "fmt"

/*
建造者模式的主要思想是将一个对象的构建过程与其表示分离，使得同样的构建过程可以创建不同的表示。
*/

// Computer 表示要构建的复杂对象
type Computer struct {
	CPU     string
	RAM     string
	Storage string
	GPU     string
}

// ComputerBuilder 定义构建Computer对象的接口
type ComputerBuilder interface {
	SetCPU(cpu string) ComputerBuilder
	SetRAM(ram string) ComputerBuilder
	SetStorage(storage string) ComputerBuilder
	SetGPU(gpu string) ComputerBuilder
	Build() Computer
}

// ConcreteComputerBuilder 实现ComputerBuilder接口的具体建造者
type ConcreteComputerBuilder struct {
	computer Computer
}

func NewComputerBuilder() *ConcreteComputerBuilder {
	return &ConcreteComputerBuilder{computer: Computer{}}
}

func (b *ConcreteComputerBuilder) SetCPU(cpu string) ComputerBuilder {
	b.computer.CPU = cpu
	return b
}

func (b *ConcreteComputerBuilder) SetRAM(ram string) ComputerBuilder {
	b.computer.RAM = ram
	return b
}

func (b *ConcreteComputerBuilder) SetStorage(storage string) ComputerBuilder {
	b.computer.Storage = storage
	return b
}

func (b *ConcreteComputerBuilder) SetGPU(gpu string) ComputerBuilder {
	b.computer.GPU = gpu
	return b
}

func (b *ConcreteComputerBuilder) Build() Computer {
	return b.computer
}

// Director 将构建过程封装在Director中
type Director struct {
	builder ComputerBuilder
}

func NewDirector(builder ComputerBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) ConstructHighEndComputer() Computer {
	return d.builder.
		SetCPU("Intel i9").
		SetRAM("32GB").
		SetStorage("1TB SSD").
		SetGPU("NVIDIA RTX 3090").
		Build()
}

func (d *Director) ConstructLowEndComputer() Computer {
	return d.builder.
		SetCPU("AMD Ryzen 5").
		SetRAM("8GB").
		SetStorage("500GB HDD").
		SetGPU("Integrated Graphics").
		Build()
}

func main() {
	builder := NewComputerBuilder()
	director := NewDirector(builder)

	highEndComputer := director.ConstructHighEndComputer()
	lowEndComputer := director.ConstructLowEndComputer()

	fmt.Println("High-end Computer:", highEndComputer)
	fmt.Println("Low-end Computer:", lowEndComputer)
}
