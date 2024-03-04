package main

import (
	"fmt"
	"sorting-service/handlers"
	"sorting-service/service/collection"
)

func main() {

	values := [][]int{
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	}

	firstOperation := handlers.SortingHandlerCMD(values[0])
	sectionOperation := handlers.SortingHandlerCMD(values[1])

	showOperations(firstOperation, sectionOperation)
}

func showOperations(ops ...*collection.Collection) {
	for i, op := range ops {
		fmt.Printf("Operação nº%d:\n", i+1)
		fmt.Println("Valor inicial:")
		fmt.Println(op.GetValue())
		op.Sort()
		fmt.Println(op.GetValue())
		fmt.Printf("--------------------\n\n")
	}
}
