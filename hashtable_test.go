package hashtable

import (
	"fmt"
	"testing"
)

var Table *Hashtable

func TestHashtable(t *testing.T) {
	Table = New()
	t.Run("test-put", testPut)
	t.Run("test-get", testGet)
	t.Run("test-delete", testDelete)
	Table = nil
}

func testPut(t *testing.T) {
	data := map[string]interface{}{
		"name": "Blaine Moser",
		"age":  31,
	}
	Table.Put("blainemoser@outlook.com", data)
	data = map[string]interface{}{
		"name": "Blaine Moser",
		"age":  32,
	}
	Table.Put("blainemoser@outlook.com", data)
}

func testGet(t *testing.T) {
	data := Table.Get("blainemoser@outlook.com")
	if profile, ok := data.(map[string]interface{}); ok {
		if profile["name"] == nil {
			panic(fmt.Errorf("name not found"))
		}
		if name, ok := profile["name"].(string); ok {
			if name != "Blaine Moser" {
				panic(fmt.Errorf("expected name to be 'Blaine Moser'"))
			}
		} else {
			panic(fmt.Errorf("name not found"))
		}
		if age, ok := profile["age"].(int); ok {
			if age != 32 {
				panic(fmt.Errorf("expected name age to be %d", 32))
			}
		} else {
			panic(fmt.Errorf("age not found"))
		}
		return
	}
	panic(fmt.Errorf("profile not a map"))
}

func testDelete(t *testing.T) {
	Table.Delete("blainemoser@outlook.com")
	data := Table.Get("blainemoser@outlook.com")
	if data != nil {
		panic(fmt.Errorf("expected '%s' to have been deleted", "blainemoser@outlook.com"))
	}
}
