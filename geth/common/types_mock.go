// Code generated by MockGen. DO NOT EDIT.
// Source: geth/common/types.go

// Package common is a generated GoMock package.
package common

import (
	context "context"
	accounts "github.com/ethereum/go-ethereum/accounts"
	keystore "github.com/ethereum/go-ethereum/accounts/keystore"
	common "github.com/ethereum/go-ethereum/common"
	les "github.com/ethereum/go-ethereum/les"
	node "github.com/ethereum/go-ethereum/node"
	whisperv6 "github.com/ethereum/go-ethereum/whisper/whisperv6"
	gomock "github.com/golang/mock/gomock"
	otto "github.com/robertkrimen/otto"
	params "github.com/status-im/status-go/geth/params"
	rpc "github.com/status-im/status-go/geth/rpc"
	reflect "reflect"
)

// MockNodeManager is a mock of NodeManager interface
type MockNodeManager struct {
	ctrl     *gomock.Controller
	recorder *MockNodeManagerMockRecorder
}

// MockNodeManagerMockRecorder is the mock recorder for MockNodeManager
type MockNodeManagerMockRecorder struct {
	mock *MockNodeManager
}

// NewMockNodeManager creates a new mock instance
func NewMockNodeManager(ctrl *gomock.Controller) *MockNodeManager {
	mock := &MockNodeManager{ctrl: ctrl}
	mock.recorder = &MockNodeManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeManager) EXPECT() *MockNodeManagerMockRecorder {
	return m.recorder
}

