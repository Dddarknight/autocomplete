package main

import (
	"io/ioutil"
	"strings"
	"sort"
	"regexp"
	"github.com/Dddarknight/autocomplete/backend/cmd/trees"
)

const (
	FILE_NAME string = "/Users/uliaegorycheva/autocomplete/backend/text.txt"
)

func GetWords() []string {
	content, _ := ioutil.ReadFile(FILE_NAME)
	content_str := string(content)
	m := regexp.MustCompile(`[\.,]`)
	content_str = m.ReplaceAllString(content_str, "")
	content_str = strings.ToLower(content_str)
	words := strings.Split(content_str, " ")
	sort.Strings(words)
	return words
}

func MakeWordsTree() *trees.TreeNode {
	words := GetWords()
	tree := new(trees.TreeNode)
	for _, word := range words {
		tree.Insert(word)
	}
	return tree
}