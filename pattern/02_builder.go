package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/
func main() {

}

type Human struct {
	age      int
	height   int
	eyeColor string
}

func newHuman() Human {
	return Human{}
}