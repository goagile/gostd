package word

import "testing"
import "fmt"

func TestIsTrue(t *testing.T) {
	if !IsTrue() {
		t.Error("IsTrue() должен вернуть true, а вернул false!")
	}
}

func TestIsFalse(t *testing.T) {
	expected := false

	result := IsFalse()
	
	if expected != result {
		t.Error(fmt.Sprintf("Метод IsFalse() должен вернуть false а вернул %v", result))
	}
}
