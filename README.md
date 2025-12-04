# LRU Cache in Go (Doubly Linked List + Hash Map)

This repository contains a simple implementation of an **LRU (Least Recently Used) Cache** in Go, built from scratch using:

- A **doubly linked list** (manual `Node` + `Queue` structs)
- A **hash map** (`map[string]*Node`) for O(1) lookup

It’s a learning-focused implementation — no external libraries, just core Go and data structure logic.

---

## 🔍 What Is LRU Cache?

An **LRU Cache** is a fixed-size cache that removes the **least recently used** item when it runs out of capacity.

- When you **access** an item:
  - If it exists → move it to the **front** (most recently used).
  - If it doesn’t exist → insert it at the front.
- When the cache is **full** and a new item comes in → remove the item at the **back** (least recently used).

In this implementation, we’re working with **strings** as values.

---

## 🧱 Data Structures Used

### `Node`

```go
type Node struct {
    Val   string
    Left  *Node
    Right *Node
}
Each node represents a cache entry and is part of a doubly linked list.

Queue (Doubly Linked List)
 
type Queue struct {
    Head   *Node
    Tail   *Node
    Length int
}
Uses sentinel Head and Tail nodes (empty nodes that mark the boundaries).

Most recently used (MRU) node is always right after Head.

Least recently used (LRU) node is always right before Tail.

Initialization:

 
func NewQueue() Queue {
    head := &Node{}
    tail := &Node{}
    head.Right = tail
    tail.Left = head
    return Queue{Head: head, Tail: tail}
}
Hash
 
type Hash map[string]*Node
A Go map that stores string values mapped to their corresponding nodes in the linked list.
This allows O(1) access during cache hits.

Cache
 
type Cache struct {
    Queue Queue
    Hash  Hash
}
Combines the Queue (for ordering) and Hash (for fast lookup).

Has a fixed capacity controlled by:

 
const SIZE = 5
⚙️ Core Operations
1. Creating a Cache
 
func NewCache() Cache {
    return Cache{Queue: NewQueue(), Hash: Hash{}}
}
2. Check(str string)
This is the main API in this example.
It simulates accessing or inserting a value into the cache.

 
func (c *Cache) Check(str string) {
    var node *Node
    if val, ok := c.Hash[str]; ok {
        // Cache hit: remove existing node from its current position
        node = c.Remove(val)
    } else {
        // Cache miss: create a new node
        node = &Node{Val: str}
    }
    // Insert (or re-insert) at the front (most recently used)
    c.Add(node)
    c.Hash[str] = node
}
Hit: node exists → remove it from current position → add to front.

Miss: create new node → add to front.

After Add, the node is always stored back into Hash.

3. Remove(n *Node)
Removes a node from the linked list and from the hash map.

 
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
Updates neighbors to skip over n.

Decrements Length.

Deletes the value from the hash.

Returns the removed node (so it can be re-used if needed).

4. Add(n *Node)
Inserts a node at the front (right after Head), marking it as the most recently used.

 
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
Always inserts at the front.

Increments Length.

If size exceeds SIZE, removes the tail’s left node, which is the LRU entry.

5. Display Functions
Cache.Display() delegates to Queue.Display():

 
func (c *Cache) Display() {
    c.Queue.Display()
}
Queue.Display() prints the current state of the cache from most recently used to least recently used:

 
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
Example output format:

text
  
3 - [{buffalo}<-->{deer}<-->{cow}]
▶️ How to Run
Save the code into a file, e.g.:

bash
  
lru.go
Run using Go:

bash
  
go run lru.go
🧪 Example Flow (from main)
 
func main() {
    fmt.Println("START CACHE")
    cache := NewCache()
    for _, word := range []string{"cow", "buffalo", "deer", "buffalo", "tiger"} {
        cache.Check(word)
        cache.Display()
    }
}
Inserts: cow, buffalo, deer

Accesses buffalo again → moves it to the front as most recently used

Inserts tiger

Since SIZE = 5, no eviction happens yet.

Each Check call prints:

Which values are being added or removed.

The full current state of the cache.

📈 Time & Space Complexity
Time Complexity

Check: O(1) average (hash lookup + linked list ops)

Add: O(1)

Remove: O(1)

Space Complexity

O(N) where N = SIZE (max number of nodes in cache + hash entries)

🚀 Possible Extensions
You can improve/extend this implementation by:

Adding separate key and value (e.g. Key string, Val string)

Turning it into a generic LRU cache using Go generics (Cache[K comparable, V any])

Exposing proper methods like:

Get(key string) (string, bool)

Put(key, val string)

Making it concurrency-safe with mutexes (sync.Mutex / sync.RWMutex)

Writing unit tests with Go’s testing package

📚 Purpose
This implementation is mainly for learning:

How LRU caches work internally

How to combine a hash map with a doubly linked list

Practicing pointers and structs in Go

Feel free to fork, modify, and play around with the code to deepen your understanding of data structures and Go.
