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

// InstanceSpecApplyConfiguration represents a declarative configuration of the InstanceSpec type for use
// with apply.
type InstanceSpecApplyConfiguration struct {
	VolumeName         *string                                 `json:"volumeName,omitempty"`
	VolumeSize         *int64                                  `json:"volumeSize,omitempty"`
	NodeID             *string                                 `json:"nodeID,omitempty"`
	EngineImage        *string                                 `json:"engineImage,omitempty"`
	Image              *string                                 `json:"image,omitempty"`
	DesireState        *longhornv1beta2.InstanceState          `json:"desireState,omitempty"`
	LogRequested       *bool                                   `json:"logRequested,omitempty"`
	SalvageRequested   *bool                                   `json:"salvageRequested,omitempty"`
	BackendStoreDriver *longhornv1beta2.BackendStoreDriverType `json:"backendStoreDriver,omitempty"`
	DataEngine         *longhornv1beta2.DataEngineType         `json:"dataEngine,omitempty"`
}

// InstanceSpecApplyConfiguration constructs a declarative configuration of the InstanceSpec type for use with
// apply.
func InstanceSpec() *InstanceSpecApplyConfiguration {
	return &InstanceSpecApplyConfiguration{}
}

// WithVolumeName sets the VolumeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeName field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithVolumeName(value string) *InstanceSpecApplyConfiguration {
	b.VolumeName = &value
	return b
}

// WithVolumeSize sets the VolumeSize field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeSize field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithVolumeSize(value int64) *InstanceSpecApplyConfiguration {
	b.VolumeSize = &value
	return b
}

// WithNodeID sets the NodeID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeID field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithNodeID(value string) *InstanceSpecApplyConfiguration {
	b.NodeID = &value
	return b
}

// WithEngineImage sets the EngineImage field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EngineImage field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithEngineImage(value string) *InstanceSpecApplyConfiguration {
	b.EngineImage = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithImage(value string) *InstanceSpecApplyConfiguration {
	b.Image = &value
	return b
}

// WithDesireState sets the DesireState field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DesireState field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithDesireState(value longhornv1beta2.InstanceState) *InstanceSpecApplyConfiguration {
	b.DesireState = &value
	return b
}

// WithLogRequested sets the LogRequested field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogRequested field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithLogRequested(value bool) *InstanceSpecApplyConfiguration {
	b.LogRequested = &value
	return b
}

// WithSalvageRequested sets the SalvageRequested field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SalvageRequested field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithSalvageRequested(value bool) *InstanceSpecApplyConfiguration {
	b.SalvageRequested = &value
	return b
}

// WithBackendStoreDriver sets the BackendStoreDriver field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BackendStoreDriver field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithBackendStoreDriver(value longhornv1beta2.BackendStoreDriverType) *InstanceSpecApplyConfiguration {
	b.BackendStoreDriver = &value
	return b
}

// WithDataEngine sets the DataEngine field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DataEngine field is set to the value of the last call.
func (b *InstanceSpecApplyConfiguration) WithDataEngine(value longhornv1beta2.DataEngineType) *InstanceSpecApplyConfiguration {
	b.DataEngine = &value
	return b
}
