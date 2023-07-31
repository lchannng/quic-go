// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/apernet/quic-go/internal/ackhandler (interfaces: SentPacketTracker)

// Package ackhandler is a generated GoMock package.
package ackhandler

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/apernet/quic-go/internal/protocol"
)

// MockSentPacketTracker is a mock of SentPacketTracker interface.
type MockSentPacketTracker struct {
	ctrl     *gomock.Controller
	recorder *MockSentPacketTrackerMockRecorder
}

// MockSentPacketTrackerMockRecorder is the mock recorder for MockSentPacketTracker.
type MockSentPacketTrackerMockRecorder struct {
	mock *MockSentPacketTracker
}

// NewMockSentPacketTracker creates a new mock instance.
func NewMockSentPacketTracker(ctrl *gomock.Controller) *MockSentPacketTracker {
	mock := &MockSentPacketTracker{ctrl: ctrl}
	mock.recorder = &MockSentPacketTrackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSentPacketTracker) EXPECT() *MockSentPacketTrackerMockRecorder {
	return m.recorder
}

// GetLowestPacketNotConfirmedAcked mocks base method.
func (m *MockSentPacketTracker) GetLowestPacketNotConfirmedAcked() protocol.PacketNumber {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLowestPacketNotConfirmedAcked")
	ret0, _ := ret[0].(protocol.PacketNumber)
	return ret0
}

// GetLowestPacketNotConfirmedAcked indicates an expected call of GetLowestPacketNotConfirmedAcked.
func (mr *MockSentPacketTrackerMockRecorder) GetLowestPacketNotConfirmedAcked() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLowestPacketNotConfirmedAcked", reflect.TypeOf((*MockSentPacketTracker)(nil).GetLowestPacketNotConfirmedAcked))
}

// ReceivedPacket mocks base method.
func (m *MockSentPacketTracker) ReceivedPacket(arg0 protocol.EncryptionLevel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReceivedPacket", arg0)
}

// ReceivedPacket indicates an expected call of ReceivedPacket.
func (mr *MockSentPacketTrackerMockRecorder) ReceivedPacket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceivedPacket", reflect.TypeOf((*MockSentPacketTracker)(nil).ReceivedPacket), arg0)
}
