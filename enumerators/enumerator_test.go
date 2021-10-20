package enumerators

import (
	"testing"
)

func CheckIntEnumerator(t *testing.T, wants []int, e Enumerator[int]) {
	for _, want := range wants {
		if m := e.MoveNext(); !m {
			t.Errorf("MoveNext(): want %t, got %t", true, m)
		}

		if got := e.Current(); want != got {
			t.Errorf("Current(): want %d, got %d", want, got)
		}
	}

	if m := e.MoveNext(); m {
		t.Errorf("MoveNext(): want %t, got %t", false, m)
	}
}
