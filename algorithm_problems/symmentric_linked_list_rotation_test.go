package algorithm_problems

import (
	"testing"

	"github.com/azraeljack/algorithm-practice/common"
)

func TestRotateOddList(t *testing.T) {
	list := []string{"A", "B", "C", "D", "E", "F", "G"}
	expectedOutput := []string{"A", "G", "B", "F", "C", "E", "D"}
	root := buildLinkedList(list)
	RotateLinkedList(root)
	result := iterateLinkedList(root)
	if !isStringSliceEqual(expectedOutput, result) {
		t.Errorf("result %v is different with expected %v", result, expectedOutput)
	}
}

func TestRotateEvenList(t *testing.T) {
	list := []string{"A", "B", "C", "D", "E", "F"}
	expectedOutput := []string{"A", "F", "B", "E", "C", "D"}
	root := buildLinkedList(list)
	RotateLinkedList(root)
	result := iterateLinkedList(root)
	if !isStringSliceEqual(expectedOutput, result) {
		t.Errorf("result %v is different with expected %v", result, expectedOutput)
	}
}

func TestRotateNilList(t *testing.T) {
	var list []string
	var expectedOutput []string
	root := buildLinkedList(list)
	RotateLinkedList(root)
	result := iterateLinkedList(root)
	if !isStringSliceEqual(expectedOutput, result) {
		t.Errorf("result %v is different with expected %v", result, expectedOutput)
	}
}

func TestRotateEmptyList(t *testing.T) {
	list := []string{}
	expectedOutput := []string{}
	root := buildLinkedList(list)
	RotateLinkedList(root)
	result := iterateLinkedList(root)
	if !isStringSliceEqual(expectedOutput, result) {
		t.Errorf("result %v is different with expected %v", result, expectedOutput)
	}
}

func TestRotateSingleList(t *testing.T) {
	list := []string{"A"}
	expectedOutput := []string{"A"}
	root := buildLinkedList(list)
	RotateLinkedList(root)
	result := iterateLinkedList(root)
	if !isStringSliceEqual(expectedOutput, result) {
		t.Errorf("result %v is different with expected %v", result, expectedOutput)
	}
}

func TestRotateDoubleList(t *testing.T) {
	list := []string{"A", "B"}
	expectedOutput := []string{"A", "B"}
	root := buildLinkedList(list)
	RotateLinkedList(root)
	result := iterateLinkedList(root)
	if !isStringSliceEqual(expectedOutput, result) {
		t.Errorf("result %v is different with expected %v", result, expectedOutput)
	}
}

func TestRotateTripleList(t *testing.T) {
	list := []string{"A", "B", "C"}
	expectedOutput := []string{"A", "C", "B"}
	root := buildLinkedList(list)
	RotateLinkedList(root)
	result := iterateLinkedList(root)
	if !isStringSliceEqual(expectedOutput, result) {
		t.Errorf("result %v is different with expected %v", result, expectedOutput)
	}
}

func isStringSliceEqual(list1, list2 []string) bool {
	if list1 == nil && list2 == nil {
		return true
	}

	if len(list1) != len(list2) {
		return false
	}

	for index, data := range list1 {
		if data != list2[index] {
			return false
		}
	}

	return true
}

func buildLinkedList(data []string) *common.ForwardDirectionalNode {
	if data == nil {
		return nil
	}
	var root, temp *common.ForwardDirectionalNode
	for index, word := range data {
		if index == 0 {
			temp = &common.ForwardDirectionalNode{Next: nil, Data: word}
			root = temp
		} else {
			temp.Next = &common.ForwardDirectionalNode{Next: nil, Data: word}
			temp = temp.Next
		}
	}
	return root
}

func iterateLinkedList(root *common.ForwardDirectionalNode) []string {
	if root == nil {
		return nil
	}
	temp := root
	result := make([]string, 0)
	for temp != nil {
		result = append(result, temp.Data.(string))
		temp = temp.Next
	}
	return result
}
