package main

import "fmt"

// global variables
var hashTableSize int = 1000

// limitation : hash function only takes string as the key
func hashFunction(key string) int {
	var sumOfAscii int = 0
	for _, v := range key {
		sumOfAscii += int(v)
	}
	return sumOfAscii % hashTableSize
}

type Pair struct {
	Key   string
	Value string
}

// HashTable Object
type HashTable struct {
	Table [][]Pair
}

func (h *HashTable) Initialize() {
	h.Table = make([][]Pair, hashTableSize)

	fmt.Println("Created an hash table of size : ", len(h.Table))
}

func (h *HashTable) Add(key string, value string) string {
	// store the value in the index
	index := hashFunction(key)
	currPair := Pair{key, value}
	h.Table[index] = append(h.Table[index], currPair)
	return value
}

func (h *HashTable) Get(key string) string {
	// get the value stored there
	index := hashFunction(key)

	for _, pair := range h.Table[index] {
		if pair.Key == key {
			return pair.Value
		}
	}

	return "Key Not Found"
}

// currently creating a hash table with these limitations
// hash_function - takes only stirng as the key
func main() {
	fmt.Println("Implementing Hash Table in Golang ...\n")

	// now create a list of resizable array of the same type - currently string
	var hm HashTable

	hm.Initialize()

	fmt.Println(len(hm.Table))

	fmt.Println("output : ", hm.Get("Gotha"))

	fmt.Println("\nEnd of implementation...")
}
