package v1alpha1

// MEMORY_VALUE_PATTERN is a regular expression that matches valid memory values.
// The expected format is a positive integer, optionally followed by a decimal point and more digits,
// followed by either "Mi" (mebibytes) or "Gi" (gibibytes).
const MEMORY_VALUE_PATTERN string = `^[0-9]+(\.[0-9]+)?(Mi|Gi)$`

// cpuPattern is a regular expression that matches strings representing CPU resources in Kubernetes.
// The format of the string is a decimal number followed by an optional "m" suffix, which stands for milli-cores.
// Examples of valid CPU resources are: "1", "1000m", "0.5", "500m", etc.
const CPU_VALUE_PATTERN1 string = `^[0-9]+(\.[0-9]+)?m?$`

// Resize rate
// Low: resize by a single incremental only
// Medium: resize steps are 1/4 of the difference between the current and optimal siz
// High: resize steps are calculated to the optomcal size
type ResizeRate string

const (
	LOW    ResizeRate = "low"
	MEDIUM ResizeRate = "medium"
	HIGH   ResizeRate = "high"
)

// The sample period ensures historical data for a minimu numbers of days before calculationg Agressiveness.
// This ensures a minimum set of data points before the action is generated.
type MinObservationPeriod string

const (
	NONE       MinObservationPeriod = "none"
	ONE_DAY    MinObservationPeriod = "1d"
	THREE_DAYS MinObservationPeriod = "3d"
	SEVEN_DAYS MinObservationPeriod = "7d"
)

type MaxObservationPeriod string

const (
	LAST_90_DAYS    MaxObservationPeriod = "90d"
	LAST_30_DAYS    MaxObservationPeriod = "30d"
	LAST_SEVEN_DAYS MaxObservationPeriod = "7d"
)

// Aggressiveness sets how agressively Turbonomic will resize in response to resource utilization.
// For example, assume a 95 percentile. The percentile utilization is the highest value that 95% of the observed
// samples fall below. By using a percentile, actions can resize to a value that is below occational utilization spikes.
type PercentileAggressiveness string

const (
	P90   PercentileAggressiveness = "p90"
	P95   PercentileAggressiveness = "p95"
	P99   PercentileAggressiveness = "p99"
	P99_1 PercentileAggressiveness = "p99.1"
	P99_5 PercentileAggressiveness = "p99.5"
	P99_9 PercentileAggressiveness = "p99.9"
	P100  PercentileAggressiveness = "p100"
)

type ActionMode string

const (
	Automatic ActionMode = "Automatic"
	Manual    ActionMode = "Manual"
	Recommend ActionMode = "Recommend"
	Disabled  ActionMode = "Disabled"
)
