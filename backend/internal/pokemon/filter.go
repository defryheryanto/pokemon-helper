package pokemon

type GetFilter struct {
	ID   int
	Name string
}

type ListFilter struct {
	Limit  int
	Offset int
}
