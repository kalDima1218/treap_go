package main

import "fmt"
import "math/rand"

type node struct{
	k, p int
	l, r *node
}

func newNode(x int) *node{
	var tmp = new(node)
	tmp.p = rand.Int()
	tmp.k = x
	return tmp
}

func merge(l *node, r *node) *node{
	if l == nil{
		return r
	}
	if r == nil{
		return l
	}
	if l.p > r.p{
		l.r = merge(l.r, r)
		return l
	}else{
		r.l = merge(l, r.l)
		return r
	}
}

func split(p *node, x int) (*node, *node){
	if p == nil{
		return nil, nil
	}
	if p.k <= x{
		var l, r = split(p.r, x)
		p.r = l
		return p, r
	}else{
		var l, r = split(p.l, x)
		p.l = r
		return l, p
	}
}

var root *node

func find(p *node, x int) bool{
	if p.k == x{
		return true
	}
	if p.r != nil && p.k < x{
		return find(p.r, x)
	}
	if p.l != nil && p.k > x{
		return find(p.l, x)
	}
	return false
}

func insert(x int){
	if root != nil && find(root, x) == true{
		return
	}
	var l, r = split(root, x)
	var t = newNode(x)
	root = merge(l, merge(t, r))
}

func read(p *node){
	if p.l != nil{
		read(p.l)
	}
	fmt.Println(p.k)
	if p.r != nil{
		read(p.r)
	}
}

func main() {
	for i := 0; i < 10; i++{
		insert(rand.Int()%100)
	}
	read(root)
}
