// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/copilot/handlers"
)

type CAPIDiegoProcessAssociationsRepo struct {
	UpsertStub        func(capiDiegoProcessAssociation handlers.CAPIDiegoProcessAssociation)
	upsertMutex       sync.RWMutex
	upsertArgsForCall []struct {
		capiDiegoProcessAssociation handlers.CAPIDiegoProcessAssociation
	}
	DeleteStub        func(capiProcessGUID handlers.CAPIProcessGUID)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		capiProcessGUID handlers.CAPIProcessGUID
	}
	ListStub        func() map[string][]string
	listMutex       sync.RWMutex
	listArgsForCall []struct{}
	listReturns     struct {
		result1 map[string][]string
	}
	listReturnsOnCall map[int]struct {
		result1 map[string][]string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CAPIDiegoProcessAssociationsRepo) Upsert(capiDiegoProcessAssociation handlers.CAPIDiegoProcessAssociation) {
	fake.upsertMutex.Lock()
	fake.upsertArgsForCall = append(fake.upsertArgsForCall, struct {
		capiDiegoProcessAssociation handlers.CAPIDiegoProcessAssociation
	}{capiDiegoProcessAssociation})
	fake.recordInvocation("Upsert", []interface{}{capiDiegoProcessAssociation})
	fake.upsertMutex.Unlock()
	if fake.UpsertStub != nil {
		fake.UpsertStub(capiDiegoProcessAssociation)
	}
}

func (fake *CAPIDiegoProcessAssociationsRepo) UpsertCallCount() int {
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	return len(fake.upsertArgsForCall)
}

func (fake *CAPIDiegoProcessAssociationsRepo) UpsertArgsForCall(i int) handlers.CAPIDiegoProcessAssociation {
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	return fake.upsertArgsForCall[i].capiDiegoProcessAssociation
}

func (fake *CAPIDiegoProcessAssociationsRepo) Delete(capiProcessGUID handlers.CAPIProcessGUID) {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		capiProcessGUID handlers.CAPIProcessGUID
	}{capiProcessGUID})
	fake.recordInvocation("Delete", []interface{}{capiProcessGUID})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		fake.DeleteStub(capiProcessGUID)
	}
}

func (fake *CAPIDiegoProcessAssociationsRepo) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *CAPIDiegoProcessAssociationsRepo) DeleteArgsForCall(i int) handlers.CAPIProcessGUID {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].capiProcessGUID
}

func (fake *CAPIDiegoProcessAssociationsRepo) List() map[string][]string {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct{}{})
	fake.recordInvocation("List", []interface{}{})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.listReturns.result1
}

func (fake *CAPIDiegoProcessAssociationsRepo) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *CAPIDiegoProcessAssociationsRepo) ListReturns(result1 map[string][]string) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 map[string][]string
	}{result1}
}

func (fake *CAPIDiegoProcessAssociationsRepo) ListReturnsOnCall(i int, result1 map[string][]string) {
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 map[string][]string
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 map[string][]string
	}{result1}
}

func (fake *CAPIDiegoProcessAssociationsRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.upsertMutex.RLock()
	defer fake.upsertMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CAPIDiegoProcessAssociationsRepo) recordInvocation(key string, args []interface{}) {
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