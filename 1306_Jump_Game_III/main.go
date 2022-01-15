package main

//the idea is this, ALWAYS START WITH THE BRUTE FORCE APPROACH ANYTHING WITH COMBOS
//DECISIONS
//DECISION TREES
//DRAW UP USUALLY HASH,SET, OR DFS OR BFS will solve many many things
//draw up a BRUTE FORCE DFS solution, look at the tree, does a BFS look for appeazing?
//does memoization fit? (usually with DFS it does/recuirsion)
//and finally from does that a DP ring a bell? can i find a sub problem and solve starting at
//beginning or end??

//in Jump Game III, you are told that you have TWO(2) decisons to be made either i + nums[i] or i - nums[i]
//that you can jump to, your target is to find a index with a value of zero, you are also given a starting index

//so brtute force, we can map okut these paths, starting at node/index X, its ADJACENT NODES ARE i + nums[i] AND i - nums[i],
//repeat that process you'll see solutions come deeper in the tree so dfs(), its a tree traversal
func canReach(arr []int, start int) bool {
	visited := map[int]bool{} //for values we've seen or computed already!, also acts like a 'visited' array/map/set
	return dfs(visited, arr, start)
}

//OUCHHHH THE GOTCHAS IS WE ARE WORKING WITH A GRAPH Or as a graph ALWAYS ALWAYS remmeber
//that GRAPHS will have cycles, unless pointed outtttt!!!!, because if you think about it

//if parent Node go to child in child child computes to parent again? you have never ending loop
//this is where the caching HELPS!!!, caching/memoization in times like some graphs act like a
//'VISITED' check to tell you we've already visited this node

func dfs(visited map[int]bool, arr []int, idx int) bool {
	//base cases,  'where we want to stop'
	//if the index is oob
	//if the index's value is our target

	if idx < 0 || idx >= len(arr) {
		return false
	}

	if arr[idx] == 0 {
		return true
	}
	//now we have 2 choices/DECISIONS , i + nums[i] OR i - nums[i]
	//search both
	if !visited[idx] {
		visited[idx] = true
		if dfs(visited, arr, idx+arr[idx]) || dfs(visited, arr, idx-arr[idx]) {

			return true
		}
	}

	//if we haven't found anything and are at the end of the road we return false
	return false

}
