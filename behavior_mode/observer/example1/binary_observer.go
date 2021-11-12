package example1

import "fmt"

type BinaryObserver struct {
	*Subject
}

func (b *BinaryObserver) Update() {
	fmt.Println("binary update")
}

func NewBinary(subject *Subject) *BinaryObserver {
	binaryOb := &BinaryObserver{}
	binaryOb.Subject = subject
	binaryOb.Subject.Attach(binaryOb)
	return binaryOb
}
