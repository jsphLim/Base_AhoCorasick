package ac

//字典树节点
type Trie_Node struct {
	fail      *Trie_Node
	next      map[rune]*Trie_Node
	isPattern bool
}
//ahoCorasick类
type AhoCorasick struct {
	root *Trie_Node
}
//new 一个字典树节点
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

//添加模式串
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
	//建立队列
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

// 扫描 返回查询结果
func (ac *AhoCorasick) ScanAhoCorasick(content string) (results []find) {
	str := []rune(content)
	pointer := ac.root
	var begin, end int  //开始位置与结束位置
	for i, c := range str {
		_, ok := pointer.next[c]
		// 失配状态更改
		for !ok && pointer != ac.root {
			pointer = pointer.fail
		}
		if _, ok = pointer.next[c]; ok {
			if pointer == ac.root { // 匹配第一个成功，记录位置
				begin = i
			}
			pointer = pointer.next[c]
			if pointer.isPattern {
				end = i // 最后一个匹配成功，记录
				results = append(results, find{begin, end})
			}
		}
	}
	return
}