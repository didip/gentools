package queue

import (
	"testing"

	"github.com/didip/gentools/tester"
)

type testStruct struct {
	Name string
}

func TestQueue(t *testing.T) {
	t.Run("integer Queue", func(t *testing.T) {
		myQueueOfInts := new(Queue[int])

		tester.AssertTrue(t, myQueueOfInts.IsEmpty())

		myQueueOfInts.Enqueue(123)
		tester.AssertFalse(t, myQueueOfInts.IsEmpty())

		myQueueOfInts.Enqueue(456)
		value := myQueueOfInts.Dequeue()
		tester.AssertEqual(t, value, 123)

		value = myQueueOfInts.Dequeue()
		tester.AssertEqual(t, value, 456)
		tester.AssertTrue(t, myQueueOfInts.IsEmpty())
	})

	t.Run("struct Queue", func(t *testing.T) {
		myQueueOfInts := new(Queue[testStruct])

		tester.AssertTrue(t, myQueueOfInts.IsEmpty())

		myQueueOfInts.Enqueue(testStruct{Name: "freddy"})
		tester.AssertFalse(t, myQueueOfInts.IsEmpty())

		myQueueOfInts.Enqueue(testStruct{Name: "mercury"})
		value := myQueueOfInts.Dequeue()
		tester.AssertEqual(t, value.Name, "freddy")

		value = myQueueOfInts.Dequeue()
		tester.AssertEqual(t, value.Name, "mercury")
		tester.AssertTrue(t, myQueueOfInts.IsEmpty())
	})
}
