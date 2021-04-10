// Code generated by counterfeiter. DO NOT EDIT.
package routingfakes

import (
	"net/http"
	"sync"

	"github.com/kieron-pivotal/menu-planner-app/routing"
)

type FakeAuthHandler struct {
	AuthGoogleStub        func(http.ResponseWriter, *http.Request)
	authGoogleMutex       sync.RWMutex
	authGoogleArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	WhoAmIStub        func(http.ResponseWriter, *http.Request)
	whoAmIMutex       sync.RWMutex
	whoAmIArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthHandler) AuthGoogle(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.authGoogleMutex.Lock()
	fake.authGoogleArgsForCall = append(fake.authGoogleArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("AuthGoogle", []interface{}{arg1, arg2})
	fake.authGoogleMutex.Unlock()
	if fake.AuthGoogleStub != nil {
		fake.AuthGoogleStub(arg1, arg2)
	}
}

func (fake *FakeAuthHandler) AuthGoogleCallCount() int {
	fake.authGoogleMutex.RLock()
	defer fake.authGoogleMutex.RUnlock()
	return len(fake.authGoogleArgsForCall)
}

func (fake *FakeAuthHandler) AuthGoogleCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.authGoogleMutex.Lock()
	defer fake.authGoogleMutex.Unlock()
	fake.AuthGoogleStub = stub
}

func (fake *FakeAuthHandler) AuthGoogleArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.authGoogleMutex.RLock()
	defer fake.authGoogleMutex.RUnlock()
	argsForCall := fake.authGoogleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAuthHandler) WhoAmI(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.whoAmIMutex.Lock()
	fake.whoAmIArgsForCall = append(fake.whoAmIArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("WhoAmI", []interface{}{arg1, arg2})
	fake.whoAmIMutex.Unlock()
	if fake.WhoAmIStub != nil {
		fake.WhoAmIStub(arg1, arg2)
	}
}

func (fake *FakeAuthHandler) WhoAmICallCount() int {
	fake.whoAmIMutex.RLock()
	defer fake.whoAmIMutex.RUnlock()
	return len(fake.whoAmIArgsForCall)
}

func (fake *FakeAuthHandler) WhoAmICalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.whoAmIMutex.Lock()
	defer fake.whoAmIMutex.Unlock()
	fake.WhoAmIStub = stub
}

func (fake *FakeAuthHandler) WhoAmIArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.whoAmIMutex.RLock()
	defer fake.whoAmIMutex.RUnlock()
	argsForCall := fake.whoAmIArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAuthHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.authGoogleMutex.RLock()
	defer fake.authGoogleMutex.RUnlock()
	fake.whoAmIMutex.RLock()
	defer fake.whoAmIMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthHandler) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ routing.AuthHandler = new(FakeAuthHandler)
