Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	fmt.Printf("Type: %s\n", reflect.TypeOf(err))
    fmt.Printf("Value: %v\n", reflect.ValueOf(err))

	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод: 
error

Произойдет несоотвествие типов (err будет иметь тип customError)

```
