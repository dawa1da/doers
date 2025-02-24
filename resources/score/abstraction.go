package score

type Score struct {
	point int64
	unit  string
}

func NewScore(unit string) *Score {
	return &Score{
		unit: unit,
	}
}

func (s *Score) Add(point int64) {
	s.point += point
}

func (s *Score) Remove(point int64) {
	s.point -= point
}
