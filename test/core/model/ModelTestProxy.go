//
//  ModelTestProxy.go
//  PureMVC Go Standard
//
//  Copyright(c) 2019 Saad Shams <saad.shams@puremvc.org>
//  Your reuse is governed by the Creative Commons Attribution 3.0 License
//

package model

import "github.com/puremvc/puremvc-go-standard-framework/src/patterns/proxy"

const MODEL_TEST_PROXY = "modelTestProxy"
const ON_REGISTER_CALLED = "onRegister Called"
const ON_REMOVE_CALLED = "onRemoveCalled"

type ModelTestProxy struct {
	proxy.Proxy
}

func (proxy *ModelTestProxy) OnRegister() {
	proxy.SetData(ON_REGISTER_CALLED)
}

func (proxy *ModelTestProxy) OnRemove() {
	proxy.SetData(ON_REMOVE_CALLED)
}
