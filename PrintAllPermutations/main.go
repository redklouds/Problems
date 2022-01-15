package main

import "fmt"

func dfs(arr []int, a int) {
	if a == len(arr)-1 {
		fmt.Println(arr)
	}
	for i := a; i < len(arr)-1; i++ {
		//swap, left equals, right equals = left equals, right equals
		arr[i], arr[i+1] = arr[i+1], arr[i]

		dfs(arr, a+1)
		arr[i+1], arr[i] = arr[i], arr[i+1]
	}
}

//start with choosing ZERO(0) as swap index
func Permutation(arr []byte, l int) {
	fmt.Printf("we have choosen index: %d of array: %c\n", l, arr)
	if l == len(arr)-1 {
		fmt.Printf("%c\n", arr)
	} else {
		//not base case there is more to search! //first time choose ZERO, (0,0) -> (0,1) -> (0,2) BAC ->(1,1) BAC -> ->(1,2)from BAC -> BCA ->(2,2)<-bt print(BCA) <-bt (2,2) Print(BAC)<-bt (1,1) -> s(1,2)ACB->(2,2) <-bt (2,2) ->ABC <-bt

		for i := l; i < len(arr); i++ {
			fmt.Printf("Swapping Choosen Index: %d with: %d\n", l, i)
			arr[i], arr[l] = arr[l], arr[i]
			Permutation(arr, l+1)
			arr[l], arr[i] = arr[i], arr[l]
		}
	}
}
func swaaap(arr string) {
	arr2 := []byte(arr)
	for i := 0; i <= len(arr2)-1; i++ {
		arr2[i], arr2[0] = arr2[0], arr2[i]

		fmt.Printf("%c\n21", arr2)
		arr2[0], arr2[i] = arr2[i], arr2[0]
	}
}

//given an index you want all values to show up in getting all 'permutations'
//of THIS particular index, print them
func permutationOfSpecificIndex(index int, array []int) {
	for i := 0; i < len(array); i++ {
		//get the permutations for JUST this specific index
		array[index], array[i] = array[i], array[index]
		fmt.Println(array)
		array[i], array[index] = array[index], array[i]
	}
}

func perm(i, j int) {
	if i > j {
		fmt.Println("I was right")
	}
	if i == j {
		fmt.Println(i)
	}
	for a := i; a <= j; a++ {

		perm(i+1, j)
	}
}
func main() {

	testData := []int{1, 3, 4, 5}

	/*
		for i := range testData {
			dfs(testData, i)

		}
	*/
	for i := 0; i < len(testData); i++ {
		permutationOfSpecificIndex(i, testData)
	}

	//perm(0, 3)
	testData2 := "ABC"
	//swaaap(testData2)
	Permutation([]byte(testData2), 0)

}
