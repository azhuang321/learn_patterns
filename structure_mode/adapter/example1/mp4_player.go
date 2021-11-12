package example1

import "fmt"

type Mp4Player struct{}

func (v *Mp4Player) PlayVlc(filename string) {
	panic("dont support play type")
}

func (v *Mp4Player) PlayMp4(filename string) {
	fmt.Println("playing mp4 type")
}
