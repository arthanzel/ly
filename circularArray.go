package main

// CircularArray is used as a log for process outputs. Only a certain number
// of lines will be stored, which prevents talkative programs (like `yes`)
// from eating memory.
type CircularArray struct {
    Capacity int
    Length int
    front, rear int
    arr []interface{}

    // rear is exclusive and gives the index which will be filled next.
}

// NewCircularArray returns a new circular array with length len. The default
// values for elements will be nil.
func NewCircularArray(cap int) *CircularArray {
    c := CircularArray { Capacity: cap, Length: 0, front: 0, rear: 0 }
    c.arr = make([]interface{}, cap)
    return &c
}

// At returns the element at index i from the front element of the array. Index
// 0 gives the first element. Indices greater than the array's length "wrap
// around."
func (ca *CircularArray) At(i int) interface{} {
    return ca.arr[(ca.front + i) % ca.Capacity]
}

// Insert puts an object at the next element in the array. If the array is full,
// the oldest element is overwritten.
func (ca *CircularArray) Insert(e interface{}) {
    ca.arr[ca.rear] = e
    ca.rear = (ca.rear + 1) % ca.Capacity

    if ca.front == ca.rear {
        ca.front = (ca.front + 1) % ca.Capacity
    }

    ca.Length++
    if (ca.Length > ca.Capacity) {
        ca.Length = ca.Capacity
    }
}

// Do calls function f on each element in the array.
func (ca *CircularArray) Do(f func(int, interface{})) {
    for i := range(ca.arr) {
        element := ca.At(i)

        // In case the element hasn't been filled yet.
        if (element != nil) {
            f(i, element)
        }
    }
}
