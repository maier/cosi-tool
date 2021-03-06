// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package registration

import (
	"sync"

	cosiapi "github.com/circonus-labs/cosi-server/api"
)

var (
	lockCosiAPIMockFetchBroker   sync.RWMutex
	lockCosiAPIMockFetchTemplate sync.RWMutex
)

// CosiAPIMock is a mock implementation of CosiAPI.
//
//     func TestSomethingThatUsesCosiAPI(t *testing.T) {
//
//         // make and configure a mocked CosiAPI
//         mockedCosiAPI := &CosiAPIMock{
//             FetchBrokerFunc: func(checkType string) (string, error) {
// 	               panic("TODO: mock out the FetchBroker method")
//             },
//             FetchTemplateFunc: func(id string) (*cosiapi.Template, error) {
// 	               panic("TODO: mock out the FetchTemplate method")
//             },
//         }
//
//         // TODO: use mockedCosiAPI in code that requires CosiAPI
//         //       and then make assertions.
//
//     }
type CosiAPIMock struct {
	// FetchBrokerFunc mocks the FetchBroker method.
	FetchBrokerFunc func(checkType string) (string, error)

	// FetchTemplateFunc mocks the FetchTemplate method.
	FetchTemplateFunc func(id string) (*cosiapi.Template, error)

	// calls tracks calls to the methods.
	calls struct {
		// FetchBroker holds details about calls to the FetchBroker method.
		FetchBroker []struct {
			// CheckType is the checkType argument value.
			CheckType string
		}
		// FetchTemplate holds details about calls to the FetchTemplate method.
		FetchTemplate []struct {
			// ID is the id argument value.
			ID string
		}
	}
}

// FetchBroker calls FetchBrokerFunc.
func (mock *CosiAPIMock) FetchBroker(checkType string) (string, error) {
	if mock.FetchBrokerFunc == nil {
		panic("moq: CosiAPIMock.FetchBrokerFunc is nil but CosiAPI.FetchBroker was just called")
	}
	callInfo := struct {
		CheckType string
	}{
		CheckType: checkType,
	}
	lockCosiAPIMockFetchBroker.Lock()
	mock.calls.FetchBroker = append(mock.calls.FetchBroker, callInfo)
	lockCosiAPIMockFetchBroker.Unlock()
	return mock.FetchBrokerFunc(checkType)
}

// FetchBrokerCalls gets all the calls that were made to FetchBroker.
// Check the length with:
//     len(mockedCosiAPI.FetchBrokerCalls())
func (mock *CosiAPIMock) FetchBrokerCalls() []struct {
	CheckType string
} {
	var calls []struct {
		CheckType string
	}
	lockCosiAPIMockFetchBroker.RLock()
	calls = mock.calls.FetchBroker
	lockCosiAPIMockFetchBroker.RUnlock()
	return calls
}

// FetchTemplate calls FetchTemplateFunc.
func (mock *CosiAPIMock) FetchTemplate(id string) (*cosiapi.Template, error) {
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
