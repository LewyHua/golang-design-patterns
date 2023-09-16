package instance

/*
单例模式（Singleton Pattern）是一种创建型设计模式，用于确保一个类只有一个实例，并提供全局访问点以获取该实例。
这意味着在整个应用程序生命周期内，只会创建一个特定类的实例，无论多少次尝试获取该实例。
单例模式通常用于以下情况：
1. 当需要控制某个类的实例数量，确保只有一个实例存在时，可以使用单例模式。
2. 当需要在整个应用程序中共享某个资源，例如数据库连接、日志记录器、配置对象等时，单例模式可以确保资源的唯一性。
*/

type singleton struct {
	name string
}

func (s *singleton) GetName() string {
	return s.name
}

var instance *singleton = new(singleton)

func NewInstance() *singleton {
	return instance
}
