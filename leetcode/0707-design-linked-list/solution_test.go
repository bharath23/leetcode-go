package leetcode0707

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolutionV0(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	have := obj.Get(1)
	assert.Equalf(t, 2, have, "%s: get at index mismatch", "singly linked list")
	obj.DeleteAtIndex(1)
	have = obj.Get(1)
	assert.Equalf(t, 3, have, "%s: get at index mismatch", "singly linked list")
}

func TestSolutionV1(t *testing.T) {
	obj := Constructor()
	obj.AddAtIndex(0, 10)
	obj.AddAtIndex(0, 20)
	obj.AddAtIndex(1, 30)
	have := obj.Get(0)
	assert.Equalf(t, 20, have, "%s: get at index mismatch", "singly linked list")
}

func TestSolutionV2(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(10)
	obj.DeleteAtIndex(0)
	have := obj.Get(0)
	assert.Equalf(t, -1, have, "%s: get at index mismatch", "singly linked list")
}

func TestSolutionV3(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(7)
	obj.AddAtHead(2)
	obj.AddAtHead(1)
	obj.AddAtIndex(3, 0)
	obj.DeleteAtIndex(2)
	obj.AddAtHead(6)
	obj.AddAtTail(4)
	obj.Get(4)
	obj.AddAtHead(4)
	obj.AddAtIndex(5, 0)
	obj.AddAtHead(6)
}
