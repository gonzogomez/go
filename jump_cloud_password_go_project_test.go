package main

import (
		// "fmt"
		"testing"
		)

func TestEncodePassword(t *testing.T) {
	encode := EncodePassword("angryMonkey")
	expected_encode := "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q=="
	if encode != expected_encode {
		t.Errorf("Encoded password was incorrect, got: %s, want: %s.", encode, expected_encode)
	}
}