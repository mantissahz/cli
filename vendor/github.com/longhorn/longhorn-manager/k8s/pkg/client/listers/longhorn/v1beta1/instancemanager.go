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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	longhornv1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// InstanceManagerLister helps list InstanceManagers.
// All objects returned here must be treated as read-only.
type InstanceManagerLister interface {
	// List lists all InstanceManagers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*longhornv1beta1.InstanceManager, err error)
	// InstanceManagers returns an object that can list and get InstanceManagers.
	InstanceManagers(namespace string) InstanceManagerNamespaceLister
	InstanceManagerListerExpansion
}

// instanceManagerLister implements the InstanceManagerLister interface.
type instanceManagerLister struct {
	listers.ResourceIndexer[*longhornv1beta1.InstanceManager]
}

// NewInstanceManagerLister returns a new InstanceManagerLister.
func NewInstanceManagerLister(indexer cache.Indexer) InstanceManagerLister {
	return &instanceManagerLister{listers.New[*longhornv1beta1.InstanceManager](indexer, longhornv1beta1.Resource("instancemanager"))}
}

// InstanceManagers returns an object that can list and get InstanceManagers.
func (s *instanceManagerLister) InstanceManagers(namespace string) InstanceManagerNamespaceLister {
	return instanceManagerNamespaceLister{listers.NewNamespaced[*longhornv1beta1.InstanceManager](s.ResourceIndexer, namespace)}
}

// InstanceManagerNamespaceLister helps list and get InstanceManagers.
// All objects returned here must be treated as read-only.
type InstanceManagerNamespaceLister interface {
	// List lists all InstanceManagers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*longhornv1beta1.InstanceManager, err error)
	// Get retrieves the InstanceManager from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*longhornv1beta1.InstanceManager, error)
	InstanceManagerNamespaceListerExpansion
}

// instanceManagerNamespaceLister implements the InstanceManagerNamespaceLister
// interface.
type instanceManagerNamespaceLister struct {
	listers.ResourceIndexer[*longhornv1beta1.InstanceManager]
}
