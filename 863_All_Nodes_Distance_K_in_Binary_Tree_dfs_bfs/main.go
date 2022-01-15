package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

//There must be a faster way, REDO too slow

//
func convertToKnaryTrie(root *TreeNode, adjList map[int][]int) {
	//using DFS, idk faster?
	//go all the way to the left
	if root == nil {
		return
	}
	if root.Left != nil {
		//go left
		adjList[root.Val] = append(adjList[root.Val], root.Left.Val)
		adjList[root.Left.Val] = append(adjList[root.Left.Val], root.Val)
		convertToKnaryTrie(root.Left, adjList)
	}
	if root.Right != nil {
		adjList[root.Val] = append(adjList[root.Val], root.Right.Val)
		adjList[root.Right.Val] = append(adjList[root.Right.Val], root.Val)
		convertToKnaryTrie(root.Right, adjList)
	}
}

func bfsLevelTraversal(adjList map[int][]int, startingRoot int, targetLevel int) []int {
	visitedSet := make(map[int]bool)
	level := 0 //start at zero since when we start, we will start at the root equal zero, and go to the target level we want
	q := []int{startingRoot}
	visitedSet[startingRoot] = true
	for len(q) != 0 {
		//if we are on our current target level lets popualte the soluton
		if level == targetLevel {
			//populate
			resultSet := make([]int, len(q))
			for i := 0; i < len(q); i++ {
				resultSet[i] = q[i]
			}

			return resultSet
		}

		qLen := len(q)
		//do i need to walk throgh every node?? when trying to reach the level?

		
		for i := 0; i < qLen; i++ {

			for _,neigh := range adjList[q[i]]{
				//for each nieghbor 
				if _, found := visitedSet[neigh]; !found {
					q = append(q, neigh)
					visitedSet[neigh] = true
				}
			}
			//lets start :D
			//if we are on our target level lets populate and push the resultset
			//queue in each node for the level
	

		}
		//update level cound
		level++
		//trim quuee
		q = q[qLen:]
	}

	//if we didn't return it means that we couple reach that level, that's why we here
	return []int{}

}
func main() {

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}




	adjList := make(map[int][]int)
	convertToKnaryTrie(root, adjList)
	bfsLevelTraversal(adjList, 5, 2)


	root1 := &TreeNode{
		Val: 1,
		Left: nil,
		Right: nil,
	}
	adjList = make(map[int][]int)
	convertToKnaryTrie(root1, adjList)
	bfsLevelTraversal(adjList, 1, 3)
}
