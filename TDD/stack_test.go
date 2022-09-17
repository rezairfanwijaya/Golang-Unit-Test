package stack_test

import (
	"testing"

	stack "github.com/rezairfanwijaya/Golang-Unit-Test/TDD"

	"github.com/stretchr/testify/assert"
)

/*
	Requirement
	* A stack is empty on construction
	* A stack has size 0 on constuction
	* After n pushes to an empty stack (n > 0), the stack is non-empty and its size equals n
	* If one pushes x then pops, the value popped is x, and the size is one less than old size
	* If one pushes x then peeks, the value return is x, and the size stay the same
	* If the size is n , then after n pops, the stack is empty and size 0
	* Popping from an empty stact return an error : ErrNoSuchElement
	* Peeking into empty stack an error : ErrNoSuchElement
*/

// TDD adalah Test Driven Development, yang kita lakukan adalah kita membuat test code nya lalu setelah itu kita buat code productionnya

func TestNewSet(t *testing.T) {
	// define object stack
	s := stack.New()
	t.Run("A stack is empty on construction", func(t *testing.T) {

		assert.True(t, s.IsEmpty())
	})

	t.Run(" A stack has size 0 on constuction", func(t *testing.T) {
		assert.Equal(t, 0, s.Size())
	})

}

func TestInsert(t *testing.T) {
	t.Run("After n pushes to an empty stack (n > 0), the stack is non-empty and its size equals n", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(3)

		assert.False(t, s.IsEmpty())
		assert.Equal(t, 3, s.Size())
	})

	t.Run("If one pushes x then pops, the value popped is x, and the size is one less than old size", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(2)
		s.Push(10)
		assert.Equal(t, 3, s.Size()) // for make sure

		val, _ := s.Pop()
		assert.Equal(t, 10, val)
		assert.Equal(t, 2, s.Size())

	})

	t.Run("If one pushes x then peeks, the value return is x, and the size stay the same ", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(90)
		assert.Equal(t, 2, s.Size()) // for make sure

		val, _ := s.Peek()
		assert.Equal(t, 90, val)
		assert.Equal(t, 2, s.Size())

	})

	t.Run("If the size is n , then after n pops, the stack is empty and size 0", func(t *testing.T) {
		s := stack.New()
		s.Push(1)
		s.Push(90)
		assert.Equal(t, 2, s.Size()) // for make sure

		tmp := 0
		for i := 0; i < 2; i++ {
			val, _ := s.Pop()
			tmp = val
		}

		assert.Equal(t, 1, tmp)
		assert.True(t, s.IsEmpty())
		assert.Equal(t, 0, s.Size())

	})
}

func TestErrorHandling(t *testing.T) {
	t.Run("Popping from an empty stact return an error : ErrNoSuchElement", func(t *testing.T) {
		s := stack.New()
		_, err := s.Pop()
		if err == nil {
			t.Fail()
			t.Logf("Expected error is not nill, but got %v", err)
		}

		assert.Equal(t, stack.ErrNoSuchElement, err)
	})

	t.Run("Peeking into empty stack an error : ErrNoSuchElement", func(t *testing.T) {
		s := stack.New()
		_, err := s.Peek()
		if err == nil {
			t.Fail()
			t.Logf("Expected error is not nill, but got %v", err)
		}

		assert.Equal(t, stack.ErrNoSuchElement, err)
	})
}
