package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gr1m0h/sizely/internal/models"
)

// OutputFormatter handles formatting and printing output
type OutputFormatter struct{}

// NewOutputFormatter creates a new OutputFormatter instance
func NewOutputFormatter() *OutputFormatter {
	return &OutputFormatter{}
}

// PrintCapacity prints sprint capacity calculation results
func (f *OutputFormatter) PrintCapacity(capacity models.SprintCapacity) {
	fmt.Printf("📊 Sprint Capacity Calculation\n")
	fmt.Printf("═══════════════════════════════\n")
	fmt.Printf("XS (1pt):   %d tasks =  %d points\n", capacity.Tasks.XS, capacity.Tasks.XS*1)
	fmt.Printf("S  (3pt):   %d tasks =  %d points\n", capacity.Tasks.S, capacity.Tasks.S*3)
	fmt.Printf("M  (5pt):   %d tasks =  %d points\n", capacity.Tasks.M, capacity.Tasks.M*5)
	fmt.Printf("L (10pt):   %d tasks = %d points\n", capacity.Tasks.L, capacity.Tasks.L*10)
	fmt.Printf("───────────────────────────────\n")
	fmt.Printf("Total:      %d tasks = %d points\n", capacity.Tasks.XS+capacity.Tasks.S+capacity.Tasks.M+capacity.Tasks.L, capacity.TotalPoints)
	fmt.Println()
}

// PrintCombinations prints reverse calculation results
func (f *OutputFormatter) PrintCombinations(result models.CombinationResult) {
	fmt.Printf("🔍 Finding combinations for %d points (max %d tasks)\n",
		result.TargetPoints, result.MaxTasks)
	fmt.Printf("═══════════════════════════════════════════════════\n")

	if result.TotalFound == 0 {
		fmt.Printf("No combinations found for %d points with max %d tasks\n",
			result.TargetPoints, result.MaxTasks)
		return
	}

	fmt.Printf("Found %d combination(s):\n\n", result.TotalFound)

	for i, combo := range result.Combinations {
		f.printCombination(i+1, combo)
	}

	// Print recommendations
	if len(result.Recommendations) > 0 {
		fmt.Printf("💡 Recommendations:\n")
		for _, rec := range result.Recommendations {
			fmt.Printf("   • %s\n", rec)
		}
		fmt.Println()
	}

	// Generate JSON output for easy integration
	fmt.Printf("📋 JSON Output:\n")
	jsonOutput, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("%s\n", jsonOutput)
}

// printCombination prints a single combination with analysis
func (f *OutputFormatter) printCombination(index int, combo models.Combination) {
	totalTasks := combo.XS + combo.S + combo.M + combo.L
	fmt.Printf("%2d. ", index)

	var parts []string
	if combo.L > 0 {
		parts = append(parts, fmt.Sprintf("L×%d", combo.L))
	}
	if combo.M > 0 {
		parts = append(parts, fmt.Sprintf("M×%d", combo.M))
	}
	if combo.S > 0 {
		parts = append(parts, fmt.Sprintf("S×%d", combo.S))
	}
	if combo.XS > 0 {
		parts = append(parts, fmt.Sprintf("XS×%d", combo.XS))
	}

	if len(parts) == 0 {
		fmt.Printf("No tasks")
	} else {
		fmt.Printf("%s", strings.Join(parts, " + "))
	}

	fmt.Printf(" = %d points (%d tasks)\n", combo.Points, totalTasks)

	// Add specific recommendations for this combination
	f.printCombinationAdvice(combo, totalTasks)
	fmt.Println()
}

// printCombinationAdvice prints advice for a specific combination
func (f *OutputFormatter) printCombinationAdvice(combo models.Combination, totalTasks int) {
	var advice []string

	// Task count analysis
	if totalTasks <= 6 {
		advice = append(advice, "💡 Low task count - excellent for focused work")
	} else if totalTasks >= 12 {
		advice = append(advice, "⚠️  High task count - may cause context switching")
	}

	// Balance analysis
	if combo.L > 0 && (combo.XS > 0 || combo.S > 0) {
		advice = append(advice, "✅ Good mix of large and small tasks")
	} else if combo.L > 2 {
		advice = append(advice, "🎯 Heavy on large tasks - ensure adequate planning")
	} else if combo.L == 0 && (combo.XS > 5 || combo.S > 4) {
		advice = append(advice, "⚡ Many small tasks - good for quick wins")
	}

	for _, tip := range advice {
		fmt.Printf("    %s\n", tip)
	}
}
