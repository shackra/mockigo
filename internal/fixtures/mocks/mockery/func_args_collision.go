// Code generated by mockigo. DO NOT EDIT.

package mockery

import match "github.com/subtle-byte/mockigo/match"
import mock "github.com/subtle-byte/mockigo/mock"

var _ = match.Any[int]

type FuncArgsCollision struct {
	mock *mock.Mock
}

func NewFuncArgsCollision(t mock.Testing) *FuncArgsCollision {
	t.Helper()
	return &FuncArgsCollision{mock: mock.NewMock(t)}
}

type _FuncArgsCollision_Expecter struct {
	mock *mock.Mock
}

func (_mock *FuncArgsCollision) EXPECT() _FuncArgsCollision_Expecter {
	 return _FuncArgsCollision_Expecter{mock: _mock.mock}
}

type _FuncArgsCollision_Foo_Call struct {
	*mock.Call
}

func (_mock *FuncArgsCollision) Foo(ret interface{}) (error) {
	_mock.mock.T.Helper()
	_results := _mock.mock.Called("Foo", ret)
	_r0 := _results.Error(0)
	return _r0
}

func (_expecter _FuncArgsCollision_Expecter) Foo(ret match.Arg[interface{}]) _FuncArgsCollision_Foo_Call {
	return _FuncArgsCollision_Foo_Call{Call: _expecter.mock.ExpectCall("Foo", ret.Arg)}
}

func (_call _FuncArgsCollision_Foo_Call) Return(_r0 error) _FuncArgsCollision_Foo_Call {
	_call.Call.Return(_r0)
	return _call
}

func (_call _FuncArgsCollision_Foo_Call) RunReturn(f func(ret interface{}) (error)) _FuncArgsCollision_Foo_Call {
	_call.Call.RunReturn(f)
	return _call
}
