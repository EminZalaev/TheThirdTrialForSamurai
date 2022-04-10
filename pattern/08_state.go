package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

//Приминяемость:
//Когда у вас есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния, причём
//типов состояний много, и их код часто меняется.
//Первоначальный объект будет постоянно ссылаться на один из объектов-состояний, делегируя ему часть своей работы.
//Для изменения состояния в контекст достаточно будет подставить другой объект-состояние.
//Паттерн Состояние позволяет реализовать иерархическую машину состояний, базирующуюся на наследовании.

//Плюсы
//Избавляет от множества больших условных операторов машины состояний.
//Концентрирует в одном месте код, связанный с определённым состоянием.
//Упрощает код контекста.
//Недостатки
//Может неоправданно усложнить код, если состояний мало и они редко меняются
//

type Context struct {
	FirstState  State
	SecondState State

	CurrentState State
}

func NewContext() *Context {
	c := &Context{}

	fs := &FirstState{
		Context: c,
	}

	ss := &SecondState{
		Context: c,
	}

	c.SetState(fs)
	c.FirstState = fs
	c.SecondState = ss
	return c
}

func (c *Context) FirstOperation() error {
	return c.CurrentState.FirstOperation()
}

func (c *Context) SecondOperation() error {
	return c.CurrentState.SecondOperation()
}

func (c *Context) ThirdOperation() error {
	return c.CurrentState.ThirdOperation()
}

func (c *Context) SetState(s State) {
	c.CurrentState = s
}

type State interface {
	FirstOperation() error
	SecondOperation() error
	ThirdOperation() error
}

type FirstState struct {
	Context *Context
}

func (firstS *FirstState) FirstOperation() error {
	fmt.Println("acces: First command ")
	return nil
}

func (firstS *FirstState) SecondOperation() error {
	fmt.Println("block: Second command ")

	return nil
}

func (firstS *FirstState) ThirdOperation() error {
	fmt.Println("State 1 -> State 2")
	firstS.Context.SetState(firstS.Context.SecondState)
	return nil
}

type SecondState struct {
	Context *Context
}

func (secondS *SecondState) FirstOperation() error {
	fmt.Println("block: First command ")
	return nil
}

func (secondS *SecondState) SecondOperation() error {
	fmt.Println("acces: Second command ")

	return nil
}

func (secondS *SecondState) ThirdOperation() error {

	return nil
}
