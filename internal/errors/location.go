package errors

type Location struct {
	file     string
	line     int
	function string
}

func (loc *Location) File() string {
	return loc.file
}
func (loc *Location) Line() int {
	return loc.line
}
func (loc *Location) Function() string {
	return loc.function
}
