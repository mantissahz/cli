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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"

	longhornv1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	applyconfigurationlonghornv1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/client/applyconfiguration/longhorn/v1beta1"
	scheme "github.com/longhorn/longhorn-manager/k8s/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// BackingImagesGetter has a method to return a BackingImageInterface.
// A group's client should implement this interface.
type BackingImagesGetter interface {
	BackingImages(namespace string) BackingImageInterface
}

// BackingImageInterface has methods to work with BackingImage resources.
type BackingImageInterface interface {
	Create(ctx context.Context, backingImage *longhornv1beta1.BackingImage, opts v1.CreateOptions) (*longhornv1beta1.BackingImage, error)
	Update(ctx context.Context, backingImage *longhornv1beta1.BackingImage, opts v1.UpdateOptions) (*longhornv1beta1.BackingImage, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, backingImage *longhornv1beta1.BackingImage, opts v1.UpdateOptions) (*longhornv1beta1.BackingImage, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*longhornv1beta1.BackingImage, error)
	List(ctx context.Context, opts v1.ListOptions) (*longhornv1beta1.BackingImageList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *longhornv1beta1.BackingImage, err error)
	Apply(ctx context.Context, backingImage *applyconfigurationlonghornv1beta1.BackingImageApplyConfiguration, opts v1.ApplyOptions) (result *longhornv1beta1.BackingImage, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, backingImage *applyconfigurationlonghornv1beta1.BackingImageApplyConfiguration, opts v1.ApplyOptions) (result *longhornv1beta1.BackingImage, err error)
	BackingImageExpansion
}

// backingImages implements BackingImageInterface
type backingImages struct {
	*gentype.ClientWithListAndApply[*longhornv1beta1.BackingImage, *longhornv1beta1.BackingImageList, *applyconfigurationlonghornv1beta1.BackingImageApplyConfiguration]
}

// newBackingImages returns a BackingImages
func newBackingImages(c *LonghornV1beta1Client, namespace string) *backingImages {
	return &backingImages{
		gentype.NewClientWithListAndApply[*longhornv1beta1.BackingImage, *longhornv1beta1.BackingImageList, *applyconfigurationlonghornv1beta1.BackingImageApplyConfiguration](
			"backingimages",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *longhornv1beta1.BackingImage { return &longhornv1beta1.BackingImage{} },
			func() *longhornv1beta1.BackingImageList { return &longhornv1beta1.BackingImageList{} },
		),
	}
}
