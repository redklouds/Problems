package main

func generateParenthesis(n int) []string {
	/*
	   the idea is that at each interval you have two choices you can either add a '(' or a ')'
	   but you know that if there are 3 open there should be 3 closed, so the rule to add a closing
	   is AS LONG AS THE NUMBER OF CLOSE IS LESS THAN THE OPENS YOU CAN ADD A CLOSING
	   if close < open -> you can add a close
	   now to add an open, you can only add opens if opens count equals N, where N tells you number of pairs
	   //and where pairs are in the form of '()' you'll have 3 open and 3 close
	   so the kicker is.. how do we backtrack ? well, we KNOW that yoyu ALWAYS have to start with an open
	   before any closing is done, why not liek PERMUTATION, where you revert the state, revert the state, when you finish
	   adding opens down?
	*/

	arr := []byte{}
	solution := make([]string, 1)
	dfs(0, 0, n, arr, solution)
	return solution

}

func dfs(openCount int, closeCount int, n int, arr []byte, sol []string) {
	if openCount == closeCount && closeCount == n {
		sol = append(sol, string(arr))
		return
	}

	if openCount < n {
		//add an open paran
		arr = append(arr, '(')
		//recurse
		dfs(openCount+1, closeCount, n, arr, sol)
		//reverse the stae
		arr = arr[:len(arr)-1]
	}

	if closeCount < openCount {
		//add an open paran
		arr = append(arr, ')')
		//recurse
		dfs(openCount, closeCount+1, n, arr, sol)
		//reverse the stae
		arr = arr[:len(arr)-1]
	}
}
