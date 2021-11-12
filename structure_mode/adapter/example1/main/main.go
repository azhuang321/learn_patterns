package main

import "go-patterns/structure_mode/adapter"

func main() {
	audioPlayer := new(adapter.AudioPlayer)
	audioPlayer.Play("mp3", "beyond the horizon.mp3")
	audioPlayer.Play("mp4", "alone.mp4")
	audioPlayer.Play("vlc", "far far away.vlc")
	audioPlayer.Play("avi", "mind me.avi")
}
