package mode

type Wrapper struct{}

func (w *Wrapper) Pack() string {
	return "wrapper"
}
