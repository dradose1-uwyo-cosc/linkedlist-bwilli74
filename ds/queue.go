package ds

import "fmt"

type Queue struct {
        list LinkedList
}

func (q *Queue) Push(value string) { //add tail node
        q.list.Insert(value)
}

func (q *Queue) Pop() (string, error) { //remove the head
        if q.list.Head == nil {
                return "", fmt.Errorf("queue is empty")
        }
	val := q.list.Head.data
        if err := q.list.RemoveAt(0); err != nil {
                return "", err
        }
	return val, nil
}
