// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory-am/fosite (interfaces: TokenEndpointHandler)

package internal

import (
	http "net/http"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory-am/fosite"
	context "golang.org/x/net/context"
)

// Mock of TokenEndpointHandler interface
type MockTokenEndpointHandler struct {
	ctrl     *gomock.Controller
	recorder *_MockTokenEndpointHandlerRecorder
}

// Recorder for MockTokenEndpointHandler (not exported)
type _MockTokenEndpointHandlerRecorder struct {
	mock *MockTokenEndpointHandler
}

func NewMockTokenEndpointHandler(ctrl *gomock.Controller) *MockTokenEndpointHandler {
	mock := &MockTokenEndpointHandler{ctrl: ctrl}
	mock.recorder = &_MockTokenEndpointHandlerRecorder{mock}
	return mock
}

func (_m *MockTokenEndpointHandler) EXPECT() *_MockTokenEndpointHandlerRecorder {
	return _m.recorder
}

func (_m *MockTokenEndpointHandler) HandleTokenEndpointRequest(_param0 context.Context, _param1 *http.Request, _param2 fosite.AccessRequester) error {
	ret := _m.ctrl.Call(_m, "HandleTokenEndpointRequest", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTokenEndpointHandlerRecorder) HandleTokenEndpointRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HandleTokenEndpointRequest", arg0, arg1, arg2)
}

func (_m *MockTokenEndpointHandler) PopulateTokenEndpointResponse(_param0 context.Context, _param1 *http.Request, _param2 fosite.AccessRequester, _param3 fosite.AccessResponder) error {
	ret := _m.ctrl.Call(_m, "PopulateTokenEndpointResponse", _param0, _param1, _param2, _param3)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTokenEndpointHandlerRecorder) PopulateTokenEndpointResponse(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PopulateTokenEndpointResponse", arg0, arg1, arg2, arg3)
}
