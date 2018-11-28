// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package check

import (
	"sync"

	circapi "github.com/circonus-labs/go-apiclient"
)

var (
	lockAPIMockCreateCheckBundle      sync.RWMutex
	lockAPIMockDeleteCheckBundleByCID sync.RWMutex
	lockAPIMockFetchCheckBundle       sync.RWMutex
	lockAPIMockSearchCheckBundles     sync.RWMutex
	lockAPIMockUpdateCheckBundle      sync.RWMutex
)

// APIMock is a mock implementation of API.
//
//     func TestSomethingThatUsesAPI(t *testing.T) {
//
//         // make and configure a mocked API
//         mockedAPI := &APIMock{
//             CreateCheckBundleFunc: func(cfg *circapi.CheckBundle) (*circapi.CheckBundle, error) {
// 	               panic("TODO: mock out the CreateCheckBundle method")
//             },
//             DeleteCheckBundleByCIDFunc: func(cid circapi.CIDType) (bool, error) {
// 	               panic("TODO: mock out the DeleteCheckBundleByCID method")
//             },
//             FetchCheckBundleFunc: func(cid circapi.CIDType) (*circapi.CheckBundle, error) {
// 	               panic("TODO: mock out the FetchCheckBundle method")
//             },
//             SearchCheckBundlesFunc: func(searchCriteria *circapi.SearchQueryType, filterCriteria *map[string][]string) (*[]circapi.CheckBundle, error) {
// 	               panic("TODO: mock out the SearchCheckBundles method")
//             },
//             UpdateCheckBundleFunc: func(cfg *circapi.CheckBundle) (*circapi.CheckBundle, error) {
// 	               panic("TODO: mock out the UpdateCheckBundle method")
//             },
//         }
//
//         // TODO: use mockedAPI in code that requires API
//         //       and then make assertions.
//
//     }
type APIMock struct {
	// CreateCheckBundleFunc mocks the CreateCheckBundle method.
	CreateCheckBundleFunc func(cfg *circapi.CheckBundle) (*circapi.CheckBundle, error)

	// DeleteCheckBundleByCIDFunc mocks the DeleteCheckBundleByCID method.
	DeleteCheckBundleByCIDFunc func(cid circapi.CIDType) (bool, error)

	// FetchCheckBundleFunc mocks the FetchCheckBundle method.
	FetchCheckBundleFunc func(cid circapi.CIDType) (*circapi.CheckBundle, error)

	// SearchCheckBundlesFunc mocks the SearchCheckBundles method.
	SearchCheckBundlesFunc func(searchCriteria *circapi.SearchQueryType, filterCriteria *circapi.SearchFilterType) (*[]circapi.CheckBundle, error)

	// UpdateCheckBundleFunc mocks the UpdateCheckBundle method.
	UpdateCheckBundleFunc func(cfg *circapi.CheckBundle) (*circapi.CheckBundle, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateCheckBundle holds details about calls to the CreateCheckBundle method.
		CreateCheckBundle []struct {
			// Cfg is the cfg argument value.
			Cfg *circapi.CheckBundle
		}
		// DeleteCheckBundleByCID holds details about calls to the DeleteCheckBundleByCID method.
		DeleteCheckBundleByCID []struct {
			// Cid is the cid argument value.
			Cid circapi.CIDType
		}
		// FetchCheckBundle holds details about calls to the FetchCheckBundle method.
		FetchCheckBundle []struct {
			// Cid is the cid argument value.
			Cid circapi.CIDType
		}
		// SearchCheckBundles holds details about calls to the SearchCheckBundles method.
		SearchCheckBundles []struct {
			// SearchCriteria is the searchCriteria argument value.
			SearchCriteria *circapi.SearchQueryType
			// FilterCriteria is the filterCriteria argument value.
			FilterCriteria *circapi.SearchFilterType
		}
		// UpdateCheckBundle holds details about calls to the UpdateCheckBundle method.
		UpdateCheckBundle []struct {
			// Cfg is the cfg argument value.
			Cfg *circapi.CheckBundle
		}
	}
}

// CreateCheckBundle calls CreateCheckBundleFunc.
func (mock *APIMock) CreateCheckBundle(cfg *circapi.CheckBundle) (*circapi.CheckBundle, error) {
	if mock.CreateCheckBundleFunc == nil {
		panic("moq: APIMock.CreateCheckBundleFunc is nil but API.CreateCheckBundle was just called")
	}
	callInfo := struct {
		Cfg *circapi.CheckBundle
	}{
		Cfg: cfg,
	}
	lockAPIMockCreateCheckBundle.Lock()
	mock.calls.CreateCheckBundle = append(mock.calls.CreateCheckBundle, callInfo)
	lockAPIMockCreateCheckBundle.Unlock()
	return mock.CreateCheckBundleFunc(cfg)
}

// CreateCheckBundleCalls gets all the calls that were made to CreateCheckBundle.
// Check the length with:
//     len(mockedAPI.CreateCheckBundleCalls())
func (mock *APIMock) CreateCheckBundleCalls() []struct {
	Cfg *circapi.CheckBundle
} {
	var calls []struct {
		Cfg *circapi.CheckBundle
	}
	lockAPIMockCreateCheckBundle.RLock()
	calls = mock.calls.CreateCheckBundle
	lockAPIMockCreateCheckBundle.RUnlock()
	return calls
}

