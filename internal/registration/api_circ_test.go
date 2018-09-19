// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package registration

import (
	"github.com/circonus-labs/circonus-gometrics/api"
	"sync"
)

var (
	lockCircAPIMockCreateCheckBundle      sync.RWMutex
	lockCircAPIMockCreateDashboard        sync.RWMutex
	lockCircAPIMockCreateGraph            sync.RWMutex
	lockCircAPIMockCreateRuleSet          sync.RWMutex
	lockCircAPIMockCreateWorksheet        sync.RWMutex
	lockCircAPIMockDeleteCheckBundleByCID sync.RWMutex
	lockCircAPIMockDeleteDashboardByCID   sync.RWMutex
	lockCircAPIMockDeleteGraphByCID       sync.RWMutex
	lockCircAPIMockDeleteWorksheetByCID   sync.RWMutex
	lockCircAPIMockFetchBroker            sync.RWMutex
	lockCircAPIMockFetchBrokers           sync.RWMutex
	lockCircAPIMockFetchCheckBundle       sync.RWMutex
	lockCircAPIMockFetchDashboard         sync.RWMutex
	lockCircAPIMockFetchGraph             sync.RWMutex
	lockCircAPIMockFetchWorksheet         sync.RWMutex
	lockCircAPIMockSearchCheckBundles     sync.RWMutex
	lockCircAPIMockSearchDashboards       sync.RWMutex
	lockCircAPIMockSearchGraphs           sync.RWMutex
	lockCircAPIMockSearchWorksheets       sync.RWMutex
	lockCircAPIMockUpdateCheckBundle      sync.RWMutex
	lockCircAPIMockUpdateDashboard        sync.RWMutex
	lockCircAPIMockUpdateGraph            sync.RWMutex
	lockCircAPIMockUpdateWorksheet        sync.RWMutex
)

