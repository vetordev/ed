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
	begin   *Coordinates
	end     *Coordinates
}

type Square struct {
	state   int
	crossed bool
}

type Coordinates struct {
	X int
	Y int
}

func (m *Maze) GetSquare(coordinates *Coordinates) (Square, error) {
	if coordinates.X > len(m.squares) || coordinates.Y > len(m.squares[coordinates.X]) {
		return *new(Square), errors.New("coordinates exceeded the limits of the maze")
	}

	return m.squares[coordinates.X][coordinates.Y], nil
}

func (m *Maze) Begin() *Coordinates {
	return m.begin
}

func (m *Maze) AtTheBeginning(coordinates *Coordinates) bool {
	return *m.begin == *coordinates
}

func (m *Maze) AtTheEnd(coordinates *Coordinates) bool {
	return *m.end == *coordinates
}

func (m *Maze) PrintMaze(file string, player *Coordinates) {

	data := ""
	for x, row := range m.squares {
		for y, square := range row {
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

	os.WriteFile(file, []byte(data), 0600)
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
		&Coordinates{1, 0},
		&Coordinates{4, 5},
	}
}
