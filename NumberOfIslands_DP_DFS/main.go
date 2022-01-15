package main

//ITS JUST LIKE searching for the highest peaks between two mountains problem
//where you call dfs on TOP,LEFT,RIGHT,DOWN of the current cell, and the base case is
//you pass the previous value to the next dfs call, and check as long as the index passed is
//NOT OUT OF BOUNDS, meaning its not less than the length of the 2D array, or MORE than the length
//AND THE CURRENT HEIGHT(CELL) IS GREATER THAN OR EQUAL TO PREVIOUS, YOU CONTINUE
//OTHERWISE YOU STOP, PROCESSING THAT CELL

//this number of island is VERY similar, in that you need to search all cells for specific things
//no matter what i need to search all cells, we can think of this as DFS search problem
//where we will search a cells neightbors (TOP, LEFT, RIGHT, DOWN) cells, and if they continue
//to have '1's then we can continue our DFS call and check each cell as vbisited

//THAT IS HOW WE CHECK A SINGLE CELL< remember that when given a problem the first example
//should give you a hint on how to solve a subset of the problem

//SO IF I KNOW HOW TO FIND ONE ISLAND, how do i find the rest? i knmow that DFS(cell[0,0]) and search all
//neighbors will continue to give me a picture of the entire issland, how do i find if there are other ones further?

//loop!!, why don't we loop thrpugh all cells, marking them as visited, if they have been visited lets ommit them

//if the matrix cannot be modified, we can just create a visited matrix as the same as
//input, and mark each box, for each box, we want to check if its been visited, then we stop our search here for
//this stack call

//otherwise if a cell has not been visited, AND its inboundm AND its a 1 we want to mark it visited, and check its neighbors

import (
	"fmt"
	"testing"
)

func main() {

	sectionMap := [][]byte{
		{'1', '1', '1', '0', '0', '0'},
		{'1', '1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0', '0'},
		{'1', '0', '0', '0', '1', '0'},
		{'1', '1', '0', '0', '0', '0'},
	}

	fmt.Println("num: ", Find(10))
	fmt.Printf("number of Islands %d\n", FindNumberOfIsland(sectionMap))
	fmt.Printf("Number of Islands %d\n", FindNumberOfIslandV2(sectionMap))

	a := testing.Benchmark(BenchMarkV1)
	b := testing.Benchmark(BenchMarkV2)

	fmt.Printf("%s\n%s", a, b)
}

func BenchMarkV1(b *testing.B) {

	//1850380	       638.8 ns/op

	sectionMap := [][]byte{
		{'1', '1', '1', '0', '0', '0'},
		{'1', '1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0', '0'},
		{'1', '0', '0', '0', '1', '0'},
		{'1', '1', '0', '0', '0', '0'},
	}
	for i := 0; i < b.N; i++ {
		FindNumberOfIsland(sectionMap)
	}
}

func BenchMarkV2(b *testing.B) {
	//	14906739	        79.63 ns/op
	sectionMap := [][]byte{
		{'1', '1', '1', '0', '0', '0'},
		{'1', '1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0', '0'},
		{'1', '0', '0', '0', '1', '0'},
		{'1', '1', '0', '0', '0', '0'},
	}
	for i := 0; i < b.N; i++ {
		FindNumberOfIslandV2(sectionMap)
	}
}
func FindNumberOfIsland(sectionMap [][]byte) int {
	visitedMap := make([][]bool, len(sectionMap))
	//initalize the visited map
	for i := range visitedMap {
		visitedMap[i] = make([]bool, len(sectionMap[0]))
	}

	//should be 3 islands
	//simple test of recursion if i understand it right

	//well uise this direction strcture to help our directions of search for islands in DFS
	directions := [][]int{
		//X,Y
		{-1, 0}, //left
		{0, -1}, //up
		{1, 0},  //right +1
		{0, 1},  //Down -1
	}
	landSymbol := byte('1')
	numIsland := 0
	for row := 0; row < len(sectionMap); row++ {
		for col := 0; col < len(sectionMap[0]); col++ {
			//so since maps are pass by reference we ONLY want to iterate when we hit a island land
			//actually we only want to search unvisited panes, since each recursive dfs call will update the visited map, we can ensure that
			//our if statement only calls dfs on landmarks taht are unvisited and sepeated
			//lets call dfs, IF it does call every time then we know this aint gonna work

			//if we find a cell that has a 1 that's LAND we search it and its not been visited
			if sectionMap[row][col] == landSymbol && !visitedMap[row][col] {
				//lets visit start at this point
				dfsSearchForIsland(row, col, visitedMap, sectionMap, directions)
				//fmt.Println("Called") //this is halted, UNTIL aftrer the recursive DFS has finished it's search
				//And BY THEN all VISITED cells would have been marked visited
				//NOW we KNOW that each time dfs is CALLED, it has found yet another island WAIT< NOT
				numIsland++
			}

		}
	}
	return numIsland
}

