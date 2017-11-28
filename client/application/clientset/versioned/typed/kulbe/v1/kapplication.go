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

package v1

import (
	v1 "github.com/enablecloud/kulbe/apis/cr/application/v1"
	scheme "github.com/enablecloud/kulbe/client/application/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KApplicationsGetter has a method to return a KApplicationInterface.
// A group's client should implement this interface.
type KApplicationsGetter interface {
	KApplications(namespace string) KApplicationInterface
}

// KApplicationInterface has methods to work with KApplication resources.
type KApplicationInterface interface {
	Create(*v1.KApplication) (*v1.KApplication, error)
	Update(*v1.KApplication) (*v1.KApplication, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.KApplication, error)
	List(opts meta_v1.ListOptions) (*v1.KApplicationList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KApplication, err error)
	KApplicationExpansion
}

// kApplications implements KApplicationInterface
type kApplications struct {
	client rest.Interface
	ns     string
}

// newKApplications returns a KApplications
func newKApplications(c *KulbeV1Client, namespace string) *kApplications {
	return &kApplications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kApplication, and returns the corresponding kApplication object, and an error if there is any.
func (c *kApplications) Get(name string, options meta_v1.GetOptions) (result *v1.KApplication, err error) {
	result = &v1.KApplication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KApplications that match those selectors.
func (c *kApplications) List(opts meta_v1.ListOptions) (result *v1.KApplicationList, err error) {
	result = &v1.KApplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kApplications.
func (c *kApplications) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a kApplication and creates it.  Returns the server's representation of the kApplication, and an error, if there is any.
func (c *kApplications) Create(kApplication *v1.KApplication) (result *v1.KApplication, err error) {
	result = &v1.KApplication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kapplications").
		Body(kApplication).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kApplication and updates it. Returns the server's representation of the kApplication, and an error, if there is any.
func (c *kApplications) Update(kApplication *v1.KApplication) (result *v1.KApplication, err error) {
	result = &v1.KApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kapplications").
		Name(kApplication.Name).
		Body(kApplication).
		Do().
		Into(result)
	return
}

// Delete takes name of the kApplication and deletes it. Returns an error if one occurs.
func (c *kApplications) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kapplications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kApplications) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kapplications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kApplication.
func (c *kApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KApplication, err error) {
	result = &v1.KApplication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kapplications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
