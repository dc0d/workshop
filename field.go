package workshop

type Field struct{ X, Y int }

func (f Field) IsValid() bool {
	if f.X > 2 || f.X < 0 || f.Y > 2 || f.Y < 0 {
		return false
	}
	return true
}
