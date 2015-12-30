package sqlparser

import (
	"bytes"
	"fmt"
	"reflect"
)

type Tree struct {
	Nodes [][]*TreeNode
	count int
	depth int
}

type TreeNode struct {
	Id     int
	Parent int
	Name   string
	IsEnd  bool
}

func NewTree() *Tree {
	return &Tree{
		Nodes: make([][]*TreeNode, 0),
	}
}

func (t *Tree) SetTree(node SQLNode) {
	t.setNode(node, 0, 1)
}

func (t *Tree) setNode(node SQLNode, parent int, depth int) {
	var (
		name  string
		nodes = getSQLNodes(node)
	)
	if nodes == nil || len(nodes) < 1 {
		name = t.GetValue(node)
		if name == "" {
			panic("not has name " + reflect.TypeOf(node).String())
			return
		}
		tn := t.getTreeNode(parent, node, name)
		tn.IsEnd = true
		t.AddNode(tn, depth)
		return
	}
	tn := t.getTreeNode(parent, node, "")
	t.AddNode(tn, depth)
	for _, n := range nodes {
		t.setNode(n, tn.Id, depth+1)
	}
}

func getSQLNodes(node SQLNode) []SQLNode {
	sns := make([]SQLNode, 0)
	nodes := node.Nodes()
	for _, n := range nodes {
		if IsNodeHasValue(n) {
			sns = append(sns, n)
		}
	}
	return sns
}

func (t *Tree) getTreeNode(parent int, node SQLNode, name string) *TreeNode {
	t.count++
	if name == "" {
		name = reflect.TypeOf(node).String()
		if name == "" {
			panic("nil name")
		}
	}
	return &TreeNode{
		Id:     t.count,
		Parent: parent,
		Name:   name,
	}
}

func (t *Tree) AddNode(tnode *TreeNode, depth int) {
	if t.depth < depth && t.depth+1 < depth {
		panic("error depths")
	}
	if t.depth < depth {
		t.depth++
		t.Nodes = append(t.Nodes, make([]*TreeNode, 0))
	}
	t.Nodes[depth-1] = append(t.Nodes[depth-1], tnode)
}

var (
	byesT  = reflect.TypeOf(make([]byte, 0))
	byes2T = reflect.TypeOf(make([][]byte, 0))
)

func (t *Tree) GetValue(node interface{}) string {
	typ := reflect.TypeOf(node)
	switch typ.Kind() {
	case reflect.String:
		return fmt.Sprintf("%s", node)
	case reflect.Int:
		return fmt.Sprintf("%d", node)
	case reflect.Bool:
		return fmt.Sprintf("%v", node)
	case reflect.Slice:
		switch {
		case typ.ConvertibleTo(byesT):
			return fmt.Sprintf("%s", node)
		case typ.ConvertibleTo(byes2T):
			buf := bytes.Buffer{}
			buf.WriteByte('{')
			bs := node.([][]byte)
			for n, b := range bs {
				buf.Write(b)
				if len(bs)-1 != n {
					buf.WriteByte(',')
				}
			}
			buf.WriteByte('}')
			return buf.String()
		default:
			return ""
		}

	case reflect.Uint8: // byte
		return fmt.Sprintf("%c", node)
	case reflect.Ptr:
		return typ.String()
	default:
		return ""
	}
}
