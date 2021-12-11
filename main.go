package main

import (
	"fmt"
	"unicode/utf8"
)

const (
	_root          = -1
	_maxCharacters = 256
)

var names = []string{"Elize",
	"Jed",
	"Rhianne",
	"Daanyaal",
	"Ammara",
	"Shyla",
	"Alice",
	"Raja",
	"Aaminah",
	"Hester",
	"Rares",
	"Kyra",
	"Aisling",
	"Avani",
	"Izaan",
	"Hajra",
	"Mia",
	"Shay",
	"Dottie",
	"Ryan",
	"Luther",
	"Atticus",
	"Milli",
	"Brandon",
	"Keon",
	"Dru",
	"Om",
	"Donnell",
	"Arjan",
	"Mila-Rose",
	"Frederic",
	"Brooke",
	"Agnes",
	"Philip",
	"Kylan",
	"Shantelle",
	"Hailie",
	"Simra",
	"Alicja",
	"Inez",
	"Farah",
	"Asma",
	"Eleni",
	"Qasim",
	"Mira",
	"Vera",
	"Amara",
	"Naeem",
	"Ritchie",
	"Laurel",
	"Rylee",
	"Maddy",
	"Hayley",
	"Amman",
	"Angelo",
	"Freya",
	"Mcauley",
	"Ayse",
	"Zaynab",
	"Carole",
	"Maleeha",
	"Carla",
	"Omar",
	"Kenzo",
	"Leticia",
	"Abigale",
	"Sama",
	"Khadija",
	"Stefania",
	"Zidan",
	"Esmay",
	"Lois",
	"Elsie-May",
	"Zac",
	"Suleman",
	"Aleena",
	"Keyaan",
	"Garry",
	"Joseff",
	"Iona",
	"Aditi",
	"Amalie",
	"Neel",
	"Astrid",
	"Mikolaj",
	"Milana",
	"Elmer",
	"Weronika",
	"Ewen",
	"Catrina",
	"Haya",
	"Tyrell",
	"Chloe",
	"Baran",
	"Camden",
	"Bella-Rose",
	"Johnny",
	"Arandeep",
	"Aniya",
	"Darcey"}

type Node struct {
	letter   rune
	children []*Node
}

func NewNode(letter rune) *Node {
	return &Node{
		letter:   letter,
		children: make([]*Node, _maxCharacters),
	}
}

func (n *Node) AddWord(word string) {
	if word == "" {
		return
	}

	firstLetter, size := utf8.DecodeRuneInString(word)
	if n.children[firstLetter] == nil {
		n.children[firstLetter] = NewNode(firstLetter)
	}

	n.children[firstLetter].AddWord(word[size:])
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

func (n *Node) GetWords() []string {
	var words []string
	for _, children := range n.children {
		if children == nil {
			continue
		}

		subWords := children.GetWords()
		if n.inRoot() {
			words = append(words, subWords...)
		} else {
			letter := string(n.letter)
			for i := range subWords {
				words = append(words, letter+subWords[i])
			}
		}
	}

	if words == nil {
		words = append(words, string(n.letter))
	}

	return words
}

func (n *Node) inRoot() bool {
	return n.letter == _root
}

func main() {
	root := NewNode(_root)

	for _, name := range names {
		root.AddWord(name)
	}

	fmt.Println(root.GetWords())
}
