package TopicalOrderingCourseSchdule

import (
	"container/list"
	"fmt"
)

type Queue struct {
	list.List
}

func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *Queue) Enqueue(val interface{}) {
	q.PushFront(val)
}

func (q *Queue) Dequeue() interface{} {
	//if i insert at the front each time i need to retrieve from the back for queue
	//FIFO
	val := q.Back()
	q.Remove(val)
	return val.Value
}

//THIS IS A DAG - Directed Acyclic Graph -> means that for every vertex u and vertex v
//that vertex u comes before vertex v
//there exits NO cycles such as A -> B -> C, C-> B,

//TOPOLOGICAL ORder/Traversal/Sort means that all dependencies are visited processed first
//used in dependecyt or task/schduling whre all dependencies must be completed first before the next
//task can start
//THINK? Booting a machine, HD, CPU, RAM must be initalized first before theBIO can use it
// ie Initalize HDD ->Start Mobo
//baking PreheatOven -> Bake
// Take Course 0 before you can take course 1

// "Array Of Type String"

//there exists MANY MANY MANY difference Toplogical order you can do
//all smallest verticies first
//largest verticies first
//top to bottom (visually)
func ToplogicalSort(digraph map[string][]string) []string {
	//see if there exists a toplogical ordering and return it

	//digraph structure:[
	//       Key: "node",
	//			val : Array if its adjacent Nodes (adjacent meaning what it points too, in a directed graph, if undirected then its all nodes A <--> B)
	//]
	//first create the ingress mapping , which will hold all the indegree/incoming edge counts for each node

	indegree := make(map[string]int)

	//ops we want to initalize the indegree map because we need them to initalize to zero
	for key, _ := range digraph {
		//insert the default node with zero count of indegreee init
		indegree[key] = 0
	}

	for _, adjList := range digraph {
		for _, n := range adjList {
			//B -> C
			//B -> D , where node is B, and N is C and D, we must increment C and D ingress value by 1 since we found that there are incoming edges pointing too it
			//if we encountered this node in an adjacency list, means that N has a incoming edge FROM NODE
			indegree[n]++
		}
	}

	//now we have a indegree map of number of imcoming edge for each we start the party
	//by creating a result topological order result and a queue to start our traversal

	var Toplogical_Order []string

	//var nodes_with_zero_indegree []string
	nodes_with_zero_indegree := Queue{}
	//populate the nodes_with_zero_ingree to get started, IF ITS A DAG (BASE REQUIREMENT) then there will always be one with a zero incoming edges
	for k, v := range indegree {
		//find zero values of incoming edges
		if v == 0 {
			nodes_with_zero_indegree.Enqueue(k)
			//nodes_with_zero_indegree = append(nodes_with_zero_indegree, k)
		}

	}

	//while there exists nodes to process let's process them!
	for !nodes_with_zero_indegree.IsEmpty() {
		//pop and add this to the topological ordering
		//type assert since an interface is returned
		node := nodes_with_zero_indegree.Dequeue().(string)
		Toplogical_Order = append(Toplogical_Order, node)

		//since we removed it from the graph' removed' we must updated the indegreee map since there is no longer this node pointing to others

		//how?
		//we have the adj nodes of the current node
		//we have the indegree map (num of edges)

		for _, adjNode := range digraph[node] {
			//for each adjacentt node, decreement its edge count
			indegree[adjNode]--

			//check if its indegrees have dropped to zero
			//don't need to worry about other nodes blocking it, if an adjacent node still have reference from another one, then its a cycle
			if indegree[adjNode] == 0 {
				//enqueue it into the group of no indegree nodes
				nodes_with_zero_indegree.Enqueue(adjNode)
			}
		}

	}
	//if we have reached here, there are two possibilities,

	//1: there exists a cycle
	//2: there exists No more nodes with zerom indegree meaning we finished processing all the nodes (depedcencies first of course)

	//to check we check if the result array contains the same length as the oroginal number of nodes
	if len(Toplogical_Order) == len(digraph) {
		return Toplogical_Order
	} else {
		panic("There exists a Cycle here")
	}
	//return make([]string, 0)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	diGraph := make(map[int][]int)
	//my directed graph map
	//populate directed graph map

	//populate the directed graph as a map

	//for each pair [[AdjacentNode, Node], ... [AdjacentNodeN-1, NodeN-1]]

	//first check is are there enough unique nodes in the map to satisfy the num courses we exit if there isnt
	for _, pair := range prerequisites {
		//if the other pair does not have adjacent nodes it won't show up in the directed graph need to fix that and initalize it as opposite
		if _, ok := diGraph[pair[0]]; !ok {
			//doesn't exists initalize it here
			diGraph[pair[0]] = make([]int, 0)
		}
		diGraph[pair[1]] = append(diGraph[pair[1]], pair[0]) //assumping that prereq[i].Length is always 2
	}

	if len(prerequisites) == 0 {
		return true
	}
	//len is O(1) operation
	if len(diGraph) >= numCourses {
		return false
	}

	return ToplogicalSort1(diGraph)

}

