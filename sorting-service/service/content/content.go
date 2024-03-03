package content

// Sorter interface Strategy
type Sorter interface {
	Sort()
}

type Content struct {
	sortAlgo Sorter
	values   []int
}

func NewContent(values []int) *Content {
	return &Content{
		values: values,
	}
}

func (c *Content) SetSortAlgo(s Sorter) {
	c.sortAlgo = s
}

func (c *Content) Sort() {
	c.sortAlgo.Sort()
}

func (c *Content) GetValue() []int {
	return c.values
}
