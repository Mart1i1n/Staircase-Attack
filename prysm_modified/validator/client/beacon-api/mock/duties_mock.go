// Code generated by MockGen. DO NOT EDIT.
// Source: validator/client/beacon-api/duties.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	shared "github.com/prysmaticlabs/prysm/v4/beacon-chain/rpc/eth/shared"
	validator "github.com/prysmaticlabs/prysm/v4/beacon-chain/rpc/eth/validator"
	primitives "github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
)

// MockdutiesProvider is a mock of dutiesProvider interface.
type MockdutiesProvider struct {
	ctrl     *gomock.Controller
	recorder *MockdutiesProviderMockRecorder
}

// MockdutiesProviderMockRecorder is the mock recorder for MockdutiesProvider.
type MockdutiesProviderMockRecorder struct {
	mock *MockdutiesProvider
}

// NewMockdutiesProvider creates a new mock instance.
func NewMockdutiesProvider(ctrl *gomock.Controller) *MockdutiesProvider {
	mock := &MockdutiesProvider{ctrl: ctrl}
	mock.recorder = &MockdutiesProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdutiesProvider) EXPECT() *MockdutiesProviderMockRecorder {
	return m.recorder
}

// GetAttesterDuties mocks base method.
func (m *MockdutiesProvider) GetAttesterDuties(ctx context.Context, epoch primitives.Epoch, validatorIndices []primitives.ValidatorIndex) ([]*validator.AttesterDuty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttesterDuties", ctx, epoch, validatorIndices)
	ret0, _ := ret[0].([]*validator.AttesterDuty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttesterDuties indicates an expected call of GetAttesterDuties.
func (mr *MockdutiesProviderMockRecorder) GetAttesterDuties(ctx, epoch, validatorIndices interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttesterDuties", reflect.TypeOf((*MockdutiesProvider)(nil).GetAttesterDuties), ctx, epoch, validatorIndices)
}

// GetCommittees mocks base method.
func (m *MockdutiesProvider) GetCommittees(ctx context.Context, epoch primitives.Epoch) ([]*shared.Committee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommittees", ctx, epoch)
	ret0, _ := ret[0].([]*shared.Committee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommittees indicates an expected call of GetCommittees.
func (mr *MockdutiesProviderMockRecorder) GetCommittees(ctx, epoch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommittees", reflect.TypeOf((*MockdutiesProvider)(nil).GetCommittees), ctx, epoch)
}

// GetProposerDuties mocks base method.
func (m *MockdutiesProvider) GetProposerDuties(ctx context.Context, epoch primitives.Epoch) ([]*validator.ProposerDuty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposerDuties", ctx, epoch)
	ret0, _ := ret[0].([]*validator.ProposerDuty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProposerDuties indicates an expected call of GetProposerDuties.
func (mr *MockdutiesProviderMockRecorder) GetProposerDuties(ctx, epoch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposerDuties", reflect.TypeOf((*MockdutiesProvider)(nil).GetProposerDuties), ctx, epoch)
}

// GetSyncDuties mocks base method.
func (m *MockdutiesProvider) GetSyncDuties(ctx context.Context, epoch primitives.Epoch, validatorIndices []primitives.ValidatorIndex) ([]*validator.SyncCommitteeDuty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSyncDuties", ctx, epoch, validatorIndices)
	ret0, _ := ret[0].([]*validator.SyncCommitteeDuty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSyncDuties indicates an expected call of GetSyncDuties.
func (mr *MockdutiesProviderMockRecorder) GetSyncDuties(ctx, epoch, validatorIndices interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSyncDuties", reflect.TypeOf((*MockdutiesProvider)(nil).GetSyncDuties), ctx, epoch, validatorIndices)
}