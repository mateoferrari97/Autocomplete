package main

import (
	"Autocomplete/prefix_tree"
	"fmt"
	"os"
)

func main() {
	// AddingWordsManually()
	// LoadingFile()
}

func LoadingFile() {
	f, err := os.Open("./example/names.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = f.Close()
	}()

	tree := prefix_tree.Load(f)

	words := tree.GetWordsByPrefix("Ar")
	for _, word := range words {
		fmt.Println(word)
	}
}

func AddingWordsManually() {
	tree := prefix_tree.New()
	tree.AddWords("matias", "fernando", "nicolas", "mateo ferrari coronel", "pepe argento", "pepino")

	words := tree.GetWordsByPrefix("pep")
	for _, word := range words {
		fmt.Println(word)
	}
}
