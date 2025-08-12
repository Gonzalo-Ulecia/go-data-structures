package linkedlist

import (
	"fmt"
	"testing"
)

func assertEqual[T comparable](t *testing.T, expected, actual T, msg string) {
	if expected != actual {
		t.Errorf("%s: expected %v, got %v", msg, expected, actual)
	}
}

func TestAppend(t *testing.T) {
	linkedList := &LinkedList[int]{}

	// Test Append to empty list
	linkedList.Append(1)
	if linkedList.size != 1 {
		t.Errorf("Expected size 1, got %d", linkedList.size)
	}
	if linkedList.f_node.value != 1 {
		t.Errorf("Expected first node value 1, got %d", linkedList.f_node.value)
	}

	// Test Append to non-empty list
	linkedList.Append(2)
	if linkedList.size != 2 {
		t.Errorf("Expected size 2, got %d", linkedList.size)
	}
	if linkedList.f_node.n_node.value != 2 {
		t.Errorf("Expected second node value 2, got %d", linkedList.f_node.n_node.value)
	}
}

func TestGet(t *testing.T) {
	linkedList := &LinkedList[int]{}
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	// Test valid Get
	if val := linkedList.Get(1); val != 2 {
		t.Errorf("Expected value 2 at index 1, got %d", val)
	}

	// Test invalid Get (index out of bounds)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for index out of bounds")
		}
	}()
	linkedList.Get(5) // Should panic
}

func TestPrepend(t *testing.T) {
	linkedList := &LinkedList[int]{}
	linkedList.Append(1)
	linkedList.Prepend(0)

	// Test Prepend
	if linkedList.size != 2 {
		t.Errorf("Expected size 2 after prepend, got %d", linkedList.size)
	}
	if linkedList.f_node.value != 0 {
		t.Errorf("Expected first node value 0, got %d", linkedList.f_node.value)
	}
}

func TestInsert(t *testing.T) {
	linkedList := &LinkedList[int]{}
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	// Test Insert at middle
	linkedList.Insert(1, 5)
	if linkedList.size != 4 {
		t.Errorf("Expected size 4 after insert, got %d", linkedList.size)
	}
	if linkedList.Get(1) != 5 {
		t.Errorf("Expected value 5 at index 1, got %d", linkedList.Get(1))
	}

	// Test Insert at beginning
	linkedList.Insert(0, 10)
	if linkedList.Get(0) != 10 {
		t.Errorf("Expected value 10 at index 0, got %d", linkedList.Get(0))
	}

	// Test Insert at end (index == size)
	linkedList.Insert(4, 15)
	if linkedList.Get(4) != 15 {
		t.Errorf("Expected value 15 at index 4, got %d", linkedList.Get(4))
	}
}

func TestContains(t *testing.T) {
	linkedList := &LinkedList[int]{}
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	// Test Contains (item exists)
	if !linkedList.Contains(2) {
		t.Errorf("Expected Contains to return true for value 2")
	}

	// Test Contains (item does not exist)
	if linkedList.Contains(5) {
		t.Errorf("Expected Contains to return false for value 5")
	}
}

func TestClone(t *testing.T) {
	linkedList := &LinkedList[int]{}
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)

	clone := linkedList.Clone()

	// Test size of clone
	if clone.size != linkedList.size {
		t.Errorf("Expected clone size %d, got %d", linkedList.size, clone.size)
	}

	// Test values in clone
	for i := 0; i < linkedList.size; i++ {
		if linkedList.Get(i) != clone.Get(i) {
			t.Errorf("Expected value %d in clone at index %d", linkedList.Get(i), i)
		}
	}
}

type Person struct {
	Name string
	Age  int
}

func TestAppendWithStruct(t *testing.T) {
	linkedList := &LinkedList[Person]{}

	// Test Append to empty list
	linkedList.Append(Person{Name: "Alice", Age: 30})
	if linkedList.size != 1 {
		t.Errorf("Expected size 1, got %d", linkedList.size)
	}
	if linkedList.f_node.value.Name != "Alice" || linkedList.f_node.value.Age != 30 {
		t.Errorf("Expected first node to be {Alice, 30}, got {%s, %d}", linkedList.f_node.value.Name, linkedList.f_node.value.Age)
	}

	// Test Append to non-empty list
	linkedList.Append(Person{Name: "Bob", Age: 25})
	if linkedList.size != 2 {
		t.Errorf("Expected size 2, got %d", linkedList.size)
	}
	if linkedList.f_node.n_node.value.Name != "Bob" || linkedList.f_node.n_node.value.Age != 25 {
		t.Errorf("Expected second node to be {Bob, 25}, got {%s, %d}", linkedList.f_node.n_node.value.Name, linkedList.f_node.n_node.value.Age)
	}
}

func TestGetWithStruct(t *testing.T) {
	linkedList := &LinkedList[Person]{}
	linkedList.Append(Person{Name: "Alice", Age: 30})
	linkedList.Append(Person{Name: "Bob", Age: 25})
	linkedList.Append(Person{Name: "Charlie", Age: 35})

	// Test valid Get
	if person := linkedList.Get(1); person.Name != "Bob" || person.Age != 25 {
		t.Errorf("Expected {Bob, 25} at index 1, got {%s, %d}", person.Name, person.Age)
	}

	// Test invalid Get (index out of bounds)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for index out of bounds")
		}
	}()
	linkedList.Get(5) // Should panic
}

