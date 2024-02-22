package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Цепочка вызовов (Chain of responsibility) - поведенческий шаблон проектирования, который позволяет избежать
// жесткой привязки отправителя запроса к получателю. Все возможные обработчики запроса образуют цепочку,
// а сам запрос перемещается по этой цепочке. Каждый объект в этой цепочке при получении запроса выбирает,
// либо закончить обработку запроса, либо передать запрос на обработку следующему по цепочке объекту.

// Паттерн цепочка вызовов применяется:
// Когда имеется более одного объекта, который может обработать определенный запрос.
// Когда надо передать запрос на выполнение одному из нескольких объект, точно не определяя, какому именно объекту.
// Когда набор объектов задается динамически.

// Преимущества паттерна цепочка вызовов:
// Ослабление связанности между объектами. Отправителю и получателю запроса ничего не известно друг о друге.
// Клиенту неизветна цепочка объектов, какие именно объекты составляют ее, как запрос в ней передается.
// В цепочку с легкостью можно добавлять новые типы объектов, которые реализуют общий интерфейс.

// Недостатки паттерна цепочка вызовов:
// никто не гарантирует, что запрос в конечном счете будет обработан. Если необходимого обработчика в цепочки
// не оказалось, то запрос просто выходит из цепочки и остается необработанным.

const (
	UNTYPED MessageType = iota
	IMPORTANT
	MEDIUM
	LOW
)

type MessageType int
type MessageBody string

//lint:ignore U1000 because
type Message struct {
	msgType MessageType
	body    MessageBody
}

type MessageHandler interface {
	execute(*Message)
	setNext(MessageHandler)
}

// Procced important message or execute next handler
// Drop execution in case of UNTYPE
type ImpMsgHandle struct {
	next MessageHandler
}

func (imh *ImpMsgHandle) execute(m *Message) {
	switch m.msgType {
	case IMPORTANT:
		fmt.Println("Bingo in Important")
		return
	case UNTYPED:
		fmt.Println("Woow wow, easy!")
		return
	default:
		imh.next.execute(m)
	}
}

func (imh *ImpMsgHandle) setNext(next MessageHandler) {
	imh.next = next
}
