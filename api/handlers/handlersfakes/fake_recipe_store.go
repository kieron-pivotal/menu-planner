// Code generated by counterfeiter. DO NOT EDIT.
package handlersfakes

import (
	"sync"

	"github.com/kieron-pivotal/menu-planner-app/handlers"
	"github.com/kieron-pivotal/menu-planner-app/models"
)

type FakeRecipeStore struct {
	ListStub        func(int) ([]models.Recipe, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 int
	}
	listReturns struct {
		result1 []models.Recipe
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []models.Recipe
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRecipeStore) List(arg1 int) ([]models.Recipe, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRecipeStore) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeRecipeStore) ListCalls(stub func(int) ([]models.Recipe, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeRecipeStore) ListArgsForCall(i int) int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRecipeStore) ListReturns(result1 []models.Recipe, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []models.Recipe
		result2 error
	}{result1, result2}
}

func (fake *FakeRecipeStore) ListReturnsOnCall(i int, result1 []models.Recipe, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []models.Recipe
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []models.Recipe
		result2 error
	}{result1, result2}
}

func (fake *FakeRecipeStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRecipeStore) recordInvocation(key string, args []interface{}) {
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

var _ handlers.RecipeStore = new(FakeRecipeStore)