func TestPrependWithStruct(t *testing.T) {
	linkedList := &LinkedList[Person]{}
	linkedList.Append(Person{Name: "Alice", Age: 30})
	linkedList.Prepend(Person{Name: "Bob", Age: 25})

	// Test Prepend
	if linkedList.size != 2 {
		t.Errorf("Expected size 2 after prepend, got %d", linkedList.size)
	}
	if linkedList.f_node.value.Name != "Bob" || linkedList.f_node.value.Age != 25 {
		t.Errorf("Expected first node value {Bob, 25}, got {%s, %d}", linkedList.f_node.value.Name, linkedList.f_node.value.Age)
	}
}

func TestInsertWithStruct(t *testing.T) {
	linkedList := &LinkedList[Person]{}
	linkedList.Append(Person{Name: "Alice", Age: 30})
	linkedList.Append(Person{Name: "Bob", Age: 25})
	linkedList.Append(Person{Name: "Charlie", Age: 35})

	// Test Insert at middle
	linkedList.Insert(1, Person{Name: "Eve", Age: 28})
	if linkedList.size != 4 {
		t.Errorf("Expected size 4 after insert, got %d", linkedList.size)
	}
	if linkedList.Get(1).Name != "Eve" || linkedList.Get(1).Age != 28 {
		t.Errorf("Expected {Eve, 28} at index 1, got {%s, %d}", linkedList.Get(1).Name, linkedList.Get(1).Age)
	}

	// Test Insert at beginning
	linkedList.Insert(0, Person{Name: "David", Age: 40})
	if linkedList.Get(0).Name != "David" || linkedList.Get(0).Age != 40 {
		t.Errorf("Expected {David, 40} at index 0, got {%s, %d}", linkedList.Get(0).Name, linkedList.Get(0).Age)
	}

	// Test Insert at end (index == size)
	linkedList.Insert(5, Person{Name: "Frank", Age: 50})
	if linkedList.Get(5).Name != "Frank" || linkedList.Get(5).Age != 50 {
		t.Errorf("Expected {Frank, 50} at index 5, got {%s, %d}", linkedList.Get(5).Name, linkedList.Get(5).Age)
	}
}

func TestContainsWithStruct(t *testing.T) {
	linkedList := &LinkedList[Person]{}
	linkedList.Append(Person{Name: "Alice", Age: 30})
	linkedList.Append(Person{Name: "Bob", Age: 25})
	linkedList.Append(Person{Name: "Charlie", Age: 35})

	// Test Contains (item exists)
	if !linkedList.Contains(Person{Name: "Bob", Age: 25}) {
		t.Errorf("Expected Contains to return true for {Bob, 25}")
	}

	// Test Contains (item does not exist)
	if linkedList.Contains(Person{Name: "David", Age: 40}) {
		t.Errorf("Expected Contains to return false for {David, 40}")
	}
}

func TestCloneWithStruct(t *testing.T) {
	linkedList := &LinkedList[Person]{}
	linkedList.Append(Person{Name: "Alice", Age: 30})
	linkedList.Append(Person{Name: "Bob", Age: 25})
	linkedList.Append(Person{Name: "Charlie", Age: 35})

	clone := linkedList.Clone()

	// Test size of clone
	if clone.size != linkedList.size {
		t.Errorf("Expected clone size %d, got %d", linkedList.size, clone.size)
	}

	// Test values in clone
	for i := 0; i < linkedList.size; i++ {
		originalPerson := linkedList.Get(i)
		clonePerson := clone.Get(i)
		if originalPerson.Name != clonePerson.Name || originalPerson.Age != clonePerson.Age {
			t.Errorf("Expected clone value {%s, %d} at index %d, got {%s, %d}", originalPerson.Name, originalPerson.Age, i, clonePerson.Name, clonePerson.Age)
		}
	}
}

func TestRemove(t *testing.T) {
	linkedList := &LinkedList[int]{}
	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)

	linkedList.Remove(0)

	// Test if item at first position was removed correctly
	if linkedList.size != 3 {
		t.Errorf("Expected LinkedList size %d, got %d", 3, linkedList.Size())
	}

	linkedList.Remove(linkedList.size - 1)

	// Test if item at last position was removed correctly
	if linkedList.size != 2 {
		t.Errorf("Expected LinkedList size %d, got %d", 2, linkedList.Size())
	}

	checkSlice := [2]int{2, 3}

	it := linkedList.Iterator()

	for _, value := range checkSlice {
		if value != it.Next() {
			t.Errorf("Expected value %d, got %d", value, linkedList.Size())
		}
	}

}

