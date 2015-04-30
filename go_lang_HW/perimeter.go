
type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) perimeter() float64 {
	return 2 * c.r * math.Pi
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}
