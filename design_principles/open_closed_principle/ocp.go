package main

import "fmt"

/*
开闭原则的核心思想是：软件实体（例如类、模块、函数等）应该对扩展开放，对修改关闭。
*/

type Payment interface {
	Pay(amount float64)
}

// CreditCardPayment 实现了 Payment 接口
type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("使用信用卡支付了 %.2f 元\n", amount)
}

// AlipayPayment 实现了 Payment 接口
type AlipayPayment struct{}

func (a *AlipayPayment) Pay(amount float64) {
	fmt.Printf("使用支付宝支付了 %.2f 元\n", amount)
}

/*
在上面的代码中，我们定义了一个 Payment 接口，它有一个 Pay 方法，代表支付操作。
然后，我们创建了两个结构体 CreditCardPayment 和 AlipayPayment，它们都实现了 Payment 接口。
接下来，我们创建一个函数来进行支付，这个函数接受一个支付方式的参数，因为我们不知道用户将使用哪种支付方式。
这里就体现了开闭原则的思想，我们只需要传递不同的支付方式而不需要修改支付函数。
*/

func ProcessPayment(paymentMethod Payment, amount float64) {
	paymentMethod.Pay(amount)
}

func main() {
	creditCardPayment := &CreditCardPayment{}
	alipayPayment := &AlipayPayment{}

	amount := 100.0

	ProcessPayment(creditCardPayment, amount)
	ProcessPayment(alipayPayment, amount)
}
