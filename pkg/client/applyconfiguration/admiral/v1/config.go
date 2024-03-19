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

// ConfigApplyConfiguration represents an declarative configuration of the Config type for use
// with apply.
type ConfigApplyConfiguration struct {
	TargetGroupSelector *string `json:"targetGroupSelector,omitempty"`
	TargetSelector      *string `json:"targetSelector,omitempty"`
}

// ConfigApplyConfiguration constructs an declarative configuration of the Config type for use with
// apply.
func Config() *ConfigApplyConfiguration {
	return &ConfigApplyConfiguration{}
}

// WithTargetGroupSelector sets the TargetGroupSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetGroupSelector field is set to the value of the last call.
func (b *ConfigApplyConfiguration) WithTargetGroupSelector(value string) *ConfigApplyConfiguration {
	b.TargetGroupSelector = &value
	return b
}

// WithTargetSelector sets the TargetSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetSelector field is set to the value of the last call.
func (b *ConfigApplyConfiguration) WithTargetSelector(value string) *ConfigApplyConfiguration {
	b.TargetSelector = &value
	return b
}
