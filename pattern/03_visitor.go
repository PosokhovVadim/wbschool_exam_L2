package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Shape interface {
	Accept(visitor ShapeVisitor)
}

type ShapeVisitor interface {
	VisitRectangle(rectangle *Rectangle)
	VisitCircle(circle *Circle)
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r *Rectangle) Accept(visitor ShapeVisitor) {
	visitor.VisitRectangle(r)
}

func (c *Circle) Accept(visitor ShapeVisitor) {
	visitor.VisitCircle(c)
}

// пример вычисления (areaCalculator здесь есть визитор)
type AreaCalculator struct {
	area float64
}

func (a *AreaCalculator) VisitRectangle(rectangle *Rectangle) {
	a.area = float64(rectangle.width * rectangle.height)
}

func (a *AreaCalculator) VisitCircle(circle *Circle) {
	a.area = float64(3.14 * circle.radius * circle.radius)
}

/*
func main() {
	rectangle := &Rectangle{width: 10, height: 5}
	circle := &Circle{radius: 5}
	areaCalculator := &AreaCalculator{}
	rectangle.Accept(areaCalculator)
	circle.Accept(areaCalculator)

}
*/

/*
Данный паттерн позволяет определить новые операции над объектом, не изменяя внутренность объекта.

Использование:
Реальный кейс - использование паттерна для различных этапов компиляции над синтаксическим деревом.

В целом используется когда для большого числа объектов нужно сделать похожие но различные методы.

*/
