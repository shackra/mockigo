// Code generated by mockigo. DO NOT EDIT.

package mockery

import match "github.com/subtle-byte/mockigo/match"
import mock "github.com/subtle-byte/mockigo/mock"

var _ = match.Any[int]

type MyReader struct {
	mock *mock.Mock
}

func NewMyReader(t mock.Testing) *MyReader {
	t.Helper()
	return &MyReader{mock: mock.NewMock(t)}
}

type _MyReader_Expecter struct {
	mock *mock.Mock
}

func (_mock *MyReader) EXPECT() _MyReader_Expecter {
	 return _MyReader_Expecter{mock: _mock.mock}
}

type _MyReader_Read_Call struct {
	*mock.Call
}

func (_mock *MyReader) Read(p []byte) (n int, err error) {
	_mock.mock.T.Helper()
	_results := _mock.mock.Called("Read", p)
	_r0 := _results.Get(0).(int)
	_r1 := _results.Error(1)
	return _r0, _r1
}

func (_expecter _MyReader_Expecter) Read(p match.Arg[[]byte]) _MyReader_Read_Call {
	return _MyReader_Read_Call{Call: _expecter.mock.ExpectCall("Read", p.Arg)}
}

func (_call _MyReader_Read_Call) Return(n int, err error) _MyReader_Read_Call {
	_call.Call.Return(n, err)
	return _call
}

func (_call _MyReader_Read_Call) RunReturn(f func(p []byte) (n int, err error)) _MyReader_Read_Call {
	_call.Call.RunReturn(f)
	return _call
}
