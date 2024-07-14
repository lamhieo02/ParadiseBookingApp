// Code generated by MockGen. DO NOT EDIT.
// Source: interactor.go

// Package accountusecase is a generated GoMock package.
package accountusecase

import (
	context "context"
	common "paradise-booking/common"
	entities "paradise-booking/entities"
	storage "paradise-booking/modules/account/storage"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccountStorage is a mock of AccountStorage interface.
type MockAccountStorage struct {
	ctrl     *gomock.Controller
	recorder *MockAccountStorageMockRecorder
}

// MockAccountStorageMockRecorder is the mock recorder for MockAccountStorage.
type MockAccountStorageMockRecorder struct {
	mock *MockAccountStorage
}

// NewMockAccountStorage creates a new mock instance.
func NewMockAccountStorage(ctrl *gomock.Controller) *MockAccountStorage {
	mock := &MockAccountStorage{ctrl: ctrl}
	mock.recorder = &MockAccountStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountStorage) EXPECT() *MockAccountStorageMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccountStorage) Create(ctx context.Context, account *entities.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, account)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAccountStorageMockRecorder) Create(ctx, account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountStorage)(nil).Create), ctx, account)
}

// CreateTx mocks base method.
func (m *MockAccountStorage) CreateTx(ctx context.Context, createUserTxParam storage.CreateUserTxParam) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTx", ctx, createUserTxParam)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTx indicates an expected call of CreateTx.
func (mr *MockAccountStorageMockRecorder) CreateTx(ctx, createUserTxParam interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockAccountStorage)(nil).CreateTx), ctx, createUserTxParam)
}

// GetAccountByEmail mocks base method.
func (m *MockAccountStorage) GetAccountByEmail(ctx context.Context, email string) (*entities.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByEmail", ctx, email)
	ret0, _ := ret[0].(*entities.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByEmail indicates an expected call of GetAccountByEmail.
func (mr *MockAccountStorageMockRecorder) GetAccountByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByEmail", reflect.TypeOf((*MockAccountStorage)(nil).GetAccountByEmail), ctx, email)
}

// GetAllAccountUserAndVendor mocks base method.
func (m *MockAccountStorage) GetAllAccountUserAndVendor(ctx context.Context, paging *common.Paging) ([]entities.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAccountUserAndVendor", ctx, paging)
	ret0, _ := ret[0].([]entities.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAccountUserAndVendor indicates an expected call of GetAllAccountUserAndVendor.
func (mr *MockAccountStorageMockRecorder) GetAllAccountUserAndVendor(ctx, paging interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAccountUserAndVendor", reflect.TypeOf((*MockAccountStorage)(nil).GetAllAccountUserAndVendor), ctx, paging)
}

// GetProfileByID mocks base method.
func (m *MockAccountStorage) GetProfileByID(ctx context.Context, id int) (*entities.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileByID", ctx, id)
	ret0, _ := ret[0].(*entities.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileByID indicates an expected call of GetProfileByID.
func (mr *MockAccountStorageMockRecorder) GetProfileByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileByID", reflect.TypeOf((*MockAccountStorage)(nil).GetProfileByID), ctx, id)
}

// UpdateAccountById mocks base method.
func (m *MockAccountStorage) UpdateAccountById(ctx context.Context, id int, accountUpdate *entities.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccountById", ctx, id, accountUpdate)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccountById indicates an expected call of UpdateAccountById.
func (mr *MockAccountStorageMockRecorder) UpdateAccountById(ctx, id, accountUpdate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountById", reflect.TypeOf((*MockAccountStorage)(nil).UpdateAccountById), ctx, id, accountUpdate)
}

// MockVerifyEmailsUseCase is a mock of VerifyEmailsUseCase interface.
type MockVerifyEmailsUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockVerifyEmailsUseCaseMockRecorder
}

// MockVerifyEmailsUseCaseMockRecorder is the mock recorder for MockVerifyEmailsUseCase.
type MockVerifyEmailsUseCaseMockRecorder struct {
	mock *MockVerifyEmailsUseCase
}

// NewMockVerifyEmailsUseCase creates a new mock instance.
func NewMockVerifyEmailsUseCase(ctrl *gomock.Controller) *MockVerifyEmailsUseCase {
	mock := &MockVerifyEmailsUseCase{ctrl: ctrl}
	mock.recorder = &MockVerifyEmailsUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerifyEmailsUseCase) EXPECT() *MockVerifyEmailsUseCaseMockRecorder {
	return m.recorder
}

// UpsertResetSetCodePassword mocks base method.
func (m *MockVerifyEmailsUseCase) UpsertResetSetCodePassword(ctx context.Context, email string) (*entities.VerifyEmail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertResetSetCodePassword", ctx, email)
	ret0, _ := ret[0].(*entities.VerifyEmail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertResetSetCodePassword indicates an expected call of UpsertResetSetCodePassword.
func (mr *MockVerifyEmailsUseCaseMockRecorder) UpsertResetSetCodePassword(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertResetSetCodePassword", reflect.TypeOf((*MockVerifyEmailsUseCase)(nil).UpsertResetSetCodePassword), ctx, email)
}