// This file was automatically generated by informer-gen

package v1alpha1

import (
	cr_v1alpha1 "github.com/kanisterio/kanister/pkg/apis/cr/v1alpha1"
	versioned "github.com/kanisterio/kanister/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kanisterio/kanister/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kanisterio/kanister/pkg/client/listers/cr/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ActionSetInformer provides access to a shared informer and lister for
// ActionSets.
type ActionSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ActionSetLister
}

type actionSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewActionSetInformer constructs a new informer for ActionSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewActionSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredActionSetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredActionSetInformer constructs a new informer for ActionSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredActionSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrV1alpha1().ActionSets(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrV1alpha1().ActionSets(namespace).Watch(options)
			},
		},
		&cr_v1alpha1.ActionSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *actionSetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredActionSetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *actionSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cr_v1alpha1.ActionSet{}, f.defaultInformer)
}

func (f *actionSetInformer) Lister() v1alpha1.ActionSetLister {
	return v1alpha1.NewActionSetLister(f.Informer().GetIndexer())
}