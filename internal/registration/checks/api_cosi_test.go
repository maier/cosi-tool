// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package checks

import (
	"github.com/circonus-labs/cosi-server/api"
	"sync"
)

var (
	lockCosiAPIMockFetchTemplate sync.RWMutex
)

// CosiAPIMock is a mock implementation of CosiAPI.
//
//     func TestSomethingThatUsesCosiAPI(t *testing.T) {
//
//         // make and configure a mocked CosiAPI
//         mockedCosiAPI := &CosiAPIMock{
//             FetchTemplateFunc: func(id string) (*api.Template, error) {
// 	               panic("TODO: mock out the FetchTemplate method")
//             },
//         }
//
//         // TODO: use mockedCosiAPI in code that requires CosiAPI
//         //       and then make assertions.
//
//     }
type CosiAPIMock struct {
	// FetchTemplateFunc mocks the FetchTemplate method.
	FetchTemplateFunc func(id string) (*api.Template, error)

	// calls tracks calls to the methods.
	calls struct {
		// FetchTemplate holds details about calls to the FetchTemplate method.
		FetchTemplate []struct {
			// ID is the id argument value.
			ID string
		}
	}
}

// FetchTemplate calls FetchTemplateFunc.
func (mock *CosiAPIMock) FetchTemplate(id string) (*api.Template, error) {
	if mock.FetchTemplateFunc == nil {
		panic("moq: CosiAPIMock.FetchTemplateFunc is nil but CosiAPI.FetchTemplate was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	lockCosiAPIMockFetchTemplate.Lock()
	mock.calls.FetchTemplate = append(mock.calls.FetchTemplate, callInfo)
	lockCosiAPIMockFetchTemplate.Unlock()
	return mock.FetchTemplateFunc(id)
}

// FetchTemplateCalls gets all the calls that were made to FetchTemplate.
// Check the length with:
//     len(mockedCosiAPI.FetchTemplateCalls())
func (mock *CosiAPIMock) FetchTemplateCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	lockCosiAPIMockFetchTemplate.RLock()
	calls = mock.calls.FetchTemplate
	lockCosiAPIMockFetchTemplate.RUnlock()
	return calls
}
