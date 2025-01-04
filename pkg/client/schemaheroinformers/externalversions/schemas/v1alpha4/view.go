/*
Copyright 2021 The SchemaHero Authors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha4

import (
	"context"
	time "time"

	schemasv1alpha4 "github.com/schemahero/schemahero/pkg/apis/schemas/v1alpha4"
	schemaheroclientset "github.com/schemahero/schemahero/pkg/client/schemaheroclientset"
	internalinterfaces "github.com/schemahero/schemahero/pkg/client/schemaheroinformers/externalversions/internalinterfaces"
	v1alpha4 "github.com/schemahero/schemahero/pkg/client/schemaherolisters/schemas/v1alpha4"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ViewInformer provides access to a shared informer and lister for
// Views.
type ViewInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha4.ViewLister
}

type viewInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewViewInformer constructs a new informer for View type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewViewInformer(client schemaheroclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredViewInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredViewInformer constructs a new informer for View type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredViewInformer(client schemaheroclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchemasV1alpha4().Views(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchemasV1alpha4().Views(namespace).Watch(context.TODO(), options)
			},
		},
		&schemasv1alpha4.View{},
		resyncPeriod,
		indexers,
	)
}

func (f *viewInformer) defaultInformer(client schemaheroclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredViewInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *viewInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&schemasv1alpha4.View{}, f.defaultInformer)
}

func (f *viewInformer) Lister() v1alpha4.ViewLister {
	return v1alpha4.NewViewLister(f.Informer().GetIndexer())
}
