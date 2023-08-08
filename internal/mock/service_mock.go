// Code generated by MockGen. DO NOT EDIT.
// Source: internal/handler/handler.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	model "github.com/anilaydinn/tdd-example/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockStudentActions is a mock of StudentActions interface.
type MockStudentActions struct {
	ctrl     *gomock.Controller
	recorder *MockStudentActionsMockRecorder
}

// MockStudentActionsMockRecorder is the mock recorder for MockStudentActions.
type MockStudentActionsMockRecorder struct {
	mock *MockStudentActions
}

// NewMockStudentActions creates a new mock instance.
func NewMockStudentActions(ctrl *gomock.Controller) *MockStudentActions {
	mock := &MockStudentActions{ctrl: ctrl}
	mock.recorder = &MockStudentActionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudentActions) EXPECT() *MockStudentActionsMockRecorder {
	return m.recorder
}

// CreateStudent mocks base method.
func (m *MockStudentActions) CreateStudent(request model.CreateStudentRequest) (model.CreateStudentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStudent", request)
	ret0, _ := ret[0].(model.CreateStudentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStudent indicates an expected call of CreateStudent.
func (mr *MockStudentActionsMockRecorder) CreateStudent(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudent", reflect.TypeOf((*MockStudentActions)(nil).CreateStudent), request)
}