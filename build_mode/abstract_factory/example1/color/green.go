package color

import "fmt"

type Green struct{}

func (r *Green) Fill() {
	fmt.Println("implements")
}
