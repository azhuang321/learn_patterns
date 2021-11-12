package example1

import "fmt"

type StartState struct{}

func (s *StartState) DoAction(ctx Context) {
	fmt.Println("start state")
	ctx.SetState(s)
}
