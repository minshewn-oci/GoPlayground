package mocks

import "github.com/stretchr/testify/mock"

type MockContext struct {
	mock.Mock
	Response any
}

func (m *MockContext) JSON(code int, object any) {
	m.Called(code, object)
	m.Response = object
}
