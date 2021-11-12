package example1

type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func NewAnd(expr1, expr2 Expression) *AndExpression {
	te := &AndExpression{}
	te.expr1 = expr1
	te.expr2 = expr2
	return te
}

func (t *AndExpression) Interpret(context string) bool {
	if t.expr1.Interpret(context) && t.expr2.Interpret(context) {
		return true
	}
	return false
}
