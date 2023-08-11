package shape

type rect struct {
	l, b float32
}

func (r *rect) area() float32 {
	return r.l * r.b
}
func areaof(r rect) float32 {
	return r.l * r.b
}
