package main

import "fmt"

func main() {
	arr := []int{25, 42, 46, 52, 64, 72, 85, 96}
	searchElement := 52
	found := false

	for i := 0; i < len(arr); i++ {

		if arr[i] == searchElement {
			fmt.Println("Element Found at:", "index", i)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Element Not Found")
	}

}