//REVISED FIND NUM ISLAND IN 2D MATRIX USING GRAPH THEORY -  is to simply ask if you can also
//edit the original input? well if you can, *THINK* if you can, you can use it as a 'visited' map
//so you don't need to allocate another matrix for visited, you can simply flip the visited cell as '0'
//making it so that you dont need to recheck another map, and every subsequent DFS call won't be impacted by cause it woul
//be searching on its latest information
//O(N^2) runtime O(1) sapce no extra allocation of space
func FindNumberOfIslandV2(sectionMap [][]byte) int {

	if len(sectionMap) == 0 {
		return 0
	}

	//WE CAN MODIFY THE ORIGINAL INPUT
	//we don't need a visited matrix
	directions := [][]int{
		//X,Y
		{0, -1}, //UP
		{1, 0},  //RIGHT
		{0, 1},  //DOWN
		{-1, 0}, //LEFT
	}
	//for every row
	land := byte('1')
	numIsland := 0
	for i := range sectionMap {
		for j := range sectionMap[i] {
			if sectionMap[i][j] == land {
				//we have hit land lets search it
				dfsSearchForIslandV2(i, j, sectionMap, directions)
				numIsland++ //REMEMBER that DFS works in this way, it goes DEEP this call stack WILL PAUSE'halt' until the above DFS finishes
				//that is BECAUSE the dfs has a for loop and searching DEEPER before coming back up, and because we only search when we hit land it only calls
				//we do this because we just want to search the entire peice of land we know e hit land but how large is the land? and we don't want to have other dfs
				//calls search the same lkand so we need to mark it as 'visited' or in this case 'water' alreadytt visited
			}
		}
	}
	//for every col
	//if the current cell is NOT a 1 then dfs into it

	//mark the current cell as 0

	return numIsland
}

func dfsSearchForIslandV2(row int, col int, sectionMap [][]byte, directions [][]int) {
	//base case, similarly, we want to continue IFF,
	//we are NOT OUT OF BOUNDS
	//and the current cell IS LAND (1)
	//REMEMBER that these are INDEX's we are comparing agains we need len(sectionMap) -1 (MINUS ONE)
	if row < 0 || row > len(sectionMap)-1 || col < 0 || col > len(sectionMap[0])-1 || sectionMap[row][col] == byte('0') {
		return //exit case, we hit water, or we out of bounds
	}
	//else WE KNOW WE ARE IN LAND BBYY lets search its neighbors
	//lets mark this as VISITED OR IN this version, we can mody the original input
	//mark it as WATER, therefore laster subsequent visits will allow it to exist earlier, instead of visiting it again
	sectionMap[row][col] = '0'
	for _, dir := range directions {
		//we can either search directions HERE NOW for neighors, if neighbor is water forget searching it further,
		//lets keep it simple
		//dir[0] = X, dir[1] = Y
		dfsSearchForIslandV2(row+dir[0], col+dir[1], sectionMap, directions)
	}
}
func Print2DMap(m [][]bool) {
	for i := range m {
		for j := range m[i] {
			fmt.Printf("%t ", m[i][j])
		}
		fmt.Printf("\n")
	}
}

func dfsSearchForIsland(row int, col int, visited [][]bool, sectionMap [][]byte, directions [][]int) {
	//ensure that we are NOT our of bounds, AND that the current cell has not been visited
	//OUR BASE CASE: Exit Case are
	//if we are out of bounds
	//if the cell has been visited
	//if the cell is 'zero'(0)
	if row < 0 || row > len(sectionMap)-1 || col < 0 || col > len(sectionMap[0])-1 || sectionMap[row][col] == byte('0') || visited[row][col] {
		return //we exit, if we reached above base case/exit case
	}

	visited[row][col] = true
	//now we KNOW that the current cell unvisited, and EQUALS '1', lets mark it visited and search our search!
	for _, dir := range directions {
		//take the current CELL, and search its four corners 'dfs'
		//dir[0] = X , dir[1] = Y
		dfsSearchForIsland(row+dir[0], col+dir[1], visited, sectionMap, directions)
	}
}

func dfs() {
	//this should print out for some time then the caller will continue, so i need
	//to BLOCK from here,
	for i := 0; i < 10; i++ {
		fmt.Println("Fuck yeah ", i)
	}
}

//one form of recursion that exists,
// if the inner recursive call doens't have any  'halt' logic, it WON't Halt
//meaning if i have a loop in the Find it would work, because loop iterates 'halting'
//that respectice call stack
//BELOW is a find example of a 'GOTCHAS' ~ now an example of 'halting' is the above dfs function
//where the recursive call CONTAINS a loop, now each recursive call, itself
//is going to 'halt and call its subsequent recursive loops' before its parent loop can continue
//therefore you are able to control the loop of recrsive call
func Find(a int) int {
	fmt.Println(a)
	if a != 5 {
		v := a - 1
		Find(v)
	}

	return a
}
