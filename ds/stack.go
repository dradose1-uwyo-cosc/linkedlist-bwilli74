package ds

type Stack struct {
        list LinkedList
}

func (s *Stack) Push(value string) {
        n := &Node{data: value}
        if s.list.Head == nil {
                s.list.Head = n
                s.list.Tail = n
                s.list.Size = 1
                return
        }
	n.Next = s.list.Head
        s.list.Head = n
        s.list.Size++
}

func (s *Stack) Pop() (string, bool) {
        if s.list.Head == nil {
                return "", false
        }
	val := s.list.Head.data
        _ = s.list.RemoveAt(0)
        return val, true
}
