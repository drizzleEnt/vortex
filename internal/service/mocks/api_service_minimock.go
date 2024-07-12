package mocks

// Code generated by http://github.com/gojuno/minimock (3.1.3). DO NOT EDIT.

//go:generate minimock -i github.com/drizzleent/vortex/internal/service.ApiService -o ./mocks/api_service_minimock.go -n ApiServiceMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/drizzleent/vortex/internal/model"
	"github.com/gojuno/minimock/v3"
)

// ApiServiceMock implements service.ApiService
type ApiServiceMock struct {
	t minimock.Tester

	funcAddClient          func(ctx context.Context, cp1 *model.Client) (i1 int64, err error)
	inspectFuncAddClient   func(ctx context.Context, cp1 *model.Client)
	afterAddClientCounter  uint64
	beforeAddClientCounter uint64
	AddClientMock          mApiServiceMockAddClient

	funcDeleteClient          func(ctx context.Context, i1 int64) (err error)
	inspectFuncDeleteClient   func(ctx context.Context, i1 int64)
	afterDeleteClientCounter  uint64
	beforeDeleteClientCounter uint64
	DeleteClientMock          mApiServiceMockDeleteClient

	funcUpdateAlgorithmStatus          func(ctx context.Context, i1 int64, ap1 *model.Algorithms) (err error)
	inspectFuncUpdateAlgorithmStatus   func(ctx context.Context, i1 int64, ap1 *model.Algorithms)
	afterUpdateAlgorithmStatusCounter  uint64
	beforeUpdateAlgorithmStatusCounter uint64
	UpdateAlgorithmStatusMock          mApiServiceMockUpdateAlgorithmStatus

	funcUpdateClient          func(ctx context.Context, i1 int64, cp1 *model.Client) (err error)
	inspectFuncUpdateClient   func(ctx context.Context, i1 int64, cp1 *model.Client)
	afterUpdateClientCounter  uint64
	beforeUpdateClientCounter uint64
	UpdateClientMock          mApiServiceMockUpdateClient
}

// NewApiServiceMock returns a mock for service.ApiService
func NewApiServiceMock(t minimock.Tester) *ApiServiceMock {
	m := &ApiServiceMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddClientMock = mApiServiceMockAddClient{mock: m}
	m.AddClientMock.callArgs = []*ApiServiceMockAddClientParams{}

	m.DeleteClientMock = mApiServiceMockDeleteClient{mock: m}
	m.DeleteClientMock.callArgs = []*ApiServiceMockDeleteClientParams{}

	m.UpdateAlgorithmStatusMock = mApiServiceMockUpdateAlgorithmStatus{mock: m}
	m.UpdateAlgorithmStatusMock.callArgs = []*ApiServiceMockUpdateAlgorithmStatusParams{}

	m.UpdateClientMock = mApiServiceMockUpdateClient{mock: m}
	m.UpdateClientMock.callArgs = []*ApiServiceMockUpdateClientParams{}

	return m
}

type mApiServiceMockAddClient struct {
	mock               *ApiServiceMock
	defaultExpectation *ApiServiceMockAddClientExpectation
	expectations       []*ApiServiceMockAddClientExpectation

	callArgs []*ApiServiceMockAddClientParams
	mutex    sync.RWMutex
}

// ApiServiceMockAddClientExpectation specifies expectation struct of the ApiService.AddClient
type ApiServiceMockAddClientExpectation struct {
	mock    *ApiServiceMock
	params  *ApiServiceMockAddClientParams
	results *ApiServiceMockAddClientResults
	Counter uint64
}

// ApiServiceMockAddClientParams contains parameters of the ApiService.AddClient
type ApiServiceMockAddClientParams struct {
	ctx context.Context
	cp1 *model.Client
}

// ApiServiceMockAddClientResults contains results of the ApiService.AddClient
type ApiServiceMockAddClientResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for ApiService.AddClient
func (mmAddClient *mApiServiceMockAddClient) Expect(ctx context.Context, cp1 *model.Client) *mApiServiceMockAddClient {
	if mmAddClient.mock.funcAddClient != nil {
		mmAddClient.mock.t.Fatalf("ApiServiceMock.AddClient mock is already set by Set")
	}

	if mmAddClient.defaultExpectation == nil {
		mmAddClient.defaultExpectation = &ApiServiceMockAddClientExpectation{}
	}

	mmAddClient.defaultExpectation.params = &ApiServiceMockAddClientParams{ctx, cp1}
	for _, e := range mmAddClient.expectations {
		if minimock.Equal(e.params, mmAddClient.defaultExpectation.params) {
			mmAddClient.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddClient.defaultExpectation.params)
		}
	}

	return mmAddClient
}

// Inspect accepts an inspector function that has same arguments as the ApiService.AddClient
func (mmAddClient *mApiServiceMockAddClient) Inspect(f func(ctx context.Context, cp1 *model.Client)) *mApiServiceMockAddClient {
	if mmAddClient.mock.inspectFuncAddClient != nil {
		mmAddClient.mock.t.Fatalf("Inspect function is already set for ApiServiceMock.AddClient")
	}

	mmAddClient.mock.inspectFuncAddClient = f

	return mmAddClient
}

