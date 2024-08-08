/*
Copyright The KubeEdge Authors.

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

	appsv1alpha1 "github.com/kubeedge/api/apis/apps/v1alpha1"
	versioned "github.com/kubeedge/kubeedge/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubeedge/kubeedge/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubeedge/kubeedge/pkg/client/listers/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NodeGroupInformer provides access to a shared informer and lister for
// NodeGroups.
type NodeGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.NodeGroupLister
}

type nodeGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNodeGroupInformer constructs a new informer for NodeGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNodeGroupInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNodeGroupInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNodeGroupInformer constructs a new informer for NodeGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNodeGroupInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1alpha1().NodeGroups().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1alpha1().NodeGroups().Watch(context.TODO(), options)
			},
		},
		&appsv1alpha1.NodeGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *nodeGroupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNodeGroupInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *nodeGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appsv1alpha1.NodeGroup{}, f.defaultInformer)
}

func (f *nodeGroupInformer) Lister() v1alpha1.NodeGroupLister {
	return v1alpha1.NewNodeGroupLister(f.Informer().GetIndexer())
}
