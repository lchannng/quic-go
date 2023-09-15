// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apernet/quic-go (interfaces: Unpacker)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"
	time "time"

	protocol "github.com/apernet/quic-go/internal/protocol"
	wire "github.com/apernet/quic-go/internal/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockUnpacker is a mock of Unpacker interface.
type MockUnpacker struct {
	ctrl     *gomock.Controller
	recorder *MockUnpackerMockRecorder
}

// MockUnpackerMockRecorder is the mock recorder for MockUnpacker.
type MockUnpackerMockRecorder struct {
	mock *MockUnpacker
}

// NewMockUnpacker creates a new mock instance.
func NewMockUnpacker(ctrl *gomock.Controller) *MockUnpacker {
	mock := &MockUnpacker{ctrl: ctrl}
	mock.recorder = &MockUnpackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnpacker) EXPECT() *MockUnpackerMockRecorder {
	return m.recorder
}

// UnpackLongHeader mocks base method.
func (m *MockUnpacker) UnpackLongHeader(arg0 *wire.Header, arg1 time.Time, arg2 []byte, arg3 protocol.VersionNumber) (*unpackedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpackLongHeader", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*unpackedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpackLongHeader indicates an expected call of UnpackLongHeader.
func (mr *MockUnpackerMockRecorder) UnpackLongHeader(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpackLongHeader", reflect.TypeOf((*MockUnpacker)(nil).UnpackLongHeader), arg0, arg1, arg2, arg3)
}

// UnpackShortHeader mocks base method.
func (m *MockUnpacker) UnpackShortHeader(arg0 time.Time, arg1 []byte) (protocol.PacketNumber, protocol.PacketNumberLen, protocol.KeyPhaseBit, []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpackShortHeader", arg0, arg1)
	ret0, _ := ret[0].(protocol.PacketNumber)
	ret1, _ := ret[1].(protocol.PacketNumberLen)
	ret2, _ := ret[2].(protocol.KeyPhaseBit)
	ret3, _ := ret[3].([]byte)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// UnpackShortHeader indicates an expected call of UnpackShortHeader.
func (mr *MockUnpackerMockRecorder) UnpackShortHeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpackShortHeader", reflect.TypeOf((*MockUnpacker)(nil).UnpackShortHeader), arg0, arg1)
}
