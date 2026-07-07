package main

import (
	"strings"
	"fmt"
)

const NUM_CHAR = 26 

type Trie struct {
	RootNode *Node
}

func newTrie() *Trie{
	root := newNode("\000")
	return &Trie{RootNode: root}
}

type Node struct {
	Char string
	End bool
	Children [NUM_CHAR]*Node 
}

func newNode(char string) *Node {
	newnode := &Node{Char : char}
	for i:=0; i< NUM_CHAR; i++{
		newnode.Children[i] = nil
	} 
	return newnode
}

func (t *Trie) insertTrieNode(word string) error{
	current := t.RootNode
	strippedword := strings.ToLower(strings.ReplaceAll(word," ",""))
	for i:=0; i<len(strippedword);i++{
		index := strippedword[i]-'a'
		if current.Children[index]==nil{
			current.Children[index] = newNode(string(strippedword[i]))
		}
		current = current.Children[index]
	}
	current.End = true
	return nil
}

func (t *Trie) searchWord(word string) bool {
	current := t.RootNode
	strippedword := strings.ToLower(strings.ReplaceAll(word," ",""))
	for i:=0; i<len(strippedword);i++{
		index := strippedword[i]-'a'
		if current == nil ||current.Children[index]==nil{
			return false
		}
		current = current.Children[index]
	}
	return true
}

func printWords(node *Node, word string) {
    if node == nil {
        return
    }

    if node.End {
        fmt.Println(word)
    }

    for _, child := range node.Children {
        if child != nil {
            printWords(child, word+child.Char)
        }
    }
}
func hasChildren(node *Node) bool {
	for _, child := range node.Children {
		if child != nil {
			return true
		}
	}
	return false
}

func (t *Trie)deleteHelper(node *Node, word string, depth int) bool {
	if node == nil {
		return false
	}

	// Đã tới cuối từ
	if depth == len(word) {

		if !node.End {
			return false // từ không tồn tại
		}

		node.End = false

		// báo cho cha biết có thể xóa node này
		return !hasChildren(node)
	}

	index := word[depth] - 'a'

	child := node.Children[index]

	if child == nil {
		return false
	}

	shouldDeleteChild := t.deleteHelper(child, word, depth+1)

	if shouldDeleteChild {
		node.Children[index] = nil

		return !node.End && !hasChildren(node)
	}

	return false
}

func main(){
	trie := newTrie()

    trie.insertTrieNode("cat")
    trie.insertTrieNode("car")
    trie.insertTrieNode("dog")

    printWords(trie.RootNode, "")
	trie.deleteHelper(trie.RootNode,"car",0)
	printWords(trie.RootNode, "")
}