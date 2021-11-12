package example1

import "fmt"

type RealImage struct {
	FileName string
}

func NewImage(filename string) *RealImage {
	ri := &RealImage{FileName: filename}
	ri.loadFromDisk(filename)
	return ri
}

func (r *RealImage) Display() {
	fmt.Printf("display %s \n", r.FileName)
}

func (r *RealImage) loadFromDisk(filename string) {
	fmt.Printf("load %s from disk \n", filename)
}
