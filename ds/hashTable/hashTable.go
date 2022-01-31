package hashTable

const ArraySize = 7

type NodeVal interface{}

type HashTable struct {
	array [ArraySize]*bucket
}

//linked list
type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key   string
	value NodeVal
	next  *bucketNode
}

func InitHashTable() *HashTable {
	h := &HashTable{}
	for i := range h.array {
		h.array[i] = &bucket{}
	}
	return h
}

func (h *HashTable) Insert(key string, value NodeVal) {
	index := hash(key)
	h.array[index].insert(key, value)
}

func (h *HashTable) Get(key string) NodeVal {
	index := hash(key)
	return h.array[index].get(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// add the key k to the bucket
// set new node as head
func (b *bucket) insert(k string, val NodeVal) {
	if b.get(k) == nil {
		newNode := &bucketNode{key: k, value: val}
		newNode.next = b.head
		b.head = newNode
	}
}

// return if a key exist on certain bucket
func (b *bucket) get(k string) NodeVal {
	currNode := b.head
	for currNode != nil {
		if currNode.key == k {
			return currNode.value
		}
		currNode = currNode.next
	}
	return nil
}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	prevNode := b.head
	if prevNode.next != nil {
		if prevNode.next.key == k {
			prevNode.next = prevNode.next.next
		}
		prevNode = prevNode.next
	}
}

//cast every the character to int and sum
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}
