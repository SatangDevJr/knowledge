package main

import (
	"fmt"
	"neversitup/src/pkg/utils/findtheoddint"
	"neversitup/src/pkg/utils/permutations"
)

func main() {
	fmt.Println("test")

	//Permutations
	fmt.Println("if input 'a' should return: ['a'], result is : ", permutations.Permutations("a"))
	fmt.Println("if input 'ab' should return: ['ab', 'ba'], result is : ", permutations.Permutations("ab"))
	fmt.Println("if input 'abc' should return: ['abc','acb','bac','bca','cab','cba'], result is : ", permutations.Permutations("abc"))
	fmt.Println("if input 'aabb' should return: ['abc','acb','bac','bca','cab','cba'], result is : ", permutations.Permutations("aabb"))

	//Find The Odd Int
	fmt.Println("[7] should return 7, because it occurs 1 time (which is odd).")
	fmt.Println("result is : ", findtheoddint.FindOddInt([]int{7}))
	fmt.Println("[0] should return 0, because it occurs 1 time (which is odd).")
	fmt.Println("result is : ", findtheoddint.FindOddInt([]int{0}))
	fmt.Println("[1,1,2] should return 2, because it occurs 1 time (which is odd).")
	fmt.Println("result is : ", findtheoddint.FindOddInt([]int{1, 1, 2}))
	fmt.Println("[0,1,0,1,0] should return 0, because it occurs 3 times (which is odd).")
	fmt.Println("result is : ", findtheoddint.FindOddInt([]int{0, 1, 0, 1, 0}))
	fmt.Println("[1,2,2,3,3,3,4,3,3,3,2,2,1] should return 4, because it occurs 1 times (which is odd).")
	fmt.Println("result is : ", findtheoddint.FindOddInt([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}))
}