// Return sets up results that will be returned by ApiService.AddClient
func (mmAddClient *mApiServiceMockAddClient) Return(i1 int64, err error) *ApiServiceMock {
	if mmAddClient.mock.funcAddClient != nil {
		mmAddClient.mock.t.Fatalf("ApiServiceMock.AddClient mock is already set by Set")
	}

	if mmAddClient.defaultExpectation == nil {
		mmAddClient.defaultExpectation = &ApiServiceMockAddClientExpectation{mock: mmAddClient.mock}
	}
	mmAddClient.defaultExpectation.results = &ApiServiceMockAddClientResults{i1, err}
	return mmAddClient.mock
}

// Set uses given function f to mock the ApiService.AddClient method
func (mmAddClient *mApiServiceMockAddClient) Set(f func(ctx context.Context, cp1 *model.Client) (i1 int64, err error)) *ApiServiceMock {
	if mmAddClient.defaultExpectation != nil {
		mmAddClient.mock.t.Fatalf("Default expectation is already set for the ApiService.AddClient method")
	}

	if len(mmAddClient.expectations) > 0 {
		mmAddClient.mock.t.Fatalf("Some expectations are already set for the ApiService.AddClient method")
	}

	mmAddClient.mock.funcAddClient = f
	return mmAddClient.mock
}

// When sets expectation for the ApiService.AddClient which will trigger the result defined by the following
// Then helper
func (mmAddClient *mApiServiceMockAddClient) When(ctx context.Context, cp1 *model.Client) *ApiServiceMockAddClientExpectation {
	if mmAddClient.mock.funcAddClient != nil {
		mmAddClient.mock.t.Fatalf("ApiServiceMock.AddClient mock is already set by Set")
	}

	expectation := &ApiServiceMockAddClientExpectation{
		mock:   mmAddClient.mock,
		params: &ApiServiceMockAddClientParams{ctx, cp1},
	}
	mmAddClient.expectations = append(mmAddClient.expectations, expectation)
	return expectation
}

// Then sets up ApiService.AddClient return parameters for the expectation previously defined by the When method
func (e *ApiServiceMockAddClientExpectation) Then(i1 int64, err error) *ApiServiceMock {
	e.results = &ApiServiceMockAddClientResults{i1, err}
	return e.mock
}

// AddClient implements service.ApiService
func (mmAddClient *ApiServiceMock) AddClient(ctx context.Context, cp1 *model.Client) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmAddClient.beforeAddClientCounter, 1)
	defer mm_atomic.AddUint64(&mmAddClient.afterAddClientCounter, 1)

	if mmAddClient.inspectFuncAddClient != nil {
		mmAddClient.inspectFuncAddClient(ctx, cp1)
	}

	mm_params := &ApiServiceMockAddClientParams{ctx, cp1}

	// Record call args
	mmAddClient.AddClientMock.mutex.Lock()
	mmAddClient.AddClientMock.callArgs = append(mmAddClient.AddClientMock.callArgs, mm_params)
	mmAddClient.AddClientMock.mutex.Unlock()

	for _, e := range mmAddClient.AddClientMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmAddClient.AddClientMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddClient.AddClientMock.defaultExpectation.Counter, 1)
		mm_want := mmAddClient.AddClientMock.defaultExpectation.params
		mm_got := ApiServiceMockAddClientParams{ctx, cp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAddClient.t.Errorf("ApiServiceMock.AddClient got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAddClient.AddClientMock.defaultExpectation.results
		if mm_results == nil {
			mmAddClient.t.Fatal("No results are set for the ApiServiceMock.AddClient")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmAddClient.funcAddClient != nil {
		return mmAddClient.funcAddClient(ctx, cp1)
	}
	mmAddClient.t.Fatalf("Unexpected call to ApiServiceMock.AddClient. %v %v", ctx, cp1)
	return
}

// AddClientAfterCounter returns a count of finished ApiServiceMock.AddClient invocations
func (mmAddClient *ApiServiceMock) AddClientAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddClient.afterAddClientCounter)
}

// AddClientBeforeCounter returns a count of ApiServiceMock.AddClient invocations
func (mmAddClient *ApiServiceMock) AddClientBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddClient.beforeAddClientCounter)
}

// Calls returns a list of arguments used in each call to ApiServiceMock.AddClient.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddClient *mApiServiceMockAddClient) Calls() []*ApiServiceMockAddClientParams {
	mmAddClient.mutex.RLock()

	argCopy := make([]*ApiServiceMockAddClientParams, len(mmAddClient.callArgs))
	copy(argCopy, mmAddClient.callArgs)

	mmAddClient.mutex.RUnlock()

	return argCopy
}

