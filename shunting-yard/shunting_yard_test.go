package shunting_yard

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetPostfixFormula(t *testing.T) {
	input := "3+4x2/(1-5)^2^3+2-2"
	expectedOutput := []string{"3", "4", "2", "x", "1", "5", "-", "2", "3", "^", "^", "/", "+", "2", "+", "2", "-"}

	output := getPostfixFormula(input)

	if diff := cmp.Diff(expectedOutput, output); diff != "" {
		t.Fatalf("(-want/+got)/n%s", diff)
	}
}
