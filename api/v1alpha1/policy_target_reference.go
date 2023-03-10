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

type PolicyTargetReference struct {
	// The Kind of the target referent
	Kind string `json:"kind"`

	// The Name of the target referent
	Name string `json:"name"`

	// The API version of the target referent
	// +optional
	APIVersion string `json:"apiVersion,omitempty"`

	// The Name of the container need to be specified if the target is workload controller
	// +optional
	Container string `json:"container,omitempty"`
}
