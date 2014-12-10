package math

type Vertex2 struct {
	X int
	Y int
}

func (v Vertex2) Add(v2 Vertex2) Vertex2 {
	return Vertex2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}