// MinimockAddClientDone returns true if the count of the AddClient invocations corresponds
// the number of defined expectations
func (m *ApiServiceMock) MinimockAddClientDone() bool {
	for _, e := range m.AddClientMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddClientMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddClientCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddClient != nil && mm_atomic.LoadUint64(&m.afterAddClientCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddClientInspect logs each unmet expectation
func (m *ApiServiceMock) MinimockAddClientInspect() {
	for _, e := range m.AddClientMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ApiServiceMock.AddClient with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddClientMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddClientCounter) < 1 {
		if m.AddClientMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ApiServiceMock.AddClient")
		} else {
			m.t.Errorf("Expected call to ApiServiceMock.AddClient with params: %#v", *m.AddClientMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddClient != nil && mm_atomic.LoadUint64(&m.afterAddClientCounter) < 1 {
		m.t.Error("Expected call to ApiServiceMock.AddClient")
	}
}

type mApiServiceMockDeleteClient struct {
	mock               *ApiServiceMock
	defaultExpectation *ApiServiceMockDeleteClientExpectation
	expectations       []*ApiServiceMockDeleteClientExpectation

	callArgs []*ApiServiceMockDeleteClientParams
	mutex    sync.RWMutex
}

// ApiServiceMockDeleteClientExpectation specifies expectation struct of the ApiService.DeleteClient
type ApiServiceMockDeleteClientExpectation struct {
	mock    *ApiServiceMock
	params  *ApiServiceMockDeleteClientParams
	results *ApiServiceMockDeleteClientResults
	Counter uint64
}

// ApiServiceMockDeleteClientParams contains parameters of the ApiService.DeleteClient
type ApiServiceMockDeleteClientParams struct {
	ctx context.Context
	i1  int64
}

// ApiServiceMockDeleteClientResults contains results of the ApiService.DeleteClient
type ApiServiceMockDeleteClientResults struct {
	err error
}

// Expect sets up expected params for ApiService.DeleteClient
func (mmDeleteClient *mApiServiceMockDeleteClient) Expect(ctx context.Context, i1 int64) *mApiServiceMockDeleteClient {
	if mmDeleteClient.mock.funcDeleteClient != nil {
		mmDeleteClient.mock.t.Fatalf("ApiServiceMock.DeleteClient mock is already set by Set")
	}

	if mmDeleteClient.defaultExpectation == nil {
		mmDeleteClient.defaultExpectation = &ApiServiceMockDeleteClientExpectation{}
	}

	mmDeleteClient.defaultExpectation.params = &ApiServiceMockDeleteClientParams{ctx, i1}
	for _, e := range mmDeleteClient.expectations {
		if minimock.Equal(e.params, mmDeleteClient.defaultExpectation.params) {
			mmDeleteClient.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeleteClient.defaultExpectation.params)
		}
	}

	return mmDeleteClient
}

// Inspect accepts an inspector function that has same arguments as the ApiService.DeleteClient
func (mmDeleteClient *mApiServiceMockDeleteClient) Inspect(f func(ctx context.Context, i1 int64)) *mApiServiceMockDeleteClient {
	if mmDeleteClient.mock.inspectFuncDeleteClient != nil {
		mmDeleteClient.mock.t.Fatalf("Inspect function is already set for ApiServiceMock.DeleteClient")
	}

	mmDeleteClient.mock.inspectFuncDeleteClient = f

	return mmDeleteClient
}

// Return sets up results that will be returned by ApiService.DeleteClient
func (mmDeleteClient *mApiServiceMockDeleteClient) Return(err error) *ApiServiceMock {
	if mmDeleteClient.mock.funcDeleteClient != nil {
		mmDeleteClient.mock.t.Fatalf("ApiServiceMock.DeleteClient mock is already set by Set")
	}

	if mmDeleteClient.defaultExpectation == nil {
		mmDeleteClient.defaultExpectation = &ApiServiceMockDeleteClientExpectation{mock: mmDeleteClient.mock}
	}
	mmDeleteClient.defaultExpectation.results = &ApiServiceMockDeleteClientResults{err}
	return mmDeleteClient.mock
}

// Set uses given function f to mock the ApiService.DeleteClient method
func (mmDeleteClient *mApiServiceMockDeleteClient) Set(f func(ctx context.Context, i1 int64) (err error)) *ApiServiceMock {
	if mmDeleteClient.defaultExpectation != nil {
		mmDeleteClient.mock.t.Fatalf("Default expectation is already set for the ApiService.DeleteClient method")
	}

	if len(mmDeleteClient.expectations) > 0 {
		mmDeleteClient.mock.t.Fatalf("Some expectations are already set for the ApiService.DeleteClient method")
	}

	mmDeleteClient.mock.funcDeleteClient = f
	return mmDeleteClient.mock
}

// When sets expectation for the ApiService.DeleteClient which will trigger the result defined by the following
// Then helper
func (mmDeleteClient *mApiServiceMockDeleteClient) When(ctx context.Context, i1 int64) *ApiServiceMockDeleteClientExpectation {
	if mmDeleteClient.mock.funcDeleteClient != nil {
		mmDeleteClient.mock.t.Fatalf("ApiServiceMock.DeleteClient mock is already set by Set")
	}

	expectation := &ApiServiceMockDeleteClientExpectation{
		mock:   mmDeleteClient.mock,
		params: &ApiServiceMockDeleteClientParams{ctx, i1},
	}
	mmDeleteClient.expectations = append(mmDeleteClient.expectations, expectation)
	return expectation
}

// Then sets up ApiService.DeleteClient return parameters for the expectation previously defined by the When method
func (e *ApiServiceMockDeleteClientExpectation) Then(err error) *ApiServiceMock {
	e.results = &ApiServiceMockDeleteClientResults{err}
	return e.mock
}

// DeleteClient implements service.ApiService
func (mmDeleteClient *ApiServiceMock) DeleteClient(ctx context.Context, i1 int64) (err error) {
	mm_atomic.AddUint64(&mmDeleteClient.beforeDeleteClientCounter, 1)
	defer mm_atomic.AddUint64(&mmDeleteClient.afterDeleteClientCounter, 1)

	if mmDeleteClient.inspectFuncDeleteClient != nil {
		mmDeleteClient.inspectFuncDeleteClient(ctx, i1)
	}

	mm_params := &ApiServiceMockDeleteClientParams{ctx, i1}

	// Record call args
	mmDeleteClient.DeleteClientMock.mutex.Lock()
	mmDeleteClient.DeleteClientMock.callArgs = append(mmDeleteClient.DeleteClientMock.callArgs, mm_params)
	mmDeleteClient.DeleteClientMock.mutex.Unlock()

	for _, e := range mmDeleteClient.DeleteClientMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDeleteClient.DeleteClientMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDeleteClient.DeleteClientMock.defaultExpectation.Counter, 1)
		mm_want := mmDeleteClient.DeleteClientMock.defaultExpectation.params
		mm_got := ApiServiceMockDeleteClientParams{ctx, i1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeleteClient.t.Errorf("ApiServiceMock.DeleteClient got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeleteClient.DeleteClientMock.defaultExpectation.results
		if mm_results == nil {
			mmDeleteClient.t.Fatal("No results are set for the ApiServiceMock.DeleteClient")
		}
		return (*mm_results).err
	}
	if mmDeleteClient.funcDeleteClient != nil {
		return mmDeleteClient.funcDeleteClient(ctx, i1)
	}
	mmDeleteClient.t.Fatalf("Unexpected call to ApiServiceMock.DeleteClient. %v %v", ctx, i1)
	return
}

// DeleteClientAfterCounter returns a count of finished ApiServiceMock.DeleteClient invocations
func (mmDeleteClient *ApiServiceMock) DeleteClientAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteClient.afterDeleteClientCounter)
}

