package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

// Паттерн "фасад" является структурным паттерном, позволяющим упрощать доступ к
// сложным системам за счёт простого интерфейса.
//
// Плюсы:
// - Упрощение взаимодействия со сложными системами
//
// Минусы:
// - При неправильной реализации может перестать играть свою роль упрощения и абстрагирования
//
//

type ComplexSystemFacade interface {
	Run()
}

type CSystemOne interface {
	MakeSomeWork()
	MakeAnotherWork()
	MakeYetAnotherWork()
}

type CSystemTwo interface {
	MakeSomeWork()
	MakeAnotherWork()
	MakeYetAnotherWork()
}

type CSystemThree interface {
	MakeSomeWork()
	MakeAnotherWork()
	MakeYetAnotherWork()
}

type SomeFacade struct {
	sOne   CSystemOne
	sTwo   CSystemTwo
	sThree CSystemThree
}

func (sf SomeFacade) Run() {
	sf.sOne.MakeSomeWork()
	sf.sOne.MakeAnotherWork()
	sf.sOne.MakeYetAnotherWork()

	sf.sTwo.MakeSomeWork()
	sf.sTwo.MakeAnotherWork()
	sf.sTwo.MakeYetAnotherWork()

	sf.sThree.MakeSomeWork()
	sf.sThree.MakeAnotherWork()
	sf.sThree.MakeYetAnotherWork()

}
