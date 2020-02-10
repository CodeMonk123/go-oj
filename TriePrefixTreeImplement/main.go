// 实现一个 Trie (前缀树)，包含 insert, search, 和 startsWith 这三个操作。

// 示例:

// Trie trie = new Trie();

// trie.insert("apple");
// trie.search("apple");   // 返回 true
// trie.search("app");     // 返回 false
// trie.startsWith("app"); // 返回 true
// trie.insert("app");
// trie.search("app");     // 返回 true
// 说明:

// 你可以假设所有的输入都是由小写字母 a-z 构成的。
// 保证所有输入均为非空字符串。

package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Left, Right *Node
	content     string
}

type Trie struct {
	root *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{nil}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this.root
	if root == nil {
		node := Node{nil, nil, word}
		this.root = &node
	} else {
		p := root
		for {
			if word < p.content {
				if p.Left == nil {
					node := Node{nil, nil, word}
					p.Left = &node
					break
				} else {
					p = p.Left
				}
			} else if word > p.content {
				if p.Right == nil {
					node := Node{nil, nil, word}
					p.Right = &node
					break
				} else {
					p = p.Right
				}
			} else {
				break
			}
		}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	p := this.root
	for {
		if p == nil {
			return false
		}
		if p.content == word {
			return true
		}
		if word < p.content {
			p = p.Left
			continue
		}
		p = p.Right
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	p := this.root
	for {
		if p == nil {
			return false
		}
		if strings.HasPrefix(p.content, prefix) {
			return true
		}
		if prefix < p.content {
			p = p.Left
			continue
		}
		p = p.Right
	}
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	tree := Constructor()
	tree.Insert("apple")
	fmt.Printf("%v\n", tree.Search("app"))
	fmt.Printf("%v\n", tree.StartsWith("app"))
}
