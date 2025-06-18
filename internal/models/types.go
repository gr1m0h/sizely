package models

// TShirtSizePoints represents the points values for each T-shirt size
var TShirtSizePoints = map[string]int{
	"XS": 1,
	"S":  3,
	"M":  5,
	"L":  10,
}

// TaskCount represents the count of each T-shirt size
type TaskCount struct {
	XS int `json:"xs" yaml:"xs"`
	S  int `json:"s" yaml:"s"`
	M  int `json:"m" yaml:"m"`
	L  int `json:"l" yaml:"l"`
}

// Combination represents a combination of T-shirt sizes with calculated points
type Combination struct {
	XS     int `json:"xs" yaml:"xs"`
	S      int `json:"s" yaml:"s"`
	M      int `json:"m" yaml:"m"`
	L      int `json:"l" yaml:"l"`
	Points int `json:"points" yaml:"points"`
}

// TaskBreakdown represents a detailed breakdown of tasks by size
type TaskBreakdown struct {
	Size   string `json:"size" yaml:"size"`
	Count  int    `json:"count" yaml:"count"`
	Points int    `json:"points" yaml:"points"`
	Total  int    `json:"total" yaml:"total"`
}

// SprintCapacity represents a complete sprint capacity calculation
type SprintCapacity struct {
	TotalPoints int
	TotalTasks  int
	Breakdown   []TaskBreakdown `json:"breakdown" yaml:"breakdown"`
	Tasks       TaskCount       `json:"tasks" yaml:"tasks"`
}

// CombinationResult represents the result of reverse calculation
type CombinationResult struct {
	TargetPoints    int           `json:"target_points" yaml:"target_points"`
	MaxTasks        int           `json:"max_tasks" yaml:"max_tasks"`
	Combinations    []Combination `json:"combinations" yaml:"combinations"`
	TotalFound      int           `json:"total_found" yaml:"total_found"`
	Recommendations []string      `json:"recommendations,omitempty" yaml:"recommendations,omitempty"`
}
