package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/Dddarknight/autocomplete/trees"
	"sort"
)

const (
	FILE_NAME string = "text.txt"
	WORD_TO_FIND string = "v"
)

func GetWords() []string {
	content, _ := ioutil.ReadFile(FILE_NAME)
	content_str := string(content)
	content_str = strings.Replace(content_str, ".", "", -1)
	content_str = strings.Replace(content_str, ",", "", -1)
	content_str = strings.ToLower(content_str)
	words := strings.Split(content_str, " ")
	sort.Strings(words)
	return words
}

func main() {
	words := GetWords()
	fmt.Println(words)
	tree := new(trees.TreeNode)
	for _, word := range words {
		tree.Insert(word)
	}
	fmt.Println(tree.Search(WORD_TO_FIND, ""))
	fmt.Println(tree.Variants(""))
}