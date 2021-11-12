package example1

import "fmt"

type HexaObserver struct {
	*Subject
}

func (b *HexaObserver) Update() {
	fmt.Println("hexa update")
}

func NewHexa(subject *Subject) *HexaObserver {
	binaryOb := &HexaObserver{}
	binaryOb.Subject = subject
	binaryOb.Subject.Attach(binaryOb)
	return binaryOb
}
