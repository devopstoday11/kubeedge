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

package v1beta1

import (
	"context"
	time "time"

	devicesv1beta1 "github.com/kubeedge/api/apis/devices/v1beta1"
	versioned "github.com/kubeedge/kubeedge/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubeedge/kubeedge/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/kubeedge/kubeedge/pkg/client/listers/devices/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DeviceModelInformer provides access to a shared informer and lister for
// DeviceModels.
type DeviceModelInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.DeviceModelLister
}

type deviceModelInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDeviceModelInformer constructs a new informer for DeviceModel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDeviceModelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDeviceModelInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDeviceModelInformer constructs a new informer for DeviceModel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDeviceModelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DevicesV1beta1().DeviceModels(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DevicesV1beta1().DeviceModels(namespace).Watch(context.TODO(), options)
			},
		},
		&devicesv1beta1.DeviceModel{},
		resyncPeriod,
		indexers,
	)
}

func (f *deviceModelInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDeviceModelInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *deviceModelInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&devicesv1beta1.DeviceModel{}, f.defaultInformer)
}

func (f *deviceModelInformer) Lister() v1beta1.DeviceModelLister {
	return v1beta1.NewDeviceModelLister(f.Informer().GetIndexer())
}
