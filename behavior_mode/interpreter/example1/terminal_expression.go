package example1

import "strings"

type TerminalExpression struct {
	data string
}

func NewTe(data string) *TerminalExpression {
	te := &TerminalExpression{}
	te.data = data
	return te
}

func (t *TerminalExpression) Interpret(context string) bool {
	if strings.Contains(context, t.data) {
		return true
	}
	return false
}
