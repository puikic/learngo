package trie

type TrieNode struct {
	Word     rune               // 当前节点存储的字符。byte只能表示英文字符，rune可以表示任意字符
	Children map[rune]*TrieNode // 孩子节点，用一个map存储
	Term     string             // 如果word是一个Term结尾，则记录下该Term
}

// Trie树
type TrieTree struct {
	root *TrieNode
}

// add把words[beginIndex:]插入到Trie树中
func (node *TrieNode) add(words []rune, term string, beginIndex int) {
	if beginIndex == len(words) {
		node.Term = term
		return
	}

	if node.Children == nil {
		node.Children = make(map[rune]*TrieNode)
	}

	word := words[beginIndex]
	if child, ok := node.Children[word]; !ok {
		child = &TrieNode{Word: word}
		node.Children[word] = child
		child.add(words, term, beginIndex+1) // 递归
	} else {
		child.add(words, term, beginIndex+1)
	}
}

// words[0]就是当前节点上存储的字符，按照words的指引顺着树往下走，最终返回words最后一个字符对应的节点
func (node *TrieNode) work(words []rune, index int) *TrieNode {
	if index == len(words)-1 {
		return node
	}
	index += 1
	word := words[index]
	if child, ok := node.Children[word]; ok {
		return child.work(words, index+1)
	} else {
		return nil
	}
	return nil
}

func (node *TrieNode) traverseTerms(terms *[]string) {

}

// 检索满足前缀的所有Term
func (tree *TrieTree) Retrieve(prefix string) []string {
	if tree.root == nil || len(tree.root.Children) == 0 {
		return nil
	}
	words := []rune(prefix)
	firstWord := words[0]
	if child, ok := tree.root.Children[firstWord]; ok {
		end := child.work(words, 0) // 找到prefix的终点--end
		if end == nil {
			return nil
		} else {
			terms := make([]string, 0, 100)
			end.traverseTerms(&terms)
			return terms
		}
	}
}
