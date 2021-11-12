package example1

import "fmt"

type StopState struct{}

func (s *StopState) DoAction(ctx Context) {
	fmt.Println("stop state")
	ctx.SetState(s)
}
