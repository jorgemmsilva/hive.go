// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	identity "github.com/iotaledger/hive.go/identity"
	mock "github.com/stretchr/testify/mock"

	net "net"

	server "github.com/iotaledger/hive.go/autopeering/server"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// HandleMessage provides a mock function with given fields: s, fromAddr, from, data
func (_m *Handler) HandleMessage(s *server.Server, fromAddr *net.UDPAddr, from *identity.Identity, data []byte) (bool, error) {
	ret := _m.Called(s, fromAddr, from, data)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*server.Server, *net.UDPAddr, *identity.Identity, []byte) bool); ok {
		r0 = rf(s, fromAddr, from, data)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*server.Server, *net.UDPAddr, *identity.Identity, []byte) error); ok {
		r1 = rf(s, fromAddr, from, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
