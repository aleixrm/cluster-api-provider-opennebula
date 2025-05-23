/*
Copyright 2024, OpenNebula Project, OpenNebula Systems.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ONEMachineTemplateSpec defines the desired state of ONEMachineTemplate
type ONEMachineTemplateSpec struct {
	// +required
	Template ONEMachineTemplateResource `json:"template"`
}

type ONEMachineTemplateResource struct {
	// +optional
	ObjectMeta clusterv1.ObjectMeta `json:"metadata.omitempty"`

	// +required
	Spec ONEMachineSpec `json:"spec"`
}

// ONEMachineTemplateStatus defines the observed state of ONEMachineTemplate
type ONEMachineTemplateStatus struct{}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ONEMachineTemplate is the Schema for the onemachinetemplates API
type ONEMachineTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ONEMachineTemplateSpec   `json:"spec,omitempty"`
	Status ONEMachineTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ONEMachineTemplateList contains a list of ONEMachineTemplate
type ONEMachineTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ONEMachineTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ONEMachineTemplate{}, &ONEMachineTemplateList{})
}
