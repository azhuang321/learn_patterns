package example1

type Context struct {
	state State
}

func New() *Context {
	return &Context{}
}

func (c *Context) SetState(state State) {
	c.state = state
}
func (c *Context) GetState() State {
	return c.state
}
