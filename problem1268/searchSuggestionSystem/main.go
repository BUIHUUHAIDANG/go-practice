package main

import (
	"fmt"
	"sort"
	"strings"
)

type Trie struct {
	Noderoot *Node
}

type Node struct {
	IsEnd    bool
	Children [26]*Node
	Char     string
}

func newNode(char string) *Node {
	return &Node{
		Char: char,
	}
}

func Constructor() Trie {
	return Trie{
		Noderoot: newNode(""),
	}
}

func (this *Trie) Insert(word string) {
	current := this.Noderoot

	word = strings.ToLower(word)

	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'

		if current.Children[index] == nil {
			current.Children[index] = newNode(string(word[i]))
		}

		current = current.Children[index]
	}

	current.IsEnd = true
}

func dfs(node *Node, prefix string, ans *[]string) {
	if node == nil || len(*ans) >= 3 {
		return
	}

	if node.IsEnd {
		*ans = append(*ans, prefix)
	}

	for i := 0; i < 26 && len(*ans) < 3; i++ {
		if node.Children[i] != nil {
			dfs(
				node.Children[i],
				prefix+node.Children[i].Char,
				ans,
			)
		}
	}
}

func (this *Trie) suggestedProducts(products []string, searchWord string) [][]string {

	sort.Strings(products)

	for _, product := range products {
		this.Insert(product)
	}

	result := [][]string{}

	current := this.Noderoot
	prefix := ""

	for i := 0; i < len(searchWord); i++ {

		index := searchWord[i] - 'a'
		prefix += string(searchWord[i])

		if current == nil || current.Children[index] == nil {

			for j := i; j < len(searchWord); j++ {
				result = append(result, []string{})
			}
			break
		}

		current = current.Children[index]

		tmp := []string{}

		dfs(current, prefix, &tmp)

		result = append(result, tmp)
	}

	return result
}

func main() {

	trie := Constructor()

	ans := trie.suggestedProducts(
		[]string{
			"mobile",
			"mouse",
			"moneypot",
			"monitor",
			"mousepad",
		},
		"mouse",
	)

	fmt.Println(ans)
}