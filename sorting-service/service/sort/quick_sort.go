package sort

import (
	"fmt"
	"sorting-service/service/content"
)

var _ content.Sorter = (*QuickSort)(nil)

// QuickSort concrete strategy
type QuickSort struct {
	v      []int
	start  int
	finish int
}

func NewQuickSortSrategy(c *content.Content) *QuickSort {
	return &QuickSort{
		v:      c.GetValue(),
		start:  0,
		finish: len(c.GetValue()) - 1,
	}
}

func (qs QuickSort) Sort() {
	fmt.Println("sorting using quick sort")
	quickSort(qs.v, qs.start, qs.finish)
}

func quickSort(v []int, start int, finish int) {
	if start < finish {
		p := partion(v, start, finish)
		quickSort(v, start, p-1)
		quickSort(v, p+1, finish)
	}
}

func partion(v []int, start int, finish int) int {
	p := v[finish]
	i := start - 1

	for j := start; j < finish; j++ {
		if v[j] < p {
			i++
			v[i], v[j] = v[j], v[i]
		}
	}

	v[i+1], v[finish] = v[finish], v[i+1]
	return i + 1
}