// StartNode mocks base method
func (m *MockNodeManager) StartNode(config *params.NodeConfig) error {
	ret := m.ctrl.Call(m, "StartNode", config)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartNode indicates an expected call of StartNode
func (mr *MockNodeManagerMockRecorder) StartNode(config interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartNode", reflect.TypeOf((*MockNodeManager)(nil).StartNode), config)
}

// EnsureSync mocks base method
func (m *MockNodeManager) EnsureSync(ctx context.Context) error {
	ret := m.ctrl.Call(m, "EnsureSync", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureSync indicates an expected call of EnsureSync
func (mr *MockNodeManagerMockRecorder) EnsureSync(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureSync", reflect.TypeOf((*MockNodeManager)(nil).EnsureSync), ctx)
}

// StopNode mocks base method
func (m *MockNodeManager) StopNode() error {
	ret := m.ctrl.Call(m, "StopNode")
	ret0, _ := ret[0].(error)
	return ret0
}

// StopNode indicates an expected call of StopNode
func (mr *MockNodeManagerMockRecorder) StopNode() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopNode", reflect.TypeOf((*MockNodeManager)(nil).StopNode))
}

// IsNodeRunning mocks base method
func (m *MockNodeManager) IsNodeRunning() bool {
	ret := m.ctrl.Call(m, "IsNodeRunning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNodeRunning indicates an expected call of IsNodeRunning
func (mr *MockNodeManagerMockRecorder) IsNodeRunning() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNodeRunning", reflect.TypeOf((*MockNodeManager)(nil).IsNodeRunning))
}

// NodeConfig mocks base method
func (m *MockNodeManager) NodeConfig() (*params.NodeConfig, error) {
	ret := m.ctrl.Call(m, "NodeConfig")
	ret0, _ := ret[0].(*params.NodeConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeConfig indicates an expected call of NodeConfig
func (mr *MockNodeManagerMockRecorder) NodeConfig() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeConfig", reflect.TypeOf((*MockNodeManager)(nil).NodeConfig))
}

// Node mocks base method
func (m *MockNodeManager) Node() (*node.Node, error) {
	ret := m.ctrl.Call(m, "Node")
	ret0, _ := ret[0].(*node.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Node indicates an expected call of Node
func (mr *MockNodeManagerMockRecorder) Node() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Node", reflect.TypeOf((*MockNodeManager)(nil).Node))
}

// PopulateStaticPeers mocks base method
func (m *MockNodeManager) PopulateStaticPeers() error {
	ret := m.ctrl.Call(m, "PopulateStaticPeers")
	ret0, _ := ret[0].(error)
	return ret0
}

// PopulateStaticPeers indicates an expected call of PopulateStaticPeers
func (mr *MockNodeManagerMockRecorder) PopulateStaticPeers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopulateStaticPeers", reflect.TypeOf((*MockNodeManager)(nil).PopulateStaticPeers))
}

// AddPeer mocks base method
func (m *MockNodeManager) AddPeer(url string) error {
	ret := m.ctrl.Call(m, "AddPeer", url)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPeer indicates an expected call of AddPeer
func (mr *MockNodeManagerMockRecorder) AddPeer(url interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPeer", reflect.TypeOf((*MockNodeManager)(nil).AddPeer), url)
}

// PeerCount mocks base method
func (m *MockNodeManager) PeerCount() int {
	ret := m.ctrl.Call(m, "PeerCount")
	ret0, _ := ret[0].(int)
	return ret0
}

// PeerCount indicates an expected call of PeerCount
func (mr *MockNodeManagerMockRecorder) PeerCount() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerCount", reflect.TypeOf((*MockNodeManager)(nil).PeerCount))
}

// LightEthereumService mocks base method
func (m *MockNodeManager) LightEthereumService() (*les.LightEthereum, error) {
	ret := m.ctrl.Call(m, "LightEthereumService")
	ret0, _ := ret[0].(*les.LightEthereum)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LightEthereumService indicates an expected call of LightEthereumService
func (mr *MockNodeManagerMockRecorder) LightEthereumService() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LightEthereumService", reflect.TypeOf((*MockNodeManager)(nil).LightEthereumService))
}

// WhisperService mocks base method
func (m *MockNodeManager) WhisperService() (*whisperv6.Whisper, error) {
	ret := m.ctrl.Call(m, "WhisperService")
	ret0, _ := ret[0].(*whisperv6.Whisper)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WhisperService indicates an expected call of WhisperService
func (mr *MockNodeManagerMockRecorder) WhisperService() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhisperService", reflect.TypeOf((*MockNodeManager)(nil).WhisperService))
}

// AccountManager mocks base method
func (m *MockNodeManager) AccountManager() (*accounts.Manager, error) {
	ret := m.ctrl.Call(m, "AccountManager")
	ret0, _ := ret[0].(*accounts.Manager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountManager indicates an expected call of AccountManager
func (mr *MockNodeManagerMockRecorder) AccountManager() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountManager", reflect.TypeOf((*MockNodeManager)(nil).AccountManager))
}

// AccountKeyStore mocks base method
func (m *MockNodeManager) AccountKeyStore() (*keystore.KeyStore, error) {
	ret := m.ctrl.Call(m, "AccountKeyStore")
	ret0, _ := ret[0].(*keystore.KeyStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AccountKeyStore indicates an expected call of AccountKeyStore
func (mr *MockNodeManagerMockRecorder) AccountKeyStore() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountKeyStore", reflect.TypeOf((*MockNodeManager)(nil).AccountKeyStore))
}

// RPCClient mocks base method
func (m *MockNodeManager) RPCClient() *rpc.Client {
	ret := m.ctrl.Call(m, "RPCClient")
	ret0, _ := ret[0].(*rpc.Client)
	return ret0
}

// RPCClient indicates an expected call of RPCClient
func (mr *MockNodeManagerMockRecorder) RPCClient() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RPCClient", reflect.TypeOf((*MockNodeManager)(nil).RPCClient))
}

// MockAccountManager is a mock of AccountManager interface
type MockAccountManager struct {
	ctrl     *gomock.Controller
	recorder *MockAccountManagerMockRecorder
}

// MockAccountManagerMockRecorder is the mock recorder for MockAccountManager
type MockAccountManagerMockRecorder struct {
	mock *MockAccountManager
}

// NewMockAccountManager creates a new mock instance
func NewMockAccountManager(ctrl *gomock.Controller) *MockAccountManager {
	mock := &MockAccountManager{ctrl: ctrl}
	mock.recorder = &MockAccountManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountManager) EXPECT() *MockAccountManagerMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method
func (m *MockAccountManager) CreateAccount(password string) (string, string, string, error) {
	ret := m.ctrl.Call(m, "CreateAccount", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// CreateAccount indicates an expected call of CreateAccount
func (mr *MockAccountManagerMockRecorder) CreateAccount(password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockAccountManager)(nil).CreateAccount), password)
}

// CreateChildAccount mocks base method
func (m *MockAccountManager) CreateChildAccount(parentAddress, password string) (string, string, error) {
	ret := m.ctrl.Call(m, "CreateChildAccount", parentAddress, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateChildAccount indicates an expected call of CreateChildAccount
func (mr *MockAccountManagerMockRecorder) CreateChildAccount(parentAddress, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChildAccount", reflect.TypeOf((*MockAccountManager)(nil).CreateChildAccount), parentAddress, password)
}

// RecoverAccount mocks base method
func (m *MockAccountManager) RecoverAccount(password, mnemonic string) (string, string, error) {
	ret := m.ctrl.Call(m, "RecoverAccount", password, mnemonic)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RecoverAccount indicates an expected call of RecoverAccount
func (mr *MockAccountManagerMockRecorder) RecoverAccount(password, mnemonic interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecoverAccount", reflect.TypeOf((*MockAccountManager)(nil).RecoverAccount), password, mnemonic)
}

// VerifyAccountPassword mocks base method
func (m *MockAccountManager) VerifyAccountPassword(keyStoreDir, address, password string) (*keystore.Key, error) {
	ret := m.ctrl.Call(m, "VerifyAccountPassword", keyStoreDir, address, password)
	ret0, _ := ret[0].(*keystore.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyAccountPassword indicates an expected call of VerifyAccountPassword
func (mr *MockAccountManagerMockRecorder) VerifyAccountPassword(keyStoreDir, address, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAccountPassword", reflect.TypeOf((*MockAccountManager)(nil).VerifyAccountPassword), keyStoreDir, address, password)
}

// SelectAccount mocks base method
func (m *MockAccountManager) SelectAccount(address, password string) error {
	ret := m.ctrl.Call(m, "SelectAccount", address, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// SelectAccount indicates an expected call of SelectAccount
func (mr *MockAccountManagerMockRecorder) SelectAccount(address, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAccount", reflect.TypeOf((*MockAccountManager)(nil).SelectAccount), address, password)
}

// ReSelectAccount mocks base method
func (m *MockAccountManager) ReSelectAccount() error {
	ret := m.ctrl.Call(m, "ReSelectAccount")
	ret0, _ := ret[0].(error)
	return ret0
}

// ReSelectAccount indicates an expected call of ReSelectAccount
func (mr *MockAccountManagerMockRecorder) ReSelectAccount() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReSelectAccount", reflect.TypeOf((*MockAccountManager)(nil).ReSelectAccount))
}

// SelectedAccount mocks base method
func (m *MockAccountManager) SelectedAccount() (*SelectedExtKey, error) {
	ret := m.ctrl.Call(m, "SelectedAccount")
	ret0, _ := ret[0].(*SelectedExtKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectedAccount indicates an expected call of SelectedAccount
func (mr *MockAccountManagerMockRecorder) SelectedAccount() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectedAccount", reflect.TypeOf((*MockAccountManager)(nil).SelectedAccount))
}

// Logout mocks base method
func (m *MockAccountManager) Logout() error {
	ret := m.ctrl.Call(m, "Logout")
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout
func (mr *MockAccountManagerMockRecorder) Logout() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAccountManager)(nil).Logout))
}

// Accounts mocks base method
func (m *MockAccountManager) Accounts() ([]common.Address, error) {
	ret := m.ctrl.Call(m, "Accounts")
	ret0, _ := ret[0].([]common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Accounts indicates an expected call of Accounts
func (mr *MockAccountManagerMockRecorder) Accounts() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accounts", reflect.TypeOf((*MockAccountManager)(nil).Accounts))
}

// AccountsRPCHandler mocks base method
func (m *MockAccountManager) AccountsRPCHandler() rpc.Handler {
	ret := m.ctrl.Call(m, "AccountsRPCHandler")
	ret0, _ := ret[0].(rpc.Handler)
	return ret0
}

// AccountsRPCHandler indicates an expected call of AccountsRPCHandler
func (mr *MockAccountManagerMockRecorder) AccountsRPCHandler() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountsRPCHandler", reflect.TypeOf((*MockAccountManager)(nil).AccountsRPCHandler))
}

// AddressToDecryptedAccount mocks base method
func (m *MockAccountManager) AddressToDecryptedAccount(address, password string) (accounts.Account, *keystore.Key, error) {
	ret := m.ctrl.Call(m, "AddressToDecryptedAccount", address, password)
	ret0, _ := ret[0].(accounts.Account)
	ret1, _ := ret[1].(*keystore.Key)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddressToDecryptedAccount indicates an expected call of AddressToDecryptedAccount
func (mr *MockAccountManagerMockRecorder) AddressToDecryptedAccount(address, password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddressToDecryptedAccount", reflect.TypeOf((*MockAccountManager)(nil).AddressToDecryptedAccount), address, password)
}

// MockJailCell is a mock of JailCell interface
type MockJailCell struct {
	ctrl     *gomock.Controller
	recorder *MockJailCellMockRecorder
}

// MockJailCellMockRecorder is the mock recorder for MockJailCell
type MockJailCellMockRecorder struct {
	mock *MockJailCell
}

// NewMockJailCell creates a new mock instance
func NewMockJailCell(ctrl *gomock.Controller) *MockJailCell {
	mock := &MockJailCell{ctrl: ctrl}
	mock.recorder = &MockJailCellMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJailCell) EXPECT() *MockJailCellMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockJailCell) Set(arg0 string, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockJailCellMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockJailCell)(nil).Set), arg0, arg1)
}

// Get mocks base method
func (m *MockJailCell) Get(arg0 string) (otto.Value, error) {
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(otto.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockJailCellMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockJailCell)(nil).Get), arg0)
}

// Run mocks base method
func (m *MockJailCell) Run(arg0 interface{}) (otto.Value, error) {
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(otto.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockJailCellMockRecorder) Run(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockJailCell)(nil).Run), arg0)
}

// Call mocks base method
func (m *MockJailCell) Call(item string, this interface{}, args ...interface{}) (otto.Value, error) {
	varargs := []interface{}{item, this}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Call", varargs...)
	ret0, _ := ret[0].(otto.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call
func (mr *MockJailCellMockRecorder) Call(item, this interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{item, this}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockJailCell)(nil).Call), varargs...)
}

