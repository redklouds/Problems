package main

/*
//given m x N matrix, find the unique paths
//this is just another permutatioin problem, as lojnjg as you figure out
//what your input is, you find all the permutations for those rules
func UniquePaths_BruteForce(m, n int) int {

}
*/

//the idea is simple with DP, THINK about the Subproblem

//THINK about the REAL QUESTION, "What are the NUMBER of unique WAYS to get from A to B"
//if you start at the cells right before the botrtom right top and Left, you can only move Right or Down
//so on the top of the target cell, thre is ONE(1) unique way to move downn
//and on the right there is only ONE(1) unique way to move right ..
//now if the quewswtion was to ask you how many unique ways does it tgake for the sleect cell on the left of the star? you'd answer 1 because you
//can onlyu go right to the star, the idea is that simple, at each cell, you can either go RIGHT or DOWN
//if you go right.. the compute for how many way to go to the star is already computed.. and if yuou go down its already computed
//just store those number of ways in a dp matrix and add them because taking RIGHT or DOWN, + those neighbors are the number of ways THEY take to get to
//the same destination, THINK

//its like asking your neighbors on a chest board, "to the guy on the right of me", have you travelled to the Queen Before? he will likely say yes(because he has)// you can simply ask him WITHOUT walking YOURSELF, 'so how many unqiue paths did you take to get to the queeen?' and he'll tell you since hes' done it
//then you look to your nieghbor standing at the bottom, and think 'i want to walk to the queen taking bottom path', and ask him how many ways hes found out
//you can infer that.. so.... from whree i've standing, if i go right there are X many times unique, and if i go bottom there are Y unique ways to walk tot he queen
//therefore from where im standing X + Y = Number of unique ways to walk to the queen, WITHOUT EVEN WALKING THERE, because people next to you have done it!

func UniquePaths(m int, n int) int {

	if m == 0 {
		return 0
	}
	//set up our dp matrix

	//dp matrix array that describes "if i land on this Row X COL how many unique paths do i have to get to the star"
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1) //so we have extra edge buffer so we don't need to do extra if conditions'you'l;l see
	}

	//set up directions, cleaner IMO
	directions := [][]int{
		//X,Y
		{0, 1}, //RIGHT
		{1, 0}, //DOWN
	}

	//remember zero based on the edges, since we added 1 extra edge (danny you'll probably forget all of this when you look again you idiot :3)
	for row := m - 1; row >= 0; row-- {

		for col := n - 1; col >= 0; col-- {

			//if the row and col is the STAR (Finish) we need to initalize our DP matrix so its 1, (inital state), remember allmDP have an inital subproblem STATE you use to solve
			//INITAL STATE is matrix[STARX][STARY] = 1
			if row == m-1 && col == n-1 {
				dp[row][col] = 1
			} else {
				//alright we are not at inital state lets do this
				for _, dir := range directions {
					//for RIGHT and BOTTOM combine the unique ways your neighbors have already computed
					dp[row][col] += dp[row+dir[0]][col+dir[1]] //dir[0]= row Dir[1] = Col(Y)
				}
			}
		}
	}
	return dp[0][0]
}
