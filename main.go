package main

func main() {
	execTable()
}

func execTable() {
	headerList := []string{"head1", "head2", "head3"}

	bodyList := [][]string{
		[]string{"A", "Good", "500"},
		[]string{"B", "BadMan", "288"},
		[]string{"C", "Ugly", "120"},
		[]string{"D", "Gopher", "800"},
	}
	t := NewTable(headerList, bodyList)
	t.Render()
}

