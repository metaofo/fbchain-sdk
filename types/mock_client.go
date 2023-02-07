// Code generated by MockGen. DO NOT EDIT.
// Source: types/client.go

// Package types is a generated GoMock package.
package types

import (
	tmtypes "github.com/FiboChain/fbc/libs/tendermint/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	codec "github.com/FiboChain/fbc/libs/cosmos-sdk/codec"
	types "github.com/FiboChain/fbc/libs/cosmos-sdk/types"
	types0 "github.com/FiboChain/fbc/libs/cosmos-sdk/x/auth/types"
	bytes "github.com/FiboChain/fbc/libs/tendermint/libs/bytes"
	types1 "github.com/FiboChain/fbc/libs/tendermint/rpc/core/types"
)

// MockBaseClient is a mock of BaseClient interface.
type MockBaseClient struct {
	ctrl     *gomock.Controller
	recorder *MockBaseClientMockRecorder
}

func (m *MockBaseClient) BlockInfo(height *int64) (*tmtypes.BlockMeta, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockBaseClient) LatestBlockNumber() (int64, error) {
	//TODO implement me
	panic("implement me")
}

// MockBaseClientMockRecorder is the mock recorder for MockBaseClient.
type MockBaseClientMockRecorder struct {
	mock *MockBaseClient
}

// NewMockBaseClient creates a new mock instance.
func NewMockBaseClient(ctrl *gomock.Controller) *MockBaseClient {
	mock := &MockBaseClient{ctrl: ctrl}
	mock.recorder = &MockBaseClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBaseClient) EXPECT() *MockBaseClientMockRecorder {
	return m.recorder
}

