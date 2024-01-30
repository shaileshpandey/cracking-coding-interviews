package interfaces

type Square struct {
	Length int
}

func (s *Square) Area() int64 {
	return int64(s.Length) * int64(s.Length)
}
