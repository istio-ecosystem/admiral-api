/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/istio-ecosystem/admiral-api/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AssetDBs returns a AssetDBInformer.
	AssetDBs() AssetDBInformer
	// Dependencies returns a DependencyInformer.
	Dependencies() DependencyInformer
	// Shards returns a ShardInformer.
	Shards() ShardInformer
	// TrafficConfigs returns a TrafficConfigInformer.
	TrafficConfigs() TrafficConfigInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AssetDBs returns a AssetDBInformer.
func (v *version) AssetDBs() AssetDBInformer {
	return &assetDBInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Dependencies returns a DependencyInformer.
func (v *version) Dependencies() DependencyInformer {
	return &dependencyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Shards returns a ShardInformer.
func (v *version) Shards() ShardInformer {
	return &shardInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TrafficConfigs returns a TrafficConfigInformer.
func (v *version) TrafficConfigs() TrafficConfigInformer {
	return &trafficConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
