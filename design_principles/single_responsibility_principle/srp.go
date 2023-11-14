package main

import "fmt"

/*
单一职责原则（Single Responsibility Principle，简称SRP）是面向对象编程中的一项基本原则，它是SOLID原则中的五个原则之一。
SRP的核心思想是一个类或模块应该有且只有一个引起它变化的原因，或者说一个类或模块应该只有一个责任。
这个原则的主要目标是将软件的各个部分分离开来，使其更加灵活、可维护和可扩展。
需要注意的是，单一职责原则并不是说一个类只能有一个方法，而是指一个类的方法和属性应该在逻辑上属于同一个责任领域。
在实际编程中，要根据具体情况来判断如何划分类的职责，以遵循这一原则。
例如Demo中的Article类就只用来负责Article相关的方法，不该有对其无关的类进行修改，不该负责其他相关的事情，例如修改author的一些属性。
*/

type Article struct {
	ID      int
	Title   string
	Author  Author
	Content string
}

type Author struct {
	Name string
}

func NewArticle(id int, title, content string, author Author) *Article {
	return &Article{
		ID:      id,
		Title:   title,
		Author:  author,
		Content: content,
	}
}

// GetTitle Getter方法用于获取文章的标题
func (a *Article) GetTitle() string {
	return a.Title
}

// GetAuthor Getter方法用于获取文章的作者
func (a *Article) GetAuthor() Author {
	return a.Author
}

// ShowArticle ShowArticle用于展示文章的内容
func ShowArticle(article *Article) {
	fmt.Println("Title:", article.GetTitle())
	fmt.Println("Author:", article.GetAuthor())
	fmt.Println("Content:")
	fmt.Println(article.Content)
}

func main() {

}
