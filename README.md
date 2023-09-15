# golang-design-patterns
Golang实现的设计模式Design Patterns

## 设计原则
> 设计原则是指在软件设计中的一组通用准则和指导原则，它们有助于编写高质量、可维护和可扩展的代码。
> 设计模式是实现设计原则的具体方式。

- <a href="https://github.com/LewyHua/golang-design-patterns/blob/main/design_principles/single_responsibility_principle/srp.go">单一职责原则（Single Responsibility Principle - SRP）</a>
- <a href="https://github.com/LewyHua/golang-design-patterns/blob/main/design_principles/open_closed_principle/ocp.go">开放封闭原则（Open-Closed Principle - OCP）</a>
- <a href="https://github.com/LewyHua/golang-design-patterns/blob/main/design_principles/liskov_substitution_principle/lsp.go">里氏替换原则（Liskov Substitution Principle - LSP）</a>
- <a href="https://github.com/LewyHua/golang-design-patterns/blob/main/design_principles/dependency_inversion_principle/dip.go">依赖倒置原则（Dependency Inversion Principle - DIP）</a>
- <a href="https://github.com/LewyHua/golang-design-patterns/blob/main/design_principles/law_of_demeter/lod.go">迪米特法则（Law of Demeter - LoD）</a>
- <a href="https://github.com/LewyHua/golang-design-patterns/blob/main/design_principles/interface_segregation_principle/isp.go">接口隔离原则（Interface Segregation Principle - ISP）</a>
- <a href="https://github.com/LewyHua/golang-design-patterns/blob/main/design_principles/composite_reuse_principle/crp.go">合成复用原则（Composite Reuse Principle - CRP）</a>


## 设计模式
### 创建型模式（Creational Patterns）
> 这些模式关注对象的创建机制，帮助我们有效地创建和管理对象实例。
- 工厂方法模式（Factory Method Pattern）
- 抽象工厂模式（Abstract Factory Pattern）
- 单例模式（Singleton Pattern）
- 建造者模式（Builder Pattern）
- 原型模式（Prototype Pattern）

### 结构型模式（Structural Patterns）
> 这些模式关注对象之间的组合，帮助我们定义如何组合类和对象以构建更大的结构。
- 适配器模式（Adapter Pattern）
- 桥接模式（Bridge Pattern）
- 组合模式（Composite Pattern）
- 装饰器模式（Decorator Pattern） 
- 外观模式（Facade Pattern）
- 享元模式（Flyweight Pattern）
- 代理模式（Proxy Pattern）

### 行为型模式（Behavioral Patterns）：
> 这些模式关注对象之间的交互，帮助我们定义对象之间的责任分配和通信方式。

- 责任链模式（Chain of Responsibility Pattern）
- 命令模式（Command Pattern）
- 解释器模式（Interpreter Pattern）
- 迭代器模式（Iterator Pattern）
- 中介者模式（Mediator Pattern）
- 备忘录模式（Memento Pattern）
- 观察者模式（Observer Pattern）
- 状态模式（State Pattern）
- 策略模式（Strategy Pattern）
- 模板方法模式（Template Method Pattern）
- 访问者模式（Visitor Pattern）