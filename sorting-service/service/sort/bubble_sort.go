package sort

import (
	"fmt"
	"sorting-service/service/collection"
)

var _ collection.Sorter = (*BubbleSort)(nil)

// BubbleSort concrete strategy
type BubbleSort struct {
	v []int
}

func NewBubbleSortStrategy(c *collection.Collection) *BubbleSort {
	return &BubbleSort{
		v: c.GetValue(),
	}
}

func (b BubbleSort) Sort() {
	fmt.Println("Sorting using bubble sort")
	for i := len(b.v) - 1; i > 0; i-- {
		swap := false
		for j := 0; j < i; j++ {
			if b.v[j] > b.v[j+1] {
				b.v[j], b.v[j+1] = b.v[j+1], b.v[j]
				swap = true
			}
		}
		if !swap {
			return
		}
	}
}
