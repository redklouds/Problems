package main

//O(n * M(for adjacent courses)) O(n) space
func canFinish(numCourses int, prerequisites [][]int) bool {
	//IDEA change the coruse scdhule into an adj List,

	//breaking it down into course, and what is the prereq nodes we need to visit,m

	//to detect a cycle is to use a "PATHSET" basically a set that tells us if we have encountered this node along our dfs traversale before

	//if on our dfs path we have, we can say there is a cycle, otherwise, we can do a "save state/revert state backtracking(dfs) method "

	//create our adj list, we need a list so that we can repersent the graph in a way that we can traverse it repersented as a graph
	//of nodes and its adjacent nodes, since prereq[0] = course prereq[1] = depedency, we can say that the course at prereq[0] has adjacent nodes at prereq[1]
	//so prereq[0] -> prereq[1]

	//we also need a PathSet , a set that recourse the course path of our dfs recursion

	pathSet := make(map[int]bool) //this will be our pathset

	//create adj list

	adjList := make(map[int][]int) //key is the course, and the value is the adjacent course/nodes

	for i := 0; i < numCourses; i++ {
		adjList[i] = []int{}
	}
	for _, coursePair := range prerequisites {
		//for each course, the idx 0 = course , the idx 1 = pre-req couse(adj node)
		adjList[coursePair[0]] = append(adjList[coursePair[0]], coursePair[1])
	}

	//GOTCHAS when creating the adj list, you also want to capture the empty values
	//that's why we use the NUMCOURSE, TO POPULATE THESE ARRAYS

	//perform dfs

	for crs := 0; crs < numCourses; crs++ {
		copyAdjL := make(map[int][]int)
		for k, v := range adjList {
			copyAdjL[k] = v
		}
		if !dfs(pathSet, copyAdjL, crs) {
			return false
		}
	}
	return true
}

func dfs(pathset map[int]bool, adjList map[int][]int, course int) bool {

	//bnase case if the current course has no req, we know it can be completed
	if len(adjList[course]) == 0 {
		return true
	}
	//we want to move our pathset insert here, because each step we want to record it
	//if the current step/node is seen before exit immediately and let the call capture it
	//for each depCourse, check if the path exists in the pathset if so return false, otherwise contuinue
	isNew := setInsert(pathset, course)
	if !isNew {
		return false
	}
	//base case if we have finished

	for _, depCourse := range adjList[course] {
		//we want to CATCH the return statement of DFS here, if this dfs call
		//tells us that its already exists and cycle is here
		//we catch the return and check it
		if !dfs(pathset, adjList, depCourse) {
			return false
		}
		remove(pathset, depCourse)
		adjList[course] = []int{}
	}

	return true
}

//returns false if the value already exists, otherwise returns true on success insert
func setInsert(pathSet map[int]bool, val int) bool {
	_, found := pathSet[val]
	pathSet[val] = true
	return !found //if true = 'found' !true = false, means we already have it here
	//if false = 'notfound' !false = true , means we inserted something new
}
func remove(pathSet map[int]bool, val int) {
	delete(pathSet, val)
}