func ToplogicalSort1(digraph map[int][]int) bool {
	//see if there exists a toplogical ordering and return it

	//digraph structure:[
	//       Key: "node",
	//			val : Array if its adjacent Nodes (adjacent meaning what it points too, in a directed graph, if undirected then its all nodes A <--> B)
	//]
	//first create the ingress mapping , which will hold all the indegree/incoming edge counts for each node

	indegree := make(map[int]int)

	//ops we want to initalize the indegree map because we need them to initalize to zero
	for key, _ := range digraph {
		//insert the default node with zero count of indegreee init
		indegree[key] = 0
	}

	for _, adjList := range digraph {
		for _, n := range adjList {
			//B -> C
			//B -> D , where node is B, and N is C and D, we must increment C and D ingress value by 1 since we found that there are incoming edges pointing too it
			//if we encountered this node in an adjacency list, means that N has a incoming edge FROM NODE
			indegree[n]++
		}
	}

	//now we have a indegree map of number of imcoming edge for each we start the party
	//by creating a result topological order result and a queue to start our traversal

	var Toplogical_Order []int

	//var nodes_with_zero_indegree []string
	nodes_with_zero_indegree := Queue{}
	//populate the nodes_with_zero_ingree to get started, IF ITS A DAG (BASE REQUIREMENT) then there will always be one with a zero incoming edges
	for k, v := range indegree {
		//find zero values of incoming edges
		if v == 0 {
			nodes_with_zero_indegree.Enqueue(k)
			//nodes_with_zero_indegree = append(nodes_with_zero_indegree, k)
		}

	}

	//while there exists nodes to process let's process them!
	for !nodes_with_zero_indegree.IsEmpty() {
		//pop and add this to the topological ordering
		//type assert since an interface is returned
		node := nodes_with_zero_indegree.Dequeue().(int)
		Toplogical_Order = append(Toplogical_Order, node)

		//since we removed it from the graph' removed' we must updated the indegreee map since there is no longer this node pointing to others

		//how?
		//we have the adj nodes of the current node
		//we have the indegree map (num of edges)

		for _, adjNode := range digraph[node] {
			//for each adjacentt node, decreement its edge count
			indegree[adjNode]--

			//check if its indegrees have dropped to zero
			//don't need to worry about other nodes blocking it, if an adjacent node still have reference from another one, then its a cycle
			if indegree[adjNode] == 0 {
				//enqueue it into the group of no indegree nodes
				nodes_with_zero_indegree.Enqueue(adjNode)
			}
		}

	}
	//if we have reached here, there are two possibilities,

	//1: there exists a cycle
	//2: there exists No more nodes with zerom indegree meaning we finished processing all the nodes (depedcencies first of course)

	//to check we check if the result array contains the same length as the oroginal number of nodes
	if len(Toplogical_Order) == len(digraph) {
		return true
	}
	return false
	//return make([]string, 0)
}

func canFinish1(numCourses int, prerequisites [][]int) bool {
	diGraph := make(map[int][]int)
	//my directed graph map
	//populate directed graph map

	//populate the directed graph as a map

	//for each pair [[AdjacentNode, Node], ... [AdjacentNodeN-1, NodeN-1]]

	indegree := make(map[int]int)

	zeroIndegreeN := -1

	//first check is are there enough unique nodes in the map to satisfy the num courses we exit if there isnt
	for _, pair := range prerequisites {
		//if the other pair does not have adjacent nodes it won't show up in the directed graph need to fix that and initalize it as opposite
		if _, ok := diGraph[pair[0]]; !ok {
			//doesn't exists initalize it here
			diGraph[pair[0]] = make([]int, 0)
		}

		if _, ok := indegree[pair[0]]; !ok {
			//doesn't exists initalize it here
			indegree[pair[0]] = 0
		}
		if _, ok := indegree[pair[1]]; !ok {
			//doesn't exists initalize it here
			indegree[pair[1]] = 0
		}
		//this tells me that node 1 has adjacent nodes to node pair 0,
		//so idx1 is the depedecy of idx0, meaning that idx[0] increments its dependencies
		indegree[pair[1]]++
		diGraph[pair[1]] = append(diGraph[pair[1]], pair[0]) //assumping that prereq[i].Length is always 2

	}

	for key, val := range indegree {
		if val == 0 {
			zeroIndegreeN = key
		}
	}

	if len(prerequisites) == 0 {
		return true
	}
	//len is O(1) operation
	if numCourses < len(diGraph) {
		return false
	}

	visited := make(map[int]bool)

	return dfsSearch(zeroIndegreeN, visited, diGraph)

	//return ToplogicalSort(diGraph)

}

func dfsSearch(n int, visited map[int]bool, diGraph map[int][]int) bool {
	visited[n] = true
	for _, adjN := range diGraph[n] {
		if visited[adjN] {
			return false
		}
		dfsSearch(adjN, visited, diGraph)
	}

	return true
}
func main() {
	diGraph := make(map[string][]string)

	//course

	//Struct: [CourseA, CourseA Dependency]
	//[ [A,B], [B,A]]
	//OR
	//[ [162, 161], [161, 162]]
	diGraph["A"] = []string{"B"}
	diGraph["B"] = []string{"A"}

	//res := ToplogicalSort(diGraph)

	//fmt.Printf("Result: %v", res)

	diGraph2 := make(map[string][]string)
	//To go to 162 -> you need 161
	// to go to 342 -> 341
	// to get to 341 -> 162
	// to get to 432 -> 342, 341
	// to get to 432
	//NODE = its adjacent Nodes
	diGraph2["161"] = []string{"162"}
	diGraph2["162"] = []string{"341"}
	diGraph2["341"] = []string{"342", "432"}
	diGraph2["342"] = []string{"432"}
	diGraph2["432"] = []string{}

	res2 := ToplogicalSort(diGraph2)

	fmt.Printf("Result: %v", res2)

	testData := [][]int{
		{1, 0},
	}

	res := canFinish(2, testData)
	if !res {
		fmt.Errorf("ERROR")
	}

	numC := 2
	td := [][]int{
		{1, 0},
		{0, 1},
	}

	res22 := canFinish1(numC, td)

	fmt.Println(res22)
}
