// Code generated by mockery v1.0.0. DO NOT EDIT.

package vendor_mocks

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import mock "github.com/stretchr/testify/mock"
import types "k8s.io/apimachinery/pkg/types"
import v1 "k8s.io/api/core/v1"
import watch "k8s.io/apimachinery/pkg/watch"

// PersistentVolumeClaimInterface is an autogenerated mock type for the PersistentVolumeClaimInterface type
type PersistentVolumeClaimInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *PersistentVolumeClaimInterface) Create(_a0 *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error) {
	ret := _m.Called(_a0)

	var r0 *v1.PersistentVolumeClaim
	if rf, ok := ret.Get(0).(func(*v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.PersistentVolumeClaim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.PersistentVolumeClaim) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name, options
func (_m *PersistentVolumeClaimInterface) Delete(name string, options *metav1.DeleteOptions) error {
	ret := _m.Called(name, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *metav1.DeleteOptions) error); ok {
		r0 = rf(name, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function with given fields: options, listOptions
func (_m *PersistentVolumeClaimInterface) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	ret := _m.Called(options, listOptions)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(options, listOptions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: name, options
func (_m *PersistentVolumeClaimInterface) Get(name string, options metav1.GetOptions) (*v1.PersistentVolumeClaim, error) {
	ret := _m.Called(name, options)

	var r0 *v1.PersistentVolumeClaim
	if rf, ok := ret.Get(0).(func(string, metav1.GetOptions) *v1.PersistentVolumeClaim); ok {
		r0 = rf(name, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.PersistentVolumeClaim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, metav1.GetOptions) error); ok {
		r1 = rf(name, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: opts
func (_m *PersistentVolumeClaimInterface) List(opts metav1.ListOptions) (*v1.PersistentVolumeClaimList, error) {
	ret := _m.Called(opts)

	var r0 *v1.PersistentVolumeClaimList
	if rf, ok := ret.Get(0).(func(metav1.ListOptions) *v1.PersistentVolumeClaimList); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.PersistentVolumeClaimList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(metav1.ListOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: name, pt, data, subresources
func (_m *PersistentVolumeClaimInterface) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v1.PersistentVolumeClaim, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name, pt, data)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.PersistentVolumeClaim
	if rf, ok := ret.Get(0).(func(string, types.PatchType, []byte, ...string) *v1.PersistentVolumeClaim); ok {
		r0 = rf(name, pt, data, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.PersistentVolumeClaim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.PatchType, []byte, ...string) error); ok {
		r1 = rf(name, pt, data, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *PersistentVolumeClaimInterface) Update(_a0 *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error) {
	ret := _m.Called(_a0)

	var r0 *v1.PersistentVolumeClaim
	if rf, ok := ret.Get(0).(func(*v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.PersistentVolumeClaim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.PersistentVolumeClaim) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: _a0
func (_m *PersistentVolumeClaimInterface) UpdateStatus(_a0 *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error) {
	ret := _m.Called(_a0)

	var r0 *v1.PersistentVolumeClaim
	if rf, ok := ret.Get(0).(func(*v1.PersistentVolumeClaim) *v1.PersistentVolumeClaim); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.PersistentVolumeClaim)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.PersistentVolumeClaim) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: opts
func (_m *PersistentVolumeClaimInterface) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(metav1.ListOptions) watch.Interface); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(metav1.ListOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}