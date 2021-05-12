// Code generated by counterfeiter. DO NOT EDIT.
package routingfakes

import (
	"net/http"
	"sync"

	"github.com/kieron-pivotal/menu-planner-app/routing"
)

type FakeRecipeHandler struct {
	GetRecipesStub        func(http.ResponseWriter, *http.Request)
	getRecipesMutex       sync.RWMutex
	getRecipesArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	NewRecipeStub        func(http.ResponseWriter, *http.Request)
	newRecipeMutex       sync.RWMutex
	newRecipeArgsForCall []struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRecipeHandler) GetRecipes(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.getRecipesMutex.Lock()
	fake.getRecipesArgsForCall = append(fake.getRecipesArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("GetRecipes", []interface{}{arg1, arg2})
	fake.getRecipesMutex.Unlock()
	if fake.GetRecipesStub != nil {
		fake.GetRecipesStub(arg1, arg2)
	}
}

func (fake *FakeRecipeHandler) GetRecipesCallCount() int {
	fake.getRecipesMutex.RLock()
	defer fake.getRecipesMutex.RUnlock()
	return len(fake.getRecipesArgsForCall)
}

func (fake *FakeRecipeHandler) GetRecipesCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.getRecipesMutex.Lock()
	defer fake.getRecipesMutex.Unlock()
	fake.GetRecipesStub = stub
}

func (fake *FakeRecipeHandler) GetRecipesArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.getRecipesMutex.RLock()
	defer fake.getRecipesMutex.RUnlock()
	argsForCall := fake.getRecipesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRecipeHandler) NewRecipe(arg1 http.ResponseWriter, arg2 *http.Request) {
	fake.newRecipeMutex.Lock()
	fake.newRecipeArgsForCall = append(fake.newRecipeArgsForCall, struct {
		arg1 http.ResponseWriter
		arg2 *http.Request
	}{arg1, arg2})
	fake.recordInvocation("NewRecipe", []interface{}{arg1, arg2})
	fake.newRecipeMutex.Unlock()
	if fake.NewRecipeStub != nil {
		fake.NewRecipeStub(arg1, arg2)
	}
}

func (fake *FakeRecipeHandler) NewRecipeCallCount() int {
	fake.newRecipeMutex.RLock()
	defer fake.newRecipeMutex.RUnlock()
	return len(fake.newRecipeArgsForCall)
}

func (fake *FakeRecipeHandler) NewRecipeCalls(stub func(http.ResponseWriter, *http.Request)) {
	fake.newRecipeMutex.Lock()
	defer fake.newRecipeMutex.Unlock()
	fake.NewRecipeStub = stub
}

func (fake *FakeRecipeHandler) NewRecipeArgsForCall(i int) (http.ResponseWriter, *http.Request) {
	fake.newRecipeMutex.RLock()
	defer fake.newRecipeMutex.RUnlock()
	argsForCall := fake.newRecipeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRecipeHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getRecipesMutex.RLock()
	defer fake.getRecipesMutex.RUnlock()
	fake.newRecipeMutex.RLock()
	defer fake.newRecipeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRecipeHandler) recordInvocation(key string, args []interface{}) {
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

var _ routing.RecipeHandler = new(FakeRecipeHandler)
