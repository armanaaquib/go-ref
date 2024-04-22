package types

type Reactangle struct {
	Width  float64
	Height float64
}

func (r Reactangle) Area() float64 {
	return r.Width * r.Height
}
