package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Стратегия (Strategy) — поведенческий шаблон проектирования, предназначенный для определения семейства алгоритмов,
// инкапсуляции каждого из них и обеспечения их взаимозаменяемости. Это позволяет выбирать алгоритм путём определения
// соответствующего класса. Шаблон Strategy позволяет менять выбранный алгоритм независимо от объектов-клиентов,
// которые его используют.

// Паттерн стратегия применяется:
// Когда есть несколько родственных классов, которые отличаются поведением. Можно задать один основной класс,
// а разные варианты поведения вынести в отдельные классы и при необходимости их применять.
// Когда необходимо обеспечить выбор из нескольких вариантов алгоритмов, которые можно легко менять
// в зависимости от условий.
// Когда необходимо менять поведение объектов на стадии выполнения программы.
// Когда класс, применяющий определенную функциональность, ничего не должен знать о ее реализации.

// Преимущества паттерна стратегия:
// Изолирует код и данные алгоритмов от остальных классов.
// Уход от наследования к делегированию.
// Горячая замена алгоритмов на лету.

// Недостатки паттерна стратегия:
// Усложняет программу за счёт дополнительных классов.
// Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

type Distance float32

type Point struct {
	X int
	Y int
}

type Geometry interface {
	GetDistance(Point, Point) Distance
}

type SpaceWithGeometry struct {
	Space
	geometry Geometry
}

func (space *SpaceWithGeometry) GetDistance(p1, p2 Point) Distance {
	return space.geometry.GetDistance(p1, p2)
}

type EuclideanGeometry struct{}

func (eg EuclideanGeometry) GetDistance(p1, p2 Point) Distance {
	fmt.Println("Euclidean metrics")
	return 0
}

type SphericalGeometry struct{}

func (sg SphericalGeometry) GetDistance(p1, p2 Point) Distance {
	fmt.Println("Spherical metrics")
	return 0
}

type HyperbolicGeometry struct{}

func (hg HyperbolicGeometry) GetDistance(p1, p2 Point) Distance {
	fmt.Println("Spherical metrics")
	return 0
}
