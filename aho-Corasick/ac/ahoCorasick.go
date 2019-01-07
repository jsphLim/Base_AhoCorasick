package ac

//node of Trie Tree
type Trie_Node struct {
	fail      *Trie_Node
	next      map[rune]*Trie_Node
	isPattern bool
}
//ahoCorasick 
type AhoCorasick struct {
	root *Trie_Node
}
//new a node
func newNode() *Trie_Node {
	return &Trie_Node{
		fail:      nil,
		next:      map[rune]*Trie_Node{},
		isPattern: false,
	}
}

func CreateAhoCorasick() *AhoCorasick {
	return &AhoCorasick{
		root: newNode(),
	}
}

type find struct {
	Begin,End int
}

//add the pattern string
func (aho *AhoCorasick) Add(pattern string) {
	str := []rune(pattern)
	pointer := aho.root
	for _, c := range str {
		if _, flag := pointer.next[c]; !flag {
			pointer.next[c] = newNode()
		}
		pointer = pointer.next[c]
	}
	pointer.isPattern = true
}

func (aho *AhoCorasick) BuildAhoCorasick() {
	//set up a queue
	queue := []*Trie_Node{}
	queue = append(queue, aho.root)
	for len(queue) != 0 {
		parent := queue[0]
		queue = append(queue[:0], queue[1:]...)
		for char, child := range parent.next {
			if parent == aho.root {
				child.fail = aho.root
			} else {
				if _, flag := parent.fail.next[char]; flag {
					child.fail = parent.fail.next[char]
				} else {
					child.fail = parent.fail
				}
			}
			queue = append(queue, child)
		}
	}
}

// scan and return the results
func (ac *AhoCorasick) ScanAhoCorasick(content string) (results []find) {
	str := []rune(content)
	pointer := ac.root
	var begin, end int  //the begin and end location
	for i, c := range str {
		_, ok := pointer.next[c]
		//match fail
		for !ok && pointer != ac.root {
			pointer = pointer.fail
		}
		if _, ok = pointer.next[c]; ok {
			if pointer == ac.root { // match the first word and return the loacation
				begin = i
			}
			pointer = pointer.next[c]
			if pointer.isPattern {
				end = i // match the last word and return the location
				results = append(results, find{begin, end})
			}
		}
	}
	return
}
