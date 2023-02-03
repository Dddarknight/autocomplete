package trees

type Leaf struct {
	value rune
	isEnd bool
}

type TreeNode struct {
	leaf Leaf
	children []*TreeNode
}

func Contains(children []*TreeNode, str rune) bool {
	for _, child := range children {
		if child.leaf.value == str {
			return true
		}
	}
	return false
}

func GetNode(children []*TreeNode, str rune) *TreeNode {
	for _, child := range children {
		if child.leaf.value == str {
			return child
		}
	}
	t := new(TreeNode)
	return t
}

func (tree *TreeNode) Insert(word string) {
    for i, char := range word {
		if Contains(tree.children, char) {
			child := GetNode(tree.children, char)
			if i + 1 == len(word) {
				child.leaf.isEnd = true
			}
			tree = child
		} else {
			new_child := new(TreeNode)
			if i + 1 == len(word) {
				new_child.leaf.value = char
				new_child.leaf.isEnd = true
			} else {
				new_child.leaf.value = char
				new_child.leaf.isEnd = false
			}
			tree.children = append(tree.children, new_child)
			tree = tree.children[len(tree.children) - 1]
		}
	} 
}

func (tree *TreeNode) Variants(chars string) []string {
	var result []string
	var newChars = chars + string(tree.leaf.value)
	if tree.leaf.isEnd && len(tree.children) != 0 {
		result = append(result, newChars)
	}
	if len(tree.children) == 0 {
		result = append(result, newChars)
		return result
	}
	for _, child := range tree.children {
		for _, variant := range child.Variants(newChars) {
		    result = append(result, variant)
		}
	}
	return result
}

func (tree *TreeNode) Search(word string, beginning string) []string {
	for i := 0; i < len(tree.children); i++ {		
		if string(tree.children[i].leaf.value) == string(word[0]) && len(word) == 1 {
			return tree.children[i].Variants(beginning)
		}
		if string(tree.children[i].leaf.value) == string(word[0]) {
			beginning = beginning + string(word[0])
			return tree.children[i].Search(word[1:], beginning)
		}
    }
	var result []string
	return result
}