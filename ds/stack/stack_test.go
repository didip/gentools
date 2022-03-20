package stack

import (
	"testing"

	"github.com/didip/gentools/tester"
)

type testStruct struct {
	Name string
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		tester.AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(123)
		tester.AssertFalse(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(456)
		value := myStackOfInts.Pop()
		tester.AssertEqual(t, value, 456)

		value = myStackOfInts.Pop()
		tester.AssertEqual(t, value, 123)
		tester.AssertTrue(t, myStackOfInts.IsEmpty())
	})

	t.Run("struct stack", func(t *testing.T) {
		myStackOfInts := new(Stack[testStruct])

		tester.AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(testStruct{Name: "freddy"})
		tester.AssertFalse(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(testStruct{Name: "mercury"})
		value := myStackOfInts.Pop()
		tester.AssertEqual(t, value.Name, "mercury")

		value = myStackOfInts.Pop()
		tester.AssertEqual(t, value.Name, "freddy")
		tester.AssertTrue(t, myStackOfInts.IsEmpty())
	})
}
