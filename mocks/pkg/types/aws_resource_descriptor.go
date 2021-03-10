// Code generated by mockery v2.2.2. DO NOT EDIT.

package mocks

import (
	compare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	mock "github.com/stretchr/testify/mock"

	runtime "k8s.io/apimachinery/pkg/runtime"

	types "github.com/aws-controllers-k8s/runtime/pkg/types"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AWSResourceDescriptor is an autogenerated mock type for the AWSResourceDescriptor type
type AWSResourceDescriptor struct {
	mock.Mock
}

// Delta provides a mock function with given fields: a, b
func (_m *AWSResourceDescriptor) Delta(a types.AWSResource, b types.AWSResource) *compare.Delta {
	ret := _m.Called(a, b)

	var r0 *compare.Delta
	if rf, ok := ret.Get(0).(func(types.AWSResource, types.AWSResource) *compare.Delta); ok {
		r0 = rf(a, b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*compare.Delta)
		}
	}

	return r0
}

// EmptyRuntimeObject provides a mock function with given fields:
func (_m *AWSResourceDescriptor) EmptyRuntimeObject() runtime.Object {
	ret := _m.Called()

	var r0 runtime.Object
	if rf, ok := ret.Get(0).(func() runtime.Object); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(runtime.Object)
		}
	}

	return r0
}

// GroupKind provides a mock function with given fields:
func (_m *AWSResourceDescriptor) GroupKind() *v1.GroupKind {
	ret := _m.Called()

	var r0 *v1.GroupKind
	if rf, ok := ret.Get(0).(func() *v1.GroupKind); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.GroupKind)
		}
	}

	return r0
}

// IsManaged provides a mock function with given fields: _a0
func (_m *AWSResourceDescriptor) IsManaged(_a0 types.AWSResource) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.AWSResource) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MarkManaged provides a mock function with given fields: _a0
func (_m *AWSResourceDescriptor) MarkManaged(_a0 types.AWSResource) {
	_m.Called(_a0)
}

// MarkUnmanaged provides a mock function with given fields: _a0
func (_m *AWSResourceDescriptor) MarkUnmanaged(_a0 types.AWSResource) {
	_m.Called(_a0)
}

// ResourceFromRuntimeObject provides a mock function with given fields: _a0
func (_m *AWSResourceDescriptor) ResourceFromRuntimeObject(_a0 runtime.Object) types.AWSResource {
	ret := _m.Called(_a0)

	var r0 types.AWSResource
	if rf, ok := ret.Get(0).(func(runtime.Object) types.AWSResource); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AWSResource)
		}
	}

	return r0
}

// UpdateCRStatus provides a mock function with given fields: _a0
func (_m *AWSResourceDescriptor) UpdateCRStatus(_a0 types.AWSResource) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(types.AWSResource) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.AWSResource) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
