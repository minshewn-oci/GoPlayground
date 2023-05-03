package router

import (
	"GoPlayground/internal/api/v1/handlers/ping"
	"GoPlayground/pkg/adapters"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MockGinWrapper struct {
	mock.Mock
	GroupPath    string
	EndpointPath string
	Handler      []gin.HandlerFunc
	basePath     string
}

type MockRouterGroup struct {
	mock.Mock
}

func (m *MockGinWrapper) Group(path string, handlers ...gin.HandlerFunc) adapters.IGinWrapper {
	m.Called(path)
	m.GroupPath = path
	m.Handler = handlers
	return m
}

func (m *MockGinWrapper) GET(path string, handlerFunc ...gin.HandlerFunc) {
	m.Called(path, handlerFunc)
	m.EndpointPath = path
	m.Handler = handlerFunc
}

type MockHandlerFactory struct {
	mock.Mock
	MockHandler ping.IHandler
}

func (m *MockHandlerFactory) Get(path string) ping.IHandler {
	m.Called(path)
	return m.MockHandler
}

type MockPingHandler struct {
	mock.Mock
}

func (m *MockPingHandler) Process(c ping.Context) {
	m.Called(c)
}

type RouterTestSuite struct {
	suite.Suite
}

func (suite *RouterTestSuite) TestRegisterGroupRegistersGroupAndEndpoints() {
	mockWrapper := new(MockGinWrapper)
	mockWrapper.On("Group", mock.Anything, mock.Anything)
	mockWrapper.On("GET", mock.Anything, mock.Anything)
	mockPingHandler := new(MockPingHandler)
	mockPingHandler.On("Process", mock.Anything)
	mockHandlerFactory := MockHandlerFactory{MockHandler: mockPingHandler}
	mockHandlerFactory.On("Get", mock.Anything)
	registrar := Registrar{HandlerFactory: &mockHandlerFactory, RouterGroup: mockWrapper}
	registrar.RegisterGroup()
	assert.Equal(suite.T(), "v1", mockWrapper.GroupPath)
	assert.Equal(suite.T(), "ping", mockWrapper.EndpointPath)
	assert.NotNil(suite.T(), mockWrapper.Handler)
	mockWrapper.Handler[0](nil)
	mockPingHandler.AssertCalled(suite.T(), "Process", mock.Anything)
}

func TestRouterTestSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}
