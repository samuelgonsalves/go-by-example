package main

import "testing"

func TestReadingProblems(t *testing.T) {
	problems := ReadProblems()

	if len(problems) == 0 {
		t.Errorf("No problems were read!")
	}

	for _, problem := range problems {
		if problem[0] == "" {
			t.Errorf("Empty problem")
		}
		if problem[1] == "" {
			t.Errorf("Empty solution")
		}
	}
}
