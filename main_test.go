package main

import (
    "testing"
)

func TestGetCounter(t *testing.T) {
	count := getCounter()

	if count <= 0 {
		t.Errorf("Counter = %d; want greater than zero", count)
	}
}