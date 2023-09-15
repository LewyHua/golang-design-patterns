package main

import "fmt"

/*
迪米特法则（Law of Demeter，简称LoD）也被称为最少知识原则（Least Knowledge Principle，简称LKP）或者“不要和陌生人说话”原则。
迪米特法则的核心思想是：
一个对象应该对其他对象保持最少的了解，不要直接与许多其他对象交互，而应该通过中介或者封装来降低对象之间的耦合度。
迪米特法则的目标是降低系统的耦合度，提高模块的独立性，从而使系统更加灵活、可维护和可扩展。
*/

/*
假设我们有一个电商系统，其中包含了购物车、商品和用户等不同的类。
购物车需要知道用户信息以及商品的信息，但不应该直接与商品和用户类进行交互。
我们可以通过封装和中介来实现迪米特法则。
*/

type Product struct {
	Name  string
	Price float64
}

type User struct {
	Name  string
	Email string
}

func NewProduct(name string, price float64) *Product {
	return &Product{Name: name, Price: price}
}

func NewUser(name, email string) *User {
	return &User{Name: name, Email: email}
}

// 接下来，定义购物车类，购物车需要知道用户和商品的信息，但不需要直接与它们交互

type ShoppingCart struct {
	User     *User
	Products []*Product
}

func NewShoppingCart(user *User) *ShoppingCart {
	return &ShoppingCart{User: user, Products: []*Product{}}
}

func (sc *ShoppingCart) AddProduct(product *Product) {
	sc.Products = append(sc.Products, product)
}

func (sc *ShoppingCart) Checkout() {
	fmt.Printf("用户 %s 购买了以下商品：\n", sc.User.Name)
	for _, product := range sc.Products {
		fmt.Printf("%s（价格：%.2f）\n", product.Name, product.Price)
	}
	fmt.Printf("总价：%.2f\n", sc.calculateTotalPrice())
}

func (sc *ShoppingCart) calculateTotalPrice() float64 {
	total := 0.0
	for _, product := range sc.Products {
		total += product.Price
	}
	return total
}

// 现在，我们可以创建用户、商品和购物车，并实现了迪米特法则，购物车只与用户和商品类有直接关系，不需要了解它们的具体细节
func main() {
	user := NewUser("Alice", "alice@example.com")
	product1 := NewProduct("Laptop", 1000.0)
	product2 := NewProduct("Phone", 500.0)

	cart := NewShoppingCart(user)
	cart.AddProduct(product1)
	cart.AddProduct(product2)

	cart.Checkout()
}
