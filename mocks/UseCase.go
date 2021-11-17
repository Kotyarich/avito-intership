// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UseCase is an autogenerated mock type for the UseCase type
type UseCase struct {
	mock.Mock
}

// ChangeBalance provides a mock function with given fields: userId, amount
func (_m *UseCase) ChangeBalance(userId int64, amount float32) error {
	ret := _m.Called(userId, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, float32) error); ok {
		r0 = rf(userId, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBalance provides a mock function with given fields: userId, currency
func (_m *UseCase) GetBalance(userId int64, currency string) (float32, error) {
	ret := _m.Called(userId, currency)

	var r0 float32
	if rf, ok := ret.Get(0).(func(int64, string) float32); ok {
		r0 = rf(userId, currency)
	} else {
		r0 = ret.Get(0).(float32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string) error); ok {
		r1 = rf(userId, currency)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferMoney provides a mock function with given fields: srcUserId, dstUserId, amount
func (_m *UseCase) TransferMoney(srcUserId int64, dstUserId int64, amount float32) error {
	ret := _m.Called(srcUserId, dstUserId, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, int64, float32) error); ok {
		r0 = rf(srcUserId, dstUserId, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
