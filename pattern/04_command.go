package pattern

import "time"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Команда (Command) — поведенческий паттерн проектирования, который превращает запросы в объекты,
// позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их,
// а также поддерживать отмену операций.

// Паттерн команда применяется:
// Когда надо передавать в качестве параметров определенные действия, вызываемые в ответ на другие действия.
// Когда необходимо обеспечить выполнение очереди запросов, а также их возможную отмену.

// Преимущества паттерна команда:
// Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
// Позволяет реализовать простую отмену и повтор операций.
// Позволяет реализовать отложенный запуск операций.
// Позволяет собирать сложные команды из простых.

// Недостаток паттерна команда это усложнение кода программы из-за введения множества дополнительных классов.

type SpaceUnit float32

type Command interface {
	execute()
}

type WorldReceiver interface {
	IncreseTime(time.Duration)
	IncreaseSpace(su SpaceUnit)
}

//lint:ignore U1000 want to disable check
type IncTimeCommand struct {
	wr WorldReceiver
	d  time.Duration
}

func (itc IncTimeCommand) execute() {
	itc.wr.IncreseTime(itc.d)
}

//lint:ignore U1000 want to disable check
type IncSpaceCommand struct {
	wr WorldReceiver
	su SpaceUnit
}

func (itc IncSpaceCommand) execute() {
	itc.wr.IncreaseSpace(itc.su)
}
