// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kamilsk/passport/server (interfaces: Service)

// Package server_test is a generated GoMock package.
package server_test

import (
	gomock "github.com/golang/mock/gomock"
	tracker "github.com/kamilsk/passport/transfer/api/v1/tracker"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// HandleTrackerFingerprintV1 mocks base method
func (m *MockService) HandleTrackerFingerprintV1(arg0 tracker.FingerprintRequest) tracker.FingerprintResponse {
	ret := m.ctrl.Call(m, "HandleTrackerFingerprintV1", arg0)
	ret0, _ := ret[0].(tracker.FingerprintResponse)
	return ret0
}

// HandleTrackerFingerprintV1 indicates an expected call of HandleTrackerFingerprintV1
func (mr *MockServiceMockRecorder) HandleTrackerFingerprintV1(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTrackerFingerprintV1", reflect.TypeOf((*MockService)(nil).HandleTrackerFingerprintV1), arg0)
}

// HandleTrackerInstructionV1 mocks base method
func (m *MockService) HandleTrackerInstructionV1(arg0 tracker.InstructionRequest) tracker.InstructionResponse {
	ret := m.ctrl.Call(m, "HandleTrackerInstructionV1", arg0)
	ret0, _ := ret[0].(tracker.InstructionResponse)
	return ret0
}

// HandleTrackerInstructionV1 indicates an expected call of HandleTrackerInstructionV1
func (mr *MockServiceMockRecorder) HandleTrackerInstructionV1(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTrackerInstructionV1", reflect.TypeOf((*MockService)(nil).HandleTrackerInstructionV1), arg0)
}
