package main

// CircularStringArray is used as a log for process outputs. Only a certain number
// of lines will be stored, which prevents talkative programs (like `yes`)
// from eating memory.
type CircularStringArray struct {
    Capacity int // Max lines that the array can hold
    Length int // Number of lines the array actually holds
    front, rear int
    arr []string

    // front is inclusive and gives the index of the first element.
    // rear is exclusive and gives the index which will be filled next.
}

// NewCircularArray returns a new circular array with capacity cap.
func NewCircularStringArray(cap int) *CircularStringArray {
    c := CircularStringArray { Capacity: cap, Length: 0, front: 0, rear: 0 }
    c.arr = make([]string, cap)
    return &c
}

// At returns the element at index i from the front element of the array. Index
// 0 gives the first element. Indices greater than the array's length "wrap
// around."
func (ca *CircularStringArray) At(i int) string {
    return ca.arr[(ca.front + i) % ca.Capacity]
}

// Insert puts an object at the next element in the array. If the array is full,
// the oldest element is overwritten.
func (ca *CircularStringArray) Insert(e string) {
    ca.arr[ca.rear] = e
    ca.rear = (ca.rear + 1) % ca.Capacity

    if ca.front == ca.rear {
        ca.front = (ca.front + 1) % ca.Capacity
    }

    if (ca.Length < ca.Capacity) {
        ca.Length++
    }
}

// Do calls function f on each element in the array.
func (ca *CircularStringArray) Do(f func(int, string)) {
    for i := range(ca.arr) {
        element := ca.At(i)

        // In case the element hasn't been filled yet.
        if (element != "") {
            f(i, element)
        }
    }
}
