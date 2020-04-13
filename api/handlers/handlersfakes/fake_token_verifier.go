// Code generated by counterfeiter. DO NOT EDIT.
package handlersfakes

import (
	"sync"

	"github.com/kieron-pivotal/menu-planner-app/handlers"
)

type FakeTokenVerifier struct {
	VerifyIDTokenStub        func(string, []string) error
	verifyIDTokenMutex       sync.RWMutex
	verifyIDTokenArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	verifyIDTokenReturns struct {
		result1 error
	}
	verifyIDTokenReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTokenVerifier) VerifyIDToken(arg1 string, arg2 []string) error {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.verifyIDTokenMutex.Lock()
	ret, specificReturn := fake.verifyIDTokenReturnsOnCall[len(fake.verifyIDTokenArgsForCall)]
	fake.verifyIDTokenArgsForCall = append(fake.verifyIDTokenArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2Copy})
	fake.recordInvocation("VerifyIDToken", []interface{}{arg1, arg2Copy})
	fake.verifyIDTokenMutex.Unlock()
	if fake.VerifyIDTokenStub != nil {
		return fake.VerifyIDTokenStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.verifyIDTokenReturns
	return fakeReturns.result1
}

func (fake *FakeTokenVerifier) VerifyIDTokenCallCount() int {
	fake.verifyIDTokenMutex.RLock()
	defer fake.verifyIDTokenMutex.RUnlock()
	return len(fake.verifyIDTokenArgsForCall)
}

func (fake *FakeTokenVerifier) VerifyIDTokenCalls(stub func(string, []string) error) {
	fake.verifyIDTokenMutex.Lock()
	defer fake.verifyIDTokenMutex.Unlock()
	fake.VerifyIDTokenStub = stub
}

func (fake *FakeTokenVerifier) VerifyIDTokenArgsForCall(i int) (string, []string) {
	fake.verifyIDTokenMutex.RLock()
	defer fake.verifyIDTokenMutex.RUnlock()
	argsForCall := fake.verifyIDTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTokenVerifier) VerifyIDTokenReturns(result1 error) {
	fake.verifyIDTokenMutex.Lock()
	defer fake.verifyIDTokenMutex.Unlock()
	fake.VerifyIDTokenStub = nil
	fake.verifyIDTokenReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTokenVerifier) VerifyIDTokenReturnsOnCall(i int, result1 error) {
	fake.verifyIDTokenMutex.Lock()
	defer fake.verifyIDTokenMutex.Unlock()
	fake.VerifyIDTokenStub = nil
	if fake.verifyIDTokenReturnsOnCall == nil {
		fake.verifyIDTokenReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyIDTokenReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTokenVerifier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.verifyIDTokenMutex.RLock()
	defer fake.verifyIDTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTokenVerifier) recordInvocation(key string, args []interface{}) {
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

var _ handlers.TokenVerifier = new(FakeTokenVerifier)
