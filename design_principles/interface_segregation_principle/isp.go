package main

import "fmt"

/*
依赖倒置原则（Dependency Inversion Principle，简称DIP）
它强调高层模块不应该依赖于低层模块，它们都应该依赖于抽象。
同时，抽象不应该依赖于具体细节，具体细节应该依赖于抽象。
这一原则有助于减少模块之间的紧耦合，提高了代码的灵活性和可维护性。
*/

/*
在下面的示例中，我们定义了一个 Notification 接口，然后实现了两个具体的通知类 EmailNotification 和 SMSNotification，它们都实现了 Notification 接口。
然后，我们有一个 App 类，它包含一个 Notification 接口的实例。
通过构造函数注入不同的通知方式，我们可以轻松地切换通知方式而不需要修改 App 类的代码，这是依赖倒置原则的体现。
这种方式提高了代码的可扩展性和可维护性，同时减少了高层模块和低层模块之间的直接依赖。
*/

// Notification 定义一个抽象的 Notification 接口
type Notification interface {
	Send(message string)
}

// EmailNotification 实现了 Notification 接口
type EmailNotification struct{}

func (n *EmailNotification) Send(message string) {
	fmt.Printf("通过电子邮件发送消息：%s\n", message)
}

// SMSNotification 实现了 Notification 接口
type SMSNotification struct{}

func (n *SMSNotification) Send(message string) {
	fmt.Printf("通过短信发送消息：%s\n", message)
}

// App 包含一个 Notification 接口的实例
type App struct {
	notifier Notification
}

// NewApp 构造函数接受一个 Notification 接口的实例
func NewApp(notifier Notification) *App {
	return &App{
		notifier: notifier,
	}
}

// SendMessage 方法使用依赖注入的方式发送消息
func (a *App) SendMessage(message string) {
	a.notifier.Send(message)
}

func main() {
	emailNotifier := &EmailNotification{}
	smsNotifier := &SMSNotification{}

	// 创建 App 实例，注入 EmailNotification 依赖
	app1 := NewApp(emailNotifier)
	app1.SendMessage("Hello, World!")

	// 创建 App 实例，注入 SMSNotification 依赖
	app2 := NewApp(smsNotifier)
	app2.SendMessage("Hi there!")

	// 无需修改 App 类的代码，可以轻松切换不同的通知方式
}
