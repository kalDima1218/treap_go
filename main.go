package main

import (
	"fmt"
	"math/rand"
)

type Item struct {
	k int
}

func newItem(x int) Item{
	var tmp Item
	tmp.k = x
	return tmp
}

type Node struct{
	i Item
	p int
	l, r, parent *Node
}

func newNode(x Item) *Node {
	var tmp Node
	tmp.p = rand.Int()
	tmp.i = x
	return &tmp
}

func (p *Node)next() *Node{
	if p.r != nil{
		p = p.r
		for p.l != nil{
			p = p.l
		}
		return p
	}
	for p.parent != nil &&  p == p.parent.r{
		p = p.parent
	}
	return p.parent
}

func (p *Node)prev() *Node{
	if p.l != nil{
		p = p.l
		for p.r != nil{
			p = p.r
		}
		return p
	}
	for p.parent != nil &&  p == p.parent.l{
		p = p.parent
	}
	return p.parent
}

func _merge(l *Node, r *Node) *Node {
	if l == nil{
		return r
	}
	if r == nil{
		return l
	}
	if l.p > r.p{
		r.parent = l
		l.r = _merge(l.r, r)
		return l
	}else{
		l.parent = r
		r.l = _merge(l, r.l)
		return r
	}
}

func _split(p *Node, x int) (*Node, *Node){
	if p == nil{
		return nil, nil
	}
	if p.i.k <= x{
		var l, r = _split(p.r, x)
		if l != nil{
			l.parent = p
		}
		p.r = l
		return p, r
	}else{
		var l, r = _split(p.l, x)
		if r != nil{
			r.parent = p
		}
		p.l = r
		return l, p
	}
}

func _print(p *Node){
	if p.l != nil{
		_print(p.l)
	}
	fmt.Println(p.i.k)
	if p.r != nil{
		_print(p.r)
	}
}

type Treap struct{
	_root, _begin, _end *Node
}

func (t *Treap)_updBegin(){
	var p = t._root
	for p.l != nil{
		p = p.l
	}
	t._begin = p
}

func (t *Treap)_updEnd(){
	var p = t._root
	for p.r != nil{
		p = p.r
	}
	t._end = p
}

func (t *Treap)begin() *Node{
	return t._begin
}

func (t *Treap)end() *Node{
	return t._end
}

func (t *Treap)count(x Item) int{
	var p = t._root
	for p.i != x{
		if p.r != nil && p.i.k < x.k{
			p = p.r
			continue
		}
		if p.l != nil && x.k < p.i.k{
			p = p.l
			continue
		}
		break
	}
	if p.i == x{
		return 1
	}else{
		return 0
	}
}

func (t *Treap)find(x Item) (*Node, bool){
	var p = t._root
	for p.i != x{
		if p.r != nil && p.i.k < x.k{
			p = p.r
			continue
		}
		if p.l != nil && x.k < p.i.k{
			p = p.l
			continue
		}
		break
	}
	return p, p.i == x
}

func (t *Treap)insert(x Item){
	if t._root != nil && t.count(x) != 0{
		return
	}
	var l, r = _split(t._root, x.k)
	t._root = _merge(l, _merge(newNode(x), r))
	t._updBegin()
	t._updEnd()
}

func (t *Treap)erase(x Item){
	if t._root == nil || t.count(x) == 0{
		return
	}
	var l, r = _split(t._root, x.k)
	l, _ = _split(l, x.k-1)
	t._root = _merge(l, r)
	t._updBegin()
	t._updEnd()
}

func (t *Treap)print(){
	_print(t._root)
}

func main() {
	var t Treap
	for i := 0; i < 10; i++{
		t.insert(newItem(rand.Int()%100))
	}
	t.print()
}
