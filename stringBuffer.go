package main

// stringBuffer implements a linked list of bytes that will be joined into a
// string.
type stringBuffer struct {
    dummy *sBufferNode
    last *sBufferNode
    Length int
}

func newStringBuffer() *stringBuffer {
    dummy := &sBufferNode{ nil, nil }
    return &stringBuffer{ dummy, dummy, 0 }
}

// Adds a new node to the stringBuffer. O(1).
func (s *stringBuffer) Add(str string) {
    node := &sBufferNode{ []byte(str), nil }
    s.last.next = node
    s.last = node
    s.Length += len(str)
}

// Copies all of the node values in a stringBuffer to a byte array and returns
// the resulting string. O(n).
func (s *stringBuffer) String() string {
    final := make([]byte, s.Length)
    cursor := 0
    currentNode := s.dummy

    for currentNode.next != nil {
        currentNode = currentNode.next
        copy(final[cursor:], currentNode.value)
        cursor += len(currentNode.value)
    }

    return string(final)
}

type sBufferNode struct {
    value []byte
    next *sBufferNode
}
