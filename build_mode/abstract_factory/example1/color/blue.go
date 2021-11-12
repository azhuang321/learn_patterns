package color

import "fmt"

type Blue struct{}

func (r *Blue) Fill() {
	fmt.Println("implements")
}
