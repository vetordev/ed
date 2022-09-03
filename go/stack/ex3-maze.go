package stack

type Maze struct {
	squares [][]int
	start   Coordinates
	end     Coordinates
}

type Coordinates struct {
	X int
	Y int
}

func (m *Maze) PrintMaze(player Coordinates) string {

	data := ""
	for x, row := range m.squares {
		for y, square := range row {
			ch := "  "

			coordinates := Coordinates{x, y}
			if coordinates == player {
				ch = "👾"
			} else if square == 1 {
				ch = "🧱"
			} else if m.isStart(coordinates) || m.isEnd(coordinates) {
				ch = "🚪"
			}

			data += ch + " "
		}
		data += "\n"
	}

	return data
}

func (m *Maze) isStart(coordinates Coordinates) bool {
	return m.start == coordinates
}

func (m *Maze) isEnd(coordinates Coordinates) bool {
	return m.end == coordinates
}

func CreateMaze() *Maze {
	return &Maze{
		[][]int{
			{1, 1, 1, 1, 1, 1},
			{0, 0, 1, 1, 1, 1},
			{1, 0, 0, 0, 1, 1},
			{1, 0, 1, 0, 1, 1},
			{1, 0, 1, 0, 0, 0},
			{1, 1, 1, 1, 1, 1},
		},
		Coordinates{1, 0},
		Coordinates{4, 5},
	}
}
