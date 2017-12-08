package aoc2017

import (
	"fmt"
	"strings"
)

type discNode struct {
	Name      string
	Weight    int
	WeightSum int
	Children  []*discNode
}

func (n *discNode) SumWeight() int {
	w := n.Weight
	for _, c := range n.Children {
		w += c.SumWeight()
	}

	n.WeightSum = w
	return w
}

func getRootNode(input string) *discNode {
	var (
		n        string
		w        int
		m        = 9999
		children = map[string][]string{}
		nodes    = []*discNode{}
		node     *discNode
		root     *discNode
	)
	rows := strings.Split(input, "\n")

	for _, row := range rows {
		p := strings.Split(row, " -> ")

		_, err := fmt.Sscanf(p[0], "%s (%d)", &n, &w)
		if err != nil {
			panic(err)
		}

		node = &discNode{
			Name:   n,
			Weight: w,
		}

		// Has children?
		if len(p) > 1 {
			children[n] = strings.Split(p[1], ", ")
		}

		nodes = append(nodes, node)
	}

	for n, c := range children {
	SubChild:
		for _, k := range c {
			for i, node := range nodes {
				if node.Name == k {
					for _, pnode := range nodes {
						if pnode.Name == n {
							pnode.Children = append(pnode.Children, nodes[i])
							continue SubChild
						}
					}
				}
			}
		}
	}

	// Find root node
	for _, node := range nodes {
		isChild := false

		for _, c := range children {
			for _, k := range c {
				if k == node.Name {
					isChild = true
					break
				}
			}
		}

		if !isChild && node.Weight < m {
			m = node.Weight
			root = node
		}
	}

	root.SumWeight()

	return root
}

func getNewWeight(input string) int {
	root := getRootNode(input)

	return findImbalance(root)
}

func findImbalance(node *discNode) int {
	sum := map[int]int{}
	idx := map[int]int{}

	for i := range node.Children {
		sum[node.Children[i].WeightSum]++
		idx[node.Children[i].WeightSum] = i
	}

	for ws, s := range sum {
		if s == 1 {
			// Found the imbalance
			w := findImbalance(node.Children[idx[ws]])

			if w == 0 {
				x := idx[ws]
				if x == 0 {
					x++
				}

				return node.Children[idx[ws]].Weight - (ws - node.Children[idx[x]].WeightSum)
			}

			return w
		}
	}

	return 0
}

func printTree(nodes []*discNode, level int) {
	for _, n := range nodes {
		fmt.Println(strings.Repeat("-", level), n.Name, "-", n.WeightSum, "(", n.Weight, ")")
		if len(n.Children) > 0 {
			printTree(n.Children, level+1)
		}
	}
}
