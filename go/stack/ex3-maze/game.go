package maze

import (
	stack2 "ed/stack"
	"errors"
	"os"
)

func findNextMove(m *Maze, player *Coordinates) (Coordinates, error) {
	movements := []Coordinates{
		{player.X + 1, player.Y},
		{player.X, player.Y + 1},
		{player.X - 1, player.Y},
		{player.X, player.Y - 1},
	}

	for _, move := range movements {
		if sq, err := m.GetSquare(&move); err == nil && sq.crossed == false && sq.state == Free {
			return move, nil
		}
	}

	return *new(Coordinates), errors.New("there's no movement available")
}

func Init(file *os.File) {
	maze := CreateMaze()
	way := stack2.NewStack[Coordinates]()

	player := maze.Begin()
	maze.PrintMaze(file, &player)

	for !maze.AtTheEnd(&player) {
		if next, err := findNextMove(maze, &player); err == nil {
			way.Push(next)
		} else {
			way.Pop()
		}

		player, _ = way.Top()

		maze.PrintMaze(file, &player)
	}
}
