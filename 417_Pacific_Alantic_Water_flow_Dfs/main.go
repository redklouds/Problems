package main

func pacificAtlantic(heights [][]int) [][]int {
    //run through the matrix checking 4running dfs top bottom and left and right into
    
    //create two ocean matrixs to track the visited/steps we make inward to the center, at the end the two matrixs will have an 'edge'/'rift' showing the heightest peaks 
    
    pacificOcean := initMatrix(len(heights), len(heights[0]))
    alanticOcean := initMatrix(len(heights), len(heights[0]))
    
    //have the oceans create the directions
    
    directions := [][]int{
        //ROW, COL = UP/DOWN | LEFT,Right
        {0,-1}, //LEFT
        {-1,0}, //UP
        {1 ,0}, //DOWN
        {0, 1}, //RIGHT
    }
    
    //now we need to setup the serach for each BOUNDARY meaning that we need to dfs on the TOP and LEFT row and COL
    //for the pacific
    
    
    //search the BOTTOM & RIGHT col and row for the alantic ocean
    
    
    
    //Since its an M X N we can search the TOP and BOTTOM rows together
    
    for col := range heights[0] {
        //search Top ROW - -1 starting since min value is 0
        dfs(heights, 0, col, -1, pacificOcean, directions )
        //Seeach Bottom ROW
        dfs(heights, len(heights)-1, col, -1, alanticOcean, directions )
    }
    
    //the RIGHT and LEFT together
    for row := range heights {
        //Search the LEFT for pacific
        dfs(heights, row, 0, -1, pacificOcean, directions)
        //Search the RIGHT column for alantic
        dfs(heights, row, len(heights[0])-1, -1, alanticOcean, directions)
    }
    
    
    //now ge the result set
    result := MatIntersect(pacificOcean, alanticOcean)
    
    return result
    
    
    
}

//get the intersect of two matrixs and returns their row,col coordinates
func MatIntersect(matA, matB [][]bool) [][]int{
    var result [][]int
    for row := range matA {
        for col := range matA[0]{
            if matA[row][col] && matB[row][col]{
                result = append(result, []int{row,col})
            }
        }
    }
    return result
}


func dfs(heights [][]int, row int , col int, previousHeight int, ocean [][]bool, directions [][]int){
    //edge cases we are out of bounds

    if row < 0 || row > len(heights)-1 || col < 0 || col > len(heights[0])-1 {
        return
    }//dont do anythingh
    
 //idea is to ahve two matrixes repersetning the respective flow of water froming from each ocean
    //pacific and alantic, we will use these as visited array, we will visit these arrays, as we walk the matrix, every step we take from each direction will fill the respective matrixs with our step to the edge, the case of making sure that the next recursive step is larger than the last value will create a 'EDGE' or 'RIFT' on each pacific matrix and alantic ocean, when we iterate over both we can find the cells that have both visited cell,s signifying that these are the rift cells that both have come in contact with , meaning these are the cells that the next negibors stopp dfs searching and because their neighbors are not larger than itself
    
    //since we are IN BOUNDS is the current cell visited OR smaller than the previous value??
    
    if heights[row][col] < previousHeight || ocean[row][col] {
        //means the the previous Height was LARGER than the curent we exit
        return
    }
    
    //now we are here, meaning that (1) we ARE WITHIN BOUNDARIES
    //(2) THE CURRENT VALUE IS LARGER OR EQUAL to the previous value we can mark this ocean cell as visited and serach its neighbors
    ocean[row][col] = true
    
    for _, dir := range directions {
        dfs(heights, row + dir[0], col + dir[1], heights[row][col], ocean, directions)
    }
}

func initMatrix(n, m int) [][]bool{
    tempM := make([][]bool, n)
    for i := range tempM{
        tempM[i] = make([]bool, m)
    }
    return tempM
}
