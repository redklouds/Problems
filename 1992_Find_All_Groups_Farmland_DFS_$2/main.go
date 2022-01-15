package main

import "fmt"

type Node struct {
	Row int
	Col int
}

func findFarmLand(land [][]int) [][]int {
	//uising BFS since we know that BFS is using a QUEUE FIFO, basically in order queuing we can assume with this method that
	//the queue of nodes will queue the nodes in a way which reads on your define, ie all right then we can assume if we
	//enqueue nodes RIGHT first, that it will read all the top levels first then the second level nodes, the last being the last farm land aval

	//IDEA search the grid when we hit a 1 farm land top left corner, record its position, then we want to
	//create a queue, and perform a bfs search until we find a node that fits our criteria of having either 0 or edge for right and bottom

	var result [][]int
	for row := range land {
		for col := range land[0] {
			if land[row][col] == 1 {
				//we hit landR				//create the queue
				startNode := Node{row, col}
				//remember that ROW and COL are the FIRST top left corner coords
				endNode := bfs(land, startNode)
				result = append(result, []int{startNode.Row, startNode.Col, endNode.Row, endNode.Col})
			}
		}
	}
	return result
}

func bfs(land [][]int, startNode Node) (endNode Node) {

	dirc := [][]int{
		{0, 1},
		{1, 0},
	}
	q := []Node{startNode}

	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		//make sure we dont have leaks

		//check the conditions
		if (n.Row+1 > len(land)-1 || land[n.Row+1][n.Col] == 0) && (n.Col+1 > len(land[0])-1 || land[n.Row][n.Col+1] == 0) {
			//we found the last end of things
			return n
		}

		//else we continue
		//if the right has not been visited, enqueue it and mark it visited
		for _, dir := range dirc {
			//if visited

			//only enqueue the actual nodes, if oob then don't queue it
			nextRow := n.Row + dir[0]
			nextCol := n.Col + dir[1]
			if nextRow < 0 || nextRow > len(land)-1 || nextCol < 0 || nextCol > len(land[0])-1 {
				continue
			}
			land[nextRow][nextCol] = 0
			q = append(q, Node{n.Row + dir[0], n.Col + dir[1]})
		}

	}

	return Node{}

}

func main() {

	land := [][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}

	res := findFarmLand(land)
	fmt.Println(res)

}
