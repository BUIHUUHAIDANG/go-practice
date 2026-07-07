package main

import "strings"

type Trie struct {
	Noderoot *Node
}

type Node struct {
	IsEnd bool
	Children [26]*Node
	Char string
}

func newNode(char string) *Node{
	newnode := &Node{Char: char}
	for i:=0 ;i<26;i++{
		newnode.Children[i] = nil 
	}
	return newnode
}

func Constructor() Trie {
	root := newNode("\000")
    return Trie{Noderoot: root}
}


func (this *Trie) Insert(word string)  {
    current := this.Noderoot
	trippedword := strings.ToLower(strings.ReplaceAll(word," ",""))
	for i:=0;i<len(trippedword);i++{
		index := trippedword[i]-'a'
		if current.Children[index] == nil{
			current.Children[index] = newNode(string(trippedword[i])) 
		}
		current = current.Children[index]
	}
	current.IsEnd = true
}


func (this *Trie) Search(word string) bool {
	current := this.Noderoot
	trippedword := strings.ToLower(strings.ReplaceAll(word," ",""))
	for i:=0;i<len(trippedword);i++{
		index := trippedword[i]-'a'
		if current == nil||current.Children[index] == nil{
			return false 
		}
		current = current.Children[index]
	}
	if current.IsEnd == true {
		return true
	}
	return false
}


func (this *Trie) StartsWith(prefix string) bool {
	current := this.Noderoot
	trippedword := strings.ToLower(strings.ReplaceAll(prefix," ",""))
	for i:=0;i<len(trippedword);i++{
		index := trippedword[i]-'a'
		if current == nil||current.Children[index] == nil{
			return false 
		}
		current = current.Children[index]
	}
	return true
}