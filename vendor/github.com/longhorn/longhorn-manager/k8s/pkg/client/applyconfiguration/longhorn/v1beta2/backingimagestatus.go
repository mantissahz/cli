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

// BackingImageStatusApplyConfiguration represents a declarative configuration of the BackingImageStatus type for use
// with apply.
type BackingImageStatusApplyConfiguration struct {
	OwnerID           *string                                                `json:"ownerID,omitempty"`
	UUID              *string                                                `json:"uuid,omitempty"`
	Size              *int64                                                 `json:"size,omitempty"`
	VirtualSize       *int64                                                 `json:"virtualSize,omitempty"`
	RealSize          *int64                                                 `json:"realSize,omitempty"`
	Checksum          *string                                                `json:"checksum,omitempty"`
	DiskFileStatusMap map[string]*longhornv1beta2.BackingImageDiskFileStatus `json:"diskFileStatusMap,omitempty"`
	DiskLastRefAtMap  map[string]string                                      `json:"diskLastRefAtMap,omitempty"`
	V2FirstCopyStatus *longhornv1beta2.BackingImageState                     `json:"v2FirstCopyStatus,omitempty"`
	V2FirstCopyDisk   *string                                                `json:"v2FirstCopyDisk,omitempty"`
}

// BackingImageStatusApplyConfiguration constructs a declarative configuration of the BackingImageStatus type for use with
// apply.
func BackingImageStatus() *BackingImageStatusApplyConfiguration {
	return &BackingImageStatusApplyConfiguration{}
}

// WithOwnerID sets the OwnerID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OwnerID field is set to the value of the last call.
func (b *BackingImageStatusApplyConfiguration) WithOwnerID(value string) *BackingImageStatusApplyConfiguration {
	b.OwnerID = &value
	return b
}

// WithUUID sets the UUID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UUID field is set to the value of the last call.
func (b *BackingImageStatusApplyConfiguration) WithUUID(value string) *BackingImageStatusApplyConfiguration {
	b.UUID = &value
	return b
}

// WithSize sets the Size field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Size field is set to the value of the last call.
func (b *BackingImageStatusApplyConfiguration) WithSize(value int64) *BackingImageStatusApplyConfiguration {
	b.Size = &value
	return b
}

// WithVirtualSize sets the VirtualSize field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VirtualSize field is set to the value of the last call.
func (b *BackingImageStatusApplyConfiguration) WithVirtualSize(value int64) *BackingImageStatusApplyConfiguration {
	b.VirtualSize = &value
	return b
}

// WithRealSize sets the RealSize field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RealSize field is set to the value of the last call.
func (b *BackingImageStatusApplyConfiguration) WithRealSize(value int64) *BackingImageStatusApplyConfiguration {
	b.RealSize = &value
	return b
}

// WithChecksum sets the Checksum field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Checksum field is set to the value of the last call.
func (b *BackingImageStatusApplyConfiguration) WithChecksum(value string) *BackingImageStatusApplyConfiguration {
	b.Checksum = &value
	return b
}

// WithDiskFileStatusMap puts the entries into the DiskFileStatusMap field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the DiskFileStatusMap field,
// overwriting an existing map entries in DiskFileStatusMap field with the same key.
func (b *BackingImageStatusApplyConfiguration) WithDiskFileStatusMap(entries map[string]*longhornv1beta2.BackingImageDiskFileStatus) *BackingImageStatusApplyConfiguration {
	if b.DiskFileStatusMap == nil && len(entries) > 0 {
		b.DiskFileStatusMap = make(map[string]*longhornv1beta2.BackingImageDiskFileStatus, len(entries))
	}
	for k, v := range entries {
		b.DiskFileStatusMap[k] = v
	}
	return b
}

// WithDiskLastRefAtMap puts the entries into the DiskLastRefAtMap field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the DiskLastRefAtMap field,
// overwriting an existing map entries in DiskLastRefAtMap field with the same key.
func (b *BackingImageStatusApplyConfiguration) WithDiskLastRefAtMap(entries map[string]string) *BackingImageStatusApplyConfiguration {
	if b.DiskLastRefAtMap == nil && len(entries) > 0 {
		b.DiskLastRefAtMap = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.DiskLastRefAtMap[k] = v
	}
	return b
}

// WithV2FirstCopyStatus sets the V2FirstCopyStatus field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V2FirstCopyStatus field is set to the value of the last call.
func (b *BackingImageStatusApplyConfiguration) WithV2FirstCopyStatus(value longhornv1beta2.BackingImageState) *BackingImageStatusApplyConfiguration {
	b.V2FirstCopyStatus = &value
	return b
}

// WithV2FirstCopyDisk sets the V2FirstCopyDisk field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the V2FirstCopyDisk field is set to the value of the last call.
func (b *BackingImageStatusApplyConfiguration) WithV2FirstCopyDisk(value string) *BackingImageStatusApplyConfiguration {
	b.V2FirstCopyDisk = &value
	return b
}
