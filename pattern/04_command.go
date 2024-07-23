package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Command interface {
	Execute()
}

// receiver
type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
}

func (l *Light) TurnOff() {
	l.isOn = false
}

type LightOnCommand struct {
	light *Light
}

func (l *LightOnCommand) Execute() {
	l.light.TurnOn()
}

type LightOffCommand struct {
	light *Light
}

func (l *LightOffCommand) Execute() {
	l.light.TurnOff()
}

//invoker

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

/*
func main() {
	light := &Light{}

	lightOn := &LightOnCommand{light: light}
	lightOff := &LightOffCommand{light: light}

	remote := &RemoteControl{}

	remote.SetCommand(lightOn)
	remote.PressButton()

	remote.SetCommand(lightOff)
	remote.PressButton()

}
*/

/*
Объяснение:

Паттерн  позволяет делегировать выполнение команды другому объекту, не изменяя его код. Command отделяет объект,
иниирующий операцию от объекта, выполняющего ее. Команда здесь - объект.

Использование:
1) Постановка в очередь запросов для выполнения в разное время.
2) Поддержка отмены операции. (позволяет моделировать транзакции) 
3) можно вести протоколирование запросов.

*/