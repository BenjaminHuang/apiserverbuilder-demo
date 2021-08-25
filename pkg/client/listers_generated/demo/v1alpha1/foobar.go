/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "steamingmind.com/apiserver-demo/pkg/apis/demo/v1alpha1"
)

// FooBarLister helps list FooBars.
type FooBarLister interface {
	// List lists all FooBars in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FooBar, err error)
	// FooBars returns an object that can list and get FooBars.
	FooBars(namespace string) FooBarNamespaceLister
	FooBarListerExpansion
}

// fooBarLister implements the FooBarLister interface.
type fooBarLister struct {
	indexer cache.Indexer
}

// NewFooBarLister returns a new FooBarLister.
func NewFooBarLister(indexer cache.Indexer) FooBarLister {
	return &fooBarLister{indexer: indexer}
}

// List lists all FooBars in the indexer.
func (s *fooBarLister) List(selector labels.Selector) (ret []*v1alpha1.FooBar, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FooBar))
	})
	return ret, err
}

// FooBars returns an object that can list and get FooBars.
func (s *fooBarLister) FooBars(namespace string) FooBarNamespaceLister {
	return fooBarNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FooBarNamespaceLister helps list and get FooBars.
type FooBarNamespaceLister interface {
	// List lists all FooBars in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FooBar, err error)
	// Get retrieves the FooBar from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FooBar, error)
	FooBarNamespaceListerExpansion
}

// fooBarNamespaceLister implements the FooBarNamespaceLister
// interface.
type fooBarNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FooBars in the indexer for a given namespace.
func (s fooBarNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FooBar, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FooBar))
	})
	return ret, err
}

// Get retrieves the FooBar from the indexer for a given namespace and name.
func (s fooBarNamespaceLister) Get(name string) (*v1alpha1.FooBar, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("foobar"), name)
	}
	return obj.(*v1alpha1.FooBar), nil
}
