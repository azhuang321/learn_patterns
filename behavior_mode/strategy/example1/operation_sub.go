package example1

type OperationSub struct{}

func (o *OperationSub) DoOperation(num1, num2 int) int {
	return num1 - num2
}
