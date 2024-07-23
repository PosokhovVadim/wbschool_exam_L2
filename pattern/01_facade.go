package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type Document struct {
	Title   string
	Content string
}

type DocumentCreator struct{}

func (dc *DocumentCreator) Create(title, content string) *Document {
	fmt.Println("Document created")
	return &Document{Title: title, Content: content}
}

type DBSaver struct{}

func (ds *DBSaver) Save(doc *Document) {
	fmt.Printf("Document '%s' saved to the database\n", doc.Title)
}

type EmailSender struct{}

func (es *EmailSender) Send(doc *Document, email string) {
	fmt.Printf("Document '%s' sent to %s via email\n", doc.Title, email)
}

type DocumentFacade struct {
	docCreator  *DocumentCreator
	dbSaver     *DBSaver
	emailSender *EmailSender
}

func NewDocumentFacade() *DocumentFacade {
	return &DocumentFacade{
		docCreator:  &DocumentCreator{},
		dbSaver:     &DBSaver{},
		emailSender: &EmailSender{},
	}
}

func (df *DocumentFacade) CreateAndSend(title, content, email string) {
	doc := df.docCreator.Create(title, content)
	df.dbSaver.Save(doc)
	df.emailSender.Send(doc, email)
}

// Пример использования

// func main() {
// 	facade := NewDocumentFacade()
// 	facade.CreateAndSend("test-title", "test-content", "test@mail.com")
// }

/*
Объяснение:

Паттерн фасад предоставляет объект, который обеспечивает доступ к сложным подсистемам, имеющим набор
зависимостей и различных компонентов. Паттерн позволяет, используя простой интерфейс, инкапсулировать
сложные процессы в один простой. Это создает удобство для клиентов ипользуеющие данный интерфейс.

К минусам можно отнести:
1) дополнительный слой абстракции
2) интрефейс может покрыть не все нужные задачи, тогда в таком случае все равно придется обращаться к внутренним компонентам подсистем.

Использование:
Можно скрыть сложность некоторого встраемого модуля/бибилотеки используя фасад.
Используется для управления слоями приложения.


*/
