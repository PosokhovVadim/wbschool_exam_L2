package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type Strategy interface {
	Sort(array []int)
}

// конкретные "стратегии"
type BubbleSort struct{}

func (b *BubbleSort) Sort(array []int) {
	// bubble sort
}

type QuickSort struct{}

func (q *QuickSort) Sort(array []int) {
	// quick sort
}

// контекст
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) Sort(array []int) {
	c.strategy.Sort(array)
}

/*
func main() {

	array := []int{1, 2, 3}

	bubbleSort := &BubbleSort{}
	quickSort := &QuickSort{}

	context := &Context{}

	context.SetStrategy(bubbleSort)
	context.Sort(array)

	context.SetStrategy(quickSort)
	context.Sort(array)
}
*/

/*
Объяснение:

Данный паттерн позволяет на основе контекста выбрать нужный алгоритм реализации.
Семейство этих алгоритмов скрыто от клиента в отдельном классе.

Плюсы:
1) гибкость системы
2) расширяемость функционала
3) поддержка принципа OCP

Минусы:
1) дополнительный слой абстракции

Применение:
Выбор нужного лагоритма поиска/cортировки/рендеринга/шифрования и тд.
*/
