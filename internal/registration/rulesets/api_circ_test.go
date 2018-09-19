// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package rulesets

import (
	"github.com/circonus-labs/circonus-gometrics/api"
	"sync"
)

var (
	lockCircAPIMockCreateRuleSet sync.RWMutex
)

// CircAPIMock is a mock implementation of CircAPI.
//
//     func TestSomethingThatUsesCircAPI(t *testing.T) {
//
//         // make and configure a mocked CircAPI
//         mockedCircAPI := &CircAPIMock{
//             CreateRuleSetFunc: func(cfg *api.RuleSet) (*api.RuleSet, error) {
// 	               panic("TODO: mock out the CreateRuleSet method")
//             },
//         }
//
//         // TODO: use mockedCircAPI in code that requires CircAPI
//         //       and then make assertions.
//
//     }
type CircAPIMock struct {
	// CreateRuleSetFunc mocks the CreateRuleSet method.
	CreateRuleSetFunc func(cfg *api.RuleSet) (*api.RuleSet, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateRuleSet holds details about calls to the CreateRuleSet method.
		CreateRuleSet []struct {
			// Cfg is the cfg argument value.
			Cfg *api.RuleSet
		}
	}
}

// CreateRuleSet calls CreateRuleSetFunc.
func (mock *CircAPIMock) CreateRuleSet(cfg *api.RuleSet) (*api.RuleSet, error) {
	if mock.CreateRuleSetFunc == nil {
		panic("moq: CircAPIMock.CreateRuleSetFunc is nil but CircAPI.CreateRuleSet was just called")
	}
	callInfo := struct {
		Cfg *api.RuleSet
	}{
		Cfg: cfg,
	}
	lockCircAPIMockCreateRuleSet.Lock()
	mock.calls.CreateRuleSet = append(mock.calls.CreateRuleSet, callInfo)
	lockCircAPIMockCreateRuleSet.Unlock()
	return mock.CreateRuleSetFunc(cfg)
}

// CreateRuleSetCalls gets all the calls that were made to CreateRuleSet.
// Check the length with:
//     len(mockedCircAPI.CreateRuleSetCalls())
func (mock *CircAPIMock) CreateRuleSetCalls() []struct {
	Cfg *api.RuleSet
} {
	var calls []struct {
		Cfg *api.RuleSet
	}
	lockCircAPIMockCreateRuleSet.RLock()
	calls = mock.calls.CreateRuleSet
	lockCircAPIMockCreateRuleSet.RUnlock()
	return calls
}
