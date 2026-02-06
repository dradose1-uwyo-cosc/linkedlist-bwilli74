package ds

import "fmt"
import "errors"

type LinkedList struct {
        Head *Node
        Tail *Node
        Size int
}

func (l *LinkedList) IsEmpty() bool {
        return l.Size == 0
}

func (l *LinkedList) GetSize() int {
        return l.Size
}

func (l *LinkedList) Insert(value string) {
        n := &Node{data: value}
        if l.Head == nil {
                l.Head = n
                l.Tail = n
                l.Size = 1
                return
        }
	      l.Tail.Next = n
        l.Tail = n
        l.Size++
}

func (l *LinkedList) InsertAt(position int, value string) error {
        if position < 0 || position > l.Size {
                return fmt.Errorf("position %d out of range", position)
        }
	n := &Node{data: value}
        // empty list
        if l.Size == 0 {
                l.Head = n
                l.Tail = n
                l.Size = 1
                return nil
        }
	// insert at head
        if position == 0 {
                n.Next = l.Head
                l.Head = n
                l.Size++
                return nil
        }
	// insert at tail
        if position == l.Size {
                l.Tail.Next = n
                l.Tail = n
                l.Size++
                return nil
        }
	// insert in middle
        prev := l.Head
        for i := 0; i < position-1; i++ {
                prev = prev.Next
       }
  n.Next = prev.Next
        prev.Next = n
        l.Size++
        return nil
}
func (l *LinkedList) RemoveAt(pos int) error {
        if pos < 0 || pos >= l.Size {
                return fmt.Errorf("position %d out of range", pos)
        }
	//if removing head
        if pos == 0 {
                l.Head = l.Head.Next
                l.Size--
                if l.Size == 0 {
                        l.Tail = nil
                }
	}
        prev := l.Head
        for i := 0; i < pos-1; i++ {
                prev = prev.Next
        }

	//unlink
        removed := prev.Next
        prev.Next = removed.Next
        l.Size--
        //new tail
        if prev.Next == nil {
                l.Tail = prev
        }
	return nil // unlink pointers
}

func (l *LinkedList) Remove(value string) error { // remove first occurrence of the value
        if l.Size == 0 {
                return errors.New("list is empty")
        }

	// if head matched
        if l.Head.data == value {
                return l.RemoveAt(0)
        }
	//prev := l.Head
        curr := l.Head.Next
        index := 1
        for curr != nil {
                if curr.data == value {
                        return l.RemoveAt(index)
                }
                //      prev = curr
                curr = curr.Next
                index++
        }

	return fmt.Errorf("value %q not found", value)
}
func (l *LinkedList) RemoveAll(value string) error {
        if l.Size == 0 {
                return errors.New("list is empty - cannot remove")
        }
	removedAny := false
        // Remove matching heads
	for l.Head != nil && l.Head.data == value {
                l.Head = l.Head.Next
                l.Size--
                removedAny = true
        }
        // If list became empty
        if l.Head == nil {
                l.Tail = nil
                if removedAny {
                        return nil
                }
                return errors.New("list is empty")
	}
        // Remove matches after head
        prev := l.Head
        curr := l.Head.Next

        for curr != nil {
                if curr.data == value {
                        prev.Next = curr.Next
                        l.Size--
                        removedAny = true
                        curr = prev.Next
                } else {
                        prev = curr
                        curr = curr.Next
                }
        }
        // Fix tail
        l.Tail = prev

        if !removedAny {
                return fmt.Errorf("value %q not found", value)
	}
        return nil
}

func (l *LinkedList) Reverse() { //reverse the list
        if l.Size <= 1 {
                return
        }
	var prev *Node = nil
        curr := l.Head
        l.Tail = l.Head

        for curr != nil {
                next := curr.Next
                curr.Next = prev
                prev = curr
                curr = next
        }
	l.Head = prev
}

func (l *LinkedList) PrintList() { // PrintList prints the list values in order
        curr := l.Head
        for curr != nil {
                fmt.Println(curr.data)
                curr = curr.Next
        }
}
