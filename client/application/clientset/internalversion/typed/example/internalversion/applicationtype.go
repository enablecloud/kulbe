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

package internalversion

import (
	application "github.com/enablecloud/kulbe/apis/cr/application"
	scheme "github.com/enablecloud/kulbe/client/application/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApplicationTypesGetter has a method to return a ApplicationTypeInterface.
// A group's client should implement this interface.
type ApplicationTypesGetter interface {
	ApplicationTypes(namespace string) ApplicationTypeInterface
}

// ApplicationTypeInterface has methods to work with ApplicationType resources.
type ApplicationTypeInterface interface {
	Create(*application.ApplicationType) (*application.ApplicationType, error)
	Update(*application.ApplicationType) (*application.ApplicationType, error)
	UpdateStatus(*application.ApplicationType) (*application.ApplicationType, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*application.ApplicationType, error)
	List(opts v1.ListOptions) (*application.ApplicationTypeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *application.ApplicationType, err error)
	ApplicationTypeExpansion
}

// applicationTypes implements ApplicationTypeInterface
type applicationTypes struct {
	client rest.Interface
	ns     string
}

// newApplicationTypes returns a ApplicationTypes
func newApplicationTypes(c *ExampleClient, namespace string) *applicationTypes {
	return &applicationTypes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the applicationType, and returns the corresponding applicationType object, and an error if there is any.
func (c *applicationTypes) Get(name string, options v1.GetOptions) (result *application.ApplicationType, err error) {
	result = &application.ApplicationType{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationtypes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApplicationTypes that match those selectors.
func (c *applicationTypes) List(opts v1.ListOptions) (result *application.ApplicationTypeList, err error) {
	result = &application.ApplicationTypeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationtypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested applicationTypes.
func (c *applicationTypes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("applicationtypes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a applicationType and creates it.  Returns the server's representation of the applicationType, and an error, if there is any.
func (c *applicationTypes) Create(applicationType *application.ApplicationType) (result *application.ApplicationType, err error) {
	result = &application.ApplicationType{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("applicationtypes").
		Body(applicationType).
		Do().
		Into(result)
	return
}

// Update takes the representation of a applicationType and updates it. Returns the server's representation of the applicationType, and an error, if there is any.
func (c *applicationTypes) Update(applicationType *application.ApplicationType) (result *application.ApplicationType, err error) {
	result = &application.ApplicationType{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationtypes").
		Name(applicationType.Name).
		Body(applicationType).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *applicationTypes) UpdateStatus(applicationType *application.ApplicationType) (result *application.ApplicationType, err error) {
	result = &application.ApplicationType{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationtypes").
		Name(applicationType.Name).
		SubResource("status").
		Body(applicationType).
		Do().
		Into(result)
	return
}

// Delete takes name of the applicationType and deletes it. Returns an error if one occurs.
func (c *applicationTypes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationtypes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *applicationTypes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationtypes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched applicationType.
func (c *applicationTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *application.ApplicationType, err error) {
	result = &application.ApplicationType{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("applicationtypes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
