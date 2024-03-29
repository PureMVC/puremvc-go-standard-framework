//
//  View.go
//  PureMVC Go Standard
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package view

import (
	"github.com/puremvc/puremvc-go-standard-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-standard-framework/src/patterns/observer"
	"sync"
)

/*
View A Singleton IView implementation.
In PureMVC, the View class assumes these responsibilities:

* Maintain a cache of IMediator instances.

* Provide methods for registering, retrieving, and removing IMediators.

* Notifying IMediators when they are registered or removed.

* Managing the observer lists for each INotification in the application.

* Providing a method for attaching IObservers to an INotification's observer list.

* Providing a method for broadcasting an INotification.

* Notifying the IObservers of a given INotification when it broadcast.
*/
type View struct {
	mediatorMap      map[string]interfaces.IMediator   // Mapping of Mediator names to Mediator instances
	observerMap      map[string][]interfaces.IObserver // Mapping of Notification names to Observer lists
	mediatorMapMutex sync.RWMutex                      // Mutex for mediatorMap
	observerMapMutex sync.RWMutex                      // Mutex for observerMap
}

var instance interfaces.IView      // The Singleton View instance.
var instanceMutex = sync.RWMutex{} // instanceMutex

/*
GetInstance View Singleton Factory method.

- parameter factory: reference that returns IView

- returns: the Singleton instance returned by executing the passed viewFunc
*/
func GetInstance(factory func() interfaces.IView) interfaces.IView {
	instanceMutex.Lock()
	defer instanceMutex.Unlock()

	if instance == nil {
		instance = factory()
		instance.InitializeView()
	}
	return instance
}

/*
InitializeView Initialize the Singleton View instance.

Called automatically by the GetInstance, this
is your opportunity to initialize the Singleton
instance in your subclass without overriding the
constructor.
*/
func (self *View) InitializeView() {
	self.mediatorMap = map[string]interfaces.IMediator{}
	self.observerMap = map[string][]interfaces.IObserver{}
}

/*
RegisterObserver Register an IObserver to be notified
of INotifications with a given name.

- parameter notificationName: the name of the INotifications to notify this IObserver of

- parameter observer: the IObserver to register
*/
func (self *View) RegisterObserver(notificationName string, observer interfaces.IObserver) {
	self.observerMapMutex.Lock()
	defer self.observerMapMutex.Unlock()

	if self.observerMap[notificationName] != nil {
		self.observerMap[notificationName] = append(self.observerMap[notificationName], observer)
	} else {
		self.observerMap[notificationName] = []interfaces.IObserver{observer}
	}
}

/*
NotifyObservers Notify the IObservers for a particular INotification.

All previously attached IObservers for this INotification's
list are notified and are passed a reference to the INotification in
the order in which they were registered.

- parameter notification: the INotification to notify IObservers of.
*/
func (self *View) NotifyObservers(notification interfaces.INotification) {
	self.observerMapMutex.RLock()

	var observers []interfaces.IObserver
	if self.observerMap[notification.Name()] != nil {
		// Get a reference to the observers list for this notification name
		observersRef := self.observerMap[notification.Name()]

		// Copy observers from reference array to working array,
		// since the reference array may change during the notification loop
		observers = make([]interfaces.IObserver, len(observersRef))
		copy(observers, observersRef)
	}

	self.observerMapMutex.RUnlock()

	// Notify Observers from the working array
	for _, observer := range observers {
		observer.NotifyObserver(notification)
	}
}

/*
RemoveObserver Remove the observer for a given notifyContext from an observer list for a given Notification name.

- parameter notificationName: which observer list to remove from

- parameter notifyContext: remove the observer with this object as its notifyContext
*/
func (self *View) RemoveObserver(notificationName string, notifyContext interface{}) {
	self.observerMapMutex.Lock()
	defer self.observerMapMutex.Unlock()

	// the observer list for the notification under inspection
	observers := self.observerMap[notificationName]

	// find the observer for the notifyContext
	for index, observer := range observers {
		if observer.CompareNotifyContext(notifyContext) == true {
			// there can only be one Observer for a given notifyContext
			// in any given Observer list, so remove it and break
			observers = append(observers[:index], observers[index+1:]...)
			break
		}
	}

	// Also, when a Notification's Observer list length falls to
	// zero, delete the notification key from the observer map
	if len(observers) == 0 {
		delete(self.observerMap, notificationName)
	}
}

/*
RegisterMediator Register an IMediator instance with the View.

Registers the IMediator so that it can be retrieved by name,
and further interrogates the IMediator for its
INotification interests.

If the IMediator returns any INotification
names to be notified about, an Observer is created encapsulating
the IMediator instance's handleNotification method
and registering it as an Observer for all INotifications the
IMediator is interested in.

- parameter mediator: a reference to the IMediator instance
*/
func (self *View) RegisterMediator(mediator interfaces.IMediator) {
	self.mediatorMapMutex.Lock()
	defer self.mediatorMapMutex.Unlock()

	// do not allow re-registration (you must removeMediator fist)
	if self.mediatorMap[mediator.GetMediatorName()] != nil {
		return
	}

	mediator.InitializeNotifier()

	// Register the Mediator for retrieval by name
	self.mediatorMap[mediator.GetMediatorName()] = mediator

	// Get Notification interests, if any.
	interests := mediator.ListNotificationInterests()

	// Register Mediator as an observer for each notification of interests
	if len(interests) > 0 {
		// Create Observer referencing this mediator's handlNotification method
		observer := &observer.Observer{Notify: mediator.HandleNotification, Context: mediator}

		// Register Mediator as Observer for its list of Notification interests
		for _, interest := range interests {
			self.RegisterObserver(interest, observer)
		}

	}
	// alert the mediator that it has been registered
	mediator.OnRegister()
}

/*
RetrieveMediator Retrieve an IMediator from the View.

- parameter mediatorName: the name of the IMediator instance to retrieve.

- returns: the IMediator instance previously registered with the given mediatorName.
*/
func (self *View) RetrieveMediator(mediatorName string) interfaces.IMediator {
	self.mediatorMapMutex.RLock()
	defer self.mediatorMapMutex.RUnlock()

	return self.mediatorMap[mediatorName]
}

/*
RemoveMediator Remove an IMediator from the View.

- parameter mediatorName: name of the IMediator instance to be removed.

- returns: the IMediator that was removed from the View
*/
func (self *View) RemoveMediator(mediatorName string) interfaces.IMediator {
	self.mediatorMapMutex.Lock()
	defer self.mediatorMapMutex.Unlock()

	// Retrieve the named mediator
	var mediator = self.mediatorMap[mediatorName]

	if mediator != nil {
		// for every notification this mediator is interested in...
		interests := mediator.ListNotificationInterests()

		for _, interest := range interests {
			// remove the observer linking the mediator
			// to the notification interest
			self.RemoveObserver(interest, mediator)
		}

		// remove the mediator from the map
		delete(self.mediatorMap, mediatorName)

		// alert the mediator that it has been removed
		mediator.OnRemove()
	}
	return mediator
}

/*
HasMediator Check if a Mediator is registered or not

- parameter mediatorName:

- returns: whether a Mediator is registered with the given mediatorName.
*/
func (self *View) HasMediator(mediatorName string) bool {
	self.mediatorMapMutex.RLock()
	self.mediatorMapMutex.RUnlock()

	return self.mediatorMap[mediatorName] != nil
}
