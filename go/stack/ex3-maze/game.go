package maze

import (
	stack2 "ed/stack"
)

func findNextSquare(m *Maze, player *Coordinates, crossed bool) *Coordinates {
	var nextCoordinates Coordinates

	if sq, err := m.GetSquare(&Coordinates{player.X + 1, player.Y}); err == nil && sq.crossed == crossed && sq.state == Free {
		nextCoordinates = Coordinates{player.X + 1, player.Y}
	} else if sq, err = m.GetSquare(&Coordinates{player.X, player.Y + 1}); err == nil && sq.crossed == crossed && sq.state == Free {
		nextCoordinates = Coordinates{player.X, player.Y + 1}
	} else if sq, err = m.GetSquare(&Coordinates{player.X - 1, player.Y}); err == nil && sq.crossed == crossed && sq.state == Free {
		nextCoordinates = Coordinates{player.X - 1, player.Y}
	} else if sq, err = m.GetSquare(&Coordinates{player.X, player.Y - 1}); err == nil && sq.crossed == crossed && sq.state == Free {
		nextCoordinates = Coordinates{player.X, player.Y - 1}
	}

	if (nextCoordinates != Coordinates{}) {
		return &nextCoordinates
	}

	return findNextSquare(m, player, false)
}

func Init() {
	maze := CreateMaze()
	way := stack2.NewStack[Coordinates]()

	player := maze.Begin()
	maze.PrintMaze("maze.txt", player)

	for !maze.AtTheEnd(player) {
		next := findNextSquare(maze, player, true)
		way.Push(*next)

		player = next

		maze.PrintMaze("maze.txt", player)
	}
}
