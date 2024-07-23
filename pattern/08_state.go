package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type TCPState interface {
	Open(*TCPConnection) string
	Close(*TCPConnection) string
	Acknowledge(*TCPConnection) string
}

// конкретные состояния
type TCPEstablished struct{}

func (t *TCPEstablished) Open(conn *TCPConnection) string {
	return "Connection is already open"
}

func (t *TCPEstablished) Close(conn *TCPConnection) string {
	conn.SetState(&TCPClosed{})
	return "Closing connection"
}

func (t *TCPEstablished) Acknowledge(conn *TCPConnection) string {
	return "Connection is acknowledged"
}

type TCPClosed struct{}

func (t *TCPClosed) Open(conn *TCPConnection) string {
	conn.SetState(&TCPEstablished{})
	return "Opening connection"
}

func (t *TCPClosed) Close(conn *TCPConnection) string {
	return "Connection is already closed"
}

func (t *TCPClosed) Acknowledge(conn *TCPConnection) string {
	return "Cannot acknowledge. Connection is closed"
}

type TCPConnection struct {
	state TCPState
}

func NewTCPConnection(state TCPState) *TCPConnection {
	return &TCPConnection{state: state}
}

func (t *TCPConnection) SetState(state TCPState) {
	t.state = state
}

func (t *TCPConnection) Open() string {
	return t.state.Open(t)
}

func (t *TCPConnection) Close() string {
	return t.state.Close(t)
}

func (t *TCPConnection) Acknowledge() string {
	return t.state.Acknowledge(t)
}

/*
func main() {
	connection := &TCPConnection{}

	connection.Open()
	connection.Acknowledge()
	connection.Close()
	connection.Acknowledge()
}
*/

/*
Объяснение:

Паттерн состояние позволяет изменять поведение объекта в зависимости от его текущего состояния.

Плюсы:
Позволяет создать систему с раздельными ярко выраженными состояниями. Все поведение, относящееся к
состоянию, выделено в отдельный объект.

Минусы:
Не подходит если объекты разных состояний должны взаимодействовать друг с другом.
*/
