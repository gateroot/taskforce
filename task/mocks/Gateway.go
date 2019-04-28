// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domaintask "taskforce/domain/task"
import mock "github.com/stretchr/testify/mock"

// Gateway is an autogenerated mock type for the Gateway type
type Gateway struct {
	mock.Mock
}

// Create provides a mock function with given fields: title, description
func (_m *Gateway) Create(title string, description string) (int, error) {
	ret := _m.Called(title, description)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(title, description)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(title, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: taskId
func (_m *Gateway) Delete(taskId int) error {
	ret := _m.Called(taskId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(taskId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Read provides a mock function with given fields: taskId
func (_m *Gateway) Read(taskId int) (domaintask.Task, error) {
	ret := _m.Called(taskId)

	var r0 domaintask.Task
	if rf, ok := ret.Get(0).(func(int) domaintask.Task); ok {
		r0 = rf(taskId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domaintask.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(taskId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *Gateway) Update(_a0 domaintask.Task) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(domaintask.Task) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
