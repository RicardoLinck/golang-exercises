package main

import (
	"fmt"
	"os"
)

type Point struct {
	X int
	Y int
}

type Square struct {
	Center Point
	Length int
}

func (s *Square) Move(dx, dy int) {
	s.Center.Move(dx, dy)
}

func (s Square) Area() int {
	return s.Length * s.Length
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func NewSquare(x, y, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0")
	}

	s := &Square{
		Center: Point{x, y},
		Length: length,
	}

	return s, nil
}

func main() {

	s, err := NewSquare(1, 2, 4)

	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	s.Move(1, 1)
	area := s.Area()

	fmt.Println(area)

	fmt.Printf("%+v\n", s)
}
