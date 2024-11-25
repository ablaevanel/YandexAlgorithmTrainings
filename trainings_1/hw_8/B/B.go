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

func (t *TreeNode) Insert(val, height int) error {
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
			fmt.Print(height+1, " ")
			return nil
		}
		return t.left.Insert(val, height+1)
	} else {
		if t.right == nil {
			t.right = &TreeNode{
				val: val,
			}
			fmt.Print(height+1, " ")
			return nil
		}
		return t.right.Insert(val, height+1)
	}
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
	fmt.Print(1, " ")
	n, _ = strconv.Atoi(arr[1])
	i := 2
	for n != 0 {
		t.Insert(n, 1)
		n, _ = strconv.Atoi(arr[i])
		i++
	}
}
