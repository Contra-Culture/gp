package gp

type (
	Node struct {
		tags     []string
		parsed   []rune
		children []*Node
	}
	pos struct {
		idx  int
		node *Node
	}
)

func (n *Node) Tags() []string {
	return n.tags
}
func (n *Node) Parsed() []rune {
	return n.parsed
}
func (n *Node) Children() []*Node {
	return n.children
}
func (root *Node) Traverse(fn func(int, int, *Node) error) (err error) {
	path := []pos{
		{
			idx:  0,
			node: root,
		},
	}
	depth := 0
	for {
		p := path[len(path)-1]
		if p.node == nil { // last child
			depth--
			nextSiblingIdx := path[len(path)-2].idx + 1
			parentsParentPath := path[:len(path)-2]
			parentsParent := parentsParentPath[len(path)-1]
			isLastChild := len(parentsParent.node.children) == nextSiblingIdx
			if isLastChild {
				path = append(parentsParentPath, pos{
					idx:  nextSiblingIdx,
					node: nil,
				})
			} else { // then continue with the next sibling
				path = append(parentsParentPath, pos{
					idx:  nextSiblingIdx,
					node: parentsParent.node.children[nextSiblingIdx],
				})
			}
		}
		err = fn(depth, p.idx, p.node)
		if err != nil {
			return err
		}
		hasChildren := len(p.node.children) > 0
		if hasChildren { // then fall down to the children nodes
			path = append(path, pos{
				idx:  0,
				node: p.node.children[0],
			})
		} else { // lift to the parent
			parentPath := path[:len(path)-1]
			if len(parentPath) == 0 {
				return
			}
			parent := parentPath[len(parentPath)-1]
			nextIdx := p.idx + 1
			isLastChild := len(parent.node.children) == nextIdx
			if isLastChild {
				path = append(parentPath, pos{
					idx:  nextIdx,
					node: nil,
				})
			} else { // then continue with the next sibling
				depth++
				path = append(parentPath, pos{
					idx:  nextIdx,
					node: parent.node.children[nextIdx],
				})
			}
		}
	}
}
