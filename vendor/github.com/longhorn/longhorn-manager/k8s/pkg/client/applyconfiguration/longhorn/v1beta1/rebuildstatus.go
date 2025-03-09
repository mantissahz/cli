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

package v1beta1

// RebuildStatusApplyConfiguration represents a declarative configuration of the RebuildStatus type for use
// with apply.
type RebuildStatusApplyConfiguration struct {
	Error              *string `json:"error,omitempty"`
	IsRebuilding       *bool   `json:"isRebuilding,omitempty"`
	Progress           *int    `json:"progress,omitempty"`
	State              *string `json:"state,omitempty"`
	FromReplicaAddress *string `json:"fromReplicaAddress,omitempty"`
}

// RebuildStatusApplyConfiguration constructs a declarative configuration of the RebuildStatus type for use with
// apply.
func RebuildStatus() *RebuildStatusApplyConfiguration {
	return &RebuildStatusApplyConfiguration{}
}

// WithError sets the Error field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Error field is set to the value of the last call.
func (b *RebuildStatusApplyConfiguration) WithError(value string) *RebuildStatusApplyConfiguration {
	b.Error = &value
	return b
}

// WithIsRebuilding sets the IsRebuilding field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IsRebuilding field is set to the value of the last call.
func (b *RebuildStatusApplyConfiguration) WithIsRebuilding(value bool) *RebuildStatusApplyConfiguration {
	b.IsRebuilding = &value
	return b
}

// WithProgress sets the Progress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Progress field is set to the value of the last call.
func (b *RebuildStatusApplyConfiguration) WithProgress(value int) *RebuildStatusApplyConfiguration {
	b.Progress = &value
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *RebuildStatusApplyConfiguration) WithState(value string) *RebuildStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithFromReplicaAddress sets the FromReplicaAddress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FromReplicaAddress field is set to the value of the last call.
func (b *RebuildStatusApplyConfiguration) WithFromReplicaAddress(value string) *RebuildStatusApplyConfiguration {
	b.FromReplicaAddress = &value
	return b
}
