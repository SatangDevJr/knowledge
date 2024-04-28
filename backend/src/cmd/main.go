package main

import (
	"fmt"
	"neversitup/src/pkg/utils/permutations"
)

func main() {
	fmt.Println("test")

	//Permutations
	fmt.Println("if input 'a' should return: ['a'], result is : ", permutations.Permutations("a"))
	fmt.Println("if input 'ab' should return: ['ab', 'ba'], result is : ", permutations.Permutations("ab"))
	fmt.Println("if input 'abc' should return: ['abc','acb','bac','bca','cab','cba'], result is : ", permutations.Permutations("abc"))
	fmt.Println("if input 'aabb' should return: ['abc','acb','bac','bca','cab','cba'], result is : ", permutations.Permutations("aabb"))
}
