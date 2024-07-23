package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type Transport interface {
	Drive() string
}

// конкретные создаваемые объекты
type Car struct{}

func (c *Car) Drive() string {
	return "Driving a car"
}

type Bike struct{}

func (b *Bike) Drive() string {
	return "Riding a bike"
}

// интерфейс создателя
type TransportFactory interface {
	CreateTransport() Transport
}

// конкретные создатели
type CarFactory struct{}

func (c *CarFactory) CreateTransport() Transport {
	return &Car{}
}

type BikeFactory struct{}

func (b *BikeFactory) CreateTransport() Transport {
	return &Bike{}
}

// func main() {
// 	var factory TransportFactory

// 	factory = &CarFactory{}
// 	transport := factory.CreateTransport()
// 	fmt.Println(transport.Drive())

// 	factory = &BikeFactory{}
// 	transport = factory.CreateTransport()
// 	fmt.Println(transport.Drive())
// }

/*
Объяснение:

Данный паттерн предоставляет интерфейс для создания объекта, но оставляет подклассам решение о том, экземпляры
какого класса должны создаваться.

Плюсы:
1) Создание объектов не привязывается к конкретным классам
2) Позволяет установить связь между параллельными иерархиями классов

Минусы:
Излишнее дублирование кода из-за классов наследников

*/
