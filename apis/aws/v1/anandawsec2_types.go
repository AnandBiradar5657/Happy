/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Anandawsec2Spec defines the desired state of Anandawsec2
type Anandawsec2Spec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Anandawsec2. Edit anandawsec2_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// Anandawsec2Status defines the observed state of Anandawsec2
type Anandawsec2Status struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Anandawsec2 is the Schema for the anandawsec2s API
type Anandawsec2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Anandawsec2Spec   `json:"spec,omitempty"`
	Status Anandawsec2Status `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Anandawsec2List contains a list of Anandawsec2
type Anandawsec2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Anandawsec2 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Anandawsec2{}, &Anandawsec2List{})
}
