package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type Request struct {
	Level   int
	Content string
}

type Handler interface {
	SetNext(handler Handler)
	Handle(request Request)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) Handle(request Request) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

// concrete handlers

type UserHandler struct {
	BaseHandler
}

func (h *UserHandler) Handle(request Request) {
	if request.Level == 1 {
		fmt.Println("User: " + request.Content)
	} else {
		h.BaseHandler.Handle(request)
	}
}

type AdminHandler struct {
	BaseHandler
}

func (h *AdminHandler) Handle(request Request) {
	if request.Level == 2 {
		fmt.Println("AdminHandler: Handling request -", request.Content)
	} else {

		h.BaseHandler.Handle(request)
	}
}

/*
func main() {
	userHandler := &UserHandler{}
	adminHandler := &AdminHandler{}

	userHandler.SetNext(adminHandler)
	userHandler.Handle(Request{Level: 1, Content: "User"})
}
*/

/*
Объяснения:

Паттерн предоставляет возможность обработать один запросам нескольким объектам. Эти объекты связываются
в цепочку, и запрос проходит по всей цепи пока не будет обработан. Важная особенность что настояий обработчик
находится по ходу прохождения по цепи (те не известен заранее ).

*/
