package even_odd

import "testing"

func testEvenOdd(t *testing.T) {
	

func TestFactorial(t *testing.T) {
	var val string
	val = even_odd(5)
	
	if val != "Odd!" {
		t.Error("Expected Odd!, got ", val)
	}
}
