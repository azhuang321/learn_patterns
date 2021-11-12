package example1

type Context struct {
	Strategy
}

func New(strategy Strategy) *Context {
	context := &Context{}
	context.Strategy = strategy
	return context
}

func (c *Context) ExecStrategy(num1, num2 int) int {
	return c.Strategy.DoOperation(num1, num2)
}
