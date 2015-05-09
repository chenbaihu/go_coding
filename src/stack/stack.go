package stack

type Stack struct {
	index int
	data  []int
}

func (s *Stack) Push(val int) bool {
	s.data = append(s.data, val)
	s.index++
	return true
}

func (s *Stack) Pop() (ret int, flag bool) {
	if s.index <= 0 {
		ret = 0
		flag = false
		return
	}
	s.index--
	ret = s.data[s.index]
	flag = true
	return
}
