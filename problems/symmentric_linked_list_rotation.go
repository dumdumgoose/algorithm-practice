package problems

import "github.com/azraeljack/algorithm-practice/common"

/*
The question was asked to rotate a single directional linked list from A->B->C->D->E->F->G to A->G->B->F->C->E->D
O(1) space and O(n) time
*/

func RotateLinkedList(root *common.ForwardDirectionalNode) {
	// check if the linked list is single or empty list
	if root == nil || root.Next == nil {
		return
	}
	// some initializations
	var (
		tailPointer        = root
		centerPointer      = root
		afterCenterPointer = root

		movingPointer       = root
		beforeMovingPointer = movingPointer
		afterMovingPointer  = movingPointer
	)

	// find all the import points in the linked list
	lengthCounter := 0
	for movingPointer.Next != nil {
		beforeMovingPointer = movingPointer
		movingPointer = movingPointer.Next
		if movingPointer.Next == nil {
			tailPointer = movingPointer
		}
		if lengthCounter%2 == 1 {
			centerPointer = centerPointer.Next
		}
		lengthCounter++
	}

	// mark the node right after center point
	afterCenterPointer = centerPointer.Next
	centerPointer.Next = nil

	// reverse the right half of the list
	movingPointer = afterCenterPointer
	beforeMovingPointer = nil
	for movingPointer != nil {
		afterMovingPointer = movingPointer.Next
		movingPointer.Next = beforeMovingPointer
		beforeMovingPointer = movingPointer
		movingPointer = afterMovingPointer
	}

	// link nodes in the requested order
	movingPointer = root
	for movingPointer != nil {
		tempPointer := movingPointer.Next
		movingPointer.Next = tailPointer
		movingPointer = tempPointer

		if movingPointer == nil {
			return
		}

		tempPointer = tailPointer.Next
		tailPointer.Next = movingPointer
		tailPointer = tempPointer
	}
}