// CircAPIMock is a mock implementation of CircAPI.
//
//     func TestSomethingThatUsesCircAPI(t *testing.T) {
//
//         // make and configure a mocked CircAPI
//         mockedCircAPI := &CircAPIMock{
//             CreateCheckBundleFunc: func(cfg *api.CheckBundle) (*api.CheckBundle, error) {
// 	               panic("TODO: mock out the CreateCheckBundle method")
//             },
//             CreateDashboardFunc: func(cfg *api.Dashboard) (*api.Dashboard, error) {
// 	               panic("TODO: mock out the CreateDashboard method")
//             },
//             CreateGraphFunc: func(cfg *api.Graph) (*api.Graph, error) {
// 	               panic("TODO: mock out the CreateGraph method")
//             },
//             CreateRuleSetFunc: func(cfg *api.RuleSet) (*api.RuleSet, error) {
// 	               panic("TODO: mock out the CreateRuleSet method")
//             },
//             CreateWorksheetFunc: func(cfg *api.Worksheet) (*api.Worksheet, error) {
// 	               panic("TODO: mock out the CreateWorksheet method")
//             },
//             DeleteCheckBundleByCIDFunc: func(cid api.CIDType) (bool, error) {
// 	               panic("TODO: mock out the DeleteCheckBundleByCID method")
//             },
//             DeleteDashboardByCIDFunc: func(cid api.CIDType) (bool, error) {
// 	               panic("TODO: mock out the DeleteDashboardByCID method")
//             },
//             DeleteGraphByCIDFunc: func(cid api.CIDType) (bool, error) {
// 	               panic("TODO: mock out the DeleteGraphByCID method")
//             },
//             DeleteWorksheetByCIDFunc: func(cid api.CIDType) (bool, error) {
// 	               panic("TODO: mock out the DeleteWorksheetByCID method")
//             },
//             FetchBrokerFunc: func(cid api.CIDType) (*api.Broker, error) {
// 	               panic("TODO: mock out the FetchBroker method")
//             },
//             FetchBrokersFunc: func() (*[]api.Broker, error) {
// 	               panic("TODO: mock out the FetchBrokers method")
//             },
//             FetchCheckBundleFunc: func(cid api.CIDType) (*api.CheckBundle, error) {
// 	               panic("TODO: mock out the FetchCheckBundle method")
//             },
//             FetchDashboardFunc: func(cid api.CIDType) (*api.Dashboard, error) {
// 	               panic("TODO: mock out the FetchDashboard method")
//             },
//             FetchGraphFunc: func(cid api.CIDType) (*api.Graph, error) {
// 	               panic("TODO: mock out the FetchGraph method")
//             },
//             FetchWorksheetFunc: func(cid api.CIDType) (*api.Worksheet, error) {
// 	               panic("TODO: mock out the FetchWorksheet method")
//             },
//             SearchCheckBundlesFunc: func(searchCriteria *api.SearchQueryType, filterCriteria *map[string][]string) (*[]api.CheckBundle, error) {
// 	               panic("TODO: mock out the SearchCheckBundles method")
//             },
//             SearchDashboardsFunc: func(searchCriteria *api.SearchQueryType, filterCriteria *api.SearchFilterType) (*[]api.Dashboard, error) {
// 	               panic("TODO: mock out the SearchDashboards method")
//             },
//             SearchGraphsFunc: func(searchCriteria *api.SearchQueryType, filterCriteria *api.SearchFilterType) (*[]api.Graph, error) {
// 	               panic("TODO: mock out the SearchGraphs method")
//             },
//             SearchWorksheetsFunc: func(searchCriteria *api.SearchQueryType, filterCriteria *api.SearchFilterType) (*[]api.Worksheet, error) {
// 	               panic("TODO: mock out the SearchWorksheets method")
//             },
//             UpdateCheckBundleFunc: func(cfg *api.CheckBundle) (*api.CheckBundle, error) {
// 	               panic("TODO: mock out the UpdateCheckBundle method")
//             },
//             UpdateDashboardFunc: func(cfg *api.Dashboard) (*api.Dashboard, error) {
// 	               panic("TODO: mock out the UpdateDashboard method")
//             },
//             UpdateGraphFunc: func(cfg *api.Graph) (*api.Graph, error) {
// 	               panic("TODO: mock out the UpdateGraph method")
//             },
//             UpdateWorksheetFunc: func(cfg *api.Worksheet) (*api.Worksheet, error) {
// 	               panic("TODO: mock out the UpdateWorksheet method")
//             },
//         }
//
//         // TODO: use mockedCircAPI in code that requires CircAPI
//         //       and then make assertions.
//
//     }
type CircAPIMock struct {
	// CreateCheckBundleFunc mocks the CreateCheckBundle method.
	CreateCheckBundleFunc func(cfg *api.CheckBundle) (*api.CheckBundle, error)

	// CreateDashboardFunc mocks the CreateDashboard method.
	CreateDashboardFunc func(cfg *api.Dashboard) (*api.Dashboard, error)

	// CreateGraphFunc mocks the CreateGraph method.
	CreateGraphFunc func(cfg *api.Graph) (*api.Graph, error)

	// CreateRuleSetFunc mocks the CreateRuleSet method.
	CreateRuleSetFunc func(cfg *api.RuleSet) (*api.RuleSet, error)

	// CreateWorksheetFunc mocks the CreateWorksheet method.
	CreateWorksheetFunc func(cfg *api.Worksheet) (*api.Worksheet, error)

	// DeleteCheckBundleByCIDFunc mocks the DeleteCheckBundleByCID method.
	DeleteCheckBundleByCIDFunc func(cid api.CIDType) (bool, error)

	// DeleteDashboardByCIDFunc mocks the DeleteDashboardByCID method.
	DeleteDashboardByCIDFunc func(cid api.CIDType) (bool, error)

	// DeleteGraphByCIDFunc mocks the DeleteGraphByCID method.
	DeleteGraphByCIDFunc func(cid api.CIDType) (bool, error)

	// DeleteWorksheetByCIDFunc mocks the DeleteWorksheetByCID method.
	DeleteWorksheetByCIDFunc func(cid api.CIDType) (bool, error)

	// FetchBrokerFunc mocks the FetchBroker method.
	FetchBrokerFunc func(cid api.CIDType) (*api.Broker, error)

	// FetchBrokersFunc mocks the FetchBrokers method.
	FetchBrokersFunc func() (*[]api.Broker, error)

	// FetchCheckBundleFunc mocks the FetchCheckBundle method.
	FetchCheckBundleFunc func(cid api.CIDType) (*api.CheckBundle, error)

	// FetchDashboardFunc mocks the FetchDashboard method.
	FetchDashboardFunc func(cid api.CIDType) (*api.Dashboard, error)

	// FetchGraphFunc mocks the FetchGraph method.
	FetchGraphFunc func(cid api.CIDType) (*api.Graph, error)

	// FetchWorksheetFunc mocks the FetchWorksheet method.
	FetchWorksheetFunc func(cid api.CIDType) (*api.Worksheet, error)

	// SearchCheckBundlesFunc mocks the SearchCheckBundles method.
	SearchCheckBundlesFunc func(searchCriteria *api.SearchQueryType, filterCriteria *map[string][]string) (*[]api.CheckBundle, error)

	// SearchDashboardsFunc mocks the SearchDashboards method.
	SearchDashboardsFunc func(searchCriteria *api.SearchQueryType, filterCriteria *api.SearchFilterType) (*[]api.Dashboard, error)

	// SearchGraphsFunc mocks the SearchGraphs method.
	SearchGraphsFunc func(searchCriteria *api.SearchQueryType, filterCriteria *api.SearchFilterType) (*[]api.Graph, error)

	// SearchWorksheetsFunc mocks the SearchWorksheets method.
	SearchWorksheetsFunc func(searchCriteria *api.SearchQueryType, filterCriteria *api.SearchFilterType) (*[]api.Worksheet, error)

	// UpdateCheckBundleFunc mocks the UpdateCheckBundle method.
	UpdateCheckBundleFunc func(cfg *api.CheckBundle) (*api.CheckBundle, error)

	// UpdateDashboardFunc mocks the UpdateDashboard method.
	UpdateDashboardFunc func(cfg *api.Dashboard) (*api.Dashboard, error)

	// UpdateGraphFunc mocks the UpdateGraph method.
	UpdateGraphFunc func(cfg *api.Graph) (*api.Graph, error)

	// UpdateWorksheetFunc mocks the UpdateWorksheet method.
	UpdateWorksheetFunc func(cfg *api.Worksheet) (*api.Worksheet, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateCheckBundle holds details about calls to the CreateCheckBundle method.
		CreateCheckBundle []struct {
			// Cfg is the cfg argument value.
			Cfg *api.CheckBundle
		}
		// CreateDashboard holds details about calls to the CreateDashboard method.
		CreateDashboard []struct {
			// Cfg is the cfg argument value.
			Cfg *api.Dashboard
		}
		// CreateGraph holds details about calls to the CreateGraph method.
		CreateGraph []struct {
			// Cfg is the cfg argument value.
			Cfg *api.Graph
		}
		// CreateRuleSet holds details about calls to the CreateRuleSet method.
		CreateRuleSet []struct {
			// Cfg is the cfg argument value.
			Cfg *api.RuleSet
		}
		// CreateWorksheet holds details about calls to the CreateWorksheet method.
		CreateWorksheet []struct {
			// Cfg is the cfg argument value.
			Cfg *api.Worksheet
		}
		// DeleteCheckBundleByCID holds details about calls to the DeleteCheckBundleByCID method.
		DeleteCheckBundleByCID []struct {
			// Cid is the cid argument value.
			Cid api.CIDType
		}
		// DeleteDashboardByCID holds details about calls to the DeleteDashboardByCID method.
		DeleteDashboardByCID []struct {
			// Cid is the cid argument value.
			Cid api.CIDType
		}
		// DeleteGraphByCID holds details about calls to the DeleteGraphByCID method.
		DeleteGraphByCID []struct {
			// Cid is the cid argument value.
			Cid api.CIDType
		}
		// DeleteWorksheetByCID holds details about calls to the DeleteWorksheetByCID method.
		DeleteWorksheetByCID []struct {
			// Cid is the cid argument value.
			Cid api.CIDType
		}
		// FetchBroker holds details about calls to the FetchBroker method.
		FetchBroker []struct {
			// Cid is the cid argument value.
			Cid api.CIDType
		}
		// FetchBrokers holds details about calls to the FetchBrokers method.
		FetchBrokers []struct {
		}
		// FetchCheckBundle holds details about calls to the FetchCheckBundle method.
		FetchCheckBundle []struct {
			// Cid is the cid argument value.
			Cid api.CIDType
		}
		// FetchDashboard holds details about calls to the FetchDashboard method.
		FetchDashboard []struct {
			// Cid is the cid argument value.
			Cid api.CIDType
		}
		// FetchGraph holds details about calls to the FetchGraph method.
		FetchGraph []struct {
			// Cid is the cid argument value.
			Cid api.CIDType
		}
		// FetchWorksheet holds details about calls to the FetchWorksheet method.
		FetchWorksheet []struct {
			// Cid is the cid argument value.
			Cid api.CIDType
		}
		// SearchCheckBundles holds details about calls to the SearchCheckBundles method.
		SearchCheckBundles []struct {
			// SearchCriteria is the searchCriteria argument value.
			SearchCriteria *api.SearchQueryType
			// FilterCriteria is the filterCriteria argument value.
			FilterCriteria *map[string][]string
		}
		// SearchDashboards holds details about calls to the SearchDashboards method.
		SearchDashboards []struct {
			// SearchCriteria is the searchCriteria argument value.
			SearchCriteria *api.SearchQueryType
			// FilterCriteria is the filterCriteria argument value.
			FilterCriteria *api.SearchFilterType
		}
		// SearchGraphs holds details about calls to the SearchGraphs method.
		SearchGraphs []struct {
			// SearchCriteria is the searchCriteria argument value.
			SearchCriteria *api.SearchQueryType
			// FilterCriteria is the filterCriteria argument value.
			FilterCriteria *api.SearchFilterType
		}
		// SearchWorksheets holds details about calls to the SearchWorksheets method.
		SearchWorksheets []struct {
			// SearchCriteria is the searchCriteria argument value.
			SearchCriteria *api.SearchQueryType
			// FilterCriteria is the filterCriteria argument value.
			FilterCriteria *api.SearchFilterType
		}
		// UpdateCheckBundle holds details about calls to the UpdateCheckBundle method.
		UpdateCheckBundle []struct {
			// Cfg is the cfg argument value.
			Cfg *api.CheckBundle
		}
		// UpdateDashboard holds details about calls to the UpdateDashboard method.
		UpdateDashboard []struct {
			// Cfg is the cfg argument value.
			Cfg *api.Dashboard
		}
		// UpdateGraph holds details about calls to the UpdateGraph method.
		UpdateGraph []struct {
			// Cfg is the cfg argument value.
			Cfg *api.Graph
		}
		// UpdateWorksheet holds details about calls to the UpdateWorksheet method.
		UpdateWorksheet []struct {
			// Cfg is the cfg argument value.
			Cfg *api.Worksheet
		}
	}
}

