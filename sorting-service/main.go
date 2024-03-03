package main

import (
	"fmt"
	"sorting-service/handlers"
)

func main() {

	values := [][]int{
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	}

	contents := handlers.SortingHandlerCMD(values)
	for _, content := range contents {
		content.Sort()
		fmt.Println(content.GetValue())
	}
}
