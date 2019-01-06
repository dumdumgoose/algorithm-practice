package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	root   *TreeNode
	values []int
	next   int
}

func Constructor(root *TreeNode) BSTIterator {
	i := BSTIterator{root: root, next: -1}
	i.inOrderTraversal(root)
	return i
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	this.next++
	return this.values[this.next]
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.next < len(this.values)-1
}

func (this *BSTIterator) inOrderTraversal(n *TreeNode) {
	if n == nil {
		return
	}

	this.inOrderTraversal(n.Left)
	this.values = append(this.values, n.Val)
	this.inOrderTraversal(n.Right)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
