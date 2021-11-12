package example1

import "fmt"

type OctalObserver struct {
	*Subject
}

func (b *OctalObserver) Update() {
	fmt.Println("octal update")
}

func NewOctal(subject *Subject) *OctalObserver {
	binaryOb := &OctalObserver{}
	binaryOb.Subject = subject
	binaryOb.Subject.Attach(binaryOb)
	return binaryOb
}
