package main

// CircularArray is used as a log for process outputs. Only a certain number
// of lines will be stored, which prevents talkative programs (like `yes`)
// from eating memory.
type CircularArray struct {
    Length int
    front, rear int
    arr []interface{}

    // rear is exclusive and gives the index which will be filled next.
}

// NewCircularArray returns a new circular array with length len. The default
// values for elements will be nil.
func NewCircularArray(len int) *CircularArray {
    c := CircularArray { Length: len }
    c.arr = make([]interface{}, len)
    c.front = 0
    c.rear = 0
    return &c
}

// At returns the element at index i from the front element of the array. Index
// 0 gives the first element. Indices greater than the array's length "wrap
// around."
func (ca *CircularArray) At(i int) interface{} {
    return ca.arr[(ca.front + i) % ca.Length]
}

// Insert puts an object at the next element in the array. If the array is full,
// the oldest element is overwritten.
func (ca *CircularArray) Insert(e interface{}) {
    ca.arr[ca.rear] = e
    ca.rear = (ca.rear + 1) % ca.Length

    if ca.front == ca.rear {
        ca.front = (ca.front + 1) % ca.Length
    }
}

// Do calls function f on each element in the array.
func (ca *CircularArray) Do(f func(interface{})) {
    for i := range(ca.arr) {
        element := ca.At(i)

        // In case the element hasn't been filled yet.
        if (element != nil) {
            f(element)
        }
    }
}
