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

type SamplePeriod struct {
	// +kubebuilder:validation:Enum=None;OneDay;ThreeDays;SevenDays
	Min *MinObservationPeriod `json:"min,omitempty"`
	// +kubebuilder:validation:Enum=Last90Days;Last30Days;Last7Days
	Max *MaxObservationPeriod `json:"max,omitempty"`
}

// Resize increment constants for CPU and Memory
type ResizeIncrements struct {
	// +kubebuilder:validation:Pattern:=`^[0-9]+(\.[0-9]+)?m?$`
	CPU *string `json:"cpu,omitempty"`
	// +kubebuilder:validation:Pattern:=`^[0-9]+(\.[0-9]+)?(Mi|Gi)$`
	Memory *string `json:"memory,omitempty"`
}

// ContainerVerticalScaleSpec defines the desired state of ContainerVerticalScale
type ContainerVerticalScaleSettings struct {
	// +kubebuilder:default:={cpu:{max:"64", min:"500m", throttlingTolerance:"20%", recommendAboveMax:true, recommendBelowMin:false}, memory:{max:"10Gi", min:"10Mi", recommendAboveMax:true, recommendBelowMin:true}}
	// +optional
	Limits *LimitResourceConstraints `json:"limits,omitempty"`

	// +kubebuilder:default:={cpu:{min:"10m", recommendBelowMin:false}, memory:{min:"10Mi", recommendBelowMin:true}}
	// +optional
	Requests *RequestResourceConstraints `json:"requests,omitempty"`

	// +kubebuilder:default:={cpu:"100m", memory:"100Mi"}
	// +optional
	Increments *ResizeIncrements `json:"increments,omitempty"`

	// +kubebuilder:default:={min:OneDay, max:Last30Days}
	// +optional
	ObservationPeriod *SamplePeriod `json:"observationPeriod,omitempty"`

	// +kubebuilder:default:=High
	// +kubebuilder:validation:Enum=Low;Medium;High
	// +optional
	RateOfResize *ResizeRate `json:"rateOfResize,omitempty"`

	// +kubebuilder:default:=P99
	// +kubebuilder:validation:Enum=P90;P95;P99;P99_1;P99_5;P99_9;P100
	// +optional
	Aggressiveness *PercentileAggressiveness `json:"aggressiveness,omitempty"`
}
