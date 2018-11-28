// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package checks

import (
	"sync"

	apiclient "github.com/circonus-labs/go-apiclient"
)

var (
	lockCircAPIMockCreateCheckBundle      sync.RWMutex
	lockCircAPIMockDeleteCheckBundleByCID sync.RWMutex
	lockCircAPIMockFetchCheckBundle       sync.RWMutex
	lockCircAPIMockSearchCheckBundles     sync.RWMutex
	lockCircAPIMockUpdateCheckBundle      sync.RWMutex
)

// CircAPIMock is a mock implementation of CircAPI.
//
//     func TestSomethingThatUsesCircAPI(t *testing.T) {
//
//         // make and configure a mocked CircAPI
//         mockedCircAPI := &CircAPIMock{
//             CreateCheckBundleFunc: func(cfg *apiclient.CheckBundle) (*apiclient.CheckBundle, error) {
// 	               panic("TODO: mock out the CreateCheckBundle method")
//             },
//             DeleteCheckBundleByCIDFunc: func(cid apiclient.CIDType) (bool, error) {
// 	               panic("TODO: mock out the DeleteCheckBundleByCID method")
//             },
//             FetchCheckBundleFunc: func(cid apiclient.CIDType) (*apiclient.CheckBundle, error) {
// 	               panic("TODO: mock out the FetchCheckBundle method")
//             },
//             SearchCheckBundlesFunc: func(searchCriteria *apiclient.SearchQueryType, filterCriteria *map[string][]string) (*[]apiclient.CheckBundle, error) {
// 	               panic("TODO: mock out the SearchCheckBundles method")
//             },
//             UpdateCheckBundleFunc: func(cfg *apiclient.CheckBundle) (*apiclient.CheckBundle, error) {
// 	               panic("TODO: mock out the UpdateCheckBundle method")
//             },
//         }
//
//         // TODO: use mockedCircAPI in code that requires CircAPI
//         //       and then make assertions.
//
//     }
type CircAPIMock struct {
	// CreateCheckBundleFunc mocks the CreateCheckBundle method.
	CreateCheckBundleFunc func(cfg *apiclient.CheckBundle) (*apiclient.CheckBundle, error)

	// DeleteCheckBundleByCIDFunc mocks the DeleteCheckBundleByCID method.
	DeleteCheckBundleByCIDFunc func(cid apiclient.CIDType) (bool, error)

	// FetchCheckBundleFunc mocks the FetchCheckBundle method.
	FetchCheckBundleFunc func(cid apiclient.CIDType) (*apiclient.CheckBundle, error)

	// SearchCheckBundlesFunc mocks the SearchCheckBundles method.
	SearchCheckBundlesFunc func(searchCriteria *apiclient.SearchQueryType, filterCriteria *apiclient.SearchFilterType) (*[]apiclient.CheckBundle, error)

	// UpdateCheckBundleFunc mocks the UpdateCheckBundle method.
	UpdateCheckBundleFunc func(cfg *apiclient.CheckBundle) (*apiclient.CheckBundle, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateCheckBundle holds details about calls to the CreateCheckBundle method.
		CreateCheckBundle []struct {
			// Cfg is the cfg argument value.
			Cfg *apiclient.CheckBundle
		}
		// DeleteCheckBundleByCID holds details about calls to the DeleteCheckBundleByCID method.
		DeleteCheckBundleByCID []struct {
			// Cid is the cid argument value.
			Cid apiclient.CIDType
		}
		// FetchCheckBundle holds details about calls to the FetchCheckBundle method.
		FetchCheckBundle []struct {
			// Cid is the cid argument value.
			Cid apiclient.CIDType
		}
		// SearchCheckBundles holds details about calls to the SearchCheckBundles method.
		SearchCheckBundles []struct {
			// SearchCriteria is the searchCriteria argument value.
			SearchCriteria *apiclient.SearchQueryType
			// FilterCriteria is the filterCriteria argument value.
			FilterCriteria *apiclient.SearchFilterType
		}
		// UpdateCheckBundle holds details about calls to the UpdateCheckBundle method.
		UpdateCheckBundle []struct {
			// Cfg is the cfg argument value.
			Cfg *apiclient.CheckBundle
		}
	}
}

