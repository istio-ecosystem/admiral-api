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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// DynamicRoutingApplyConfiguration represents an declarative configuration of the DynamicRouting type for use
// with apply.
type DynamicRoutingApplyConfiguration struct {
	Name              *string `json:"name,omitempty"`
	Url               *string `json:"url,omitempty"`
	CacheKeyAlgorithm *string `json:"cacheKeyAlgorithm,omitempty"`
	TtlSec            *int    `json:"ttlSec,omitempty"`
	Local             *bool   `json:"local,omitempty"`
}

// DynamicRoutingApplyConfiguration constructs an declarative configuration of the DynamicRouting type for use with
// apply.
func DynamicRouting() *DynamicRoutingApplyConfiguration {
	return &DynamicRoutingApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *DynamicRoutingApplyConfiguration) WithName(value string) *DynamicRoutingApplyConfiguration {
	b.Name = &value
	return b
}

// WithUrl sets the Url field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Url field is set to the value of the last call.
func (b *DynamicRoutingApplyConfiguration) WithUrl(value string) *DynamicRoutingApplyConfiguration {
	b.Url = &value
	return b
}

// WithCacheKeyAlgorithm sets the CacheKeyAlgorithm field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CacheKeyAlgorithm field is set to the value of the last call.
func (b *DynamicRoutingApplyConfiguration) WithCacheKeyAlgorithm(value string) *DynamicRoutingApplyConfiguration {
	b.CacheKeyAlgorithm = &value
	return b
}

// WithTtlSec sets the TtlSec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TtlSec field is set to the value of the last call.
func (b *DynamicRoutingApplyConfiguration) WithTtlSec(value int) *DynamicRoutingApplyConfiguration {
	b.TtlSec = &value
	return b
}

// WithLocal sets the Local field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Local field is set to the value of the last call.
func (b *DynamicRoutingApplyConfiguration) WithLocal(value bool) *DynamicRoutingApplyConfiguration {
	b.Local = &value
	return b
}
