package example1

type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func NewOr(expr1, expr2 Expression) *OrExpression {
	te := &OrExpression{}
	te.expr1 = expr1
	te.expr2 = expr2
	return te
}

func (t *OrExpression) Interpret(context string) bool {
	if t.expr1.Interpret(context) || t.expr2.Interpret(context) {
		return true
	}
	return false
}