// CreateCheckBundle calls CreateCheckBundleFunc.
func (mock *CircAPIMock) CreateCheckBundle(cfg *apiclient.CheckBundle) (*apiclient.CheckBundle, error) {
	if mock.CreateCheckBundleFunc == nil {
		panic("moq: CircAPIMock.CreateCheckBundleFunc is nil but CircAPI.CreateCheckBundle was just called")
	}
	callInfo := struct {
		Cfg *apiclient.CheckBundle
	}{
		Cfg: cfg,
	}
	lockCircAPIMockCreateCheckBundle.Lock()
	mock.calls.CreateCheckBundle = append(mock.calls.CreateCheckBundle, callInfo)
	lockCircAPIMockCreateCheckBundle.Unlock()
	return mock.CreateCheckBundleFunc(cfg)
}

// CreateCheckBundleCalls gets all the calls that were made to CreateCheckBundle.
// Check the length with:
//     len(mockedCircAPI.CreateCheckBundleCalls())
func (mock *CircAPIMock) CreateCheckBundleCalls() []struct {
	Cfg *apiclient.CheckBundle
} {
	var calls []struct {
		Cfg *apiclient.CheckBundle
	}
	lockCircAPIMockCreateCheckBundle.RLock()
	calls = mock.calls.CreateCheckBundle
	lockCircAPIMockCreateCheckBundle.RUnlock()
	return calls
}

// DeleteCheckBundleByCID calls DeleteCheckBundleByCIDFunc.
func (mock *CircAPIMock) DeleteCheckBundleByCID(cid apiclient.CIDType) (bool, error) {
	if mock.DeleteCheckBundleByCIDFunc == nil {
		panic("moq: CircAPIMock.DeleteCheckBundleByCIDFunc is nil but CircAPI.DeleteCheckBundleByCID was just called")
	}
	callInfo := struct {
		Cid apiclient.CIDType
	}{
		Cid: cid,
	}
	lockCircAPIMockDeleteCheckBundleByCID.Lock()
	mock.calls.DeleteCheckBundleByCID = append(mock.calls.DeleteCheckBundleByCID, callInfo)
	lockCircAPIMockDeleteCheckBundleByCID.Unlock()
	return mock.DeleteCheckBundleByCIDFunc(cid)
}

// DeleteCheckBundleByCIDCalls gets all the calls that were made to DeleteCheckBundleByCID.
// Check the length with:
//     len(mockedCircAPI.DeleteCheckBundleByCIDCalls())
func (mock *CircAPIMock) DeleteCheckBundleByCIDCalls() []struct {
	Cid apiclient.CIDType
} {
	var calls []struct {
		Cid apiclient.CIDType
	}
	lockCircAPIMockDeleteCheckBundleByCID.RLock()
	calls = mock.calls.DeleteCheckBundleByCID
	lockCircAPIMockDeleteCheckBundleByCID.RUnlock()
	return calls
}

// FetchCheckBundle calls FetchCheckBundleFunc.
func (mock *CircAPIMock) FetchCheckBundle(cid apiclient.CIDType) (*apiclient.CheckBundle, error) {
	if mock.FetchCheckBundleFunc == nil {
		panic("moq: CircAPIMock.FetchCheckBundleFunc is nil but CircAPI.FetchCheckBundle was just called")
	}
	callInfo := struct {
		Cid apiclient.CIDType
	}{
		Cid: cid,
	}
	lockCircAPIMockFetchCheckBundle.Lock()
	mock.calls.FetchCheckBundle = append(mock.calls.FetchCheckBundle, callInfo)
	lockCircAPIMockFetchCheckBundle.Unlock()
	return mock.FetchCheckBundleFunc(cid)
}

