package example1

type Expression interface {
	Interpret(context string) bool
}
