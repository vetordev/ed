package maze

import (
	stack2 "ed/stack"
)

func findNextMove(m *Maze, player *Coordinates, crossed bool) *Coordinates {
	movements := []Coordinates{
		{player.X + 1, player.Y},
		{player.X, player.Y + 1},
		{player.X - 1, player.Y},
		{player.X, player.Y - 1},
	}

	for _, move := range movements {
		if sq, err := m.GetSquare(&move); err == nil && sq.crossed == crossed && sq.state == Free {
			return &move
		}
	}

	return findNextMove(m, player, !crossed)
}

func Init() {
	maze := CreateMaze()
	way := stack2.NewStack[Coordinates]()

	player := maze.Begin()
	maze.PrintMaze("maze.txt", player)

	for !maze.AtTheEnd(player) {
		next := findNextMove(maze, player, false)
		way.Push(*next)

		player = next

		maze.PrintMaze("maze.txt", player)
	}
}
