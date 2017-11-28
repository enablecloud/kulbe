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

// KApplicationListsGetter has a method to return a KApplicationListInterface.
// A group's client should implement this interface.
type KApplicationListsGetter interface {
	KApplicationLists(namespace string) KApplicationListInterface
}

// KApplicationListInterface has methods to work with KApplicationList resources.
type KApplicationListInterface interface {
	Create(*v1.KApplicationList) (*v1.KApplicationList, error)
	Update(*v1.KApplicationList) (*v1.KApplicationList, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.KApplicationList, error)
	List(opts meta_v1.ListOptions) (*v1.KApplicationListList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KApplicationList, err error)
	KApplicationListExpansion
}

// kApplicationLists implements KApplicationListInterface
type kApplicationLists struct {
	client rest.Interface
	ns     string
}

// newKApplicationLists returns a KApplicationLists
func newKApplicationLists(c *KulbeV1Client, namespace string) *kApplicationLists {
	return &kApplicationLists{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kApplicationList, and returns the corresponding kApplicationList object, and an error if there is any.
func (c *kApplicationLists) Get(name string, options meta_v1.GetOptions) (result *v1.KApplicationList, err error) {
	result = &v1.KApplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kapplicationlists").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KApplicationLists that match those selectors.
func (c *kApplicationLists) List(opts meta_v1.ListOptions) (result *v1.KApplicationListList, err error) {
	result = &v1.KApplicationListList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kapplicationlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kApplicationLists.
func (c *kApplicationLists) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kapplicationlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a kApplicationList and creates it.  Returns the server's representation of the kApplicationList, and an error, if there is any.
func (c *kApplicationLists) Create(kApplicationList *v1.KApplicationList) (result *v1.KApplicationList, err error) {
	result = &v1.KApplicationList{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kapplicationlists").
		Body(kApplicationList).
		Do().
		Into(result)
	return
}

// Update takes the representation of a kApplicationList and updates it. Returns the server's representation of the kApplicationList, and an error, if there is any.
func (c *kApplicationLists) Update(kApplicationList *v1.KApplicationList) (result *v1.KApplicationList, err error) {
	result = &v1.KApplicationList{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kapplicationlists").
		Name(kApplicationList.Name).
		Body(kApplicationList).
		Do().
		Into(result)
	return
}

// Delete takes name of the kApplicationList and deletes it. Returns an error if one occurs.
func (c *kApplicationLists) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kapplicationlists").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kApplicationLists) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kapplicationlists").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kApplicationList.
func (c *kApplicationLists) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.KApplicationList, err error) {
	result = &v1.KApplicationList{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kapplicationlists").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
