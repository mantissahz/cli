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

import (
	longhornv1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
)

// EngineImageStatusApplyConfiguration represents a declarative configuration of the EngineImageStatus type for use
// with apply.
type EngineImageStatusApplyConfiguration struct {
	OwnerID           *string                                `json:"ownerID,omitempty"`
	State             *longhornv1beta1.EngineImageState      `json:"state,omitempty"`
	RefCount          *int                                   `json:"refCount,omitempty"`
	NoRefSince        *string                                `json:"noRefSince,omitempty"`
	Conditions        map[string]ConditionApplyConfiguration `json:"conditions,omitempty"`
	NodeDeploymentMap map[string]bool                        `json:"nodeDeploymentMap,omitempty"`
}

// EngineImageStatusApplyConfiguration constructs a declarative configuration of the EngineImageStatus type for use with
// apply.
func EngineImageStatus() *EngineImageStatusApplyConfiguration {
	return &EngineImageStatusApplyConfiguration{}
}

// WithOwnerID sets the OwnerID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OwnerID field is set to the value of the last call.
func (b *EngineImageStatusApplyConfiguration) WithOwnerID(value string) *EngineImageStatusApplyConfiguration {
	b.OwnerID = &value
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *EngineImageStatusApplyConfiguration) WithState(value longhornv1beta1.EngineImageState) *EngineImageStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithRefCount sets the RefCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RefCount field is set to the value of the last call.
func (b *EngineImageStatusApplyConfiguration) WithRefCount(value int) *EngineImageStatusApplyConfiguration {
	b.RefCount = &value
	return b
}

// WithNoRefSince sets the NoRefSince field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NoRefSince field is set to the value of the last call.
func (b *EngineImageStatusApplyConfiguration) WithNoRefSince(value string) *EngineImageStatusApplyConfiguration {
	b.NoRefSince = &value
	return b
}

// WithConditions puts the entries into the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Conditions field,
// overwriting an existing map entries in Conditions field with the same key.
func (b *EngineImageStatusApplyConfiguration) WithConditions(entries map[string]ConditionApplyConfiguration) *EngineImageStatusApplyConfiguration {
	if b.Conditions == nil && len(entries) > 0 {
		b.Conditions = make(map[string]ConditionApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.Conditions[k] = v
	}
	return b
}

// WithNodeDeploymentMap puts the entries into the NodeDeploymentMap field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NodeDeploymentMap field,
// overwriting an existing map entries in NodeDeploymentMap field with the same key.
func (b *EngineImageStatusApplyConfiguration) WithNodeDeploymentMap(entries map[string]bool) *EngineImageStatusApplyConfiguration {
	if b.NodeDeploymentMap == nil && len(entries) > 0 {
		b.NodeDeploymentMap = make(map[string]bool, len(entries))
	}
	for k, v := range entries {
		b.NodeDeploymentMap[k] = v
	}
	return b
}