// DeleteCheckBundleByCID calls DeleteCheckBundleByCIDFunc.
func (mock *APIMock) DeleteCheckBundleByCID(cid circapi.CIDType) (bool, error) {
	if mock.DeleteCheckBundleByCIDFunc == nil {
		panic("moq: APIMock.DeleteCheckBundleByCIDFunc is nil but API.DeleteCheckBundleByCID was just called")
	}
	callInfo := struct {
		Cid circapi.CIDType
	}{
		Cid: cid,
	}
	lockAPIMockDeleteCheckBundleByCID.Lock()
	mock.calls.DeleteCheckBundleByCID = append(mock.calls.DeleteCheckBundleByCID, callInfo)
	lockAPIMockDeleteCheckBundleByCID.Unlock()
	return mock.DeleteCheckBundleByCIDFunc(cid)
}

// DeleteCheckBundleByCIDCalls gets all the calls that were made to DeleteCheckBundleByCID.
// Check the length with:
//     len(mockedAPI.DeleteCheckBundleByCIDCalls())
func (mock *APIMock) DeleteCheckBundleByCIDCalls() []struct {
	Cid circapi.CIDType
} {
	var calls []struct {
		Cid circapi.CIDType
	}
	lockAPIMockDeleteCheckBundleByCID.RLock()
	calls = mock.calls.DeleteCheckBundleByCID
	lockAPIMockDeleteCheckBundleByCID.RUnlock()
	return calls
}

// FetchCheckBundle calls FetchCheckBundleFunc.
func (mock *APIMock) FetchCheckBundle(cid circapi.CIDType) (*circapi.CheckBundle, error) {
	if mock.FetchCheckBundleFunc == nil {
		panic("moq: APIMock.FetchCheckBundleFunc is nil but API.FetchCheckBundle was just called")
	}
	callInfo := struct {
		Cid circapi.CIDType
	}{
		Cid: cid,
	}
	lockAPIMockFetchCheckBundle.Lock()
	mock.calls.FetchCheckBundle = append(mock.calls.FetchCheckBundle, callInfo)
	lockAPIMockFetchCheckBundle.Unlock()
	return mock.FetchCheckBundleFunc(cid)
}

// FetchCheckBundleCalls gets all the calls that were made to FetchCheckBundle.
// Check the length with:
//     len(mockedAPI.FetchCheckBundleCalls())
func (mock *APIMock) FetchCheckBundleCalls() []struct {
	Cid circapi.CIDType
} {
	var calls []struct {
		Cid circapi.CIDType
	}
	lockAPIMockFetchCheckBundle.RLock()
	calls = mock.calls.FetchCheckBundle
	lockAPIMockFetchCheckBundle.RUnlock()
	return calls
}

// SearchCheckBundles calls SearchCheckBundlesFunc.
func (mock *APIMock) SearchCheckBundles(searchCriteria *circapi.SearchQueryType, filterCriteria *circapi.SearchFilterType) (*[]circapi.CheckBundle, error) {
	if mock.SearchCheckBundlesFunc == nil {
		panic("moq: APIMock.SearchCheckBundlesFunc is nil but API.SearchCheckBundles was just called")
	}
	callInfo := struct {
		SearchCriteria *circapi.SearchQueryType
		FilterCriteria *circapi.SearchFilterType
	}{
		SearchCriteria: searchCriteria,
		FilterCriteria: filterCriteria,
	}
	lockAPIMockSearchCheckBundles.Lock()
	mock.calls.SearchCheckBundles = append(mock.calls.SearchCheckBundles, callInfo)
	lockAPIMockSearchCheckBundles.Unlock()
	return mock.SearchCheckBundlesFunc(searchCriteria, filterCriteria)
}

// SearchCheckBundlesCalls gets all the calls that were made to SearchCheckBundles.
// Check the length with:
//     len(mockedAPI.SearchCheckBundlesCalls())
func (mock *APIMock) SearchCheckBundlesCalls() []struct {
	SearchCriteria *circapi.SearchQueryType
	FilterCriteria *circapi.SearchFilterType
} {
	var calls []struct {
		SearchCriteria *circapi.SearchQueryType
		FilterCriteria *circapi.SearchFilterType
	}
	lockAPIMockSearchCheckBundles.RLock()
	calls = mock.calls.SearchCheckBundles
	lockAPIMockSearchCheckBundles.RUnlock()
	return calls
}

// UpdateCheckBundle calls UpdateCheckBundleFunc.
func (mock *APIMock) UpdateCheckBundle(cfg *circapi.CheckBundle) (*circapi.CheckBundle, error) {
	if mock.UpdateCheckBundleFunc == nil {
		panic("moq: APIMock.UpdateCheckBundleFunc is nil but API.UpdateCheckBundle was just called")
	}
	callInfo := struct {
		Cfg *circapi.CheckBundle
	}{
		Cfg: cfg,
	}
	lockAPIMockUpdateCheckBundle.Lock()
	mock.calls.UpdateCheckBundle = append(mock.calls.UpdateCheckBundle, callInfo)
	lockAPIMockUpdateCheckBundle.Unlock()
	return mock.UpdateCheckBundleFunc(cfg)
}

// UpdateCheckBundleCalls gets all the calls that were made to UpdateCheckBundle.
// Check the length with:
//     len(mockedAPI.UpdateCheckBundleCalls())
func (mock *APIMock) UpdateCheckBundleCalls() []struct {
	Cfg *circapi.CheckBundle
} {
	var calls []struct {
		Cfg *circapi.CheckBundle
	}
	lockAPIMockUpdateCheckBundle.RLock()
	calls = mock.calls.UpdateCheckBundle
	lockAPIMockUpdateCheckBundle.RUnlock()
	return calls
}