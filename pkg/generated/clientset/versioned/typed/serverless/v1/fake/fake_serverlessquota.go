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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/SUMMERLm/serverlessquota/pkg/apis/serverless/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServerlessquotas implements ServerlessquotaInterface
type FakeServerlessquotas struct {
	Fake *FakeServerlessV1
	ns   string
}

var serverlessquotasResource = v1.SchemeGroupVersion.WithResource("serverlessquotas")

var serverlessquotasKind = v1.SchemeGroupVersion.WithKind("Serverlessquota")

// Get takes name of the serverlessquota, and returns the corresponding serverlessquota object, and an error if there is any.
func (c *FakeServerlessquotas) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Serverlessquota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serverlessquotasResource, c.ns, name), &v1.Serverlessquota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Serverlessquota), err
}

// List takes label and field selectors, and returns the list of Serverlessquotas that match those selectors.
func (c *FakeServerlessquotas) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ServerlessquotaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serverlessquotasResource, serverlessquotasKind, c.ns, opts), &v1.ServerlessquotaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ServerlessquotaList{ListMeta: obj.(*v1.ServerlessquotaList).ListMeta}
	for _, item := range obj.(*v1.ServerlessquotaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serverlessquotas.
func (c *FakeServerlessquotas) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serverlessquotasResource, c.ns, opts))

}

// Create takes the representation of a serverlessquota and creates it.  Returns the server's representation of the serverlessquota, and an error, if there is any.
func (c *FakeServerlessquotas) Create(ctx context.Context, serverlessquota *v1.Serverlessquota, opts metav1.CreateOptions) (result *v1.Serverlessquota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serverlessquotasResource, c.ns, serverlessquota), &v1.Serverlessquota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Serverlessquota), err
}

// Update takes the representation of a serverlessquota and updates it. Returns the server's representation of the serverlessquota, and an error, if there is any.
func (c *FakeServerlessquotas) Update(ctx context.Context, serverlessquota *v1.Serverlessquota, opts metav1.UpdateOptions) (result *v1.Serverlessquota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serverlessquotasResource, c.ns, serverlessquota), &v1.Serverlessquota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Serverlessquota), err
}

// Delete takes name of the serverlessquota and deletes it. Returns an error if one occurs.
func (c *FakeServerlessquotas) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(serverlessquotasResource, c.ns, name, opts), &v1.Serverlessquota{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServerlessquotas) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serverlessquotasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ServerlessquotaList{})
	return err
}

// Patch applies the patch and returns the patched serverlessquota.
func (c *FakeServerlessquotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Serverlessquota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serverlessquotasResource, c.ns, name, pt, data, subresources...), &v1.Serverlessquota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Serverlessquota), err
}