// CreateCheckBundle calls CreateCheckBundleFunc.
func (mock *CircAPIMock) CreateCheckBundle(cfg *api.CheckBundle) (*api.CheckBundle, error) {
	if mock.CreateCheckBundleFunc == nil {
		panic("moq: CircAPIMock.CreateCheckBundleFunc is nil but CircAPI.CreateCheckBundle was just called")
	}
	callInfo := struct {
		Cfg *api.CheckBundle
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
	Cfg *api.CheckBundle
} {
	var calls []struct {
		Cfg *api.CheckBundle
	}
	lockCircAPIMockCreateCheckBundle.RLock()
	calls = mock.calls.CreateCheckBundle
	lockCircAPIMockCreateCheckBundle.RUnlock()
	return calls
}

// CreateDashboard calls CreateDashboardFunc.
func (mock *CircAPIMock) CreateDashboard(cfg *api.Dashboard) (*api.Dashboard, error) {
	if mock.CreateDashboardFunc == nil {
		panic("moq: CircAPIMock.CreateDashboardFunc is nil but CircAPI.CreateDashboard was just called")
	}
	callInfo := struct {
		Cfg *api.Dashboard
	}{
		Cfg: cfg,
	}
	lockCircAPIMockCreateDashboard.Lock()
	mock.calls.CreateDashboard = append(mock.calls.CreateDashboard, callInfo)
	lockCircAPIMockCreateDashboard.Unlock()
	return mock.CreateDashboardFunc(cfg)
}

// CreateDashboardCalls gets all the calls that were made to CreateDashboard.
// Check the length with:
//     len(mockedCircAPI.CreateDashboardCalls())
func (mock *CircAPIMock) CreateDashboardCalls() []struct {
	Cfg *api.Dashboard
} {
	var calls []struct {
		Cfg *api.Dashboard
	}
	lockCircAPIMockCreateDashboard.RLock()
	calls = mock.calls.CreateDashboard
	lockCircAPIMockCreateDashboard.RUnlock()
	return calls
}

// CreateGraph calls CreateGraphFunc.
func (mock *CircAPIMock) CreateGraph(cfg *api.Graph) (*api.Graph, error) {
	if mock.CreateGraphFunc == nil {
		panic("moq: CircAPIMock.CreateGraphFunc is nil but CircAPI.CreateGraph was just called")
	}
	callInfo := struct {
		Cfg *api.Graph
	}{
		Cfg: cfg,
	}
	lockCircAPIMockCreateGraph.Lock()
	mock.calls.CreateGraph = append(mock.calls.CreateGraph, callInfo)
	lockCircAPIMockCreateGraph.Unlock()
	return mock.CreateGraphFunc(cfg)
}

// CreateGraphCalls gets all the calls that were made to CreateGraph.
// Check the length with:
//     len(mockedCircAPI.CreateGraphCalls())
func (mock *CircAPIMock) CreateGraphCalls() []struct {
	Cfg *api.Graph
} {
	var calls []struct {
		Cfg *api.Graph
	}
	lockCircAPIMockCreateGraph.RLock()
	calls = mock.calls.CreateGraph
	lockCircAPIMockCreateGraph.RUnlock()
	return calls
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

// CreateWorksheet calls CreateWorksheetFunc.
func (mock *CircAPIMock) CreateWorksheet(cfg *api.Worksheet) (*api.Worksheet, error) {
	if mock.CreateWorksheetFunc == nil {
		panic("moq: CircAPIMock.CreateWorksheetFunc is nil but CircAPI.CreateWorksheet was just called")
	}
	callInfo := struct {
		Cfg *api.Worksheet
	}{
		Cfg: cfg,
	}
	lockCircAPIMockCreateWorksheet.Lock()
	mock.calls.CreateWorksheet = append(mock.calls.CreateWorksheet, callInfo)
	lockCircAPIMockCreateWorksheet.Unlock()
	return mock.CreateWorksheetFunc(cfg)
}

// CreateWorksheetCalls gets all the calls that were made to CreateWorksheet.
// Check the length with:
//     len(mockedCircAPI.CreateWorksheetCalls())
func (mock *CircAPIMock) CreateWorksheetCalls() []struct {
	Cfg *api.Worksheet
} {
	var calls []struct {
		Cfg *api.Worksheet
	}
	lockCircAPIMockCreateWorksheet.RLock()
	calls = mock.calls.CreateWorksheet
	lockCircAPIMockCreateWorksheet.RUnlock()
	return calls
}

// DeleteCheckBundleByCID calls DeleteCheckBundleByCIDFunc.
func (mock *CircAPIMock) DeleteCheckBundleByCID(cid api.CIDType) (bool, error) {
	if mock.DeleteCheckBundleByCIDFunc == nil {
		panic("moq: CircAPIMock.DeleteCheckBundleByCIDFunc is nil but CircAPI.DeleteCheckBundleByCID was just called")
	}
	callInfo := struct {
		Cid api.CIDType
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
	Cid api.CIDType
} {
	var calls []struct {
		Cid api.CIDType
	}
	lockCircAPIMockDeleteCheckBundleByCID.RLock()
	calls = mock.calls.DeleteCheckBundleByCID
	lockCircAPIMockDeleteCheckBundleByCID.RUnlock()
	return calls
}

// DeleteDashboardByCID calls DeleteDashboardByCIDFunc.
func (mock *CircAPIMock) DeleteDashboardByCID(cid api.CIDType) (bool, error) {
	if mock.DeleteDashboardByCIDFunc == nil {
		panic("moq: CircAPIMock.DeleteDashboardByCIDFunc is nil but CircAPI.DeleteDashboardByCID was just called")
	}
	callInfo := struct {
		Cid api.CIDType
	}{
		Cid: cid,
	}
	lockCircAPIMockDeleteDashboardByCID.Lock()
	mock.calls.DeleteDashboardByCID = append(mock.calls.DeleteDashboardByCID, callInfo)
	lockCircAPIMockDeleteDashboardByCID.Unlock()
	return mock.DeleteDashboardByCIDFunc(cid)
}

// DeleteDashboardByCIDCalls gets all the calls that were made to DeleteDashboardByCID.
// Check the length with:
//     len(mockedCircAPI.DeleteDashboardByCIDCalls())
func (mock *CircAPIMock) DeleteDashboardByCIDCalls() []struct {
	Cid api.CIDType
} {
	var calls []struct {
		Cid api.CIDType
	}
	lockCircAPIMockDeleteDashboardByCID.RLock()
	calls = mock.calls.DeleteDashboardByCID
	lockCircAPIMockDeleteDashboardByCID.RUnlock()
	return calls
}

// DeleteGraphByCID calls DeleteGraphByCIDFunc.
func (mock *CircAPIMock) DeleteGraphByCID(cid api.CIDType) (bool, error) {
	if mock.DeleteGraphByCIDFunc == nil {
		panic("moq: CircAPIMock.DeleteGraphByCIDFunc is nil but CircAPI.DeleteGraphByCID was just called")
	}
	callInfo := struct {
		Cid api.CIDType
	}{
		Cid: cid,
	}
	lockCircAPIMockDeleteGraphByCID.Lock()
	mock.calls.DeleteGraphByCID = append(mock.calls.DeleteGraphByCID, callInfo)
	lockCircAPIMockDeleteGraphByCID.Unlock()
	return mock.DeleteGraphByCIDFunc(cid)
}

// DeleteGraphByCIDCalls gets all the calls that were made to DeleteGraphByCID.
// Check the length with:
//     len(mockedCircAPI.DeleteGraphByCIDCalls())
func (mock *CircAPIMock) DeleteGraphByCIDCalls() []struct {
	Cid api.CIDType
} {
	var calls []struct {
		Cid api.CIDType
	}
	lockCircAPIMockDeleteGraphByCID.RLock()
	calls = mock.calls.DeleteGraphByCID
	lockCircAPIMockDeleteGraphByCID.RUnlock()
	return calls
}

// DeleteWorksheetByCID calls DeleteWorksheetByCIDFunc.
func (mock *CircAPIMock) DeleteWorksheetByCID(cid api.CIDType) (bool, error) {
	if mock.DeleteWorksheetByCIDFunc == nil {
		panic("moq: CircAPIMock.DeleteWorksheetByCIDFunc is nil but CircAPI.DeleteWorksheetByCID was just called")
	}
	callInfo := struct {
		Cid api.CIDType
	}{
		Cid: cid,
	}
	lockCircAPIMockDeleteWorksheetByCID.Lock()
	mock.calls.DeleteWorksheetByCID = append(mock.calls.DeleteWorksheetByCID, callInfo)
	lockCircAPIMockDeleteWorksheetByCID.Unlock()
	return mock.DeleteWorksheetByCIDFunc(cid)
}

// DeleteWorksheetByCIDCalls gets all the calls that were made to DeleteWorksheetByCID.
// Check the length with:
//     len(mockedCircAPI.DeleteWorksheetByCIDCalls())
func (mock *CircAPIMock) DeleteWorksheetByCIDCalls() []struct {
	Cid api.CIDType
} {
	var calls []struct {
		Cid api.CIDType
	}
	lockCircAPIMockDeleteWorksheetByCID.RLock()
	calls = mock.calls.DeleteWorksheetByCID
	lockCircAPIMockDeleteWorksheetByCID.RUnlock()
	return calls
}

// FetchBroker calls FetchBrokerFunc.
func (mock *CircAPIMock) FetchBroker(cid api.CIDType) (*api.Broker, error) {
	if mock.FetchBrokerFunc == nil {
		panic("moq: CircAPIMock.FetchBrokerFunc is nil but CircAPI.FetchBroker was just called")
	}
	callInfo := struct {
		Cid api.CIDType
	}{
		Cid: cid,
	}
	lockCircAPIMockFetchBroker.Lock()
	mock.calls.FetchBroker = append(mock.calls.FetchBroker, callInfo)
	lockCircAPIMockFetchBroker.Unlock()
	return mock.FetchBrokerFunc(cid)
}

// FetchBrokerCalls gets all the calls that were made to FetchBroker.
// Check the length with:
//     len(mockedCircAPI.FetchBrokerCalls())
func (mock *CircAPIMock) FetchBrokerCalls() []struct {
	Cid api.CIDType
} {
	var calls []struct {
		Cid api.CIDType
	}
	lockCircAPIMockFetchBroker.RLock()
	calls = mock.calls.FetchBroker
	lockCircAPIMockFetchBroker.RUnlock()
	return calls
}

// FetchBrokers calls FetchBrokersFunc.
func (mock *CircAPIMock) FetchBrokers() (*[]api.Broker, error) {
	if mock.FetchBrokersFunc == nil {
		panic("moq: CircAPIMock.FetchBrokersFunc is nil but CircAPI.FetchBrokers was just called")
	}
	callInfo := struct {
	}{}
	lockCircAPIMockFetchBrokers.Lock()
	mock.calls.FetchBrokers = append(mock.calls.FetchBrokers, callInfo)
	lockCircAPIMockFetchBrokers.Unlock()
	return mock.FetchBrokersFunc()
}

// FetchBrokersCalls gets all the calls that were made to FetchBrokers.
// Check the length with:
//     len(mockedCircAPI.FetchBrokersCalls())
func (mock *CircAPIMock) FetchBrokersCalls() []struct {
} {
	var calls []struct {
	}
	lockCircAPIMockFetchBrokers.RLock()
	calls = mock.calls.FetchBrokers
	lockCircAPIMockFetchBrokers.RUnlock()
	return calls
}

// FetchCheckBundle calls FetchCheckBundleFunc.
func (mock *CircAPIMock) FetchCheckBundle(cid api.CIDType) (*api.CheckBundle, error) {
	if mock.FetchCheckBundleFunc == nil {
		panic("moq: CircAPIMock.FetchCheckBundleFunc is nil but CircAPI.FetchCheckBundle was just called")
	}
	callInfo := struct {
		Cid api.CIDType
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
	Cid api.CIDType
} {
	var calls []struct {
		Cid api.CIDType
	}
	lockCircAPIMockFetchCheckBundle.RLock()
	calls = mock.calls.FetchCheckBundle
	lockCircAPIMockFetchCheckBundle.RUnlock()
	return calls
}

// FetchDashboard calls FetchDashboardFunc.
func (mock *CircAPIMock) FetchDashboard(cid api.CIDType) (*api.Dashboard, error) {
	if mock.FetchDashboardFunc == nil {
		panic("moq: CircAPIMock.FetchDashboardFunc is nil but CircAPI.FetchDashboard was just called")
	}
	callInfo := struct {
		Cid api.CIDType
	}{
		Cid: cid,
	}
	lockCircAPIMockFetchDashboard.Lock()
	mock.calls.FetchDashboard = append(mock.calls.FetchDashboard, callInfo)
	lockCircAPIMockFetchDashboard.Unlock()
	return mock.FetchDashboardFunc(cid)
}

// FetchDashboardCalls gets all the calls that were made to FetchDashboard.
// Check the length with:
//     len(mockedCircAPI.FetchDashboardCalls())
func (mock *CircAPIMock) FetchDashboardCalls() []struct {
	Cid api.CIDType
} {
	var calls []struct {
		Cid api.CIDType
	}
	lockCircAPIMockFetchDashboard.RLock()
	calls = mock.calls.FetchDashboard
	lockCircAPIMockFetchDashboard.RUnlock()
	return calls
}

// FetchGraph calls FetchGraphFunc.
func (mock *CircAPIMock) FetchGraph(cid api.CIDType) (*api.Graph, error) {
	if mock.FetchGraphFunc == nil {
		panic("moq: CircAPIMock.FetchGraphFunc is nil but CircAPI.FetchGraph was just called")
	}
	callInfo := struct {
		Cid api.CIDType
	}{
		Cid: cid,
	}
	lockCircAPIMockFetchGraph.Lock()
	mock.calls.FetchGraph = append(mock.calls.FetchGraph, callInfo)
	lockCircAPIMockFetchGraph.Unlock()
	return mock.FetchGraphFunc(cid)
}

// FetchGraphCalls gets all the calls that were made to FetchGraph.
// Check the length with:
//     len(mockedCircAPI.FetchGraphCalls())
func (mock *CircAPIMock) FetchGraphCalls() []struct {
	Cid api.CIDType
} {
	var calls []struct {
		Cid api.CIDType
	}
	lockCircAPIMockFetchGraph.RLock()
	calls = mock.calls.FetchGraph
	lockCircAPIMockFetchGraph.RUnlock()
	return calls
}

// FetchWorksheet calls FetchWorksheetFunc.
func (mock *CircAPIMock) FetchWorksheet(cid api.CIDType) (*api.Worksheet, error) {
	if mock.FetchWorksheetFunc == nil {
		panic("moq: CircAPIMock.FetchWorksheetFunc is nil but CircAPI.FetchWorksheet was just called")
	}
	callInfo := struct {
		Cid api.CIDType
	}{
		Cid: cid,
	}
	lockCircAPIMockFetchWorksheet.Lock()
	mock.calls.FetchWorksheet = append(mock.calls.FetchWorksheet, callInfo)
	lockCircAPIMockFetchWorksheet.Unlock()
	return mock.FetchWorksheetFunc(cid)
}

// FetchWorksheetCalls gets all the calls that were made to FetchWorksheet.
// Check the length with:
//     len(mockedCircAPI.FetchWorksheetCalls())
func (mock *CircAPIMock) FetchWorksheetCalls() []struct {
	Cid api.CIDType
} {
	var calls []struct {
		Cid api.CIDType
	}
	lockCircAPIMockFetchWorksheet.RLock()
	calls = mock.calls.FetchWorksheet
	lockCircAPIMockFetchWorksheet.RUnlock()
	return calls
}

// SearchCheckBundles calls SearchCheckBundlesFunc.
func (mock *CircAPIMock) SearchCheckBundles(searchCriteria *api.SearchQueryType, filterCriteria *map[string][]string) (*[]api.CheckBundle, error) {
	if mock.SearchCheckBundlesFunc == nil {
		panic("moq: CircAPIMock.SearchCheckBundlesFunc is nil but CircAPI.SearchCheckBundles was just called")
	}
	callInfo := struct {
		SearchCriteria *api.SearchQueryType
		FilterCriteria *map[string][]string
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
	SearchCriteria *api.SearchQueryType
	FilterCriteria *map[string][]string
} {
	var calls []struct {
		SearchCriteria *api.SearchQueryType
		FilterCriteria *map[string][]string
	}
	lockCircAPIMockSearchCheckBundles.RLock()
	calls = mock.calls.SearchCheckBundles
	lockCircAPIMockSearchCheckBundles.RUnlock()
	return calls
}

// SearchDashboards calls SearchDashboardsFunc.
func (mock *CircAPIMock) SearchDashboards(searchCriteria *api.SearchQueryType, filterCriteria *api.SearchFilterType) (*[]api.Dashboard, error) {
	if mock.SearchDashboardsFunc == nil {
		panic("moq: CircAPIMock.SearchDashboardsFunc is nil but CircAPI.SearchDashboards was just called")
	}
	callInfo := struct {
		SearchCriteria *api.SearchQueryType
		FilterCriteria *api.SearchFilterType
	}{
		SearchCriteria: searchCriteria,
		FilterCriteria: filterCriteria,
	}
	lockCircAPIMockSearchDashboards.Lock()
	mock.calls.SearchDashboards = append(mock.calls.SearchDashboards, callInfo)
	lockCircAPIMockSearchDashboards.Unlock()
	return mock.SearchDashboardsFunc(searchCriteria, filterCriteria)
}

// SearchDashboardsCalls gets all the calls that were made to SearchDashboards.
// Check the length with:
//     len(mockedCircAPI.SearchDashboardsCalls())
func (mock *CircAPIMock) SearchDashboardsCalls() []struct {
	SearchCriteria *api.SearchQueryType
	FilterCriteria *api.SearchFilterType
} {
	var calls []struct {
		SearchCriteria *api.SearchQueryType
		FilterCriteria *api.SearchFilterType
	}
	lockCircAPIMockSearchDashboards.RLock()
	calls = mock.calls.SearchDashboards
	lockCircAPIMockSearchDashboards.RUnlock()
	return calls
}

// SearchGraphs calls SearchGraphsFunc.
func (mock *CircAPIMock) SearchGraphs(searchCriteria *api.SearchQueryType, filterCriteria *api.SearchFilterType) (*[]api.Graph, error) {
	if mock.SearchGraphsFunc == nil {
		panic("moq: CircAPIMock.SearchGraphsFunc is nil but CircAPI.SearchGraphs was just called")
	}
	callInfo := struct {
		SearchCriteria *api.SearchQueryType
		FilterCriteria *api.SearchFilterType
	}{
		SearchCriteria: searchCriteria,
		FilterCriteria: filterCriteria,
	}
	lockCircAPIMockSearchGraphs.Lock()
	mock.calls.SearchGraphs = append(mock.calls.SearchGraphs, callInfo)
	lockCircAPIMockSearchGraphs.Unlock()
	return mock.SearchGraphsFunc(searchCriteria, filterCriteria)
}

// SearchGraphsCalls gets all the calls that were made to SearchGraphs.
// Check the length with:
//     len(mockedCircAPI.SearchGraphsCalls())
func (mock *CircAPIMock) SearchGraphsCalls() []struct {
	SearchCriteria *api.SearchQueryType
	FilterCriteria *api.SearchFilterType
} {
	var calls []struct {
		SearchCriteria *api.SearchQueryType
		FilterCriteria *api.SearchFilterType
	}
	lockCircAPIMockSearchGraphs.RLock()
	calls = mock.calls.SearchGraphs
	lockCircAPIMockSearchGraphs.RUnlock()
	return calls
}

// SearchWorksheets calls SearchWorksheetsFunc.
func (mock *CircAPIMock) SearchWorksheets(searchCriteria *api.SearchQueryType, filterCriteria *api.SearchFilterType) (*[]api.Worksheet, error) {
	if mock.SearchWorksheetsFunc == nil {
		panic("moq: CircAPIMock.SearchWorksheetsFunc is nil but CircAPI.SearchWorksheets was just called")
	}
	callInfo := struct {
		SearchCriteria *api.SearchQueryType
		FilterCriteria *api.SearchFilterType
	}{
		SearchCriteria: searchCriteria,
		FilterCriteria: filterCriteria,
	}
	lockCircAPIMockSearchWorksheets.Lock()
	mock.calls.SearchWorksheets = append(mock.calls.SearchWorksheets, callInfo)
	lockCircAPIMockSearchWorksheets.Unlock()
	return mock.SearchWorksheetsFunc(searchCriteria, filterCriteria)
}

// SearchWorksheetsCalls gets all the calls that were made to SearchWorksheets.
// Check the length with:
//     len(mockedCircAPI.SearchWorksheetsCalls())
func (mock *CircAPIMock) SearchWorksheetsCalls() []struct {
	SearchCriteria *api.SearchQueryType
	FilterCriteria *api.SearchFilterType
} {
	var calls []struct {
		SearchCriteria *api.SearchQueryType
		FilterCriteria *api.SearchFilterType
	}
	lockCircAPIMockSearchWorksheets.RLock()
	calls = mock.calls.SearchWorksheets
	lockCircAPIMockSearchWorksheets.RUnlock()
	return calls
}

// UpdateCheckBundle calls UpdateCheckBundleFunc.
func (mock *CircAPIMock) UpdateCheckBundle(cfg *api.CheckBundle) (*api.CheckBundle, error) {
	if mock.UpdateCheckBundleFunc == nil {
		panic("moq: CircAPIMock.UpdateCheckBundleFunc is nil but CircAPI.UpdateCheckBundle was just called")
	}
	callInfo := struct {
		Cfg *api.CheckBundle
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
	Cfg *api.CheckBundle
} {
	var calls []struct {
		Cfg *api.CheckBundle
	}
	lockCircAPIMockUpdateCheckBundle.RLock()
	calls = mock.calls.UpdateCheckBundle
	lockCircAPIMockUpdateCheckBundle.RUnlock()
	return calls
}

// UpdateDashboard calls UpdateDashboardFunc.
func (mock *CircAPIMock) UpdateDashboard(cfg *api.Dashboard) (*api.Dashboard, error) {
	if mock.UpdateDashboardFunc == nil {
		panic("moq: CircAPIMock.UpdateDashboardFunc is nil but CircAPI.UpdateDashboard was just called")
	}
	callInfo := struct {
		Cfg *api.Dashboard
	}{
		Cfg: cfg,
	}
	lockCircAPIMockUpdateDashboard.Lock()
	mock.calls.UpdateDashboard = append(mock.calls.UpdateDashboard, callInfo)
	lockCircAPIMockUpdateDashboard.Unlock()
	return mock.UpdateDashboardFunc(cfg)
}

// UpdateDashboardCalls gets all the calls that were made to UpdateDashboard.
// Check the length with:
//     len(mockedCircAPI.UpdateDashboardCalls())
func (mock *CircAPIMock) UpdateDashboardCalls() []struct {
	Cfg *api.Dashboard
} {
	var calls []struct {
		Cfg *api.Dashboard
	}
	lockCircAPIMockUpdateDashboard.RLock()
	calls = mock.calls.UpdateDashboard
	lockCircAPIMockUpdateDashboard.RUnlock()
	return calls
}

// UpdateGraph calls UpdateGraphFunc.
func (mock *CircAPIMock) UpdateGraph(cfg *api.Graph) (*api.Graph, error) {
	if mock.UpdateGraphFunc == nil {
		panic("moq: CircAPIMock.UpdateGraphFunc is nil but CircAPI.UpdateGraph was just called")
	}
	callInfo := struct {
		Cfg *api.Graph
	}{
		Cfg: cfg,
	}
	lockCircAPIMockUpdateGraph.Lock()
	mock.calls.UpdateGraph = append(mock.calls.UpdateGraph, callInfo)
	lockCircAPIMockUpdateGraph.Unlock()
	return mock.UpdateGraphFunc(cfg)
}

// UpdateGraphCalls gets all the calls that were made to UpdateGraph.
// Check the length with:
//     len(mockedCircAPI.UpdateGraphCalls())
func (mock *CircAPIMock) UpdateGraphCalls() []struct {
	Cfg *api.Graph
} {
	var calls []struct {
		Cfg *api.Graph
	}
	lockCircAPIMockUpdateGraph.RLock()
	calls = mock.calls.UpdateGraph
	lockCircAPIMockUpdateGraph.RUnlock()
	return calls
}

// UpdateWorksheet calls UpdateWorksheetFunc.
func (mock *CircAPIMock) UpdateWorksheet(cfg *api.Worksheet) (*api.Worksheet, error) {
	if mock.UpdateWorksheetFunc == nil {
		panic("moq: CircAPIMock.UpdateWorksheetFunc is nil but CircAPI.UpdateWorksheet was just called")
	}
	callInfo := struct {
		Cfg *api.Worksheet
	}{
		Cfg: cfg,
	}
	lockCircAPIMockUpdateWorksheet.Lock()
	mock.calls.UpdateWorksheet = append(mock.calls.UpdateWorksheet, callInfo)
	lockCircAPIMockUpdateWorksheet.Unlock()
	return mock.UpdateWorksheetFunc(cfg)
}

// UpdateWorksheetCalls gets all the calls that were made to UpdateWorksheet.
// Check the length with:
//     len(mockedCircAPI.UpdateWorksheetCalls())
func (mock *CircAPIMock) UpdateWorksheetCalls() []struct {
	Cfg *api.Worksheet
} {
	var calls []struct {
		Cfg *api.Worksheet
	}
	lockCircAPIMockUpdateWorksheet.RLock()
	calls = mock.calls.UpdateWorksheet
	lockCircAPIMockUpdateWorksheet.RUnlock()
	return calls
}
