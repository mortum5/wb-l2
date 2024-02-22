package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
	Паттерн позволяет менять объектам свое поведение, исходя из состояния
	в котором они находятся. Возможные состояния объекта выделяются
	в отдельные объекты-состояния и базовому объекту потребуется ссылаться
	только на текущее состояния и иметь возможность его изменять на другое.
	Так можно удобно добавлять новые состояния и изменять старые не
	внося изменения в код базового объекта. Большой условный оператор
	заменяется разными объектами-состояниями и ссылкой на них в базовом объекте.
	Плюс:
	- Уход от больших громоздких условных операторов
	- Код, относящийся к определенному состоянию находится в одном месте
	Минусы:
	- Может неоправдано усложнять код из-за введения дополнительных объектов,
	если состояний небольшое количество
	паттерн Состояние изменяет функциональность одних и тех же элементов управления музыкальным
	проигрывателем, в зависимости от того, в каком состоянии находится сейчас проигрыватель.
*/

type State interface {
	ToAState()
	ToBState()
	ToCState()
	ToDState()
}

type StateMachine struct {
	aState  State
	bState  State
	cState  State
	dState  State
	curStat State
}

func (sm *StateMachine) SetState(st State) {
	sm.curStat = st
}

type AState struct {
	sm *StateMachine
}

type BState struct {
	sm *StateMachine
}

type CState struct {
	sm *StateMachine
}

type DState struct {
	sm *StateMachine
}

func (as *AState) ToAState() {}

func (as *AState) ToBState() {
	as.sm.curStat = as.sm.bState
}
func (as *AState) ToCState() {
	as.sm.curStat = as.sm.cState
}
func (as *AState) ToDState() {}

func (as *BState) ToAState() {}
func (as *BState) ToBState() {}
func (as *BState) ToCState() {}
func (as *BState) ToDState() {
	as.sm.curStat = as.sm.dState

}
