/*
Copyright by suman sarker.
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

package v1

import (
	v1 "github.com/Sumon009838/CRD/pkg/apis/reader.com/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// BookStoreLister helps list BookStores.
// All objects returned here must be treated as read-only.
type BookStoreLister interface {
	// List lists all BookStores in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.BookStore, err error)
	// BookStores returns an object that can list and get BookStores.
	BookStores(namespace string) BookStoreNamespaceLister
	BookStoreListerExpansion
}

// bookStoreLister implements the BookStoreLister interface.
type bookStoreLister struct {
	listers.ResourceIndexer[*v1.BookStore]
}

// NewBookStoreLister returns a new BookStoreLister.
func NewBookStoreLister(indexer cache.Indexer) BookStoreLister {
	return &bookStoreLister{listers.New[*v1.BookStore](indexer, v1.Resource("bookstore"))}
}

// BookStores returns an object that can list and get BookStores.
func (s *bookStoreLister) BookStores(namespace string) BookStoreNamespaceLister {
	return bookStoreNamespaceLister{listers.NewNamespaced[*v1.BookStore](s.ResourceIndexer, namespace)}
}

// BookStoreNamespaceLister helps list and get BookStores.
// All objects returned here must be treated as read-only.
type BookStoreNamespaceLister interface {
	// List lists all BookStores in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.BookStore, err error)
	// Get retrieves the BookStore from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.BookStore, error)
	BookStoreNamespaceListerExpansion
}

// bookStoreNamespaceLister implements the BookStoreNamespaceLister
// interface.
type bookStoreNamespaceLister struct {
	listers.ResourceIndexer[*v1.BookStore]
}
