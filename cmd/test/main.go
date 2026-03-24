package main

import "fmt"


func main() {
	
	s := [2]struct{x, y int}{{1, 0}, {2, 1}}

	for i, v := range s {
		fmt.Println(i, v)
	}
}