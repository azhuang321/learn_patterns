package shape

import "fmt"

type Square struct{}

func (r *Square) Draw() {
	fmt.Println("draw square")
}