// FetchCheckBundleCalls gets all the calls that were made to FetchCheckBundle.
// Check the length with:
//     len(mockedCircAPI.FetchCheckBundleCalls())
func (mock *CircAPIMock) FetchCheckBundleCalls() []struct {
	Cid apiclient.CIDType
} {
	var calls []struct {
		Cid apiclient.CIDType
	}
	lockCircAPIMockFetchCheckBundle.RLock()
	calls = mock.calls.FetchCheckBundle
	lockCircAPIMockFetchCheckBundle.RUnlock()
	return calls
}

// SearchCheckBundles calls SearchCheckBundlesFunc.
func (mock *CircAPIMock) SearchCheckBundles(searchCriteria *apiclient.SearchQueryType, filterCriteria *apiclient.SearchFilterType) (*[]apiclient.CheckBundle, error) {
	if mock.SearchCheckBundlesFunc == nil {
		panic("moq: CircAPIMock.SearchCheckBundlesFunc is nil but CircAPI.SearchCheckBundles was just called")
	}
	callInfo := struct {
		SearchCriteria *apiclient.SearchQueryType
		FilterCriteria *apiclient.SearchFilterType
	}{
		SearchCriteria: searchCriteria,
		FilterCriteria: filterCriteria,
	}
	lockCircAPIMockSearchCheckBundles.Lock()
	mock.calls.SearchCheckBundles = append(mock.calls.SearchCheckBundles, callInfo)
	lockCircAPIMockSearchCheckBundles.Unlock()
	return mock.SearchCheckBundlesFunc(searchCriteria, filterCriteria)
}

// SearchCheckBundlesCalls gets all the calls that were made to SearchCheckBundles.
// Check the length with:
//     len(mockedCircAPI.SearchCheckBundlesCalls())
func (mock *CircAPIMock) SearchCheckBundlesCalls() []struct {
	SearchCriteria *apiclient.SearchQueryType
	FilterCriteria *apiclient.SearchFilterType
} {
	var calls []struct {
		SearchCriteria *apiclient.SearchQueryType
		FilterCriteria *apiclient.SearchFilterType
	}
	lockCircAPIMockSearchCheckBundles.RLock()
	calls = mock.calls.SearchCheckBundles
	lockCircAPIMockSearchCheckBundles.RUnlock()
	return calls
}

// UpdateCheckBundle calls UpdateCheckBundleFunc.
func (mock *CircAPIMock) UpdateCheckBundle(cfg *apiclient.CheckBundle) (*apiclient.CheckBundle, error) {
	if mock.UpdateCheckBundleFunc == nil {
		panic("moq: CircAPIMock.UpdateCheckBundleFunc is nil but CircAPI.UpdateCheckBundle was just called")
	}
	callInfo := struct {
		Cfg *apiclient.CheckBundle
	}{
		Cfg: cfg,
	}
	lockCircAPIMockUpdateCheckBundle.Lock()
	mock.calls.UpdateCheckBundle = append(mock.calls.UpdateCheckBundle, callInfo)
	lockCircAPIMockUpdateCheckBundle.Unlock()
	return mock.UpdateCheckBundleFunc(cfg)
}

// UpdateCheckBundleCalls gets all the calls that were made to UpdateCheckBundle.
// Check the length with:
//     len(mockedCircAPI.UpdateCheckBundleCalls())
func (mock *CircAPIMock) UpdateCheckBundleCalls() []struct {
	Cfg *apiclient.CheckBundle
} {
	var calls []struct {
		Cfg *apiclient.CheckBundle
	}
	lockCircAPIMockUpdateCheckBundle.RLock()
	calls = mock.calls.UpdateCheckBundle
	lockCircAPIMockUpdateCheckBundle.RUnlock()
	return calls
}
