package prefixtree

// PrefixTree is a prefix tree datastruct
type PrefixTree struct {
	nodes map[interface{}]*Node
}

// NewPrefixTree new a prefix tree
func NewPrefixTree() *PrefixTree {
	return &PrefixTree{
		nodes: make(map[interface{}]*Node),
	}
}

// Add a node into this prefix tree
func (p *PrefixTree) Add(keys []interface{}, value interface{}) {
	if len(keys) > 0 {
		if p.nodes[keys[0]] == nil {
			p.nodes[keys[0]] = NewNode(keys[0], nil, nil)
		}
		n := p.nodes[keys[0]]
		for _, v := range keys[1:] {
			if temp := n.children[v]; temp == nil {
				nt := NewNode(v, nil, n)
				n.children[v] = nt
				n = nt
			} else {
				n = temp
			}
		}
		n.value = value
	}
}

// Del a node of the prefix tree
func (p *PrefixTree) Del(keys []interface{}) {
	if len(keys) > 0 {
		if len(keys) == 1 {
			delete(p.nodes, keys[0])
		}
		n := p.nodes[keys[0]]
		if n == nil {
			return
		}
		for _, v := range keys[1:] {
			n = n.children[v]
			if n == nil {
				return
			}
		}
		if n.parent != nil {
			delete(n.parent.children, keys[len(keys)-1])
			p.delParen(n.parent)
		}
	}
}

func (p *PrefixTree) delParen(parent *Node) {
	if parent != nil && parent.value == nil && len(parent.children) == 0 {
		if parent.parent == nil {
			delete(p.nodes, parent.key)
			return
		}
		delete(parent.parent.children, parent.key)
		p.delParen(parent.parent)
	}
}

// Get a node's value of prefix tree
func (p *PrefixTree) Get(keys []interface{}) interface{} {
	if len(keys) > 0 {
		n := p.nodes[keys[0]]
		if n == nil {
			return nil
		}
		for _, k := range keys[1:] {
			n = n.children[k]
			if n == nil {
				return nil
			}
		}
		return n.value
	}
	return nil
}

// Update a node of prefix tree.
func (p *PrefixTree) Update(keys []interface{}, value interface{}) {
	if len(keys) > 0 {
		n := p.nodes[keys[0]]
		if n == nil {
			return
		}
		for _, v := range keys[1:] {
			n = n.children[v]
			if n == nil {
				return
			}
		}
		n.value = value
	}
}

// IsEmpty return prefix tree is empty
func (p *PrefixTree) IsEmpty() bool {
	if len(p.nodes) > 0 {
		println(len(p.nodes))
		return false
	}
	return true
}

// Node is a node of a prefix tree
type Node struct {
	parent   *Node
	key      interface{}
	value    interface{}
	children map[interface{}]*Node
}

// NewNode will create a prefix tree node struct
func NewNode(key, value interface{}, parent *Node) *Node {
	return &Node{
		key:      key,
		value:    value,
		parent:   parent,
		children: make(map[interface{}]*Node),
	}
}
