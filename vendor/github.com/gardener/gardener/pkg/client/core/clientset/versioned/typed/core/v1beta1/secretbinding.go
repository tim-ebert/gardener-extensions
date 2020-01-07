/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
	"time"

	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	scheme "github.com/gardener/gardener/pkg/client/core/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SecretBindingsGetter has a method to return a SecretBindingInterface.
// A group's client should implement this interface.
type SecretBindingsGetter interface {
	SecretBindings(namespace string) SecretBindingInterface
}

// SecretBindingInterface has methods to work with SecretBinding resources.
type SecretBindingInterface interface {
	Create(*v1beta1.SecretBinding) (*v1beta1.SecretBinding, error)
	Update(*v1beta1.SecretBinding) (*v1beta1.SecretBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.SecretBinding, error)
	List(opts v1.ListOptions) (*v1beta1.SecretBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.SecretBinding, err error)
	SecretBindingExpansion
}

// secretBindings implements SecretBindingInterface
type secretBindings struct {
	client rest.Interface
	ns     string
}

// newSecretBindings returns a SecretBindings
func newSecretBindings(c *CoreV1beta1Client, namespace string) *secretBindings {
	return &secretBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the secretBinding, and returns the corresponding secretBinding object, and an error if there is any.
func (c *secretBindings) Get(name string, options v1.GetOptions) (result *v1beta1.SecretBinding, err error) {
	result = &v1beta1.SecretBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secretbindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecretBindings that match those selectors.
func (c *secretBindings) List(opts v1.ListOptions) (result *v1beta1.SecretBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.SecretBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secretbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested secretBindings.
func (c *secretBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("secretbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a secretBinding and creates it.  Returns the server's representation of the secretBinding, and an error, if there is any.
func (c *secretBindings) Create(secretBinding *v1beta1.SecretBinding) (result *v1beta1.SecretBinding, err error) {
	result = &v1beta1.SecretBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("secretbindings").
		Body(secretBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a secretBinding and updates it. Returns the server's representation of the secretBinding, and an error, if there is any.
func (c *secretBindings) Update(secretBinding *v1beta1.SecretBinding) (result *v1beta1.SecretBinding, err error) {
	result = &v1beta1.SecretBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("secretbindings").
		Name(secretBinding.Name).
		Body(secretBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the secretBinding and deletes it. Returns an error if one occurs.
func (c *secretBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secretbindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *secretBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secretbindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched secretBinding.
func (c *secretBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.SecretBinding, err error) {
	result = &v1beta1.SecretBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("secretbindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}