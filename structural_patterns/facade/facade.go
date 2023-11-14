package main

import "fmt"

/*
外观模式（Facade Pattern）是一种结构型设计模式，提供了一个统一的接口，用于访问一组相关的接口，简化了复杂系统的操作。
外观模式通过定义一个高层接口，将子系统的一组接口整合为一个更高级的接口，使客户端更容易使用。
*/

// Amplifier 子系统中的一个组件
type Amplifier struct{}

func (a *Amplifier) On() {
	fmt.Println("Amplifier is ON")
}

func (a *Amplifier) Off() {
	fmt.Println("Amplifier is OFF")
}

func (a *Amplifier) SetVolume(volume int) {
	fmt.Printf("Setting Amplifier volume to %d\n", volume)
}

// DVDPlayer 子系统中的一个组件
type DVDPlayer struct{}

func (d *DVDPlayer) On() {
	fmt.Println("DVD Player is ON")
}

func (d *DVDPlayer) Off() {
	fmt.Println("DVD Player is OFF")
}

func (d *DVDPlayer) Play(movie string) {
	fmt.Printf("Playing movie: %s\n", movie)
}

// HomeTheaterFacade 外观类，整合了子系统的一组接口
type HomeTheaterFacade struct {
	amplifier *Amplifier
	dvdPlayer *DVDPlayer
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		amplifier: &Amplifier{},
		dvdPlayer: &DVDPlayer{},
	}
}

// WatchMovie 提供一个简化的接口，使客户端更容易使用
func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	h.amplifier.On()
	h.amplifier.SetVolume(10)
	h.dvdPlayer.On()
	h.dvdPlayer.Play(movie)
}

// EndMovie 提供一个简化的接口，使客户端更容易使用
func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting down the home theater...")
	h.amplifier.Off()
	h.dvdPlayer.Off()
}

func main() {
	// 使用外观模式
	homeTheater := NewHomeTheaterFacade()
	homeTheater.WatchMovie("Inception")
	// Output:
	// Get ready to watch a movie...
	// Amplifier is ON
	// Setting Amplifier volume to 10
	// DVD Player is ON
	// Playing movie: Inception

	homeTheater.EndMovie()
	// Output:
	// Shutting down the home theater...
	// Amplifier is OFF
	// DVD Player is OFF
}
