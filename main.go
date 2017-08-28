package main

import (
	"flag"
	"fmt"
)

var (
	intOpt  = flag.Int("int", 1234, "help message for i option")
	boolOpt = flag.Bool("bool", false, "help message for b option")
	strOpt  = flag.String("str", "default", "help message for s option")
)

func main() {
	//exec()
	execTable()

}

func execTable() {
	t := NewTable()
	t.SetHeader([]string{"head1", "head2", "head3"})
	data := [][]string{
		[]string{"A", "Good", "500"},
		[]string{"B", "BadMan", "288"},
		[]string{"C", "Ugly", "120"},
		[]string{"D", "Gopher", "800"},
	}
	t.SetBody(data)
	t.Render()
}

func exec() {
	//fmt.Println("hello")
	flag.Parse()

	fmt.Println(*intOpt)
	fmt.Println(*boolOpt)
	fmt.Println(*strOpt)
}
