package v1alpha1

// CpuLimitResourceConstraint defines the resize constraint for CPU limit.
type CpuLimitResourceConstraint struct {
	// +kubebuilder:validation:Pattern:=`^[0-9]+(\.[0-9]+)?m?$`
	Max *string `json:"max,omitempty"`
	// +kubebuilder:validation:Pattern:=`^[0-9]+(\.[0-9]+)?m?$`
	Min               *string `json:"min,omitempty"`
	RecommendAboveMax *bool   `json:"recommendAboveMax,omitempty"`
	RecommendBelowMin *bool   `json:"recommendBelowMin,omitempty"`
}

// MemoryResourceConstraint defines the resize constraint for Memory limit.
type MemoryLimitResourceConstraint struct {
	// +kubebuilder:validation:Pattern:=`^[0-9]+(\.[0-9]+)?(Mi|Gi)$`
	Max *string `json:"max,omitempty"`
	// +kubebuilder:validation:Pattern:=`^[0-9]+(\.[0-9]+)?(Mi|Gi)$`
	Min               *string `json:"min,omitempty"`
	RecommendAboveMax *bool   `json:"recommendAboveMax,omitempty"`
	RecommendBelowMin *bool   `json:"recommendBelowMin,omitempty"`
}

// LimitResourceConstraints defines the resource constraints for CPU limit and Memory limit.
type LimitResourceConstraints struct {
	CPU    *CpuLimitResourceConstraint    `json:"cpu,omitempty"`
	Memory *MemoryLimitResourceConstraint `json:"memory,omitempty"`
}

// CpuRequestResourceConstraint defines the resize constraint for CPU request.
// For now Turbo only generate resize down for CPU request.
type CpuRequestResourceConstraint struct {
	// +kubebuilder:validation:Pattern:=`^[0-9]+(\.[0-9]+)?m?$`
	Min               *string `json:"min,omitempty"`
	RecommendBelowMin *bool   `json:"recommendBelowMin,omitempty"`
}

// CpuRequestResourceConstraint defines the resize constraint for Memory request.
// For now Turbo only generate resize down for Memory request.
type MemoryRequestResourceConstraint struct {
	// +kubebuilder:validation:Pattern:=`^[0-9]+(\.[0-9]+)?(Mi|Gi)$`
	Min               *string `json:"min,omitempty"`
	RecommendBelowMin *bool   `json:"recommendBelowMin,omitempty"`
}

// RequestResourceConstraints defines the resource constraints for CPU request and Memory request
type RequestResourceConstraints struct {
	CPU    *CpuRequestResourceConstraint    `json:"cpu,omitempty"`
	Memory *MemoryRequestResourceConstraint `json:"memory,omitempty"`
}
