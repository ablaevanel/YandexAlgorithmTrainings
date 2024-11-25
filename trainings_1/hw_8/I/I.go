package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TreeNode struct {
	val      string
	children []*TreeNode
}

func (t *TreeNode) Find(val string) *TreeNode {
	if t == nil {
		return nil
	}
	if t.val == val {
		return t
	} else {
		var parent *TreeNode
		for _, child := range t.children {
			tn := child.Find(val)
			if tn != nil {
				parent = tn
				break
			}
		}
		return parent
	}
}

func (t *TreeNode) HowMany(ans []string, i int) (int, int) {
	cnt := len(t.children)
	var add int
	for _, child := range t.children {
		add, i = child.HowMany(ans, i)
		ans[i] = child.val + " " + strconv.Itoa(add)
		i++
		cnt += add
	}
	return cnt, i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 1024*1024*10), 1024*1024*10)
	scanner.Scan()
	n, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	scanner.Scan()
	node := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	t := []*TreeNode{&TreeNode{
		val: node[1],
		children: []*TreeNode{
			&TreeNode{val: node[0]},
		},
	}}
	m := make(map[string]*TreeNode)
	m[node[0]] = t[0].children[0]
	m[node[1]] = t[0]
	for i := 1; i < n-1; i++ {
		scanner.Scan()
		node = strings.Split(strings.TrimSpace(scanner.Text()), " ")
		parent, ok := m[node[1]]
		if !ok {
			for _, n := range t {
				ref := n.Find(node[1])
				if ref != nil {
					parent = ref
					break
				}
			}
			if parent == nil {
				t = append(t, &TreeNode{val: node[1]})
				parent = t[len(t)-1]
				m[node[1]] = parent
			}
		}
		child, ok := m[node[0]]
		if ok {
			parent.children = append(parent.children, child)
			for i := 0; i < len(t); i++ {
				if t[i].val == node[0] {
					t = append(t[:i], t[i+1:]...)
				}
			}
		} else {
			parent.children = append(parent.children, &TreeNode{val: node[0]})
		}
	}
	ans := make([]string, n)
	i, cnt := 0, 0
	for _, n := range t {
		cnt, i = n.HowMany(ans, i)
		ans[i] = n.val + " " + strconv.Itoa(cnt)
	}
	sort.Strings(ans)
	for _, str := range ans {
		fmt.Println(str)
	}
}