func TestEmptyList(t *testing.T) {
	linkedList := &LinkedList[int]{}

	if linkedList.Size() != 0 {
		t.Error("Expected size 0 for new list")
	}
	if linkedList.f_node != nil || linkedList.c_node != nil {
		t.Error("Expected head and tail to be nil in empty list")
	}
	if linkedList.IsEmpty() {
		t.Error("Expected IsEmpty to return false on empty list")
	}
}

func TestGetOutOfBounds(t *testing.T) {
	list := &LinkedList[int]{}
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when accessing index in empty list")
		}
	}()
	list.Get(0)
}

func TestAppendMultipleTypes(t *testing.T) {
	listInt := &LinkedList[int]{}
	listFloat := &LinkedList[float64]{}
	listString := &LinkedList[string]{}

	listInt.Append(100)
	listFloat.Append(3.14)
	listString.Append("hello")

	assertEqual(t, 100, listInt.Get(0), "Append int")
	assertEqual(t, 3.14, listFloat.Get(0), "Append float64")
	assertEqual(t, "hello", listString.Get(0), "Append string")
}

func TestInsertInAllPositions(t *testing.T) {
	list := &LinkedList[int]{}
	list.Append(1)
	list.Append(3)

	// Insert at head
	list.Insert(0, 0)
	assertEqual(t, 0, list.Get(0), "Insert at head")

	// Insert at middle
	list.Insert(2, 2)
	assertEqual(t, 2, list.Get(2), "Insert at middle")

	// Insert at end
	list.Insert(list.Size(), 4)
	assertEqual(t, 4, list.Get(4), "Insert at end")

	expected := []int{0, 1, 2, 3, 4}
	for i, val := range expected {
		assertEqual(t, val, list.Get(i), fmt.Sprintf("Insert sequence at index %d", i))
	}
}

func TestRemoveAll(t *testing.T) {
	list := &LinkedList[int]{}
	for i := 1; i <= 5; i++ {
		list.Append(i)
	}
	for i := 0; i < 5; i++ {
		list.Remove(0)
	}
	if list.Size() != 0 {
		t.Error("Expected size 0 after removing all elements")
	}
}

func TestCloneIndependence(t *testing.T) {
	list := &LinkedList[int]{}
	list.Append(1)
	list.Append(2)

	clone := list.Clone()
	clone.Append(3)
	list.Remove(1)

	if list.Size() != 1 {
		t.Error("Original list should have size 1 after remove")
	}
	if clone.Size() != 3 {
		t.Error("Clone should have size 3 after append")
	}
}

func TestIteratorExhaustion(t *testing.T) {
	list := &LinkedList[string]{}
	words := []string{"go", "is", "fun"}
	for _, w := range words {
		list.Append(w)
	}

	it := list.Iterator()
	var results []string

	for it.HasNext() {
		results = append(results, it.Next())
	}

	for i, word := range words {
		assertEqual(t, word, results[i], fmt.Sprintf("Iterator value at index %d", i))
	}
}

func TestComplexTypeModification(t *testing.T) {
	type Task struct {
		Title string
		Done  bool
	}
	list := &LinkedList[Task]{}
	task := Task{Title: "Write tests", Done: false}
	list.Append(task)

	// Modify from list directly
	ptr := &list.f_node.value
	ptr.Done = true

	if !list.Get(0).Done {
		t.Error("Expected task to be marked as done")
	}
}

func TestRemoveInvalidIndex(t *testing.T) {
	list := &LinkedList[int]{}
	list.Append(1)
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when removing invalid index")
		}
	}()
	list.Remove(5)
}

func TestRemoveOnlyElement(t *testing.T) {
	list := &LinkedList[int]{}
	list.Append(10)
	list.Remove(0)

	if list.Size() != 0 || list.f_node != nil || list.c_node != nil {
		t.Error("Expected completely empty list after removing only element")
	}
}

func TestRemoveFirstElement(t *testing.T) {
	list := &LinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Remove(0)

	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
	if list.f_node.value != 2 {
		t.Errorf("Expected first node to be 2, got %d", list.f_node.value)
	}
}

func TestRemoveLastElement(t *testing.T) {
	list := &LinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Remove(2)

	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}
	if list.c_node.value != 2 {
		t.Errorf("Expected last node to be 2, got %d", list.c_node.value)
	}
}

func TestRemoveMiddleElement(t *testing.T) {
	list := &LinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Remove(1)

	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}

	values := []int{list.f_node.value, list.c_node.value}
	expected := []int{1, 3}

	for i, v := range values {
		if v != expected[i] {
			t.Errorf("Expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestRemoveNegativeIndex(t *testing.T) {
	list := &LinkedList[int]{}
	list.Append(1)
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when removing with negative index")
		}
	}()
	list.Remove(-1)
}

func TestRemoveFromEmptyList(t *testing.T) {
	list := &LinkedList[int]{}
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when removing from empty list")
		}
	}()
	list.Remove(0)
}


func TestContainsEdgeCases(t *testing.T) {
	list := &LinkedList[string]{}
	list.Append("a")
	list.Append("b")
	list.Append("")

	if !list.Contains("") {
		t.Error("Expected to find empty string in list")
	}

	if list.Contains("c") {
		t.Error("Expected not to find value 'c'")
	}
}
