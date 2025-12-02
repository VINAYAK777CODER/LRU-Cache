package main

import "fmt"

const SIZE = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Hash map[string]*Node

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head
	return Queue{Head: head, Tail: tail}
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

//check
func (c *Cache) Check(str string) {
    var node *Node
    if val, ok := c.Hash[str]; ok {
        node = c.Remove(val)
    } else {
        node = &Node{Val: str}
    }
    c.Add(node)
    c.Hash[str] = node
}

//Remove
func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove:%s\n", n.Val)
	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left

	c.Queue.Length -= 1
	delete(c.Hash, n.Val)
	return n

}

//Add
func (c *Cache) Add(n *Node) {
	fmt.Printf("add at first :%s\n", n.Val)
	temp := c.Queue.Head.Right
	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = temp
	temp.Left = n
	c.Queue.Length += 1
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}
func (q *Queue) Display() {
    fmt.Printf("%d - [", q.Length)
    for node := q.Head.Right; node != q.Tail; node = node.Right {
        fmt.Printf("{%s}", node.Val)
        if node.Right != q.Tail {
            fmt.Printf("<-->")
        }
    }
    fmt.Println("]")
}


func main() {
	fmt.Println("START CACHE")
	cache := NewCache()
	for _, word := range []string{"cow", "buffalo", "deer","buffalo" ,"tiger",} {
		cache.Check(word)
		cache.Display()
	}
}
