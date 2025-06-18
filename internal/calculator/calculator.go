package calculator

import (
	"sort"

	"github.com/gr1m0h/sizely/internal/models"
)

// Calculator handles sprint capacity calculations
type Calculator struct{}

// NewCalculator creates a new Calculator instance
func NewCalculator() *Calculator {
	return &Calculator{}
}

// CalculatePoints calculates total points from task counts
func (c *Calculator) CalculatePoints(tasks models.TaskCount) int {
	return tasks.XS*models.TShirtSizePoints["XS"] +
		tasks.S*models.TShirtSizePoints["S"] +
		tasks.M*models.TShirtSizePoints["M"] +
		tasks.L*models.TShirtSizePoints["L"]
}

// CalculateSprintCapacity calculates complete sprint capacity with assessment
func (c *Calculator) CalculateSprintCapacity(tasks models.TaskCount) models.SprintCapacity {
	// Create breakdown
	breakdown := []models.TaskBreakdown{
		{
			Size:   "XS",
			Count:  tasks.XS,
			Points: models.TShirtSizePoints["XS"],
			Total:  tasks.XS * models.TShirtSizePoints["XS"],
		},
		{
			Size:   "S",
			Count:  tasks.S,
			Points: models.TShirtSizePoints["S"],
			Total:  tasks.S * models.TShirtSizePoints["S"],
		},
		{
			Size:   "M",
			Count:  tasks.M,
			Points: models.TShirtSizePoints["M"],
			Total:  tasks.M * models.TShirtSizePoints["M"],
		},
		{
			Size:   "L",
			Count:  tasks.L,
			Points: models.TShirtSizePoints["L"],
			Total:  tasks.L * models.TShirtSizePoints["L"],
		},
	}

	totalPoints := c.CalculatePoints(tasks)
	totalTasks := tasks.XS + tasks.S + tasks.M + tasks.L

	return models.SprintCapacity{
		TotalPoints: totalPoints,
		TotalTasks:  totalTasks,
		Breakdown:   breakdown,
		Tasks:       tasks,
	}
}

// FindCombinations finds all task combinations for target points
func (c *Calculator) FindCombinations(targetPoints, maxTasks int) models.CombinationResult {
	combinations := c.generateCombinations(targetPoints, maxTasks)

	return models.CombinationResult{
		TargetPoints: targetPoints,
		MaxTasks:     maxTasks,
		Combinations: combinations,
		TotalFound:   len(combinations),
	}
}

// generateCombinations generates all valid combinations for target points
func (c *Calculator) generateCombinations(targetPoints, maxTasks int) []models.Combination {
	var combinations []models.Combination

	maxL := min(targetPoints/models.TShirtSizePoints["L"], maxTasks)

	for l := 0; l <= maxL; l++ {
		remainingAfterL := targetPoints - l*models.TShirtSizePoints["L"]
		maxM := min(remainingAfterL/models.TShirtSizePoints["M"], maxTasks-l)

		for m := 0; m <= maxM; m++ {
			remainingAfterM := remainingAfterL - m*models.TShirtSizePoints["M"]
			maxS := min(remainingAfterM/models.TShirtSizePoints["S"], maxTasks-l-m)

			for s := 0; s <= maxS; s++ {
				remainingAfterS := remainingAfterM - s*models.TShirtSizePoints["S"]

				// XS must match exactly and not exceed max tasks
				if remainingAfterS >= 0 && remainingAfterS%models.TShirtSizePoints["XS"] == 0 {
					xs := remainingAfterS / models.TShirtSizePoints["XS"]
					totalTasks := l + m + s + xs

					if totalTasks <= maxTasks {
						combo := models.Combination{
							XS:     xs,
							S:      s,
							M:      m,
							L:      l,
							Points: targetPoints,
						}
						combinations = append(combinations, combo)
					}
				}
			}
		}
	}

	// Sort combinations by total tasks (ascending)
	sort.Slice(combinations, func(i, j int) bool {
		totalI := combinations[i].XS + combinations[i].S + combinations[i].M + combinations[i].L
		totalJ := combinations[j].XS + combinations[j].S + combinations[j].M + combinations[j].L
		return totalI < totalJ
	})

	return combinations
}
