package collection

// Sorter interface Strategy
type Sorter interface {
	Sort()
}

type Collection struct {
	sortAlgo Sorter
	values   []int
}

func NewContent(values []int) *Collection {
	return &Collection{
		values: values,
	}
}

func (c *Collection) SetSortAlgo(s Sorter) {
	c.sortAlgo = s
}

func (c *Collection) Sort() {
	c.sortAlgo.Sort()
}

func (c *Collection) GetValue() []int {
	return c.values
}
