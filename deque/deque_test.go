package deque

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	list := NewDeque[int]()
	expected := make([]int, 10)
	for i := 0; i < 10; i++ {
		list.Add(i)
		expected[i] = i
	}
	storedValues := list.ListValues()
	if !reflect.DeepEqual(storedValues, expected) {
		t.Errorf("\nExpected %v\nGot %v", expected, storedValues)
	}
}

func TestAddFront(t *testing.T) {
	expected := []string{"three", "two", "one"}
	inputs := []string{"one", "two", "three"}
	list := NewDeque[string]()
	for _, item := range inputs {
		list.AddFront(item)
	}
	assert.Equal(t, len(expected), int(list.Length))
	storedValues := list.ListValues()
	if !reflect.DeepEqual(storedValues, expected) {
		t.Errorf("\nExpected %v\nGot %v", expected, storedValues)
	}
}

func TestMixedAdd(t *testing.T) {
	list := NewDeque[int]()
	list.Add(9)
	list.AddFront(8)
	list.Add(10)
	list.Add(12)
	list.AddFront(20)
	expected := []int{20, 8, 9, 10, 12}
	storedValues := list.ListValues()
	assert.Equal(t, len(expected), list.Length)
	if !reflect.DeepEqual(storedValues, expected) {
		t.Errorf("\nExpected %v\nGot %v", expected, storedValues)
	}
}

func TestPop(t *testing.T) {
	list := NewDeque[int]()
	for i := 0; i < 10; i++ {
		list.Add(i)
	}
	v, _ := list.Pop()
	assert.Equal(t, 9, v)
	v, _ = list.PopFirst()
	assert.Equal(t, 0, v)
	v, _ = list.Pop()
	assert.Equal(t, 8, v)
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	assert.Equal(t, len(expected), list.Length)
	storedValues := list.ListValues()
	if !reflect.DeepEqual(storedValues, expected) {
		t.Errorf("\nExpected %v\nGot %v", expected, storedValues)
	}
}
