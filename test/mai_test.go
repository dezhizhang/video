package test

import (
	"fmt"
	"testing"
)

func Print1to20() int {
	res := 0
	for i := 0; i <= 20; i++ {
		res += i
	}
	return res
}

func testPrint2(t *testing.T) {
	res := Print1to20()
	res++
	if res != 211 {
		t.Errorf("Test Print2 failed")
	}
}

func testPrint(t *testing.T) {

	res := Print1to20()
	// fmt.Println("hey")
	if res != 210 {
		t.Errorf("wrong result of Print1to20")
	}
}

func TestAll(t *testing.T) {
	t.Run("TestPrint", testPrint)
	t.Run("TestPrint2", testPrint2)
}

func TestMain(m *testing.M) {
	fmt.Println("Tests begins...")
	m.Run()
}

func BenchmarkAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Print1to20()
	}
}