// DeleteClientBeforeCounter returns a count of ApiServiceMock.DeleteClient invocations
func (mmDeleteClient *ApiServiceMock) DeleteClientBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteClient.beforeDeleteClientCounter)
}

// Calls returns a list of arguments used in each call to ApiServiceMock.DeleteClient.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeleteClient *mApiServiceMockDeleteClient) Calls() []*ApiServiceMockDeleteClientParams {
	mmDeleteClient.mutex.RLock()

	argCopy := make([]*ApiServiceMockDeleteClientParams, len(mmDeleteClient.callArgs))
	copy(argCopy, mmDeleteClient.callArgs)

	mmDeleteClient.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteClientDone returns true if the count of the DeleteClient invocations corresponds
// the number of defined expectations
func (m *ApiServiceMock) MinimockDeleteClientDone() bool {
	for _, e := range m.DeleteClientMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteClientMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteClientCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteClient != nil && mm_atomic.LoadUint64(&m.afterDeleteClientCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteClientInspect logs each unmet expectation
func (m *ApiServiceMock) MinimockDeleteClientInspect() {
	for _, e := range m.DeleteClientMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ApiServiceMock.DeleteClient with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteClientMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteClientCounter) < 1 {
		if m.DeleteClientMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ApiServiceMock.DeleteClient")
		} else {
			m.t.Errorf("Expected call to ApiServiceMock.DeleteClient with params: %#v", *m.DeleteClientMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteClient != nil && mm_atomic.LoadUint64(&m.afterDeleteClientCounter) < 1 {
		m.t.Error("Expected call to ApiServiceMock.DeleteClient")
	}
}

type mApiServiceMockUpdateAlgorithmStatus struct {
	mock               *ApiServiceMock
	defaultExpectation *ApiServiceMockUpdateAlgorithmStatusExpectation
	expectations       []*ApiServiceMockUpdateAlgorithmStatusExpectation

	callArgs []*ApiServiceMockUpdateAlgorithmStatusParams
	mutex    sync.RWMutex
}

// ApiServiceMockUpdateAlgorithmStatusExpectation specifies expectation struct of the ApiService.UpdateAlgorithmStatus
type ApiServiceMockUpdateAlgorithmStatusExpectation struct {
	mock    *ApiServiceMock
	params  *ApiServiceMockUpdateAlgorithmStatusParams
	results *ApiServiceMockUpdateAlgorithmStatusResults
	Counter uint64
}

// ApiServiceMockUpdateAlgorithmStatusParams contains parameters of the ApiService.UpdateAlgorithmStatus
type ApiServiceMockUpdateAlgorithmStatusParams struct {
	ctx context.Context
	i1  int64
	ap1 *model.Algorithms
}

// ApiServiceMockUpdateAlgorithmStatusResults contains results of the ApiService.UpdateAlgorithmStatus
type ApiServiceMockUpdateAlgorithmStatusResults struct {
	err error
}

// Expect sets up expected params for ApiService.UpdateAlgorithmStatus
func (mmUpdateAlgorithmStatus *mApiServiceMockUpdateAlgorithmStatus) Expect(ctx context.Context, i1 int64, ap1 *model.Algorithms) *mApiServiceMockUpdateAlgorithmStatus {
	if mmUpdateAlgorithmStatus.mock.funcUpdateAlgorithmStatus != nil {
		mmUpdateAlgorithmStatus.mock.t.Fatalf("ApiServiceMock.UpdateAlgorithmStatus mock is already set by Set")
	}

	if mmUpdateAlgorithmStatus.defaultExpectation == nil {
		mmUpdateAlgorithmStatus.defaultExpectation = &ApiServiceMockUpdateAlgorithmStatusExpectation{}
	}

	mmUpdateAlgorithmStatus.defaultExpectation.params = &ApiServiceMockUpdateAlgorithmStatusParams{ctx, i1, ap1}
	for _, e := range mmUpdateAlgorithmStatus.expectations {
		if minimock.Equal(e.params, mmUpdateAlgorithmStatus.defaultExpectation.params) {
			mmUpdateAlgorithmStatus.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdateAlgorithmStatus.defaultExpectation.params)
		}
	}

	return mmUpdateAlgorithmStatus
}

// Inspect accepts an inspector function that has same arguments as the ApiService.UpdateAlgorithmStatus
func (mmUpdateAlgorithmStatus *mApiServiceMockUpdateAlgorithmStatus) Inspect(f func(ctx context.Context, i1 int64, ap1 *model.Algorithms)) *mApiServiceMockUpdateAlgorithmStatus {
	if mmUpdateAlgorithmStatus.mock.inspectFuncUpdateAlgorithmStatus != nil {
		mmUpdateAlgorithmStatus.mock.t.Fatalf("Inspect function is already set for ApiServiceMock.UpdateAlgorithmStatus")
	}

	mmUpdateAlgorithmStatus.mock.inspectFuncUpdateAlgorithmStatus = f

	return mmUpdateAlgorithmStatus
}

// Return sets up results that will be returned by ApiService.UpdateAlgorithmStatus
func (mmUpdateAlgorithmStatus *mApiServiceMockUpdateAlgorithmStatus) Return(err error) *ApiServiceMock {
	if mmUpdateAlgorithmStatus.mock.funcUpdateAlgorithmStatus != nil {
		mmUpdateAlgorithmStatus.mock.t.Fatalf("ApiServiceMock.UpdateAlgorithmStatus mock is already set by Set")
	}

	if mmUpdateAlgorithmStatus.defaultExpectation == nil {
		mmUpdateAlgorithmStatus.defaultExpectation = &ApiServiceMockUpdateAlgorithmStatusExpectation{mock: mmUpdateAlgorithmStatus.mock}
	}
	mmUpdateAlgorithmStatus.defaultExpectation.results = &ApiServiceMockUpdateAlgorithmStatusResults{err}
	return mmUpdateAlgorithmStatus.mock
}

// Set uses given function f to mock the ApiService.UpdateAlgorithmStatus method
func (mmUpdateAlgorithmStatus *mApiServiceMockUpdateAlgorithmStatus) Set(f func(ctx context.Context, i1 int64, ap1 *model.Algorithms) (err error)) *ApiServiceMock {
	if mmUpdateAlgorithmStatus.defaultExpectation != nil {
		mmUpdateAlgorithmStatus.mock.t.Fatalf("Default expectation is already set for the ApiService.UpdateAlgorithmStatus method")
	}

	if len(mmUpdateAlgorithmStatus.expectations) > 0 {
		mmUpdateAlgorithmStatus.mock.t.Fatalf("Some expectations are already set for the ApiService.UpdateAlgorithmStatus method")
	}

	mmUpdateAlgorithmStatus.mock.funcUpdateAlgorithmStatus = f
	return mmUpdateAlgorithmStatus.mock
}

// When sets expectation for the ApiService.UpdateAlgorithmStatus which will trigger the result defined by the following
// Then helper
func (mmUpdateAlgorithmStatus *mApiServiceMockUpdateAlgorithmStatus) When(ctx context.Context, i1 int64, ap1 *model.Algorithms) *ApiServiceMockUpdateAlgorithmStatusExpectation {
	if mmUpdateAlgorithmStatus.mock.funcUpdateAlgorithmStatus != nil {
		mmUpdateAlgorithmStatus.mock.t.Fatalf("ApiServiceMock.UpdateAlgorithmStatus mock is already set by Set")
	}

	expectation := &ApiServiceMockUpdateAlgorithmStatusExpectation{
		mock:   mmUpdateAlgorithmStatus.mock,
		params: &ApiServiceMockUpdateAlgorithmStatusParams{ctx, i1, ap1},
	}
	mmUpdateAlgorithmStatus.expectations = append(mmUpdateAlgorithmStatus.expectations, expectation)
	return expectation
}

// Then sets up ApiService.UpdateAlgorithmStatus return parameters for the expectation previously defined by the When method
func (e *ApiServiceMockUpdateAlgorithmStatusExpectation) Then(err error) *ApiServiceMock {
	e.results = &ApiServiceMockUpdateAlgorithmStatusResults{err}
	return e.mock
}

// UpdateAlgorithmStatus implements service.ApiService
func (mmUpdateAlgorithmStatus *ApiServiceMock) UpdateAlgorithmStatus(ctx context.Context, i1 int64, ap1 *model.Algorithms) (err error) {
	mm_atomic.AddUint64(&mmUpdateAlgorithmStatus.beforeUpdateAlgorithmStatusCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdateAlgorithmStatus.afterUpdateAlgorithmStatusCounter, 1)

	if mmUpdateAlgorithmStatus.inspectFuncUpdateAlgorithmStatus != nil {
		mmUpdateAlgorithmStatus.inspectFuncUpdateAlgorithmStatus(ctx, i1, ap1)
	}

	mm_params := &ApiServiceMockUpdateAlgorithmStatusParams{ctx, i1, ap1}

	// Record call args
	mmUpdateAlgorithmStatus.UpdateAlgorithmStatusMock.mutex.Lock()
	mmUpdateAlgorithmStatus.UpdateAlgorithmStatusMock.callArgs = append(mmUpdateAlgorithmStatus.UpdateAlgorithmStatusMock.callArgs, mm_params)
	mmUpdateAlgorithmStatus.UpdateAlgorithmStatusMock.mutex.Unlock()

	for _, e := range mmUpdateAlgorithmStatus.UpdateAlgorithmStatusMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpdateAlgorithmStatus.UpdateAlgorithmStatusMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdateAlgorithmStatus.UpdateAlgorithmStatusMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdateAlgorithmStatus.UpdateAlgorithmStatusMock.defaultExpectation.params
		mm_got := ApiServiceMockUpdateAlgorithmStatusParams{ctx, i1, ap1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdateAlgorithmStatus.t.Errorf("ApiServiceMock.UpdateAlgorithmStatus got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdateAlgorithmStatus.UpdateAlgorithmStatusMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdateAlgorithmStatus.t.Fatal("No results are set for the ApiServiceMock.UpdateAlgorithmStatus")
		}
		return (*mm_results).err
	}
	if mmUpdateAlgorithmStatus.funcUpdateAlgorithmStatus != nil {
		return mmUpdateAlgorithmStatus.funcUpdateAlgorithmStatus(ctx, i1, ap1)
	}
	mmUpdateAlgorithmStatus.t.Fatalf("Unexpected call to ApiServiceMock.UpdateAlgorithmStatus. %v %v %v", ctx, i1, ap1)
	return
}

// UpdateAlgorithmStatusAfterCounter returns a count of finished ApiServiceMock.UpdateAlgorithmStatus invocations
func (mmUpdateAlgorithmStatus *ApiServiceMock) UpdateAlgorithmStatusAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateAlgorithmStatus.afterUpdateAlgorithmStatusCounter)
}

