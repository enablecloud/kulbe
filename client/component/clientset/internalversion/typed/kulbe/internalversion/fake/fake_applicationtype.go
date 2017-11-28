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
	component "github.com/enablecloud/kulbe/apis/cr/component"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplicationTypes implements ApplicationTypeInterface
type FakeApplicationTypes struct {
	Fake *FakeKulbe
	ns   string
}

var applicationtypesResource = schema.GroupVersionResource{Group: "kulbe.enablecloud.github.com", Version: "", Resource: "applicationtypes"}

var applicationtypesKind = schema.GroupVersionKind{Group: "kulbe.enablecloud.github.com", Version: "", Kind: "ApplicationType"}

// Get takes name of the applicationType, and returns the corresponding applicationType object, and an error if there is any.
func (c *FakeApplicationTypes) Get(name string, options v1.GetOptions) (result *component.ApplicationType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationtypesResource, c.ns, name), &component.ApplicationType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*component.ApplicationType), err
}

// List takes label and field selectors, and returns the list of ApplicationTypes that match those selectors.
func (c *FakeApplicationTypes) List(opts v1.ListOptions) (result *component.ApplicationTypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationtypesResource, applicationtypesKind, c.ns, opts), &component.ApplicationTypeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &component.ApplicationTypeList{}
	for _, item := range obj.(*component.ApplicationTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applicationTypes.
func (c *FakeApplicationTypes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationtypesResource, c.ns, opts))

}

// Create takes the representation of a applicationType and creates it.  Returns the server's representation of the applicationType, and an error, if there is any.
func (c *FakeApplicationTypes) Create(applicationType *component.ApplicationType) (result *component.ApplicationType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationtypesResource, c.ns, applicationType), &component.ApplicationType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*component.ApplicationType), err
}

// Update takes the representation of a applicationType and updates it. Returns the server's representation of the applicationType, and an error, if there is any.
func (c *FakeApplicationTypes) Update(applicationType *component.ApplicationType) (result *component.ApplicationType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationtypesResource, c.ns, applicationType), &component.ApplicationType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*component.ApplicationType), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApplicationTypes) UpdateStatus(applicationType *component.ApplicationType) (*component.ApplicationType, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(applicationtypesResource, "status", c.ns, applicationType), &component.ApplicationType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*component.ApplicationType), err
}

// Delete takes name of the applicationType and deletes it. Returns an error if one occurs.
func (c *FakeApplicationTypes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(applicationtypesResource, c.ns, name), &component.ApplicationType{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplicationTypes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationtypesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &component.ApplicationTypeList{})
	return err
}

// Patch applies the patch and returns the patched applicationType.
func (c *FakeApplicationTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *component.ApplicationType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationtypesResource, c.ns, name, data, subresources...), &component.ApplicationType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*component.ApplicationType), err
}
