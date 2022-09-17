package stack

import "fmt"

var ErrNoSuchElement error = fmt.Errorf("no such element")

// struct stack
type Stack struct {
	arr []int
}

// function new
func New() *Stack {
	return &Stack{
		arr: make([]int, 0),
	}
}

// function IsEmpty
func (s Stack) IsEmpty() bool {
	return s.Size() == 0
}

// function cek size stack
func (s Stack) Size() int {
	return len(s.arr)
}

// function push untuk memasukan data ke stack
func (s *Stack) Push(number int) error {
	s.arr = append(s.arr, number)
	return nil
}

// function pop untuk mengeluarkan nilai stack terakhir (size arr -1)
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrNoSuchElement
	}
	lastStack := s.lastElement()
	s.arr = s.arr[:s.Size()-1]
	return lastStack, nil
}

// function peek untuk memunculkan nilai terakhir (size arr stay same)
func (s Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrNoSuchElement
	}

	return s.lastElement(), nil
}

// funcion private untuk mendapatkan last element
func (s Stack) lastElement() int {
	return s.arr[s.Size()-1]
}
