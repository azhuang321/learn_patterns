package example1

import "fmt"

type VlcPlayer struct{}

func (v *VlcPlayer) PlayVlc(filename string) {
	fmt.Println("playing vlc type")
}

func (v *VlcPlayer) PlayMp4(filename string) {
	panic("dont support play type")
}
