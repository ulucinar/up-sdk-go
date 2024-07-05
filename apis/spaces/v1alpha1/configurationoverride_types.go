// Copyright 2024 Upbound Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,categories=spaces

type ConfigurationOverride struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigurationOverrideSpec   `json:"spec"`
	Status ConfigurationOverrideStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

type ConfigurationOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigurationOverride `json:"items"`
}

type PropagationMode string

const (
	PropagateAscend  PropagationMode = "Ascending"
	PropagateDescend PropagationMode = "Descending"
	PropagateNone    PropagationMode = "None"
)

type Target struct {
	*corev1.TypedObjectReference `json:",inline"`
	// +kubebuilder:validation:Enum=managed
	Category *string `json:"category,omitempty"`
}

type Metadata struct {
	Annotations map[string]string `json:"annotations,omitempty"`
}

type Patch struct {
	Metadata *Metadata `json:"metadata,omitempty"`
}

type ConfigurationOverrideSpec struct {
	// +kubebuilder:validation:MinLength=1
	ControlPlane string `json:"controlPlane"`

	Target Target `json:"target"`
	// +kubebuilder:validation:Enum=None;Ascending;Descending
	// +kubebuilder:default=None
	PropagationMode PropagationMode `json:"propagationMode"`
	Patch           Patch           `json:"patch"`
}

type ConfigurationOverrideStatus struct {
	xpv1.ResourceStatus `json:",inline"`
}

func init() {
	SchemeBuilder.Register(&ConfigurationOverride{}, &ConfigurationOverrideList{})
}
