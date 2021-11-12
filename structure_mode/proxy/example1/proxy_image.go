package example1

type ProxyImage struct {
	RealImage *RealImage
	FileName  string
}

func (p *ProxyImage) ProxyImage(filename string) {
	p.FileName = filename
}

func (p *ProxyImage) Display() {
	if p.RealImage == nil {
		p.RealImage = NewImage(p.FileName)
	}
	p.RealImage.Display()
}
