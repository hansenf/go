package main

import "testing"

func TestIsOddPositive(t *testing.T) {
	t.Log("Input: 5")
	t.Log("Expected: true")

	if IsOdd(5) != true {
		t.Error("Salah! 5 seharusnya ganjil dan function return true")
	}
}

func TestIsOddNegative(t *testing.T) {
	t.Log("Input: 6")
	t.Log("Expected: false")

	if IsOdd(6) != false {
		t.Error("Salah! 6 seharusnya genap dan function return false")
	}
}

func TestIsOddBadInput(t *testing.T) {
	t.Log("Input: -5")
	t.Log("Expected: false")

	if IsOdd(-5) != false {
		t.Error("Salah! -5 seharusnya dianggap bad input dan function return false")
	}
}

