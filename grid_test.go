package grid_test

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/elight/grid"
)



func Test3x3Transform(t *testing.T) {
	cells := grid.Grid{
		{1, 0, 1},
		{0, 1, 0},
		{1, 0, 1}}

	new_cells := grid.Transform(cells)

	plus_cells := grid.Grid{
		{0, 1, 0},
		{1, 0, 1},
		{0, 1, 0}}

	assert.Equal(t, plus_cells, new_cells)
}

func TestPlusTransform(t *testing.T) {
	cells := grid.Grid{
		{0, 1, 0},
		{1, 1, 1},
		{0, 1, 0}}


	new_cells := grid.Transform(cells)


	outer_cells := grid.Grid{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1}}

	assert.Equal(t, outer_cells, new_cells)
}

func TestFindingNeighbors(t* testing.T) {
	cells := grid.Grid{
		{1, 0, 1},
		{0, 1, 0},
		{1, 0, 1}}

	var coord grid.Coord

	coord = grid.Coord{0, 0}
	expected := []grid.Coord{{1,1}, {0,1}, {1,0}}
	assertNeighborsFor(t, coord, cells, expected)

	coord = grid.Coord{1, 1}
	expected = []grid.Coord{
		{0, 0}, {0, 1}, {0, 2},
	        {1, 0}, {1, 2},
		{2, 0}, {2, 1}, {2, 2}}
	assertNeighborsFor(t, coord, cells, expected)
}

func TestIsCoordInBounds(t* testing.T) {
	cells := grid.Grid{
		{1, 0, 1},
		{0, 1, 0},
		{1, 0, 1}}

	assert.False(t, grid.IsCoordInBounds(grid.Coord{0,0}, grid.Coord{0,0}, cells))
	assert.True(t, grid.IsCoordInBounds(grid.Coord{0,0}, grid.Coord{0,1}, cells))
	assert.True(t, grid.IsCoordInBounds(grid.Coord{0,0}, grid.Coord{1,0}, cells))
	assert.True(t, grid.IsCoordInBounds(grid.Coord{0,0}, grid.Coord{1,1}, cells))
}

func TestNextStateFor(t* testing.T) {
	assert.Equal(t, 0, grid.NextStateFor(1, 0))
	assert.Equal(t, 0, grid.NextStateFor(1, 1))
	assert.Equal(t, 1, grid.NextStateFor(1, 2))
	assert.Equal(t, 1, grid.NextStateFor(1, 3))
	assert.Equal(t, 0, grid.NextStateFor(1, 4))

	assert.Equal(t, 0, grid.NextStateFor(0, 2))
	assert.Equal(t, 1, grid.NextStateFor(0, 3))
	assert.Equal(t, 0, grid.NextStateFor(0, 4))
}

func assertNeighborsFor(t* testing.T, origin grid.Coord, cells grid.Grid, expected []grid.Coord) {
	actual := grid.NeighborsFor(origin, cells)

	var found bool
	for _, expected_coord := range expected {
		found = false
		for _, coord := range actual {
			if expected_coord == coord { 
				found = true
			}
		}
		assert.True(t, found, "For %v, expected to find neighbor %v but didn't", origin, expected_coord)
	}
}

