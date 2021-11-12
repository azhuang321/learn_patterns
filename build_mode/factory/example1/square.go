package example1

import "fmt"

type Square struct{}

func (r *Square) Draw() {
	fmt.Println("implements")
}
