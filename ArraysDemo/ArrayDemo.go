// This program will give overview of Arrays.
// How to create blank array
// Multidimensional Array
package main

import "fmt"

func main() {
	//creating an array of n int
	a := [2]int{4, 3}

	//If you want the compiler to derive on the size of the array
	_ = [...]int{1, 2, 3, 4, 5, 56, 3}

	//If you want to create an array with default size
	zeros := [8]int{}
	fmt.Println("Zeros Array:", zeros)

	//Pointers with nil values initialized
	ptrs := [8]*int{} // a list of int pointers, filled with 8 nil references ( <nil> )
	fmt.Println("Pointer initialized array:", ptrs)

	emptystr := [8]string{} // a list of string filled with 8 times ""
	fmt.Println("Empty String array:", emptystr)

	fmt.Println(a)

	var arr1 = []int{1, 3, 4, 5}
	var arr2 = []int{1, 2, 4, 5, 6, 7}
	arr3 := findCommonBetweenTwoArrays(arr1, arr2)
	fmt.Println(arr3)

	//Working with personalized type of Array
	type Data struct {
		Number int
		Text   string
	}

	structs := [8]Data{}
	fmt.Println("Personalized type:", structs) // All the 8 elements will be like {0, ""} (Number = 0, Text = "")

	//MUTIDIMENSIONAL ARRAY
	multiDimArray := [2][3]int{{1, 2, 3}, {4, 5, 6}} // Here [2][3]int represents a 2 item array where each item is an array of 3 elements.
	fmt.Println("Multi Dimensional Array:", multiDimArray)
	//NOTE: There are no limits for the number of dimensions, so you can have e.g. [3][4][8][2]string array.

}

// func to find common elements between two arrays
func findCommonBetweenTwoArrays(arr1 []int, arr2 []int) []int {
	i := 0
	j := 0
	var arr3 []int
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			arr3 = append(arr3, arr1[i])
			i++
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}
	return arr3
}
