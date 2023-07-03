package main

type stack []string

func (s *stack) Push(str string) {
	*s = append(*s, str)
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) pop(str string) (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]
	return element, true
}

func (s *stack) getTop() string {
	return (*s)[len(*s)-1]
}