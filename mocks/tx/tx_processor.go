// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocktx

import (
	client "github.com/hyperledger-labs/firefly-fabconnect/internal/fabric/client"
	mock "github.com/stretchr/testify/mock"

	tx "github.com/hyperledger-labs/firefly-fabconnect/internal/tx"
)

// TxProcessor is an autogenerated mock type for the TxProcessor type
type TxProcessor struct {
	mock.Mock
}

// Init provides a mock function with given fields: _a0
func (_m *TxProcessor) Init(_a0 client.RPCClient) {
	_m.Called(_a0)
}

// OnMessage provides a mock function with given fields: _a0
func (_m *TxProcessor) OnMessage(_a0 tx.TxContext) {
	_m.Called(_a0)
}
