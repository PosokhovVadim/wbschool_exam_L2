package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// builder объект
type TextConverter interface {
	ConvertCharacter(char byte)
	ConvertParagraph()
}

type ASCIIConverter struct {
	text string
}

func (a *ASCIIConverter) ConvertCharacter(char byte) {
	a.text += string(char)
}

func (a *ASCIIConverter) ConvertParagraph() {
	a.text += "\n"
}

func (a *ASCIIConverter) GetText() string {
	return a.text
}

type HTMLConverter struct {
	html string
}

func (h *HTMLConverter) ConvertCharacter(char byte) {
	h.html += string(char)
}

func (h *HTMLConverter) ConvertParagraph() {
	h.html += "<p></p>"
}

func (h *HTMLConverter) GetHTML() string {
	return h.html
}

// директор
type RTFReader struct {
	builder TextConverter
}

func NewRTFReader(builder TextConverter) *RTFReader {
	return &RTFReader{builder: builder}
}

func (r *RTFReader) Parse(text string) {

	for _, char := range text {
		r.builder.ConvertCharacter(byte(char))
	}
	r.builder.ConvertParagraph()
}

// Пример использования

/*
func main() {
	text := "some text"

	ASCIIConverter := &ASCIIConverter{}
	RTFReader := NewRTFReader(ASCIIConverter)
	RTFReader.Parse(text)

	HTMLConverter := &HTMLConverter{}
	RTFReader = NewRTFReader(HTMLConverter)
	RTFReader.Parse(text)
}
*/

/*
Объяснение:

Паттерн строитель позволяет отделить сложный объект с несколькими независимыми составляющими на мелкие части и определить единую точку использования.
Происходит это таким образом, что при запуске одного и того же процесса (в примере RTFReader.Parse(text)) можно получить разные представления объекта.

К плюсам: 
1) позволяет легко масштабироваться.
2) изолирует код, реализующий конструированние и масшатибрование.
*/
