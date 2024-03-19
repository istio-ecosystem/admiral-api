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

import (
	v1 "github.com/istio-ecosystem/admiral-api/pkg/apis/admiral/v1"
)

// EdgeServiceApplyConfiguration represents an declarative configuration of the EdgeService type for use
// with apply.
type EdgeServiceApplyConfiguration struct {
	DynamicRouting []*v1.DynamicRouting `json:"dynamicRouting,omitempty"`
	Filters        []*v1.Filter         `json:"filters,omitempty"`
	Routes         []*v1.Route          `json:"routes,omitempty"`
	Targets        []*v1.Target         `json:"targets,omitempty"`
	TargetGroups   []*v1.TargetGroup    `json:"targetGroups,omitempty"`
}

// EdgeServiceApplyConfiguration constructs an declarative configuration of the EdgeService type for use with
// apply.
func EdgeService() *EdgeServiceApplyConfiguration {
	return &EdgeServiceApplyConfiguration{}
}

// WithDynamicRouting adds the given value to the DynamicRouting field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the DynamicRouting field.
func (b *EdgeServiceApplyConfiguration) WithDynamicRouting(values ...**v1.DynamicRouting) *EdgeServiceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithDynamicRouting")
		}
		b.DynamicRouting = append(b.DynamicRouting, *values[i])
	}
	return b
}

// WithFilters adds the given value to the Filters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Filters field.
func (b *EdgeServiceApplyConfiguration) WithFilters(values ...**v1.Filter) *EdgeServiceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithFilters")
		}
		b.Filters = append(b.Filters, *values[i])
	}
	return b
}

// WithRoutes adds the given value to the Routes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Routes field.
func (b *EdgeServiceApplyConfiguration) WithRoutes(values ...**v1.Route) *EdgeServiceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRoutes")
		}
		b.Routes = append(b.Routes, *values[i])
	}
	return b
}

// WithTargets adds the given value to the Targets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Targets field.
func (b *EdgeServiceApplyConfiguration) WithTargets(values ...**v1.Target) *EdgeServiceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTargets")
		}
		b.Targets = append(b.Targets, *values[i])
	}
	return b
}

// WithTargetGroups adds the given value to the TargetGroups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TargetGroups field.
func (b *EdgeServiceApplyConfiguration) WithTargetGroups(values ...**v1.TargetGroup) *EdgeServiceApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTargetGroups")
		}
		b.TargetGroups = append(b.TargetGroups, *values[i])
	}
	return b
}
