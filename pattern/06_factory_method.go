package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Фабричный метод (Factory Method) — это порождающий паттерн проектирования, который решает проблему создания
// различных продуктов, без указания конкретных классов продуктов.

// Паттерн фабричный метод применяется:
// Когда заранее неизвестно, объекты каких типов необходимо создавать.
// Когда система должна быть независимой от процесса создания новых объектов и расширяемой: в нее можно
// легко вводить новые классы, объекты которых система должна создавать.
// Когда создание новых объектов необходимо делегировать из базового класса классам наследникам.

// Преимущества паттерна фабричный метод:
// Избавляет класс от привязки к конкретным классам продуктов.
// Выделяет код производства продуктов в одно место, упрощая поддержку кода.
// Упрощает добавление новых продуктов в программу.

// Недостатки паттерна фабричный метод:
// Для каждого нового продукта необходимо создавать свой класс создателя.

type CelestialObject interface {
	IsCelestialObject() bool
}

type Star struct{}

func (s Star) IsCelestialObject() bool { return true }

type Planet struct{}

func (s Planet) IsCelestialObject() bool { return true }

type Asteroid struct{}

func (s Asteroid) IsCelestialObject() bool { return true }

func NewStar() CelestialObject {
	return Star{}
}

func NewPlanet() CelestialObject {
	return Planet{}
}

func NewAsteroid() CelestialObject {
	return Planet{}
}

func GetCelestialObject(celType string) CelestialObject {
	switch celType {
	case "star":
		return NewStar()
	case "planet":
		return NewPlanet()
	case "asteroid":
		return NewAsteroid()
	default:
		return nil
	}
}
