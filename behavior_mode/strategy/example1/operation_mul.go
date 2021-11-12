package example1

type OperationMul struct{}

func (o *OperationMul) DoOperation(num1, num2 int) int {
	return num1 * num2
}
