/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
	v1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AWSMachineClassLister helps list AWSMachineClasses.
type AWSMachineClassLister interface {
	// List lists all AWSMachineClasses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AWSMachineClass, err error)
	// AWSMachineClasses returns an object that can list and get AWSMachineClasses.
	AWSMachineClasses(namespace string) AWSMachineClassNamespaceLister
	AWSMachineClassListerExpansion
}

// aWSMachineClassLister implements the AWSMachineClassLister interface.
type aWSMachineClassLister struct {
	indexer cache.Indexer
}

// NewAWSMachineClassLister returns a new AWSMachineClassLister.
func NewAWSMachineClassLister(indexer cache.Indexer) AWSMachineClassLister {
	return &aWSMachineClassLister{indexer: indexer}
}

// List lists all AWSMachineClasses in the indexer.
func (s *aWSMachineClassLister) List(selector labels.Selector) (ret []*v1alpha1.AWSMachineClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AWSMachineClass))
	})
	return ret, err
}

// AWSMachineClasses returns an object that can list and get AWSMachineClasses.
func (s *aWSMachineClassLister) AWSMachineClasses(namespace string) AWSMachineClassNamespaceLister {
	return aWSMachineClassNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AWSMachineClassNamespaceLister helps list and get AWSMachineClasses.
type AWSMachineClassNamespaceLister interface {
	// List lists all AWSMachineClasses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AWSMachineClass, err error)
	// Get retrieves the AWSMachineClass from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AWSMachineClass, error)
	AWSMachineClassNamespaceListerExpansion
}

// aWSMachineClassNamespaceLister implements the AWSMachineClassNamespaceLister
// interface.
type aWSMachineClassNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AWSMachineClasses in the indexer for a given namespace.
func (s aWSMachineClassNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AWSMachineClass, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AWSMachineClass))
	})
	return ret, err
}

// Get retrieves the AWSMachineClass from the indexer for a given namespace and name.
func (s aWSMachineClassNamespaceLister) Get(name string) (*v1alpha1.AWSMachineClass, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("awsmachineclass"), name)
	}
	return obj.(*v1alpha1.AWSMachineClass), nil
}
