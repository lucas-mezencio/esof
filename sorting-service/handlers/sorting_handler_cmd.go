package handlers

import (
	"sorting-service/service/content"
	"sorting-service/service/sort"
)

func SortingHandlerCMD(input [][]int) []*content.Content {
	var contents []*content.Content
	for _, values := range input {
		content := content.NewContent(values)

		// NOTE: Setting the sorting strategy based on content vector length
		if len(values) <= 10 {
			content.SetSortAlgo(sort.NewBubbleSortStrategy(content))
		} else {
			content.SetSortAlgo(sort.NewQuickSortSrategy(content))
		}

		contents = append(contents, content)
	}
	return contents
}
