package grid

type Coord struct {
	X int
	Y int
}

type Grid [][]int

func Transform(grid Grid) Grid {
	new_grid := make(Grid, len(grid))
	for y, row := range grid {
		new_grid[y] = make([]int, len(grid[y]))
		for x, v := range row {
			new_grid[y][x] = NextStateFor(v, LivingNeighborsFor(Coord{x,y}, grid))
		}
	}
	return new_grid
}

func LivingNeighborsFor(coord Coord, grid Grid) (numAlive int) {
	numAlive = 0
	for _, coord := range NeighborsFor(coord, grid) {
		if grid[coord.Y][coord.X] == 1 {
			numAlive++
		}
	}
	return
}

func NeighborsFor(coord Coord, grid Grid) (neighbors []Coord) {
	neighbors = make([]Coord, 0)

	var candidate Coord
	for _, dx := range([]int{-1, 0, 1}) {
		for _, dy := range([]int{-1, 0, 1}) {
			candidate =  Coord{coord.X + dx, coord.Y + dy}
			if IsCoordInBounds(candidate, coord, grid) {
				neighbors = append(neighbors, candidate)
			}
		
		}
	}		
	return
}

func NextStateFor(alive int, numLivingNeighbors int) (aliveAfter int) {
	aliveAfter = 0
	if (alive == 1 && (numLivingNeighbors == 2 || numLivingNeighbors == 3)) ||
		numLivingNeighbors == 3 {
		aliveAfter = 1
	}
	return
}

func IsCoordInBounds(coord Coord, origin Coord, grid Grid) bool {
	return  coord.X >= 0 && coord.X < gridWidth(grid) &&
		coord.Y >= 0 && coord.Y < gridHeight(grid) &&
		!(coord.X == origin.X && coord.Y == origin.Y)
}

func gridWidth(grid Grid) int {
	return len(grid[0])
}

func gridHeight(grid Grid) int {
	return len(grid)
}

