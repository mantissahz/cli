/*
Copyright The Longhorn Authors.

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

package v1beta2

import (
	longhornv1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
)

// InstanceProcessSpecApplyConfiguration represents a declarative configuration of the InstanceProcessSpec type for use
// with apply.
type InstanceProcessSpecApplyConfiguration struct {
	Name               *string                                 `json:"name,omitempty"`
	BackendStoreDriver *longhornv1beta2.BackendStoreDriverType `json:"backendStoreDriver,omitempty"`
	DataEngine         *longhornv1beta2.DataEngineType         `json:"dataEngine,omitempty"`
}

// InstanceProcessSpecApplyConfiguration constructs a declarative configuration of the InstanceProcessSpec type for use with
// apply.
func InstanceProcessSpec() *InstanceProcessSpecApplyConfiguration {
	return &InstanceProcessSpecApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *InstanceProcessSpecApplyConfiguration) WithName(value string) *InstanceProcessSpecApplyConfiguration {
	b.Name = &value
	return b
}

// WithBackendStoreDriver sets the BackendStoreDriver field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BackendStoreDriver field is set to the value of the last call.
func (b *InstanceProcessSpecApplyConfiguration) WithBackendStoreDriver(value longhornv1beta2.BackendStoreDriverType) *InstanceProcessSpecApplyConfiguration {
	b.BackendStoreDriver = &value
	return b
}

// WithDataEngine sets the DataEngine field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DataEngine field is set to the value of the last call.
func (b *InstanceProcessSpecApplyConfiguration) WithDataEngine(value longhornv1beta2.DataEngineType) *InstanceProcessSpecApplyConfiguration {
	b.DataEngine = &value
	return b
}
