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
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKApplicationLists implements KApplicationListInterface
type FakeKApplicationLists struct {
	Fake *FakeKulbeV1
	ns   string
}

var kapplicationlistsResource = schema.GroupVersionResource{Group: "kulbe.enablecloud.github.com", Version: "v1", Resource: "kapplicationlists"}

var kapplicationlistsKind = schema.GroupVersionKind{Group: "kulbe.enablecloud.github.com", Version: "v1", Kind: "KApplicationList"}

// Get takes name of the kApplicationList, and returns the corresponding kApplicationList object, and an error if there is any.
func (c *FakeKApplicationLists) Get(name string, options v1.GetOptions) (result *application_v1.KApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kapplicationlistsResource, c.ns, name), &application_v1.KApplicationList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*application_v1.KApplicationList), err
}

// List takes label and field selectors, and returns the list of KApplicationLists that match those selectors.
func (c *FakeKApplicationLists) List(opts v1.ListOptions) (result *application_v1.KApplicationListList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kapplicationlistsResource, kapplicationlistsKind, c.ns, opts), &application_v1.KApplicationListList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*application_v1.KApplicationListList), err
}

// Watch returns a watch.Interface that watches the requested kApplicationLists.
func (c *FakeKApplicationLists) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kapplicationlistsResource, c.ns, opts))

}

// Create takes the representation of a kApplicationList and creates it.  Returns the server's representation of the kApplicationList, and an error, if there is any.
func (c *FakeKApplicationLists) Create(kApplicationList *application_v1.KApplicationList) (result *application_v1.KApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kapplicationlistsResource, c.ns, kApplicationList), &application_v1.KApplicationList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*application_v1.KApplicationList), err
}

// Update takes the representation of a kApplicationList and updates it. Returns the server's representation of the kApplicationList, and an error, if there is any.
func (c *FakeKApplicationLists) Update(kApplicationList *application_v1.KApplicationList) (result *application_v1.KApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kapplicationlistsResource, c.ns, kApplicationList), &application_v1.KApplicationList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*application_v1.KApplicationList), err
}

// Delete takes name of the kApplicationList and deletes it. Returns an error if one occurs.
func (c *FakeKApplicationLists) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kapplicationlistsResource, c.ns, name), &application_v1.KApplicationList{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKApplicationLists) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kapplicationlistsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &application_v1.KApplicationListList{})
	return err
}

// Patch applies the patch and returns the patched kApplicationList.
func (c *FakeKApplicationLists) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *application_v1.KApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kapplicationlistsResource, c.ns, name, data, subresources...), &application_v1.KApplicationList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*application_v1.KApplicationList), err
}
