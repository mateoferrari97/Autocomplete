package prefix_tree

import (
	"bufio"
	"io"
	"unicode/utf8"
)

const (
	_root          = -1
	_maxCharacters = 256
)

type Node struct {
	letter   rune
	children []*Node
}

func Load(r io.Reader) *Node {
	root := New()
	if r != nil {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			word := scanner.Text()
			root.addWord(word)
		}
	}

	return root
}

func New() *Node {
	return &Node{
		letter:   _root,
		children: make([]*Node, _maxCharacters),
	}
}

func (n *Node) AddWords(words ...string) {
	for _, word := range words {
		n.addWord(word)
	}
}

func (n *Node) CheckIfWordExists(word string) bool {
	if word == "" {
		return true
	}

	firstLetter, size := utf8.DecodeRuneInString(word)
	if n.children[firstLetter] == nil {
		return false
	}

	return n.children[firstLetter].CheckIfWordExists(word[size:])
}

func (n *Node) GetWordsByPrefix(prefix string) []string {
	if prefix == "" {
		return n.getWords()
	}

	start := n.traverseToPrefix(prefix)
	if start == nil {
		return nil
	}

	lengthPrefix := len(prefix)
	words := start.getWords()
	for i := range words {
		words[i] = prefix[:lengthPrefix-1] + words[i]
	}

	return words
}

func (n *Node) getWords() []string {
	var words []string
	for _, children := range n.children {
		if children == nil {
			continue
		}

		subWords := children.getWords()
		if n.inRoot() {
			words = append(words, subWords...)
		} else {
			letter := string(n.letter)
			for i := range subWords {
				words = append(words, letter+subWords[i])
			}
		}
	}

	if words == nil && !n.inRoot() {
		words = append(words, string(n.letter))
	}

	return words
}

func (n *Node) traverseToPrefix(prefix string) *Node {
	if prefix == "" {
		return n
	}

	firstLetter, size := utf8.DecodeRuneInString(prefix)
	for _, children := range n.children {
		if children == nil || firstLetter != children.letter {
			continue
		}

		return children.traverseToPrefix(prefix[size:])
	}

	return nil
}

func (n *Node) inRoot() bool {
	return n.letter == _root
}

func (n *Node) addWord(word string) {
	if word == "" {
		return
	}

	firstLetter, size := utf8.DecodeRuneInString(word)
	if n.children[firstLetter] == nil {
		n.children[firstLetter] = newNode(firstLetter)
	}

	n.children[firstLetter].addWord(word[size:])
}

func newNode(letter rune) *Node {
	return &Node{
		letter:   letter,
		children: make([]*Node, _maxCharacters),
	}
}
