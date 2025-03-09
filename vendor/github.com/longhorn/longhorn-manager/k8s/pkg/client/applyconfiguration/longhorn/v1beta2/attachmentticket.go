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

// AttachmentTicketApplyConfiguration represents a declarative configuration of the AttachmentTicket type for use
// with apply.
type AttachmentTicketApplyConfiguration struct {
	ID         *string                       `json:"id,omitempty"`
	Type       *longhornv1beta2.AttacherType `json:"type,omitempty"`
	NodeID     *string                       `json:"nodeID,omitempty"`
	Parameters map[string]string             `json:"parameters,omitempty"`
	Generation *int64                        `json:"generation,omitempty"`
}

// AttachmentTicketApplyConfiguration constructs a declarative configuration of the AttachmentTicket type for use with
// apply.
func AttachmentTicket() *AttachmentTicketApplyConfiguration {
	return &AttachmentTicketApplyConfiguration{}
}

// WithID sets the ID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ID field is set to the value of the last call.
func (b *AttachmentTicketApplyConfiguration) WithID(value string) *AttachmentTicketApplyConfiguration {
	b.ID = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *AttachmentTicketApplyConfiguration) WithType(value longhornv1beta2.AttacherType) *AttachmentTicketApplyConfiguration {
	b.Type = &value
	return b
}

// WithNodeID sets the NodeID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeID field is set to the value of the last call.
func (b *AttachmentTicketApplyConfiguration) WithNodeID(value string) *AttachmentTicketApplyConfiguration {
	b.NodeID = &value
	return b
}

// WithParameters puts the entries into the Parameters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Parameters field,
// overwriting an existing map entries in Parameters field with the same key.
func (b *AttachmentTicketApplyConfiguration) WithParameters(entries map[string]string) *AttachmentTicketApplyConfiguration {
	if b.Parameters == nil && len(entries) > 0 {
		b.Parameters = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Parameters[k] = v
	}
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *AttachmentTicketApplyConfiguration) WithGeneration(value int64) *AttachmentTicketApplyConfiguration {
	b.Generation = &value
	return b
}
