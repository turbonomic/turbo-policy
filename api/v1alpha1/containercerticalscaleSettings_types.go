/*
Copyright 2022.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

// Resize rate
// Low: resize by a single incremental only
// Medium: resize steps are 1/4 of the difference between the current and optimal siz
// High: resize steps are calculated to the optomcal size
type ResizeRate string

const (
	Low    ResizeRate = "Low"
	Medium ResizeRate = "Medium"
	High   ResizeRate = "High"
)

// The sample period ensures historical data for a minimu numbers of days before calculationg Agressiveness.
// This ensures a minimum set of data points before the action is generated.
type MinObservationPeriod string

const (
	None      MinObservationPeriod = "None"
	OneDay    MinObservationPeriod = "1 Day"
	ThreeDays MinObservationPeriod = "30 Days"
	SevenDays MinObservationPeriod = "7 Days"
)

type MaxObservationPeriod string

const (
	Last90Days MaxObservationPeriod = "Last 90 Days"
	Last30Days MaxObservationPeriod = "1 Day"
	Last7Days  MaxObservationPeriod = "Last 7 Days"
)

type SamplePeriod struct {
	Min MinObservationPeriod `json:"min,omitempty"`
	Max MaxObservationPeriod `json:"max,omitempty"`
}

// Aggressiveness sets how agressively Turbonomic will resize in response to resource utilization.
// For example, assume a 95 percentile. The percentile utilization is the highest value that 95% of the observed
// samples fall below. By using a percentile, actions can resize to a value that is below occational utilization spikes.
type PercentileAggressiveness string

const (
	P90   PercentileAggressiveness = "P90"
	P95   PercentileAggressiveness = "P95"
	P99   PercentileAggressiveness = "P99"
	P99_1 PercentileAggressiveness = "P99.1"
	P99_5 PercentileAggressiveness = "P99.5"
	P99_9 PercentileAggressiveness = "P99.9"
	P100  PercentileAggressiveness = "P100"
)

// LimitResourceConstraints defines the resize constraint for resource like CPU limit or Memory limit.
type LimitResourceConstraint struct {
	Max               string `json:"max,omitempty"`
	Min               string `json:"min,omitempty"`
	RecommendAboveMax bool   `json:"recommendAboveMax,omitempty"`
	RecommendBelowMin bool   `json:"recommendBelowMin,omitempty"`
}

// RequestResourceConstraint defines the resize constraint for resource like CPU request or Memory request
type RequestResourceConstraint struct {
	Min               string `json:"min,omitempty"`
	RecommendBelowMin bool   `json:"recommendBelowMin,omitempty"`
}

// LimitResourceConstraints defines the resource constraints for CPU limit and Memory limit.
type LimitResourceConstraints struct {
	CPU    LimitResourceConstraint `json:"cpu,omitempty"`
	Memory LimitResourceConstraint `json:"memory,omitempty"`
}

// RequestResourceConstraints defines the resource constraints for CPU request and Memory request
type RequestResourceConstraints struct {
	CPU    RequestResourceConstraint `json:"cpu,omitempty"`
	Memory RequestResourceConstraint `json:"memory,omitempty"`
}

// Resize increment constants for CPU and Memory
type ResizeIncrements struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

// ContainerVerticalScaleSpec defines the desired state of ContainerVerticalScale
type ContainerVerticalScaleSettings struct {
	// +kubebuilder:default:={cpu:{max:"64", min:"500m", recommendAboveMax:true, recommendBelowMin:false}, memory:{max:"10Gi", min:"10Mi", recommendAboveMax:true, recommendBelowMin:true}}
	// +optional
	Limits LimitResourceConstraints `json:"limits,omitempty"`

	// +kubebuilder:default:={cpu:{min:"10m", recommendBelowMin:false}, memory:{min:"10Mi", recommendBelowMin:true}}
	// +optional
	Requests RequestResourceConstraints `json:"requests,omitempty"`

	// +kubebuilder:default:={cpu:"100m", memory:"100Mi"}
	// +optional
	Increments ResizeIncrements `json:"increments,omitempty"`

	// +kubebuilder:default:={min:OneDay, max:Last30Days}
	// +optional
	ObservationPeriod SamplePeriod `json:"observationPeriod,omitempty"`

	// +kubebuilder:default:=High
	// +optional
	RateOfResize ResizeRate `json:"rateOfResize,omitempty"`

	// +kubebuilder:default:=P99
	// +optional
	Aggressiveness PercentileAggressiveness `json:"aggressiveness,omitempty"`
}