// Stop mocks base method
func (m *MockJailCell) Stop() error {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockJailCellMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockJailCell)(nil).Stop))
}

// MockJailManager is a mock of JailManager interface
type MockJailManager struct {
	ctrl     *gomock.Controller
	recorder *MockJailManagerMockRecorder
}

// MockJailManagerMockRecorder is the mock recorder for MockJailManager
type MockJailManagerMockRecorder struct {
	mock *MockJailManager
}

// NewMockJailManager creates a new mock instance
func NewMockJailManager(ctrl *gomock.Controller) *MockJailManager {
	mock := &MockJailManager{ctrl: ctrl}
	mock.recorder = &MockJailManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJailManager) EXPECT() *MockJailManagerMockRecorder {
	return m.recorder
}

// Call mocks base method
func (m *MockJailManager) Call(chatID, this, args string) string {
	ret := m.ctrl.Call(m, "Call", chatID, this, args)
	ret0, _ := ret[0].(string)
	return ret0
}

// Call indicates an expected call of Call
func (mr *MockJailManagerMockRecorder) Call(chatID, this, args interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockJailManager)(nil).Call), chatID, this, args)
}

// CreateCell mocks base method
func (m *MockJailManager) CreateCell(chatID string) (JailCell, error) {
	ret := m.ctrl.Call(m, "CreateCell", chatID)
	ret0, _ := ret[0].(JailCell)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCell indicates an expected call of CreateCell
func (mr *MockJailManagerMockRecorder) CreateCell(chatID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCell", reflect.TypeOf((*MockJailManager)(nil).CreateCell), chatID)
}

// Parse mocks base method
func (m *MockJailManager) Parse(chatID, js string) string {
	ret := m.ctrl.Call(m, "Parse", chatID, js)
	ret0, _ := ret[0].(string)
	return ret0
}

// Parse indicates an expected call of Parse
func (mr *MockJailManagerMockRecorder) Parse(chatID, js interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockJailManager)(nil).Parse), chatID, js)
}

