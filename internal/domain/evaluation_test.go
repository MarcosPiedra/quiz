package domain

import "testing"

func TestCalculateComparation(t *testing.T) {
	tests := []struct {
		others       []float32
		score        float32
		expectedComp int
	}{
		{[]float32{0.1, 0.2, 0.3}, 0.5, 100},     // Better than all
		{[]float32{0.5, 0.7, 0.9}, 0.5, 0},       // Equal to one, worse than others
		{[]float32{}, 0.5, -1},                   // Empty array
		{[]float32{0.3, 0.4, 0.5, 0.6}, 0.5, 50}, // Equal case
	}

	for _, tt := range tests {
		result := CalculateComparation(tt.others, tt.score)
		if result != tt.expectedComp {
			t.Errorf("For score %v, expected %d, got %d", tt.score, tt.expectedComp, result)
		}
	}
}
