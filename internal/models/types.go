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

// CapacityAssessment represents the assessment of a sprint capacity
type CapacityAssessment struct {
	Status      string `json:"status" yaml:"status"`
	Message     string `json:"message" yaml:"message"`
	TotalPoints int    `json:"total_points" yaml:"total_points"`
	TotalTasks  int    `json:"total_tasks" yaml:"total_tasks"`
	Optimal     bool   `json:"optimal" yaml:"optimal"`
}

// CapacityStatus represents different capacity status levels
type CapacityStatus string

const (
	StatusOptimal      CapacityStatus = "OPTIMAL"
	StatusConservative CapacityStatus = "CONSERVATIVE"
	StatusAggressive   CapacityStatus = "AGGRESSIVE"
	StatusTooLow       CapacityStatus = "TOO_LOW"
	StatusTooHigh      CapacityStatus = "TOO_HIGH"
)

// TaskBreakdown represents a detailed breakdown of tasks by size
type TaskBreakdown struct {
	Size   string `json:"size" yaml:"size"`
	Count  int    `json:"count" yaml:"count"`
	Points int    `json:"points" yaml:"points"`
	Total  int    `json:"total" yaml:"total"`
}

// SprintCapacity represents a complete sprint capacity calculation
type SprintCapacity struct {
	Assessment CapacityAssessment `json:"assessment" yaml:"assessment"`
	Breakdown  []TaskBreakdown    `json:"breakdown" yaml:"breakdown"`
	Tasks      TaskCount          `json:"tasks" yaml:"tasks"`
}

// CombinationResult represents the result of reverse calculation
type CombinationResult struct {
	TargetPoints    int           `json:"target_points" yaml:"target_points"`
	MaxTasks        int           `json:"max_tasks" yaml:"max_tasks"`
	Combinations    []Combination `json:"combinations" yaml:"combinations"`
	TotalFound      int           `json:"total_found" yaml:"total_found"`
	Recommendations []string      `json:"recommendations,omitempty" yaml:"recommendations,omitempty"`
}
