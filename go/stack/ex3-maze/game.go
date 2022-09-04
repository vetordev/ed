package maze

import stack2 "ed/stack"

// deve retornar a proxima coordenada
// verifica se ha um quadrado possivel de ser caminhado
func findNextSquare(m *Maze, player *Coordinates, crossed bool) Coordinates {
	var nextCoordinates Coordinates

	if sq, err := m.GetSquare(Coordinates{player.X + 1, player.Y + 1}); err != nil && sq.crossed == crossed {
		nextCoordinates = Coordinates{player.X + 1, player.Y + 1}
	} else if sq, err := m.GetSquare(Coordinates{player.X, player.Y + 1}); err != nil && sq.crossed == crossed {
		nextCoordinates = Coordinates{player.X, player.Y + 1}
	} else if sq, err := m.GetSquare(Coordinates{player.X, player.Y - 1}); err != nil && sq.crossed == crossed {
		nextCoordinates = Coordinates{player.X, player.Y - 1}
	} else if sq, err := m.GetSquare(Coordinates{player.X - 1, player.Y}); err != nil && sq.crossed == crossed {
		nextCoordinates = Coordinates{player.X - 1, player.Y}
	}

	if (nextCoordinates != Coordinates{}) {
		return nextCoordinates
	}

	return findNextSquare(m, player, false)
}

func Init() {
	maze := CreateMaze()
	way := stack2.NewStack[Square]()

	player := maze.Begin()
	maze.PrintMaze("maze.txt", *player)

	for maze.AtTheEnd(*player) {

	}
}
