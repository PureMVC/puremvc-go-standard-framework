//
//  Model.go
//  PureMVC Go Standard
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package model

import (
	"github.com/puremvc/puremvc-go-standard-framework/src/interfaces"
	"sync"
)

/*
Model A Singleton IModel implementation.

In PureMVC, the Model class provides
access to model objects (Proxies) by named lookup.

The Model assumes these responsibilities:

* Maintain a cache of IProxy instances.

* Provide methods for registering, retrieving, and removing IProxy instances.

Your application must register IProxy instances
with the Model. Typically, you use an
ICommand to create and register IProxy
instances once the Facade has initialized the Core
actors.
*/
type Model struct {
	proxyMap      map[string]interfaces.IProxy // Mapping of proxyNames to IProxy instances
	proxyMapMutex sync.RWMutex                 // Mutex for proxyMap
}

var instance interfaces.IModel // The Singleton Model instance.
var instanceMutex sync.RWMutex // instanceMutex for thread safety

/*
GetInstance Model Singleton Factory method.

- parameter factory: reference that returns IModel

- returns: the instance returned by the passed modelFunc
*/
func GetInstance(factory func() interfaces.IModel) interfaces.IModel {
	instanceMutex.Lock()
	defer instanceMutex.Unlock()

	if instance == nil {
		instance = factory()
		instance.InitializeModel()
	}
	return instance
}

/*
InitializeModel Initialize the Model instance.

Called automatically by the GetInstance, this
is your opportunity to initialize the Singleton
instance in your subclass without overriding the
constructor.
*/
func (self *Model) InitializeModel() {
	self.proxyMap = map[string]interfaces.IProxy{}
}

/*
RegisterProxy Register an IProxy with the Model.

- parameter proxy: an IProxy to be held by the Model.
*/
func (self *Model) RegisterProxy(proxy interfaces.IProxy) {
	self.proxyMapMutex.Lock()
	defer self.proxyMapMutex.Unlock()

	proxy.InitializeNotifier()
	self.proxyMap[proxy.GetProxyName()] = proxy
	proxy.OnRegister()
}

/*
RetrieveProxy Retrieve an IProxy from the Model.

- parameter proxyName:

- returns: the IProxy instance previously registered with the given proxyName.
*/
func (self *Model) RetrieveProxy(proxyName string) interfaces.IProxy {
	self.proxyMapMutex.RLock()
	defer self.proxyMapMutex.RUnlock()

	return self.proxyMap[proxyName]
}

/*
RemoveProxy Remove an IProxy from the Model.

- parameter proxyName: name of the IProxy instance to be removed.

- returns: the IProxy that was removed from the Model
*/
func (self *Model) RemoveProxy(proxyName string) interfaces.IProxy {
	self.proxyMapMutex.Lock()
	defer self.proxyMapMutex.Unlock()

	var proxy = self.proxyMap[proxyName]
	if proxy != nil {
		delete(self.proxyMap, proxyName)
		proxy.OnRemove()
	}
	return proxy
}

/*
HasProxy Check if a Proxy is registered

- parameter proxyName:

- returns: whether a Proxy is currently registered with the given proxyName.
*/
func (self *Model) HasProxy(proxyName string) bool {
	self.proxyMapMutex.RLock()
	defer self.proxyMapMutex.RUnlock()

	return self.proxyMap[proxyName] != nil
}
