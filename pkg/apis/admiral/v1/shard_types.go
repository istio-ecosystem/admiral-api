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
// Shard is the Schema for the shards API
type Shard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ShardSpec   `json:"spec,omitempty"`
	Status ShardStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ShardList contains a list of Shard
type ShardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Shard `json:"items"`
}

// ShardSpec defines the desired state of Shard
type ShardSpec struct {
	Clusters []ClusterShards `json:"clusters,omitempty"`
}

// clusters defines list of clusters monitored by the shard
// This includes cluster name, locality and list of cluster identities for which resources need to be managed
type ClusterShards struct {
	Name       string         `json:"name,omitempty"`
	Locality   string         `json:"locality,omitempty"`
	Identities []IdentityItem `json:"identities,omitempty"`
}

type IdentityItem struct {
	Name        string `json:"name,omitempty"`
	Environment string `json:"environment,omitempty"`
}

// ShardStatus defines the observed state of Shard
type ShardStatus struct {
	ClustersMonitored int                    `json:"clustersMonitored,omitempty"`
	Conditions        []ShardStatusCondition `json:"conditions,omitempty"`
	FailureDetails    FailureDetails         `json:"failureDetails,omitempty"`
	LastUpdatedTime   metav1.Time            `json:"lastUpdateTime,omitempty"`
}

type ConditionStatus string

const (
	TrueConditionStatus  ConditionStatus = "true"
	FalseConditionStatus ConditionStatus = "false"
)

type ConditionType string

const (
	SyncComplete ConditionType = "SyncComplete"
	SyncFailed   ConditionType = "SyncFailed"
)

type ConditionReason string

const (
	Processing    ConditionReason = "stillProcessing"
	Processed     ConditionReason = "processed"
	ErrorOccurred ConditionReason = "errorOccurred"
)

/* condition defines details for status condition including type, when it was updates and reason for the update

	Possible condition type are -
	SyncComplete - Set to provide update on the sync state
	SyncFailed - Set to provide update on the sync state if failure occurred

	Possible condition reason are -
	stillProcessing - set when resources for the provided identities are getting processed
    processed - set when all the clusters and related identity resource are processes
    errorOccurred - set when error occurred while processing the resources, more details for which identities failed and why will be provided in failureDetails section
*/
type ShardStatusCondition struct {
	Message         string          `json:"message,omitempty"`
	Reason          ConditionReason `json:"reason,omitempty"`
	Status          ConditionStatus `json:"status,omitempty"`
	Type            ConditionType   `json:"type,omitempty"`
	LastUpdatedTime metav1.Time     `json:"lastUpdateTime,omitempty"`
}

// failureDetails define details of which clusters and identities observed failures while processing resources
type FailureDetails struct {
	LastUpdatedTime metav1.Time     `json:"lastUpdateTime,omitempty"`
	FailedClusters  []FailedCluster `json:"clusters,omitempty"`
}

type FailedCluster struct {
	Name             string           `json:"name,omitempty"`
	FailedIdentities []FailedIdentity `json:"identities,omitempty"`
}

type FailedIdentity struct {
	Name    string `json:"name,omitempty"`
	Message string `json:"errorMessage,omitempty"`
}