// CreateAndInitCell mocks base method
func (m *MockJailManager) CreateAndInitCell(chatID string, code ...string) string {
	varargs := []interface{}{chatID}
	for _, a := range code {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateAndInitCell", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateAndInitCell indicates an expected call of CreateAndInitCell
func (mr *MockJailManagerMockRecorder) CreateAndInitCell(chatID interface{}, code ...interface{}) *gomock.Call {
	varargs := append([]interface{}{chatID}, code...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAndInitCell", reflect.TypeOf((*MockJailManager)(nil).CreateAndInitCell), varargs...)
}

// Cell mocks base method
func (m *MockJailManager) Cell(chatID string) (JailCell, error) {
	ret := m.ctrl.Call(m, "Cell", chatID)
	ret0, _ := ret[0].(JailCell)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cell indicates an expected call of Cell
func (mr *MockJailManagerMockRecorder) Cell(chatID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cell", reflect.TypeOf((*MockJailManager)(nil).Cell), chatID)
}

// Execute mocks base method
func (m *MockJailManager) Execute(chatID, code string) string {
	ret := m.ctrl.Call(m, "Execute", chatID, code)
	ret0, _ := ret[0].(string)
	return ret0
}

// Execute indicates an expected call of Execute
func (mr *MockJailManagerMockRecorder) Execute(chatID, code interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockJailManager)(nil).Execute), chatID, code)
}

// SetBaseJS mocks base method
func (m *MockJailManager) SetBaseJS(js string) {
	m.ctrl.Call(m, "SetBaseJS", js)
}

// SetBaseJS indicates an expected call of SetBaseJS
func (mr *MockJailManagerMockRecorder) SetBaseJS(js interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBaseJS", reflect.TypeOf((*MockJailManager)(nil).SetBaseJS), js)
}

// Stop mocks base method
func (m *MockJailManager) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockJailManagerMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockJailManager)(nil).Stop))
}
