/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	application_v1 "github.com/enablecloud/kulbe/apis/cr/application/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKApplications implements KApplicationInterface
type FakeKApplications struct {
	Fake *FakeKulbeV1
	ns   string
}

var kapplicationsResource = schema.GroupVersionResource{Group: "kulbe.enablecloud.github.com", Version: "v1", Resource: "kapplications"}

var kapplicationsKind = schema.GroupVersionKind{Group: "kulbe.enablecloud.github.com", Version: "v1", Kind: "KApplication"}

// Get takes name of the kApplication, and returns the corresponding kApplication object, and an error if there is any.
func (c *FakeKApplications) Get(name string, options v1.GetOptions) (result *application_v1.KApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kapplicationsResource, c.ns, name), &application_v1.KApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*application_v1.KApplication), err
}

// List takes label and field selectors, and returns the list of KApplications that match those selectors.
func (c *FakeKApplications) List(opts v1.ListOptions) (result *application_v1.KApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kapplicationsResource, kapplicationsKind, c.ns, opts), &application_v1.KApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &application_v1.KApplicationList{}
	for _, item := range obj.(*application_v1.KApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kApplications.
func (c *FakeKApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kapplicationsResource, c.ns, opts))

}

// Create takes the representation of a kApplication and creates it.  Returns the server's representation of the kApplication, and an error, if there is any.
func (c *FakeKApplications) Create(kApplication *application_v1.KApplication) (result *application_v1.KApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kapplicationsResource, c.ns, kApplication), &application_v1.KApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*application_v1.KApplication), err
}

// Update takes the representation of a kApplication and updates it. Returns the server's representation of the kApplication, and an error, if there is any.
func (c *FakeKApplications) Update(kApplication *application_v1.KApplication) (result *application_v1.KApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kapplicationsResource, c.ns, kApplication), &application_v1.KApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*application_v1.KApplication), err
}

// Delete takes name of the kApplication and deletes it. Returns an error if one occurs.
func (c *FakeKApplications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kapplicationsResource, c.ns, name), &application_v1.KApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kapplicationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &application_v1.KApplicationList{})
	return err
}

// Patch applies the patch and returns the patched kApplication.
func (c *FakeKApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *application_v1.KApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kapplicationsResource, c.ns, name, data, subresources...), &application_v1.KApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*application_v1.KApplication), err
}
