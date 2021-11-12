package example1

type MediaAdapter struct {
	advanceMusicPlayer AdvancedMediaPlayer
}

func (a *MediaAdapter) MediaAdapter(audioType string) {
	switch audioType {
	case "vlc":
		a.advanceMusicPlayer = new(VlcPlayer)
	case "mp4":
		a.advanceMusicPlayer = new(Mp4Player)
	default:
		panic("not support audio type")
	}
}

func (a *MediaAdapter) Play(audioType, filename string) {
	switch audioType {
	case "vlc":
		a.advanceMusicPlayer.PlayVlc(filename)
	case "mp4":
		a.advanceMusicPlayer.PlayMp4(filename)
	default:
		panic("not support audio type")
	}
}
