package main

type CircularArray struct {
    Length int
    front, rear int
    arr []interface{}
}

func NewCircularArray(len int) *CircularArray {
    c := CircularArray { Length: len }
    c.arr = make([]interface{}, len)
    c.front = 0
    c.rear = 0
    return &c
}

func (ca *CircularArray) Insert(e interface{}) {
    ca.arr[ca.rear] = e
    ca.rear = (ca.rear + 1) % ca.Length

    if ca.front == ca.rear {
        ca.front = (ca.front + 1) % ca.Length
    }
}

func (ca *CircularArray) Do(f func(interface{})) {
    for i := 0; i < ca.Length; i++ {
        el := ca.arr[(ca.front + i) % ca.Length]
        if (el != nil) {
            f(el)
        }
    }
}
