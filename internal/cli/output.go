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

	// Capacity assessment
	f.printCapacityAssessment(capacity.Assessment)
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

// printCapacityAssessment prints the capacity assessment section
func (f *OutputFormatter) printCapacityAssessment(assessment models.CapacityAssessment) {
	fmt.Printf("\n🎯 Capacity Assessment\n")

	var emoji string
	switch assessment.Status {
	case string(models.StatusOptimal):
		emoji = "✅"
	case string(models.StatusConservative), string(models.StatusAggressive):
		emoji = "⚠️ "
	default:
		emoji = "🔴"
	}

	fmt.Printf("%s %s: %d points - %s\n",
		emoji,
		assessment.Status,
		assessment.TotalPoints,
		assessment.Message)

	// Additional SRE-specific advice
	f.printSREAdvice(assessment)
}

// printSREAdvice prints SRE-specific capacity advice
func (f *OutputFormatter) printSREAdvice(assessment models.CapacityAssessment) {
	fmt.Printf("\n💼 SRE Team Considerations:\n")

	if assessment.TotalPoints >= 35 {
		fmt.Printf("   • High capacity - ensure 30%% reserved for incident response\n")
		fmt.Printf("   • Consider extending timeline to maintain reliability\n")
	} else if assessment.TotalPoints >= 28 {
		fmt.Printf("   • Good capacity - reserve 20-25%% for operational tasks\n")
		fmt.Printf("   • Monitor error budget consumption during sprint\n")
	} else {
		fmt.Printf("   • Lower capacity - opportunity for proactive improvements\n")
		fmt.Printf("   • Good time for technical debt reduction\n")
	}

	if assessment.TotalTasks >= 12 {
		fmt.Printf("   • High task count may cause context switching\n")
		fmt.Printf("   • Consider combining smaller tasks or reducing scope\n")
	} else if assessment.TotalTasks <= 6 {
		fmt.Printf("   • Low task count - good for deep focus work\n")
		fmt.Printf("   • Ideal for complex system improvements\n")
	}
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

	// SRE-specific analysis
	if combo.L >= 2 {
		advice = append(advice, "🔧 Strategic work focused - good for system improvements")
	}

	if totalTasks >= 8 && combo.XS >= 3 {
		advice = append(advice, "🚨 Mix includes rapid response capacity")
	}

	for _, tip := range advice {
		fmt.Printf("    %s\n", tip)
	}
}

// PrintJSON prints any data structure as formatted JSON
func (f *OutputFormatter) PrintJSON(data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling JSON: %w", err)
	}

	fmt.Println(string(jsonData))
	return nil
}

// PrintError prints error messages in a consistent format
func (f *OutputFormatter) PrintError(err error) {
	fmt.Printf("❌ Error: %v\n", err)
}

// PrintSuccess prints success messages
func (f *OutputFormatter) PrintSuccess(message string) {
	fmt.Printf("✅ %s\n", message)
}

// PrintWarning prints warning messages
func (f *OutputFormatter) PrintWarning(message string) {
	fmt.Printf("⚠️  %s\n", message)
}

// PrintInfo prints informational messages
func (f *OutputFormatter) PrintInfo(message string) {
	fmt.Printf("ℹ️  %s\n", message)
}

