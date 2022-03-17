/*
Copyright 2020 The Helm Broker Authors.

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
	v1alpha1 "github.com/kyma-project/helm-broker/pkg/apis/addons/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AddonsConfigurationLister helps list AddonsConfigurations.
type AddonsConfigurationLister interface {
	// List lists all AddonsConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AddonsConfiguration, err error)
	// AddonsConfigurations returns an object that can list and get AddonsConfigurations.
	AddonsConfigurations(namespace string) AddonsConfigurationNamespaceLister
	AddonsConfigurationListerExpansion
}

// addonsConfigurationLister implements the AddonsConfigurationLister interface.
type addonsConfigurationLister struct {
	indexer cache.Indexer
}

// NewAddonsConfigurationLister returns a new AddonsConfigurationLister.
func NewAddonsConfigurationLister(indexer cache.Indexer) AddonsConfigurationLister {
	return &addonsConfigurationLister{indexer: indexer}
}

// List lists all AddonsConfigurations in the indexer.
func (s *addonsConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.AddonsConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AddonsConfiguration))
	})
	return ret, err
}

// AddonsConfigurations returns an object that can list and get AddonsConfigurations.
func (s *addonsConfigurationLister) AddonsConfigurations(namespace string) AddonsConfigurationNamespaceLister {
	return addonsConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AddonsConfigurationNamespaceLister helps list and get AddonsConfigurations.
type AddonsConfigurationNamespaceLister interface {
	// List lists all AddonsConfigurations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AddonsConfiguration, err error)
	// Get retrieves the AddonsConfiguration from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AddonsConfiguration, error)
	AddonsConfigurationNamespaceListerExpansion
}

// addonsConfigurationNamespaceLister implements the AddonsConfigurationNamespaceLister
// interface.
type addonsConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AddonsConfigurations in the indexer for a given namespace.
func (s addonsConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AddonsConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AddonsConfiguration))
	})
	return ret, err
}

// Get retrieves the AddonsConfiguration from the indexer for a given namespace and name.
func (s addonsConfigurationNamespaceLister) Get(name string) (*v1alpha1.AddonsConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("addonsconfiguration"), name)
	}
	return obj.(*v1alpha1.AddonsConfiguration), nil
}