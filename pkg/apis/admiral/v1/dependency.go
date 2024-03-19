package v1

import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Dependency struct {
	v1.TypeMeta   `json:",inline"`
	v1.ObjectMeta `json:"metadata"`
	Spec          DependencySpec   `json:"spec"`
	Status        DependencyStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DependencyList struct {
	v1.TypeMeta `json:",inline"`
	v1.ListMeta `json:"metadata"`

	Items []Dependency `json:"items"`
}

type DependencySpec struct {
	Source        string   `json:"source,omitempty"`
	IdentityLabel string   `json:"identityLabel,omitempty"`
	Destinations  []string `json:"destinations,omitempty"`
}

type DependencyStatus struct {
	ClusterSynced int32  `json:"clustersSynced"`
	State         string `json:"state"`
}