// PrintTable prints data in a formatted table
func (f *OutputFormatter) PrintTable(headers []string, rows [][]string) {
	if len(headers) == 0 || len(rows) == 0 {
		return
	}

	// Calculate column widths
	colWidths := make([]int, len(headers))
	for i, header := range headers {
		colWidths[i] = len(header)
	}

	for _, row := range rows {
		for i, cell := range row {
			if i < len(colWidths) && len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	// Print table
	f.printTableRow(headers, colWidths)
	f.printTableSeparator(colWidths)
	for _, row := range rows {
		f.printTableRow(row, colWidths)
	}
}

// printTableRow prints a single table row
func (f *OutputFormatter) printTableRow(cells []string, widths []int) {
	fmt.Print("│")
	for i, cell := range cells {
		if i < len(widths) {
			fmt.Printf(" %-*s │", widths[i], cell)
		}
	}
	fmt.Println()
}

// printTableSeparator prints table separator line
func (f *OutputFormatter) printTableSeparator(widths []int) {
	fmt.Print("├")
	for i, width := range widths {
		fmt.Print(strings.Repeat("─", width+2))
		if i < len(widths)-1 {
			fmt.Print("┼")
		}
	}
	fmt.Println("┤")
}

// PrintDivider prints a visual divider line
func (f *OutputFormatter) PrintDivider(char string, length int) {
	fmt.Println(strings.Repeat(char, length))
}

// PrintHeader prints a formatted header with decorative borders
func (f *OutputFormatter) PrintHeader(title string) {
	border := strings.Repeat("═", len(title)+4)
	fmt.Printf("╔%s╗\n", border)
	fmt.Printf("║  %s  ║\n", title)
	fmt.Printf("╚%s╝\n", border)
}

// PrintSubHeader prints a simple sub-header
func (f *OutputFormatter) PrintSubHeader(title string) {
	fmt.Printf("\n📋 %s\n", title)
	fmt.Printf("%s\n", strings.Repeat("─", len(title)+5))
}

// PrintProgress prints a simple progress indicator
func (f *OutputFormatter) PrintProgress(current, total int, message string) {
	percentage := float64(current) / float64(total) * 100
	fmt.Printf("⏳ %s: %d/%d (%.1f%%)\n", message, current, total, percentage)
}

// PrintSummary prints a summary box with key-value pairs
func (f *OutputFormatter) PrintSummary(title string, data map[string]interface{}) {
	fmt.Printf("\n📊 %s\n", title)
	fmt.Printf("┌%s┐\n", strings.Repeat("─", len(title)+5))

	for key, value := range data {
		fmt.Printf("│ %-15s: %v\n", key, value)
	}

	fmt.Printf("└%s┘\n", strings.Repeat("─", len(title)+5))
}

// PrintBulletList prints a bulleted list with optional emoji
func (f *OutputFormatter) PrintBulletList(title string, items []string, emoji string) {
	if emoji == "" {
		emoji = "•"
	}

	if title != "" {
		fmt.Printf("\n%s\n", title)
	}

	for _, item := range items {
		fmt.Printf("  %s %s\n", emoji, item)
	}
}

// PrintKeyValue prints a simple key-value pair
func (f *OutputFormatter) PrintKeyValue(key string, value interface{}) {
	fmt.Printf("%-20s: %v\n", key, value)
}

// PrintSeparator prints a visual separator
func (f *OutputFormatter) PrintSeparator() {
	fmt.Println("────────────────────────────────────────────────")
}

// PrintBanner prints a prominent banner message
func (f *OutputFormatter) PrintBanner(message string) {
	width := len(message) + 6
	border := strings.Repeat("*", width)

	fmt.Printf("\n%s\n", border)
	fmt.Printf("*  %s  *\n", message)
	fmt.Printf("%s\n\n", border)
}

// PrintBox prints text inside a box with borders
func (f *OutputFormatter) PrintBox(title string, content []string) {
	maxWidth := len(title)
	for _, line := range content {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	width := maxWidth + 4
	topBorder := "┌" + strings.Repeat("─", width-2) + "┐"
	bottomBorder := "└" + strings.Repeat("─", width-2) + "┘"

	fmt.Printf("\n%s\n", topBorder)

	if title != "" {
		padding := (width - len(title) - 2) / 2
		fmt.Printf("│%s%s%s│\n",
			strings.Repeat(" ", padding),
			title,
			strings.Repeat(" ", width-len(title)-padding-2))

		if len(content) > 0 {
			separator := "├" + strings.Repeat("─", width-2) + "┤"
			fmt.Printf("%s\n", separator)
		}
	}

	for _, line := range content {
		padding := width - len(line) - 2
		fmt.Printf("│ %s%s│\n", line, strings.Repeat(" ", padding-1))
	}

	fmt.Printf("%s\n\n", bottomBorder)
}

// PrintStatusIcon returns appropriate status icon
func (f *OutputFormatter) PrintStatusIcon(status string) string {
	switch strings.ToUpper(status) {
	case "SUCCESS", "OPTIMAL", "COMPLETED":
		return "✅"
	case "WARNING", "CONSERVATIVE", "AGGRESSIVE":
		return "⚠️"
	case "ERROR", "FAILED", "TOO_HIGH", "TOO_LOW":
		return "❌"
	case "INFO", "PENDING":
		return "ℹ️"
	case "QUESTION":
		return "❓"
	default:
		return "•"
	}
}

// PrintColoredText prints text with color codes (for terminals that support ANSI)
func (f *OutputFormatter) PrintColoredText(text, color string) {
	colorCodes := map[string]string{
		"reset":  "\033[0m",
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"cyan":   "\033[36m",
		"white":  "\033[37m",
		"bold":   "\033[1m",
	}

	if code, exists := colorCodes[strings.ToLower(color)]; exists {
		fmt.Printf("%s%s%s", code, text, colorCodes["reset"])
	} else {
		fmt.Print(text)
	}
}

// PrintFormattedList prints a formatted list with numbering or bullets
func (f *OutputFormatter) PrintFormattedList(title string, items []string, numbered bool) {
	if title != "" {
		fmt.Printf("\n%s\n", title)
		fmt.Printf("%s\n", strings.Repeat("─", len(title)))
	}

	for i, item := range items {
		if numbered {
			fmt.Printf("%2d. %s\n", i+1, item)
		} else {
			fmt.Printf("   • %s\n", item)
		}
	}
	fmt.Println()
}

// PrintStats prints statistical information in a formatted way
func (f *OutputFormatter) PrintStats(title string, stats map[string]float64, unit string) {
	if title != "" {
		fmt.Printf("\n📊 %s\n", title)
		fmt.Printf("%s\n", strings.Repeat("─", len(title)+5))
	}

	for key, value := range stats {
		if unit != "" {
			fmt.Printf("%-20s: %8.2f %s\n", key, value, unit)
		} else {
			fmt.Printf("%-20s: %8.2f\n", key, value)
		}
	}
	fmt.Println()
}

// PrintProgressBar prints a simple ASCII progress bar
func (f *OutputFormatter) PrintProgressBar(current, total int, width int, label string) {
	if width <= 0 {
		width = 40
	}

	percentage := float64(current) / float64(total)
	filled := int(percentage * float64(width))

	bar := strings.Repeat("█", filled) + strings.Repeat("░", width-filled)

	if label != "" {
		fmt.Printf("%s: [%s] %d/%d (%.1f%%)\n", label, bar, current, total, percentage*100)
	} else {
		fmt.Printf("[%s] %d/%d (%.1f%%)\n", bar, current, total, percentage*100)
	}
}

// PrintFooter prints a footer message with timestamp
func (f *OutputFormatter) PrintFooter(message string) {
	fmt.Printf("\n%s\n", strings.Repeat("─", 50))
	if message != "" {
		fmt.Printf("%s\n", message)
	}
	fmt.Printf("Generated at: %s\n", getCurrentTimestamp())
}

func getCurrentTimestamp() string {
	// In a real implementation, this would use:
	// return time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s", "2024-01-XX XX:XX:XX") // Placeholder for artifact
}
