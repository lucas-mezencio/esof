package handlers

import (
	"sorting-service/service/collection"
	"sorting-service/service/sort"
)

func SortingHandlerCMD(values []int) *collection.Collection {

	col := collection.NewContent(values)

	// NOTE: Setting the sorting strategy based on content vector length
	if len(values) <= 10 {
		col.SetSortAlgo(sort.NewBubbleSortStrategy(col))
	} else {
		col.SetSortAlgo(sort.NewQuickSortSrategy(col))
	}
	return col
}