// UpdateAlgorithmStatusBeforeCounter returns a count of ApiServiceMock.UpdateAlgorithmStatus invocations
func (mmUpdateAlgorithmStatus *ApiServiceMock) UpdateAlgorithmStatusBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateAlgorithmStatus.beforeUpdateAlgorithmStatusCounter)
}

// Calls returns a list of arguments used in each call to ApiServiceMock.UpdateAlgorithmStatus.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdateAlgorithmStatus *mApiServiceMockUpdateAlgorithmStatus) Calls() []*ApiServiceMockUpdateAlgorithmStatusParams {
	mmUpdateAlgorithmStatus.mutex.RLock()

	argCopy := make([]*ApiServiceMockUpdateAlgorithmStatusParams, len(mmUpdateAlgorithmStatus.callArgs))
	copy(argCopy, mmUpdateAlgorithmStatus.callArgs)

	mmUpdateAlgorithmStatus.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateAlgorithmStatusDone returns true if the count of the UpdateAlgorithmStatus invocations corresponds
// the number of defined expectations
func (m *ApiServiceMock) MinimockUpdateAlgorithmStatusDone() bool {
	for _, e := range m.UpdateAlgorithmStatusMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateAlgorithmStatusMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateAlgorithmStatusCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateAlgorithmStatus != nil && mm_atomic.LoadUint64(&m.afterUpdateAlgorithmStatusCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateAlgorithmStatusInspect logs each unmet expectation
func (m *ApiServiceMock) MinimockUpdateAlgorithmStatusInspect() {
	for _, e := range m.UpdateAlgorithmStatusMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ApiServiceMock.UpdateAlgorithmStatus with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateAlgorithmStatusMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateAlgorithmStatusCounter) < 1 {
		if m.UpdateAlgorithmStatusMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ApiServiceMock.UpdateAlgorithmStatus")
		} else {
			m.t.Errorf("Expected call to ApiServiceMock.UpdateAlgorithmStatus with params: %#v", *m.UpdateAlgorithmStatusMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateAlgorithmStatus != nil && mm_atomic.LoadUint64(&m.afterUpdateAlgorithmStatusCounter) < 1 {
		m.t.Error("Expected call to ApiServiceMock.UpdateAlgorithmStatus")
	}
}

type mApiServiceMockUpdateClient struct {
	mock               *ApiServiceMock
	defaultExpectation *ApiServiceMockUpdateClientExpectation
	expectations       []*ApiServiceMockUpdateClientExpectation

	callArgs []*ApiServiceMockUpdateClientParams
	mutex    sync.RWMutex
}

// ApiServiceMockUpdateClientExpectation specifies expectation struct of the ApiService.UpdateClient
type ApiServiceMockUpdateClientExpectation struct {
	mock    *ApiServiceMock
	params  *ApiServiceMockUpdateClientParams
	results *ApiServiceMockUpdateClientResults
	Counter uint64
}

// ApiServiceMockUpdateClientParams contains parameters of the ApiService.UpdateClient
type ApiServiceMockUpdateClientParams struct {
	ctx context.Context
	i1  int64
	cp1 *model.Client
}

// ApiServiceMockUpdateClientResults contains results of the ApiService.UpdateClient
type ApiServiceMockUpdateClientResults struct {
	err error
}

// Expect sets up expected params for ApiService.UpdateClient
func (mmUpdateClient *mApiServiceMockUpdateClient) Expect(ctx context.Context, i1 int64, cp1 *model.Client) *mApiServiceMockUpdateClient {
	if mmUpdateClient.mock.funcUpdateClient != nil {
		mmUpdateClient.mock.t.Fatalf("ApiServiceMock.UpdateClient mock is already set by Set")
	}

	if mmUpdateClient.defaultExpectation == nil {
		mmUpdateClient.defaultExpectation = &ApiServiceMockUpdateClientExpectation{}
	}

	mmUpdateClient.defaultExpectation.params = &ApiServiceMockUpdateClientParams{ctx, i1, cp1}
	for _, e := range mmUpdateClient.expectations {
		if minimock.Equal(e.params, mmUpdateClient.defaultExpectation.params) {
			mmUpdateClient.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdateClient.defaultExpectation.params)
		}
	}

	return mmUpdateClient
}

// Inspect accepts an inspector function that has same arguments as the ApiService.UpdateClient
func (mmUpdateClient *mApiServiceMockUpdateClient) Inspect(f func(ctx context.Context, i1 int64, cp1 *model.Client)) *mApiServiceMockUpdateClient {
	if mmUpdateClient.mock.inspectFuncUpdateClient != nil {
		mmUpdateClient.mock.t.Fatalf("Inspect function is already set for ApiServiceMock.UpdateClient")
	}

	mmUpdateClient.mock.inspectFuncUpdateClient = f

	return mmUpdateClient
}

// Return sets up results that will be returned by ApiService.UpdateClient
func (mmUpdateClient *mApiServiceMockUpdateClient) Return(err error) *ApiServiceMock {
	if mmUpdateClient.mock.funcUpdateClient != nil {
		mmUpdateClient.mock.t.Fatalf("ApiServiceMock.UpdateClient mock is already set by Set")
	}

	if mmUpdateClient.defaultExpectation == nil {
		mmUpdateClient.defaultExpectation = &ApiServiceMockUpdateClientExpectation{mock: mmUpdateClient.mock}
	}
	mmUpdateClient.defaultExpectation.results = &ApiServiceMockUpdateClientResults{err}
	return mmUpdateClient.mock
}

// Set uses given function f to mock the ApiService.UpdateClient method
func (mmUpdateClient *mApiServiceMockUpdateClient) Set(f func(ctx context.Context, i1 int64, cp1 *model.Client) (err error)) *ApiServiceMock {
	if mmUpdateClient.defaultExpectation != nil {
		mmUpdateClient.mock.t.Fatalf("Default expectation is already set for the ApiService.UpdateClient method")
	}

	if len(mmUpdateClient.expectations) > 0 {
		mmUpdateClient.mock.t.Fatalf("Some expectations are already set for the ApiService.UpdateClient method")
	}

	mmUpdateClient.mock.funcUpdateClient = f
	return mmUpdateClient.mock
}

// When sets expectation for the ApiService.UpdateClient which will trigger the result defined by the following
// Then helper
func (mmUpdateClient *mApiServiceMockUpdateClient) When(ctx context.Context, i1 int64, cp1 *model.Client) *ApiServiceMockUpdateClientExpectation {
	if mmUpdateClient.mock.funcUpdateClient != nil {
		mmUpdateClient.mock.t.Fatalf("ApiServiceMock.UpdateClient mock is already set by Set")
	}

	expectation := &ApiServiceMockUpdateClientExpectation{
		mock:   mmUpdateClient.mock,
		params: &ApiServiceMockUpdateClientParams{ctx, i1, cp1},
	}
	mmUpdateClient.expectations = append(mmUpdateClient.expectations, expectation)
	return expectation
}

// Then sets up ApiService.UpdateClient return parameters for the expectation previously defined by the When method
func (e *ApiServiceMockUpdateClientExpectation) Then(err error) *ApiServiceMock {
	e.results = &ApiServiceMockUpdateClientResults{err}
	return e.mock
}

// UpdateClient implements service.ApiService
func (mmUpdateClient *ApiServiceMock) UpdateClient(ctx context.Context, i1 int64, cp1 *model.Client) (err error) {
	mm_atomic.AddUint64(&mmUpdateClient.beforeUpdateClientCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdateClient.afterUpdateClientCounter, 1)

	if mmUpdateClient.inspectFuncUpdateClient != nil {
		mmUpdateClient.inspectFuncUpdateClient(ctx, i1, cp1)
	}

	mm_params := &ApiServiceMockUpdateClientParams{ctx, i1, cp1}

	// Record call args
	mmUpdateClient.UpdateClientMock.mutex.Lock()
	mmUpdateClient.UpdateClientMock.callArgs = append(mmUpdateClient.UpdateClientMock.callArgs, mm_params)
	mmUpdateClient.UpdateClientMock.mutex.Unlock()

	for _, e := range mmUpdateClient.UpdateClientMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpdateClient.UpdateClientMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdateClient.UpdateClientMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdateClient.UpdateClientMock.defaultExpectation.params
		mm_got := ApiServiceMockUpdateClientParams{ctx, i1, cp1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdateClient.t.Errorf("ApiServiceMock.UpdateClient got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdateClient.UpdateClientMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdateClient.t.Fatal("No results are set for the ApiServiceMock.UpdateClient")
		}
		return (*mm_results).err
	}
	if mmUpdateClient.funcUpdateClient != nil {
		return mmUpdateClient.funcUpdateClient(ctx, i1, cp1)
	}
	mmUpdateClient.t.Fatalf("Unexpected call to ApiServiceMock.UpdateClient. %v %v %v", ctx, i1, cp1)
	return
}

// UpdateClientAfterCounter returns a count of finished ApiServiceMock.UpdateClient invocations
func (mmUpdateClient *ApiServiceMock) UpdateClientAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateClient.afterUpdateClientCounter)
}

