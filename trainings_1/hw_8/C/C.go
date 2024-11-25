package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) Insert(val int) error {
	if t == nil {
		return errors.New("tree not exists")
	}
	if t.val == val {
		return nil
	} else if t.val > val {
		if t.left == nil {
			t.left = &TreeNode{
				val: val,
			}
			return nil
		}
		return t.left.Insert(val)
	} else {
		if t.right == nil {
			t.right = &TreeNode{
				val: val,
			}
			return nil
		}
		return t.right.Insert(val)
	}
}

func (t *TreeNode) DescTraversal(a []int, cnt int) int {
	if t == nil {
		return cnt
	}
	cnt = t.right.DescTraversal(a, cnt)
	a[cnt] = t.val
	cnt++
	cnt = t.left.DescTraversal(a, cnt)
	return cnt
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 1024*1024*10), 1024*1024*10)
	scanner.Scan()
	arr := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	n, _ := strconv.Atoi(arr[0])
	t := &TreeNode{
		val: n,
	}
	n, _ = strconv.Atoi(arr[1])
	i := 2
	for n != 0 {
		t.Insert(n)
		n, _ = strconv.Atoi(arr[i])
		i++
	}
	a := make([]int, len(arr)-1)
	cnt := 0
	t.DescTraversal(a, cnt)
	fmt.Print(a[1])
}
