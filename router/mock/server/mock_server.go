// Code generated by MockGen. DO NOT EDIT.
// Source: ./router/server/server.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	pgproto3 "github.com/jackc/pgx/v5/pgproto3"
	config "github.com/pg-sharding/spqr/pkg/config"
	kr "github.com/pg-sharding/spqr/pkg/models/kr"
	shard "github.com/pg-sharding/spqr/pkg/shard"
	txstatus "github.com/pg-sharding/spqr/pkg/txstatus"
	reflect "reflect"
)

// MockServer is a mock of Server interface
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// HasPrepareStatement mocks base method
func (m *MockServer) HasPrepareStatement(hash uint64) (bool, shard.PreparedStatementDescriptor) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPrepareStatement", hash)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(shard.PreparedStatementDescriptor)
	return ret0, ret1
}

// HasPrepareStatement indicates an expected call of HasPrepareStatement
func (mr *MockServerMockRecorder) HasPrepareStatement(hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPrepareStatement", reflect.TypeOf((*MockServer)(nil).HasPrepareStatement), hash)
}

// PrepareStatement mocks base method
func (m *MockServer) PrepareStatement(hash uint64, rd shard.PreparedStatementDescriptor) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PrepareStatement", hash, rd)
}

// PrepareStatement indicates an expected call of PrepareStatement
func (mr *MockServerMockRecorder) PrepareStatement(hash, rd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareStatement", reflect.TypeOf((*MockServer)(nil).PrepareStatement), hash, rd)
}

// SetTxStatus mocks base method
func (m *MockServer) SetTxStatus(status txstatus.TXStatus) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTxStatus", status)
}

// SetTxStatus indicates an expected call of SetTxStatus
func (mr *MockServerMockRecorder) SetTxStatus(status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTxStatus", reflect.TypeOf((*MockServer)(nil).SetTxStatus), status)
}

// TxStatus mocks base method
func (m *MockServer) TxStatus() txstatus.TXStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxStatus")
	ret0, _ := ret[0].(txstatus.TXStatus)
	return ret0
}

// TxStatus indicates an expected call of TxStatus
func (mr *MockServerMockRecorder) TxStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxStatus", reflect.TypeOf((*MockServer)(nil).TxStatus))
}

// Name mocks base method
func (m *MockServer) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockServerMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockServer)(nil).Name))
}

// Send mocks base method
func (m *MockServer) Send(query pgproto3.FrontendMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", query)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockServerMockRecorder) Send(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockServer)(nil).Send), query)
}

// Receive mocks base method
func (m *MockServer) Receive() (pgproto3.BackendMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Receive")
	ret0, _ := ret[0].(pgproto3.BackendMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Receive indicates an expected call of Receive
func (mr *MockServerMockRecorder) Receive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockServer)(nil).Receive))
}

// AddDataShard mocks base method
func (m *MockServer) AddDataShard(clid string, shardKey kr.ShardKey, tsa string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDataShard", clid, shardKey, tsa)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddDataShard indicates an expected call of AddDataShard
func (mr *MockServerMockRecorder) AddDataShard(clid, shardKey, tsa interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDataShard", reflect.TypeOf((*MockServer)(nil).AddDataShard), clid, shardKey, tsa)
}

// UnRouteShard mocks base method
func (m *MockServer) UnRouteShard(sh kr.ShardKey, rule *config.FrontendRule) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnRouteShard", sh, rule)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnRouteShard indicates an expected call of UnRouteShard
func (mr *MockServerMockRecorder) UnRouteShard(sh, rule interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnRouteShard", reflect.TypeOf((*MockServer)(nil).UnRouteShard), sh, rule)
}

// Datashards mocks base method
func (m *MockServer) Datashards() []shard.Shard {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Datashards")
	ret0, _ := ret[0].([]shard.Shard)
	return ret0
}

// Datashards indicates an expected call of Datashards
func (mr *MockServerMockRecorder) Datashards() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Datashards", reflect.TypeOf((*MockServer)(nil).Datashards))
}

// Cancel mocks base method
func (m *MockServer) Cancel() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cancel")
	ret0, _ := ret[0].(error)
	return ret0
}

// Cancel indicates an expected call of Cancel
func (mr *MockServerMockRecorder) Cancel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cancel", reflect.TypeOf((*MockServer)(nil).Cancel))
}

// Reset mocks base method
func (m *MockServer) Reset() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reset")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset
func (mr *MockServerMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockServer)(nil).Reset))
}

// Sync mocks base method
func (m *MockServer) Sync() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Sync indicates an expected call of Sync
func (mr *MockServerMockRecorder) Sync() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockServer)(nil).Sync))
}