// UpdateClientBeforeCounter returns a count of ApiServiceMock.UpdateClient invocations
func (mmUpdateClient *ApiServiceMock) UpdateClientBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdateClient.beforeUpdateClientCounter)
}

// Calls returns a list of arguments used in each call to ApiServiceMock.UpdateClient.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdateClient *mApiServiceMockUpdateClient) Calls() []*ApiServiceMockUpdateClientParams {
	mmUpdateClient.mutex.RLock()

	argCopy := make([]*ApiServiceMockUpdateClientParams, len(mmUpdateClient.callArgs))
	copy(argCopy, mmUpdateClient.callArgs)

	mmUpdateClient.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateClientDone returns true if the count of the UpdateClient invocations corresponds
// the number of defined expectations
func (m *ApiServiceMock) MinimockUpdateClientDone() bool {
	for _, e := range m.UpdateClientMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateClientMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateClientCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateClient != nil && mm_atomic.LoadUint64(&m.afterUpdateClientCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateClientInspect logs each unmet expectation
func (m *ApiServiceMock) MinimockUpdateClientInspect() {
	for _, e := range m.UpdateClientMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ApiServiceMock.UpdateClient with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateClientMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateClientCounter) < 1 {
		if m.UpdateClientMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ApiServiceMock.UpdateClient")
		} else {
			m.t.Errorf("Expected call to ApiServiceMock.UpdateClient with params: %#v", *m.UpdateClientMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdateClient != nil && mm_atomic.LoadUint64(&m.afterUpdateClientCounter) < 1 {
		m.t.Error("Expected call to ApiServiceMock.UpdateClient")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ApiServiceMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAddClientInspect()

		m.MinimockDeleteClientInspect()

		m.MinimockUpdateAlgorithmStatusInspect()

		m.MinimockUpdateClientInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ApiServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ApiServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddClientDone() &&
		m.MinimockDeleteClientDone() &&
		m.MinimockUpdateAlgorithmStatusDone() &&
		m.MinimockUpdateClientDone()
}
