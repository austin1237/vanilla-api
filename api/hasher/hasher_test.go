package hasher

import "testing"

func TestHashString(t *testing.T) {
	expected := "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="
	result := HashString("angryMonkey")

	if result != expected {
		t.Errorf("Hash was incorrect, got: %s, expected: %s.", result, expected)
	}
}
