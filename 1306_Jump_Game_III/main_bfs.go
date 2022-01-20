package main

func canReach_bfs(arr []int, start int) bool {
	//this is honestly a SEARCH PROBLEM, with searching we have DFS, BFS, Binary Search,
	//in this problem DFS or BFS make most since sense we want to search ALL combinations and possibly visit all without skipping any items

	//think, you can simply use DFS and a visited map to see which index was visited,

	//conditions are if you are out of bounds or visited, then don't visit again,

	//you can also perform a BFS-esk thinking, where you queue in each item in a linear fashion,
	//making the run time O(N+E) and space O(n)

	//lets do bfs
	return bfs(start, arr)
}

func bfs(i int, arr []int) bool {
	visited := make(map[int]bool)
	q := []int{i}
	visited[i] = true

	for len(q) != 0 {
		//for each node, lets enqueue its left ( i + num[i]) and its right (i - num[i]), and repeat, if we are able to enqueue, we are honestly looking for a PATH
		//it doesn't matter how many ways there is, as long as a way exists we can return
		index := q[0]
		q = q[1:] //includes the first but not the last

		if arr[index] == 0 {
			return true
		}
		//check the left
		if (index-arr[index]) >= 0 && !visited[index-arr[index]] {
			//within left boundary
			visited[index - arr[index]] = true
			q = append(q, index-arr[index])
		}
		if (index+arr[index]) <= len(arr)-1 && !visited[index+arr[index]] {
			//within right boundary
			visited[index-arr[index]] = true
			q = append(q, index+arr[index])
		}
	}

	return false
}
