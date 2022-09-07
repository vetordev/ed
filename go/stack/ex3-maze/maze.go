package maze

import (
	"errors"
	"os"
)

const (
	Free = 0
	Wall = 1
)

type Maze struct {
	squares [][]Square
	begin   Coordinates
	end     Coordinates
}

type Square struct {
	state   int
	crossed bool
}

type Coordinates struct {
	X int
	Y int
}

func (m *Maze) GetSquare(coordinates *Coordinates) (*Square, error) {
	if coordinates.X < 0 || coordinates.Y < 0 || coordinates.Y >= len(m.squares) || coordinates.X >= len(m.squares[coordinates.Y]) {
		return new(Square), errors.New("coordinates exceeded the limits of the maze")
	}

	return &m.squares[coordinates.Y][coordinates.X], nil
}

func (m *Maze) Begin() Coordinates {
	return m.begin
}

func (m *Maze) AtTheBeginning(coordinates *Coordinates) bool {
	return m.begin == *coordinates
}

func (m *Maze) AtTheEnd(coordinates *Coordinates) bool {
	return m.end == *coordinates
}

func (m *Maze) PrintMaze(file *os.File, player *Coordinates) {

	data := ""
	for y, row := range m.squares {
		for x, square := range row {
			ch := "  "

			coordinates := Coordinates{x, y}
			if coordinates == *player {
				ch = "👾"
			} else if square.state == Wall {
				ch = "🧱"
			} else if m.AtTheBeginning(&coordinates) || m.AtTheEnd(&coordinates) {
				ch = "🚪"
			}

			data += ch + " "
		}
		data += "\n"
	}

	file.WriteString(data)
	file.Sync()
}

func CreateMaze() *Maze {

	defaultMaze := [][]int{
		{Wall, Wall, Wall, Wall, Wall, Wall},
		{Free, Free, Wall, Wall, Wall, Wall},
		{Wall, Free, Free, Free, Wall, Wall},
		{Wall, Free, Wall, Free, Wall, Wall},
		{Wall, Free, Wall, Free, Free, Free},
		{Wall, Wall, Wall, Wall, Wall, Wall},
	}

	squares := make([][]Square, len(defaultMaze))

	for x, row := range defaultMaze {
		squares[x] = make([]Square, len(defaultMaze[0]))
		for y, state := range row {
			squares[x][y] = Square{state, false}
		}
	}

	return &Maze{
		squares,
		Coordinates{0, 1},
		Coordinates{5, 4},
	}
}