// Block mocks base method.
func (m *MockBaseClient) Block(height *int64) (*types1.ResultBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Block", height)
	ret0, _ := ret[0].(*types1.ResultBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Block indicates an expected call of Block.
func (mr *MockBaseClientMockRecorder) Block(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Block", reflect.TypeOf((*MockBaseClient)(nil).Block), height)
}

// BlockResults mocks base method.
func (m *MockBaseClient) BlockResults(height *int64) (*types1.ResultBlockResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockResults", height)
	ret0, _ := ret[0].(*types1.ResultBlockResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockResults indicates an expected call of BlockResults.
func (mr *MockBaseClientMockRecorder) BlockResults(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockResults", reflect.TypeOf((*MockBaseClient)(nil).BlockResults), height)
}

// BlockchainInfo mocks base method.
func (m *MockBaseClient) BlockchainInfo(minHeight, maxHeight int64) (*types1.ResultBlockchainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockchainInfo", minHeight, maxHeight)
	ret0, _ := ret[0].(*types1.ResultBlockchainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockchainInfo indicates an expected call of BlockchainInfo.
func (mr *MockBaseClientMockRecorder) BlockchainInfo(minHeight, maxHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockchainInfo", reflect.TypeOf((*MockBaseClient)(nil).BlockchainInfo), minHeight, maxHeight)
}

// Broadcast mocks base method.
func (m *MockBaseClient) Broadcast(txBytes []byte, broadcastMode string) (types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broadcast", txBytes, broadcastMode)
	ret0, _ := ret[0].(types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Broadcast indicates an expected call of Broadcast.
func (mr *MockBaseClientMockRecorder) Broadcast(txBytes, broadcastMode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockBaseClient)(nil).Broadcast), txBytes, broadcastMode)
}

// BuildAndBroadcast mocks base method.
func (m *MockBaseClient) BuildAndBroadcast(fromName, passphrase, memo string, msgs []types.Msg, accNumber, seqNumber uint64) (types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildAndBroadcast", fromName, passphrase, memo, msgs, accNumber, seqNumber)
	ret0, _ := ret[0].(types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildAndBroadcast indicates an expected call of BuildAndBroadcast.
func (mr *MockBaseClientMockRecorder) BuildAndBroadcast(fromName, passphrase, memo, msgs, accNumber, seqNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildAndBroadcast", reflect.TypeOf((*MockBaseClient)(nil).BuildAndBroadcast), fromName, passphrase, memo, msgs, accNumber, seqNumber)
}

// BuildStdTx mocks base method.
func (m *MockBaseClient) BuildStdTx(fromName, passphrase, memo string, msgs []types.Msg, accNumber, seqNumber uint64) (*types0.StdTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildStdTx", fromName, passphrase, memo, msgs, accNumber, seqNumber)
	ret0, _ := ret[0].(types0.StdTx)
	ret1, _ := ret[1].(error)
	return &ret0, ret1
}

// BuildStdTx indicates an expected call of BuildStdTx.
func (mr *MockBaseClientMockRecorder) BuildStdTx(fromName, passphrase, memo, msgs, accNumber, seqNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildStdTx", reflect.TypeOf((*MockBaseClient)(nil).BuildStdTx), fromName, passphrase, memo, msgs, accNumber, seqNumber)
}

// BuildTxForSim mocks base method.
func (m *MockBaseClient) BuildTxForSim(msgs []types.Msg, memo string, accNumber, seqNumber uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildTxForSim", msgs, memo, accNumber, seqNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildTxForSim indicates an expected call of BuildTxForSim.
func (mr *MockBaseClientMockRecorder) BuildTxForSim(msgs, memo, accNumber, seqNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildTxForSim", reflect.TypeOf((*MockBaseClient)(nil).BuildTxForSim), msgs, memo, accNumber, seqNumber)
}

// BuildUnsignedStdTxOffline mocks base method.
func (m *MockBaseClient) BuildUnsignedStdTxOffline(msgs []types.Msg, memo string) *types0.StdTx {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildUnsignedStdTxOffline", msgs, memo)
	ret0, _ := ret[0].(types0.StdTx)
	return &ret0
}

// BuildUnsignedStdTxOffline indicates an expected call of BuildUnsignedStdTxOffline.
func (mr *MockBaseClientMockRecorder) BuildUnsignedStdTxOffline(msgs, memo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildUnsignedStdTxOffline", reflect.TypeOf((*MockBaseClient)(nil).BuildUnsignedStdTxOffline), msgs, memo)
}

// CalculateGas mocks base method.
func (m *MockBaseClient) CalculateGas(txBytes []byte) (types0.StdFee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateGas", txBytes)
	ret0, _ := ret[0].(types0.StdFee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateGas indicates an expected call of CalculateGas.
func (mr *MockBaseClientMockRecorder) CalculateGas(txBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateGas", reflect.TypeOf((*MockBaseClient)(nil).CalculateGas), txBytes)
}

// Commit mocks base method.
func (m *MockBaseClient) Commit(height *int64) (*types1.ResultCommit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", height)
	ret0, _ := ret[0].(*types1.ResultCommit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit.
func (mr *MockBaseClientMockRecorder) Commit(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockBaseClient)(nil).Commit), height)
}

// Genesis mocks base method.
func (m *MockBaseClient) Genesis() (*types1.ResultGenesis, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Genesis")
	ret0, _ := ret[0].(*types1.ResultGenesis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Genesis indicates an expected call of Genesis.
func (mr *MockBaseClientMockRecorder) Genesis() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Genesis", reflect.TypeOf((*MockBaseClient)(nil).Genesis))
}

// GetCodec mocks base method.
func (m *MockBaseClient) GetCodec() *codec.Codec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCodec")
	ret0, _ := ret[0].(*codec.Codec)
	return ret0
}

// GetCodec indicates an expected call of GetCodec.
func (mr *MockBaseClientMockRecorder) GetCodec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCodec", reflect.TypeOf((*MockBaseClient)(nil).GetCodec))
}

// GetConfig mocks base method.
func (m *MockBaseClient) GetConfig() ClientConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(ClientConfig)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockBaseClientMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockBaseClient)(nil).GetConfig))
}

// Query mocks base method.
func (m *MockBaseClient) Query(path string, key bytes.HexBytes) ([]byte, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", path, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Query indicates an expected call of Query.
func (mr *MockBaseClientMockRecorder) Query(path, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockBaseClient)(nil).Query), path, key)
}

// QueryStore mocks base method.
func (m *MockBaseClient) QueryStore(key bytes.HexBytes, storeName, endPath string) ([]byte, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryStore", key, storeName, endPath)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryStore indicates an expected call of QueryStore.
func (mr *MockBaseClientMockRecorder) QueryStore(key, storeName, endPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryStore", reflect.TypeOf((*MockBaseClient)(nil).QueryStore), key, storeName, endPath)
}

// Tx mocks base method.
func (m *MockBaseClient) Tx(hash []byte, prove bool) (*types1.ResultTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx", hash, prove)
	ret0, _ := ret[0].(*types1.ResultTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tx indicates an expected call of Tx.
func (mr *MockBaseClientMockRecorder) Tx(hash, prove interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockBaseClient)(nil).Tx), hash, prove)
}

// TxSearch mocks base method.
func (m *MockBaseClient) TxSearch(query string, prove bool, page, perPage int, orderBy string) (*types1.ResultTxSearch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxSearch", query, prove, page, perPage, orderBy)
	ret0, _ := ret[0].(*types1.ResultTxSearch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxSearch indicates an expected call of TxSearch.
func (mr *MockBaseClientMockRecorder) TxSearch(query, prove, page, perPage, orderBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxSearch", reflect.TypeOf((*MockBaseClient)(nil).TxSearch), query, prove, page, perPage, orderBy)
}

// Validators mocks base method.
func (m *MockBaseClient) Validators(height *int64, page, perPage int) (*types1.ResultValidators, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validators", height, page, perPage)
	ret0, _ := ret[0].(*types1.ResultValidators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validators indicates an expected call of Validators.
func (mr *MockBaseClientMockRecorder) Validators(height, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validators", reflect.TypeOf((*MockBaseClient)(nil).Validators), height, page, perPage)
}

// MockTxHandler is a mock of TxHandler interface.
type MockTxHandler struct {
	ctrl     *gomock.Controller
	recorder *MockTxHandlerMockRecorder
}

// MockTxHandlerMockRecorder is the mock recorder for MockTxHandler.
type MockTxHandlerMockRecorder struct {
	mock *MockTxHandler
}

// NewMockTxHandler creates a new mock instance.
func NewMockTxHandler(ctrl *gomock.Controller) *MockTxHandler {
	mock := &MockTxHandler{ctrl: ctrl}
	mock.recorder = &MockTxHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTxHandler) EXPECT() *MockTxHandlerMockRecorder {
	return m.recorder
}

// BuildAndBroadcast mocks base method.
func (m *MockTxHandler) BuildAndBroadcast(fromName, passphrase, memo string, msgs []types.Msg, accNumber, seqNumber uint64) (types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildAndBroadcast", fromName, passphrase, memo, msgs, accNumber, seqNumber)
	ret0, _ := ret[0].(types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildAndBroadcast indicates an expected call of BuildAndBroadcast.
func (mr *MockTxHandlerMockRecorder) BuildAndBroadcast(fromName, passphrase, memo, msgs, accNumber, seqNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildAndBroadcast", reflect.TypeOf((*MockTxHandler)(nil).BuildAndBroadcast), fromName, passphrase, memo, msgs, accNumber, seqNumber)
}

// BuildStdTx mocks base method.
func (m *MockTxHandler) BuildStdTx(fromName, passphrase, memo string, msgs []types.Msg, accNumber, seqNumber uint64) (types0.StdTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildStdTx", fromName, passphrase, memo, msgs, accNumber, seqNumber)
	ret0, _ := ret[0].(types0.StdTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildStdTx indicates an expected call of BuildStdTx.
func (mr *MockTxHandlerMockRecorder) BuildStdTx(fromName, passphrase, memo, msgs, accNumber, seqNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildStdTx", reflect.TypeOf((*MockTxHandler)(nil).BuildStdTx), fromName, passphrase, memo, msgs, accNumber, seqNumber)
}

// BuildUnsignedStdTxOffline mocks base method.
func (m *MockTxHandler) BuildUnsignedStdTxOffline(msgs []types.Msg, memo string) types0.StdTx {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildUnsignedStdTxOffline", msgs, memo)
	ret0, _ := ret[0].(types0.StdTx)
	return ret0
}

// BuildUnsignedStdTxOffline indicates an expected call of BuildUnsignedStdTxOffline.
func (mr *MockTxHandlerMockRecorder) BuildUnsignedStdTxOffline(msgs, memo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildUnsignedStdTxOffline", reflect.TypeOf((*MockTxHandler)(nil).BuildUnsignedStdTxOffline), msgs, memo)
}

// MockSimulationHandler is a mock of SimulationHandler interface.
type MockSimulationHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSimulationHandlerMockRecorder
}

// MockSimulationHandlerMockRecorder is the mock recorder for MockSimulationHandler.
type MockSimulationHandlerMockRecorder struct {
	mock *MockSimulationHandler
}

// NewMockSimulationHandler creates a new mock instance.
func NewMockSimulationHandler(ctrl *gomock.Controller) *MockSimulationHandler {
	mock := &MockSimulationHandler{ctrl: ctrl}
	mock.recorder = &MockSimulationHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSimulationHandler) EXPECT() *MockSimulationHandlerMockRecorder {
	return m.recorder
}

// BuildTxForSim mocks base method.
func (m *MockSimulationHandler) BuildTxForSim(msgs []types.Msg, memo string, accNumber, seqNumber uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildTxForSim", msgs, memo, accNumber, seqNumber)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildTxForSim indicates an expected call of BuildTxForSim.
func (mr *MockSimulationHandlerMockRecorder) BuildTxForSim(msgs, memo, accNumber, seqNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildTxForSim", reflect.TypeOf((*MockSimulationHandler)(nil).BuildTxForSim), msgs, memo, accNumber, seqNumber)
}

// CalculateGas mocks base method.
func (m *MockSimulationHandler) CalculateGas(txBytes []byte) (types0.StdFee, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateGas", txBytes)
	ret0, _ := ret[0].(types0.StdFee)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateGas indicates an expected call of CalculateGas.
func (mr *MockSimulationHandlerMockRecorder) CalculateGas(txBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateGas", reflect.TypeOf((*MockSimulationHandler)(nil).CalculateGas), txBytes)
}

// MockClientQuery is a mock of ClientQuery interface.
type MockClientQuery struct {
	ctrl     *gomock.Controller
	recorder *MockClientQueryMockRecorder
}

// MockClientQueryMockRecorder is the mock recorder for MockClientQuery.
type MockClientQueryMockRecorder struct {
	mock *MockClientQuery
}

// NewMockClientQuery creates a new mock instance.
func NewMockClientQuery(ctrl *gomock.Controller) *MockClientQuery {
	mock := &MockClientQuery{ctrl: ctrl}
	mock.recorder = &MockClientQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientQuery) EXPECT() *MockClientQueryMockRecorder {
	return m.recorder
}

// Block mocks base method.
func (m *MockClientQuery) Block(height *int64) (*types1.ResultBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Block", height)
	ret0, _ := ret[0].(*types1.ResultBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Block indicates an expected call of Block.
func (mr *MockClientQueryMockRecorder) Block(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Block", reflect.TypeOf((*MockClientQuery)(nil).Block), height)
}

// BlockResults mocks base method.
func (m *MockClientQuery) BlockResults(height *int64) (*types1.ResultBlockResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockResults", height)
	ret0, _ := ret[0].(*types1.ResultBlockResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockResults indicates an expected call of BlockResults.
func (mr *MockClientQueryMockRecorder) BlockResults(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockResults", reflect.TypeOf((*MockClientQuery)(nil).BlockResults), height)
}

// BlockchainInfo mocks base method.
func (m *MockClientQuery) BlockchainInfo(minHeight, maxHeight int64) (*types1.ResultBlockchainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockchainInfo", minHeight, maxHeight)
	ret0, _ := ret[0].(*types1.ResultBlockchainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockchainInfo indicates an expected call of BlockchainInfo.
func (mr *MockClientQueryMockRecorder) BlockchainInfo(minHeight, maxHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockchainInfo", reflect.TypeOf((*MockClientQuery)(nil).BlockchainInfo), minHeight, maxHeight)
}

// Commit mocks base method.
func (m *MockClientQuery) Commit(height *int64) (*types1.ResultCommit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", height)
	ret0, _ := ret[0].(*types1.ResultCommit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit.
func (mr *MockClientQueryMockRecorder) Commit(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockClientQuery)(nil).Commit), height)
}

// Genesis mocks base method.
func (m *MockClientQuery) Genesis() (*types1.ResultGenesis, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Genesis")
	ret0, _ := ret[0].(*types1.ResultGenesis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Genesis indicates an expected call of Genesis.
func (mr *MockClientQueryMockRecorder) Genesis() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Genesis", reflect.TypeOf((*MockClientQuery)(nil).Genesis))
}

// Query mocks base method.
func (m *MockClientQuery) Query(path string, key bytes.HexBytes) ([]byte, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", path, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Query indicates an expected call of Query.
func (mr *MockClientQueryMockRecorder) Query(path, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockClientQuery)(nil).Query), path, key)
}

// QueryStore mocks base method.
func (m *MockClientQuery) QueryStore(key bytes.HexBytes, storeName, endPath string) ([]byte, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryStore", key, storeName, endPath)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryStore indicates an expected call of QueryStore.
func (mr *MockClientQueryMockRecorder) QueryStore(key, storeName, endPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryStore", reflect.TypeOf((*MockClientQuery)(nil).QueryStore), key, storeName, endPath)
}

// Tx mocks base method.
func (m *MockClientQuery) Tx(hash []byte, prove bool) (*types1.ResultTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx", hash, prove)
	ret0, _ := ret[0].(*types1.ResultTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tx indicates an expected call of Tx.
func (mr *MockClientQueryMockRecorder) Tx(hash, prove interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockClientQuery)(nil).Tx), hash, prove)
}

// TxSearch mocks base method.
func (m *MockClientQuery) TxSearch(query string, prove bool, page, perPage int, orderBy string) (*types1.ResultTxSearch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxSearch", query, prove, page, perPage, orderBy)
	ret0, _ := ret[0].(*types1.ResultTxSearch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxSearch indicates an expected call of TxSearch.
func (mr *MockClientQueryMockRecorder) TxSearch(query, prove, page, perPage, orderBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxSearch", reflect.TypeOf((*MockClientQuery)(nil).TxSearch), query, prove, page, perPage, orderBy)
}

// Validators mocks base method.
func (m *MockClientQuery) Validators(height *int64, page, perPage int) (*types1.ResultValidators, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validators", height, page, perPage)
	ret0, _ := ret[0].(*types1.ResultValidators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validators indicates an expected call of Validators.
func (mr *MockClientQueryMockRecorder) Validators(height, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validators", reflect.TypeOf((*MockClientQuery)(nil).Validators), height, page, perPage)
}

// MockClientTx is a mock of ClientTx interface.
type MockClientTx struct {
	ctrl     *gomock.Controller
	recorder *MockClientTxMockRecorder
}

// MockClientTxMockRecorder is the mock recorder for MockClientTx.
type MockClientTxMockRecorder struct {
	mock *MockClientTx
}

// NewMockClientTx creates a new mock instance.
func NewMockClientTx(ctrl *gomock.Controller) *MockClientTx {
	mock := &MockClientTx{ctrl: ctrl}
	mock.recorder = &MockClientTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientTx) EXPECT() *MockClientTxMockRecorder {
	return m.recorder
}

// Broadcast mocks base method.
func (m *MockClientTx) Broadcast(txBytes []byte, broadcastMode string) (types.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broadcast", txBytes, broadcastMode)
	ret0, _ := ret[0].(types.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Broadcast indicates an expected call of Broadcast.
func (mr *MockClientTxMockRecorder) Broadcast(txBytes, broadcastMode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockClientTx)(nil).Broadcast), txBytes, broadcastMode)
}

func (m *MockBaseClient) Status() (*types1.ResultStatus, error) {
	//TODO implement me
	panic("implement me")
}