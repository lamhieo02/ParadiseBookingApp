// Code generated by MockGen. DO NOT EDIT.
// Source: distributor.go

// Package worker is a generated GoMock package.
package worker

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	asynq "github.com/hibiken/asynq"
)

// MockTaskDistributor is a mock of TaskDistributor interface.
type MockTaskDistributor struct {
	ctrl     *gomock.Controller
	recorder *MockTaskDistributorMockRecorder
}

// MockTaskDistributorMockRecorder is the mock recorder for MockTaskDistributor.
type MockTaskDistributorMockRecorder struct {
	mock *MockTaskDistributor
}

// NewMockTaskDistributor creates a new mock instance.
func NewMockTaskDistributor(ctrl *gomock.Controller) *MockTaskDistributor {
	mock := &MockTaskDistributor{ctrl: ctrl}
	mock.recorder = &MockTaskDistributorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskDistributor) EXPECT() *MockTaskDistributorMockRecorder {
	return m.recorder
}

// DistributeTaskSendConfirmBooking mocks base method.
func (m *MockTaskDistributor) DistributeTaskSendConfirmBooking(ctx context.Context, payload *PayloadSendConfirmBooking, opts ...asynq.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, payload}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DistributeTaskSendConfirmBooking", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DistributeTaskSendConfirmBooking indicates an expected call of DistributeTaskSendConfirmBooking.
func (mr *MockTaskDistributorMockRecorder) DistributeTaskSendConfirmBooking(ctx, payload interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, payload}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistributeTaskSendConfirmBooking", reflect.TypeOf((*MockTaskDistributor)(nil).DistributeTaskSendConfirmBooking), varargs...)
}

// DistributeTaskSendConfirmBookingGuider mocks base method.
func (m *MockTaskDistributor) DistributeTaskSendConfirmBookingGuider(ctx context.Context, payload *PayloadSendConfirmBookingGuider, opts ...asynq.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, payload}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DistributeTaskSendConfirmBookingGuider", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DistributeTaskSendConfirmBookingGuider indicates an expected call of DistributeTaskSendConfirmBookingGuider.
func (mr *MockTaskDistributorMockRecorder) DistributeTaskSendConfirmBookingGuider(ctx, payload interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, payload}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistributeTaskSendConfirmBookingGuider", reflect.TypeOf((*MockTaskDistributor)(nil).DistributeTaskSendConfirmBookingGuider), varargs...)
}

// DistributeTaskSendVerifyEmail mocks base method.
func (m *MockTaskDistributor) DistributeTaskSendVerifyEmail(ctx context.Context, payload *PayloadSendVerifyEmail, opts ...asynq.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, payload}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DistributeTaskSendVerifyEmail", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DistributeTaskSendVerifyEmail indicates an expected call of DistributeTaskSendVerifyEmail.
func (mr *MockTaskDistributorMockRecorder) DistributeTaskSendVerifyEmail(ctx, payload interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, payload}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistributeTaskSendVerifyEmail", reflect.TypeOf((*MockTaskDistributor)(nil).DistributeTaskSendVerifyEmail), varargs...)
}

// DistributeTaskSendVerifyResetCodePassword mocks base method.
func (m *MockTaskDistributor) DistributeTaskSendVerifyResetCodePassword(ctx context.Context, payload *PayloadSendVerifyResetCodePassword, opts ...asynq.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, payload}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DistributeTaskSendVerifyResetCodePassword", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DistributeTaskSendVerifyResetCodePassword indicates an expected call of DistributeTaskSendVerifyResetCodePassword.
func (mr *MockTaskDistributorMockRecorder) DistributeTaskSendVerifyResetCodePassword(ctx, payload interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, payload}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistributeTaskSendVerifyResetCodePassword", reflect.TypeOf((*MockTaskDistributor)(nil).DistributeTaskSendVerifyResetCodePassword), varargs...)
}

// DistributeTaskUpdateStatusBooking mocks base method.
func (m *MockTaskDistributor) DistributeTaskUpdateStatusBooking(ctx context.Context, opts ...asynq.Option) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DistributeTaskUpdateStatusBooking", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DistributeTaskUpdateStatusBooking indicates an expected call of DistributeTaskUpdateStatusBooking.
func (mr *MockTaskDistributorMockRecorder) DistributeTaskUpdateStatusBooking(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistributeTaskUpdateStatusBooking", reflect.TypeOf((*MockTaskDistributor)(nil).DistributeTaskUpdateStatusBooking), varargs...)
}