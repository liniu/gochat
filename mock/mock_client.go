// Code generated by MockGen. DO NOT EDIT.
// Source: wx/http.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	multipart "mime/multipart"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	wx "github.com/liniu/gochat/wx"
)

// MockUploadForm is a mock of UploadForm interface.
type MockUploadForm struct {
	ctrl     *gomock.Controller
	recorder *MockUploadFormMockRecorder
}

// MockUploadFormMockRecorder is the mock recorder for MockUploadForm.
type MockUploadFormMockRecorder struct {
	mock *MockUploadForm
}

// NewMockUploadForm creates a new mock instance.
func NewMockUploadForm(ctrl *gomock.Controller) *MockUploadForm {
	mock := &MockUploadForm{ctrl: ctrl}
	mock.recorder = &MockUploadFormMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploadForm) EXPECT() *MockUploadFormMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockUploadForm) Write(w *multipart.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", w)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockUploadFormMockRecorder) Write(w interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockUploadForm)(nil).Write), w)
}

// MockHTTPClient is a mock of HTTPClient interface.
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientMockRecorder
}

// MockHTTPClientMockRecorder is the mock recorder for MockHTTPClient.
type MockHTTPClientMockRecorder struct {
	mock *MockHTTPClient
}

// NewMockHTTPClient creates a new mock instance.
func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &MockHTTPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPClient) EXPECT() *MockHTTPClientMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockHTTPClient) Do(ctx context.Context, method, reqURL string, body []byte, options ...wx.HTTPOption) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, method, reqURL, body}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Do", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockHTTPClientMockRecorder) Do(ctx, method, reqURL, body interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, method, reqURL, body}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHTTPClient)(nil).Do), varargs...)
}

// Upload mocks base method.
func (m *MockHTTPClient) Upload(ctx context.Context, reqURL string, form wx.UploadForm, options ...wx.HTTPOption) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, reqURL, form}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Upload", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockHTTPClientMockRecorder) Upload(ctx, reqURL, form interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, reqURL, form}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockHTTPClient)(nil).Upload), varargs...)
}
