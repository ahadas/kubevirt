// Automatically generated by MockGen. DO NOT EDIT!
// Source: common.go

package device_manager

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of DeviceHandler interface
type MockDeviceHandler struct {
	ctrl     *gomock.Controller
	recorder *_MockDeviceHandlerRecorder
}

// Recorder for MockDeviceHandler (not exported)
type _MockDeviceHandlerRecorder struct {
	mock *MockDeviceHandler
}

func NewMockDeviceHandler(ctrl *gomock.Controller) *MockDeviceHandler {
	mock := &MockDeviceHandler{ctrl: ctrl}
	mock.recorder = &_MockDeviceHandlerRecorder{mock}
	return mock
}

func (_m *MockDeviceHandler) EXPECT() *_MockDeviceHandlerRecorder {
	return _m.recorder
}

func (_m *MockDeviceHandler) GetDeviceIOMMUGroup(basepath string, pciAddress string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetDeviceIOMMUGroup", basepath, pciAddress)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDeviceHandlerRecorder) GetDeviceIOMMUGroup(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDeviceIOMMUGroup", arg0, arg1)
}

func (_m *MockDeviceHandler) GetDeviceDriver(basepath string, pciAddress string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetDeviceDriver", basepath, pciAddress)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDeviceHandlerRecorder) GetDeviceDriver(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDeviceDriver", arg0, arg1)
}

func (_m *MockDeviceHandler) GetDeviceNumaNode(basepath string, pciAddress string) int {
	ret := _m.ctrl.Call(_m, "GetDeviceNumaNode", basepath, pciAddress)
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockDeviceHandlerRecorder) GetDeviceNumaNode(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDeviceNumaNode", arg0, arg1)
}

func (_m *MockDeviceHandler) GetDevicePCIID(basepath string, pciAddress string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetDevicePCIID", basepath, pciAddress)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDeviceHandlerRecorder) GetDevicePCIID(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDevicePCIID", arg0, arg1)
}

func (_m *MockDeviceHandler) GetMdevParentPCIAddr(mdevUUID string) (string, error) {
	ret := _m.ctrl.Call(_m, "GetMdevParentPCIAddr", mdevUUID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDeviceHandlerRecorder) GetMdevParentPCIAddr(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMdevParentPCIAddr", arg0)
}

func (_m *MockDeviceHandler) CreateMDEVType(mdevType string, parentID string) error {
	ret := _m.ctrl.Call(_m, "CreateMDEVType", mdevType, parentID)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDeviceHandlerRecorder) CreateMDEVType(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateMDEVType", arg0, arg1)
}

func (_m *MockDeviceHandler) RemoveMDEVType(mdevUUID string) error {
	ret := _m.ctrl.Call(_m, "RemoveMDEVType", mdevUUID)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDeviceHandlerRecorder) RemoveMDEVType(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoveMDEVType", arg0)
}

func (_m *MockDeviceHandler) ReadMDEVAvailableInstances(mdevType string, parentID string) (int, error) {
	ret := _m.ctrl.Call(_m, "ReadMDEVAvailableInstances", mdevType, parentID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDeviceHandlerRecorder) ReadMDEVAvailableInstances(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadMDEVAvailableInstances", arg0, arg1)
}
