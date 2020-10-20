/**
* @File:          tree.go
* @Aim:
* @Author:        foyon
* @Created Time:  2020-10-15
 */

/*
[
{id:1, name:xx1, pid:0},
{id:2, name:xx2, pid:0},
{id:3, name:xx3, pid:1},
{id:4, name:xx4, pid:3},
{id:5, name:xx5, pid:4},
]
*/
package main

import (
	"fmt"
)

type TreeOpt interface {
	createNode(id int) *Tree
	delNode()
}

type Tree struct {
	Id    int
	Name  string
	Pid   int
	Left  *Tree
	Right *Tree
}

func (t *Tree) createNode(id int) *Tree {
	t.Id = id
	return t
}

func (t *Tree) delNode() {

	return
}

type param struct {
	Id   int
	Name string
	Pid  int
}

func main() {
	nodes := []param{{Id: 1, Name: "a", Pid: 0}}
	buildTree(nodes)
}

func createRootNode() (root *Tree) {
	root = &Tree{}
	return
}

func addNode(root *Tree, p *param) {

	tmpNode := Tree{
		Id:   p.Id,
		Name: p.Name,
		Pid:  p.Pid,
	}

	if p.Pid == 0 {
		root.Child = append(root.Child, tmpNode)
	} else {
		tmpNode := root
		for {
		}

	}

	fmt.Printf("%+v", root)
}

func buildTree(nodes []param) {
	rootNode := createRootNode()
	for _, v := range nodes {
		addNode(rootNode, &v)
	}
}
