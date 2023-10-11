package utils

import (
	"testing"
)

func TestMapWrapper_Put(t *testing.T) {
	// Create a new MapWrapper
	m := NewMapWrapper[int, string]()

	// Test Put
	key := 1
	value := "One"
	m.Put(key, value)

	// Verify that the key is mapped to the value
	result := m.Get(key)
	if result != value {
		t.Errorf("Expected '%s' for key %d, but got '%s'", value, key, result)
	}
}

func TestMapWrapper_PutWithNil(t *testing.T) {
	// Create a new MapWrapper
	m := NewMapWrapper[int, string]()

	// Test Put with nil value
	key := 1
	value := ""
	m.Put(key, value)

	// Verify that the key is mapped to the nil value of the value type
	result := m.Get(key)
	if result != value {
		t.Errorf("Expected empty string for key %d, but got '%s'", key, result)
	}
}

func TestMapWrapper_PutOverride(t *testing.T) {
	// Create a new MapWrapper
	m := NewMapWrapper[int, string]()

	// Test Put to override an existing key
	key := 1
	originalValue := "One"
	newValue := "NewOne"
	m.Put(key, originalValue)
	m.Put(key, newValue)

	// Verify that the key is now mapped to the new value
	result := m.Get(key)
	if result != newValue {
		t.Errorf("Expected '%s' for key %d, but got '%s'", newValue, key, result)
	}
}

func TestMapWrapper_PutMultiple(t *testing.T) {
	// Create a new MapWrapper
	m := NewMapWrapper[int, string]()

	// Test Put for multiple keys
	keys := []int{1, 2, 3}
	values := []string{"One", "Two", "Three"}

	for i, key := range keys {
		value := values[i]
		m.Put(key, value)
	}

	// Verify that each key is mapped to the corresponding value
	for i, key := range keys {
		value := values[i]
		result := m.Get(key)
		if result != value {
			t.Errorf("Expected '%s' for key %d, but got '%s'", value, key, result)
		}
	}
}

func TestMapWrapper_GetExistingKey(t *testing.T) {
	// Create a new MapWrapper
	m := NewMapWrapper[int, string]()

	// Test Get with an existing key
	key := 1
	value := "One"
	m.Put(key, value)

	// Verify that the function returns the associated value
	result := m.Get(key)
	if result != value {
		t.Errorf("Expected '%s' for key %d, but got '%s'", value, key, result)
	}
}

func TestMapWrapper_GetNonExistingKey(t *testing.T) {
	// Create a new MapWrapper
	m := NewMapWrapper[int, string]()

	// Test Get with a non-existing key
	key := 2

	// Verify that the function returns the nil value of the value type
	result := m.Get(key)
	if result != "" {
		t.Errorf("Expected an empty string for key %d, but got '%s'", key, result)
	}
}

func TestMapWrapper_RemoveExistingKey(t *testing.T) {
	// Create a new MapWrapper
	m := NewMapWrapper[int, string]()

	// Test Remove with an existing key
	key := 1
	value := "One"
	m.Put(key, value)

	// Remove the mapping
	m.Remove(key)

	// Verify that the mapping has been removed
	if m.ContainsKey(key) {
		t.Errorf("Expected key %d to be removed, but it still exists.", key)
	}
}

func TestMapWrapper_RemoveNonExistingKey(t *testing.T) {
	// Create a new MapWrapper
	m := NewMapWrapper[int, string]()

	// Test Remove with a non-existing key
	key := 2

	// Remove the non-existing key
	m.Remove(key)

	// Verify that no-op has occurred, and the map is still empty
	if size := m.Size(); size != 0 {
		t.Errorf("Expected size 0 for a non-existing key, but got %d", size)
	}
}

func TestMapWrapper_ContainsKeyExistingKey(t *testing.T) {
	// Create a new MapWrapper
	m := NewMapWrapper[int, string]()

	// Test ContainsKey with an existing key
	key := 1
	value := "One"
	m.Put(key, value)

	// Verify that the function returns true for an existing key
	result := m.ContainsKey(key)
	if !result {
		t.Errorf("Expected true for key %d, but got false", key)
	}
}

func TestMapWrapper_ContainsKeyNonExistingKey(t *testing.T) {
	// Create a new MapWrapper
	m := NewMapWrapper[int, string]()

	// Test ContainsKey with a non-existing key
	key := 2

	// Verify that the function returns false for a non-existing key
	result := m.ContainsKey(key)
	if result {
		t.Errorf("Expected false for key %d, but got true", key)
	}
}

func TestMapWrapper_MergeExistingKey(t *testing.T) {
	// Create a new MapWrapper
	m := NewMapWrapper[int, int]()

	// Test Merge with an existing key
	key := 1
	oldValue := 5
	newValue := 10
	mergeOp := func(a, b int) int {
		return a + b
	}

	// Put an initial value
	m.Put(key, oldValue)

	// Merge the new value
	m.Merge(key, newValue, mergeOp)

	// Verify that the merge operation was applied
	result := m.Get(key)
	if result != oldValue+newValue {
		t.Errorf("Expected merged value '%d' for key %d, but got '%d'", oldValue+newValue, key, result)
	}
}

func TestMapWrapper_MergeNonExistingKey(t *testing.T) {
	// Create a new MapWrapper
	m := NewMapWrapper[int, int]()

	// Test Merge with a non-existing key
	key := 2
	newValue := 10
	mergeOp := func(a, b int) int {
		return a + b
	}

	// Merge the new value for a non-existing key
	m.Merge(key, newValue, mergeOp)

	// Verify that the key and value were added
	result := m.Get(key)
	if result != newValue {
		t.Errorf("Expected merged value '%d' for key %d, but got '%d'", newValue, key, result)
	}
}

func TestMapWrapper_SizeEmptyMap(t *testing.T) {
	// Create a new empty MapWrapper
	m := NewMapWrapper[int, string]()

	// Verify that the size is 0 for an empty map
	size := m.Size()
	if size != 0 {
		t.Errorf("Expected size 0 for an empty map, but got %d", size)
	}
}

func TestMapWrapper_SizeNonEmptyMap(t *testing.T) {
	// Create a new MapWrapper with mappings
	m := NewMapWrapper[int, string]()

	// Add some mappings
	m.Put(1, "One")
	m.Put(2, "Two")

	// Verify that the size reflects the number of mappings
	size := m.Size()
	if size != 2 {
		t.Errorf("Expected size 2 for a map with 2 mappings, but got %d", size)
	}
}
