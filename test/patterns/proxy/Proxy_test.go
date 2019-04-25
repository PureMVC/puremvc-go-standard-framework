//
//  Proxy_test.go
//  PureMVC Go Standard
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package proxy

import (
	"github.com/puremvc/puremvc-go-standard-framework/src/interfaces"
	"github.com/puremvc/puremvc-go-standard-framework/src/patterns/proxy"
	"testing"
)

/**
Test the PureMVC Proxy class.
*/

/**
  Tests getting the name using Proxy class accessor method. Setting can only be done in constructor.
*/
func TestNameAccessor(t *testing.T) {
	// Create a new Proxy and use accessors to set the proxy name
	var proxy interfaces.IProxy = &proxy.Proxy{Name: "TestProxy", Data: nil}

	// test assertions
	if proxy.GetProxyName() != "TestProxy" {
		t.Error("Expecting proxy.GetProxyName() == 'TestProxy'")
	}
}

/**
  Tests setting and getting the data using Proxy class accessor methods.
*/
func TestDataAccessor(t *testing.T) {
	// Create a new Proxy and use accessors to set the data
	var proxy interfaces.IProxy = &proxy.Proxy{Name: "colors"}
	proxy.SetData([]string{"red", "green", "blue"})

	var data = proxy.GetData().([]string)

	// test assertions
	if len(data) != 3 {
		t.Error("Expecting len(data) == 3")
	}
	if data[0] != "red" {
		t.Error("Expecting data[0] == 'red'")
	}
	if data[1] != "green" {
		t.Error("Expecting data[1] == 'green'")
	}
	if data[2] != "blue" {
		t.Error("Expecting data[2] == 'blue'")
	}
}

/**
  Tests setting the name and body using the Notification class Constructor.
*/
func TestConstructor(t *testing.T) {
	// Create a new Proxy using the Constructor to set the name and data
	var proxy interfaces.IProxy = &proxy.Proxy{Name: "colors", Data: []string{"red", "green", "blue"}}

	var data = proxy.GetData().([]string)

	// test assertions
	if proxy == nil {
		t.Error("Expecting proxy not nil")
	}
	if proxy.GetProxyName() != "colors" {
		t.Error("Expecting proxy.GetProxyName() == 'colors'")
	}
	if len(data) != 3 {
		t.Error("Expecting data.length == 3")
	}
	if data[0] != "red" {
		t.Error("Expecting data[0] == 'red'")
	}
	if data[1] != "green" {
		t.Error("Expecting data[1] == 'green'")
	}
	if data[2] != "blue" {
		t.Error("Expecting data[2] == 'blue'")
	}
}
