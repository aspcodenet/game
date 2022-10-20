package concurrency

import "testing"

func TestVerifyAllBlablaIsCorrect(t *testing.T) {
	result := Run4()

	if result != 23 {
		t.Fail()
	}

}

func BenchmarkRun4(t *testing.B) {
	Run4()
}
