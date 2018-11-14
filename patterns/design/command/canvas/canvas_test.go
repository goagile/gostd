package canvas

import (
	"fmt"
	"testing"
)

func Test_Move_One(t *testing.T) {
	want := "->Point(1, 1)\n"

	command := Canvas(
		Move(1, 1))

	got := command.Do()

	if want != got {
		t.Fatalf("\nwant:\n%v\n got:\n%v\n", want, got)
	}
}

func Test_Move_Two(t *testing.T) {
	want := "->Point(1, 1)\n" +
		"->Point(2, 2)\n"

	command := Canvas(
		Move(1, 1),
		Move(2, 2))

	got := command.Do()

	if want != got {
		t.Fatalf("\nwant:\n%v\n got:\n%v\n", want, got)
	}
}

func Test_Move_Undo(t *testing.T) {
	want := "->Point(1, 1)\n"

	command := Canvas(
		Move(1, 1),
		Move(2, 2))

	got := command.Undo()

	if want != got {
		t.Fatalf("\nwant:\n%v\n got:\n%v\n", want, got)
	}
}

//
// Command
//
type Command interface {
	Do() string
}

//
// Canvas
//
func Canvas(commands ...Command) *canvas {
	return &canvas{commands}
}

type canvas struct {
	commands []Command
}

func (c *canvas) Do() string {
	result := ""
	for _, command := range c.commands {
		result += command.Do()
	}
	return result
}

func (c *canvas) Undo() string {
	c.commands = c.commands[:len(c.commands)-1]
	return c.Do()
}

//
// Move
//
func Move(x, y int) *move {
	return &move{x, y}
}

type move struct {
	x, y int
}

func (m *move) Do() string {
	return fmt.Sprintf("->Point(%v, %v)\n", m.x, m.y)
}
