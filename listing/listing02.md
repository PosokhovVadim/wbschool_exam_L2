Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
defer выполняет отложенные вычисления в обратном порядке относительно их вызовов. 
Вывод программы:
2
1

Функция test: переменная x является именнованным возвращаемым значением. После последней строки функции срабатывает defer который модифицирует возвращаемую переменную. Поэтому вывод 2
Функциия anotherTest: на момент возврата x зафиксированно как 1, и дальнейший defer поменяет лишь локальное значение x, т.е. внутри defer x будет равен 2, но возврат функции - 1.
```
