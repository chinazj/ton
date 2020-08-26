/*
Copyright 2020 The Krole Authors

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

package v1alpha1

import (
	"context"
	time "time"
	tonv1alpha1 "ton/pkg/apis/ton/v1alpha1"
	versioned "ton/pkg/client/clientset/versioned"
	internalinterfaces "ton/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "ton/pkg/client/listers/ton/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LineInformer provides access to a shared informer and lister for
// Lines.
type LineInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.LineLister
}

type lineInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLineInformer constructs a new informer for Line type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLineInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLineInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLineInformer constructs a new informer for Line type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLineInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TonV1alpha1().Lines(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TonV1alpha1().Lines(namespace).Watch(context.TODO(), options)
			},
		},
		&tonv1alpha1.Line{},
		resyncPeriod,
		indexers,
	)
}

func (f *lineInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLineInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *lineInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tonv1alpha1.Line{}, f.defaultInformer)
}

func (f *lineInformer) Lister() v1alpha1.LineLister {
	return v1alpha1.NewLineLister(f.Informer().GetIndexer())
}
