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

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AssetDB is the Schema for the assetdbs API
type AssetDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AssetDBSpec   `json:"spec,omitempty"`
	Status AssetDBStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AssetDBList contains a list of AssetDB
type AssetDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AssetDB `json:"items"`
}

// AssetDBSpec defines the desired state of AssetDB
type AssetDBSpec struct {
	AssetName string        `json:"assetname,omitempty"`
	Clusters  []ClusterItem `json:"clusters,omitempty"`
}

type ClusterItem struct {
	Name        string   `json:"name,omitempty"`
	Environment []string `json:"environment,omitempty"`
}

// AssetDBStatus defines the observed state of AssetDB
type AssetDBStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}
