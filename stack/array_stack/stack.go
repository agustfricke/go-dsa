package arraystack

type Stack struct {
	items [8]int
	front int
	rear  int
	size  int
}

func NewStack() *Stack {
	return &Stack{
		rear: -1,
		size: 0,
	}
}

// insert
func (s *Stack) Push(item int) {
	s.size++
	if s.size == 0 {
		s.front = item
		s.rear = s.rear + 1
	}
}

// delete
func (s *Stack) GetSize() int {
  return s.size
}
