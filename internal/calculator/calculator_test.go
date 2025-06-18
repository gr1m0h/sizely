package calculator

import (
	"testing"

	"github.com/gr1m0h/sizely/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePoints(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name     string
		tasks    models.TaskCount
		expected int
	}{
		{
			name:     "Empty tasks",
			tasks:    models.TaskCount{XS: 0, S: 0, M: 0, L: 0},
			expected: 0,
		},
		{
			name:     "Only XS tasks",
			tasks:    models.TaskCount{XS: 5, S: 0, M: 0, L: 0},
			expected: 5,
		},
		{
			name:     "Mixed tasks",
			tasks:    models.TaskCount{XS: 3, S: 2, M: 1, L: 1},
			expected: 24, // 3*1 + 2*3 + 1*5 + 1*10 = 3 + 6 + 5 + 10 = 24
		},
		{
			name:     "All large tasks",
			tasks:    models.TaskCount{XS: 0, S: 0, M: 0, L: 3},
			expected: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.CalculatePoints(tt.tasks)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCalculateSprintCapacity(t *testing.T) {
	calc := NewCalculator()
	tasks := models.TaskCount{XS: 3, S: 2, M: 1, L: 1}

	result := calc.CalculateSprintCapacity(tasks)

	// Check total points calculation
	expectedPoints := 3*1 + 2*3 + 1*5 + 1*10 // 3 + 6 + 5 + 10 = 24
	actualPoints := 0
	for _, breakdown := range result.Breakdown {
		actualPoints += breakdown.Total
	}
	assert.Equal(t, expectedPoints, actualPoints)

	// Check breakdown
	assert.Len(t, result.Breakdown, 4)

	// Check XS breakdown
	xsBreakdown := result.Breakdown[0]
	assert.Equal(t, "XS", xsBreakdown.Size)
	assert.Equal(t, 3, xsBreakdown.Count)
	assert.Equal(t, 1, xsBreakdown.Points)
	assert.Equal(t, 3, xsBreakdown.Total)

	// Check L breakdown
	lBreakdown := result.Breakdown[3]
	assert.Equal(t, "L", lBreakdown.Size)
	assert.Equal(t, 1, lBreakdown.Count)
	assert.Equal(t, 10, lBreakdown.Points)
	assert.Equal(t, 10, lBreakdown.Total)

	// Check tasks
	assert.Equal(t, tasks, result.Tasks)
}

func TestFindCombinations(t *testing.T) {
	calc := NewCalculator()

	tests := []struct {
		name         string
		targetPoints int
		maxTasks     int
		minExpected  int // minimum number of combinations expected
	}{
		{
			name:         "Small target",
			targetPoints: 10,
			maxTasks:     10,
			minExpected:  1, // At least L×1 should be possible
		},
		{
			name:         "Medium target",
			targetPoints: 24,
			maxTasks:     15,
			minExpected:  1,
		},
		{
			name:         "Large target",
			targetPoints: 33,
			maxTasks:     15,
			minExpected:  5, // Multiple combinations should be possible
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.FindCombinations(tt.targetPoints, tt.maxTasks)

			assert.Equal(t, tt.targetPoints, result.TargetPoints)
			assert.Equal(t, tt.maxTasks, result.MaxTasks)
			assert.GreaterOrEqual(t, result.TotalFound, tt.minExpected)
			assert.Equal(t, len(result.Combinations), result.TotalFound)

			// Verify all combinations are valid
			for _, combo := range result.Combinations {
				totalPoints := combo.XS*1 + combo.S*3 + combo.M*5 + combo.L*10
				totalTasks := combo.XS + combo.S + combo.M + combo.L

				assert.Equal(t, tt.targetPoints, totalPoints, "Combination points should match target")
				assert.LessOrEqual(t, totalTasks, tt.maxTasks, "Combination tasks should not exceed max")
				assert.Equal(t, tt.targetPoints, combo.Points, "Combination.Points should match target")
			}
		})
	}
}

func TestGenerateCombinationsEdgeCases(t *testing.T) {
	calc := NewCalculator()

	t.Run("No valid combinations", func(t *testing.T) {
		result := calc.FindCombinations(2, 1) // 2 points with max 1 task - XS×2 = 2 tasks, exceeds max
		assert.Equal(t, 0, result.TotalFound) // No valid combinations possible
	})

	t.Run("Zero points", func(t *testing.T) {
		result := calc.FindCombinations(0, 5)
		assert.Equal(t, 1, result.TotalFound) // Only empty combination
		assert.Equal(t, 0, result.Combinations[0].XS)
		assert.Equal(t, 0, result.Combinations[0].S)
		assert.Equal(t, 0, result.Combinations[0].M)
		assert.Equal(t, 0, result.Combinations[0].L)
	})

	t.Run("Very restrictive max tasks", func(t *testing.T) {
		result := calc.FindCombinations(30, 3) // 30 points with max 3 tasks (L×3 = 30 points)
		assert.Greater(t, result.TotalFound, 0)

		for _, combo := range result.Combinations {
			totalTasks := combo.XS + combo.S + combo.M + combo.L
			assert.LessOrEqual(t, totalTasks, 3)
		}
	})
}

func TestCombinationsSorting(t *testing.T) {
	calc := NewCalculator()
	result := calc.FindCombinations(15, 10)

	// Check that combinations are sorted by total tasks (ascending)
	for i := 1; i < len(result.Combinations); i++ {
		prevTotal := result.Combinations[i-1].XS + result.Combinations[i-1].S + result.Combinations[i-1].M + result.Combinations[i-1].L
		currTotal := result.Combinations[i].XS + result.Combinations[i].S + result.Combinations[i].M + result.Combinations[i].L
		assert.LessOrEqual(t, prevTotal, currTotal, "Combinations should be sorted by total tasks")
	}
}

func BenchmarkCalculatePoints(b *testing.B) {
	calc := NewCalculator()
	tasks := models.TaskCount{XS: 5, S: 3, M: 2, L: 1}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calc.CalculatePoints(tasks)
	}
}

func BenchmarkFindCombinations(b *testing.B) {
	calc := NewCalculator()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calc.FindCombinations(33, 15)
	}
}
