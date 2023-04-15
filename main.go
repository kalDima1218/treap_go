package main

import (
	"fmt"
	"math/rand"
)

type Item struct {
	k int // k is the key of the item
}

// newItem creates a new item with the given key
func newItem(x int) Item {
	var tmp Item
	tmp.k = x
	return tmp
}

type Node struct {
	i            Item // i is the item stored in the node
	p            int  // p is the priority of the node
	l, r, parent *Node
}

// newNode creates a new node with the given item
func newNode(x Item) *Node {
	var tmp Node
	tmp.p = rand.Int()
	tmp.i = x
	return &tmp
}

// next returns the next node in the treap
func (p *Node) next() *Node {
	if p.r != nil {
		p = p.r
		for p.l != nil {
			p = p.l
		}
		return p
	}
	for p.parent != nil && p == p.parent.r {
		p = p.parent
	}
	return p.parent
}

// prev returns the previous node in the treap
func (p *Node) prev() *Node {
	if p.l != nil {
		p = p.l
		for p.r != nil {
			p = p.r
		}
		return p
	}
	for p.parent != nil && p == p.parent.l {
		p = p.parent
	}
	return p.parent
}

// _merge merges two nodes in the treap
func _merge(l *Node, r *Node) *Node {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	if l.p > r.p {
		r.parent = l
		l.r = _merge(l.r, r)
		return l
	} else {
		l.parent = r
		r.l = _merge(l, r.l)
		return r
	}
}

// _split splits the treap at the given key
func _split(p *Node, x int) (*Node, *Node) {
	if p == nil {
		return nil, nil
	}
	if p.i.k <= x {
		var l, r = _split(p.r, x)
		if l != nil {
			l.parent = p
		}
		p.r = l
		return p, r
	} else {
		var l, r = _split(p.l, x)
		if r != nil {
			r.parent = p
		}
		p.l = r
		return l, p
	}
}

// _print prints the treap in order
func _print(p *Node) {
	if p.l != nil {
		_print(p.l)
	}
	fmt.Println(p.i.k)
	if p.r != nil {
		_print(p.r)
	}
}

type Treap struct {
	_root, _begin, _end *Node // _root is the root node of the treap, _begin is the node with the smallest key, _end is the node with the largest key
}

// _updBegin updates the _begin node of the treap
func (t *Treap) _updBegin() {
	var p = t._root
	for p.l != nil {
		p = p.l
	}
	t._begin = p
}

// _updEnd updates the _end node of the treap
func (t *Treap) _updEnd() {
	var p = t._root
	for p.r != nil {
		p = p.r
	}
	t._end = p
}

// begin returns the node with the smallest key in the treap
func (t *Treap) begin() *Node {
	return t._begin
}

// end returns the node with the largest key in the treap
func (t *Treap) end() *Node {
	return t._end
}

// count returns the number of occurrences of the given item in the treap
func (t *Treap) count(x Item) int {
	var p = t._root
	for p.i != x {
		if p.r != nil && p.i.k < x.k {
			p = p.r
			continue
		}
		if p.l != nil && x.k < p.i.k {
			p = p.l
			continue
		}
		break
	}
	if p.i == x {
		return 1
	} else {
		return 0
	}
}

// find returns the node containing the given item and a boolean indicating whether the item was found
func (t *Treap) find(x Item) (*Node, bool) {
	var p = t._root
	for p.i != x {
		if p.r != nil && p.i.k < x.k {
			p = p.r
			continue
		}
		if p.l != nil && x.k < p.i.k {
			p = p.l
			continue
		}
		break
	}
	return p, p.i == x
}

// insert inserts the given item into the treap
func (t *Treap) insert(x Item) {
	if t._root != nil && t.count(x) != 0 {
		return
	}
	var l, r = _split(t._root, x.k)
	t._root = _merge(l, _merge(newNode(x), r))
	t._updBegin()
	t._updEnd()
}

// erase removes the given item from the treap
func (t *Treap) erase(x Item) {
	if t._root == nil || t.count(x) == 0 {
		return
	}
	var l, r = _split(t._root, x.k)
	l, _ = _split(l, x.k-1)
	t._root = _merge(l, r)
	t._updBegin()
	t._updEnd()
}

// print prints the treap in order
func (t *Treap) print() {
	_print(t._root)
}

func main() {
	var t Treap
	for i := 0; i < 10; i++ {
		t.insert(newItem(rand.Int() % 100))
	}
	t.print()
}
