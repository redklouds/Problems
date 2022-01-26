package main

func main() {
}

type Coords struct {
	Row int
	Col int
}

func shortestDistance(grid [][]int, targetisland1 int, targetisland2 int) int {

	shortestDistance := 0
	return shortestDistance

}

//given two coordinates return the distance between them
func ShortestDistance(coords1, coords2 Coords) int {
	//manhatten distance formula
	// distance = Max(|x1 - x2|, |y1 - y2|)
	dist := Abs(coords1.Row-coords2.Row) + Abs(coords1.Col-coords2.Col)
	return dist - 1
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
