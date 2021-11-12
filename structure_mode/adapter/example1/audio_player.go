package example1

import "fmt"

type AudioPlayer struct {
	//MediaAdapter MediaAdapter
}

func (a *AudioPlayer) Play(audioType, filename string) {
	switch audioType {
	case "mp3":
		fmt.Println("play mp3")
	case "vlc":
		fallthrough
	case "mp4":
		mediaAdapter := &MediaAdapter{}
		mediaAdapter.MediaAdapter(audioType)
		mediaAdapter.Play(audioType, filename)
	default:
		panic("audio type not support")
	}
}
