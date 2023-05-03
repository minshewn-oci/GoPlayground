package router

import (
	"GoPlayground/internal/api/v1/handlers/ping"
	"bufio"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net"
	"net/http"
	"testing"
)

type MockRouterGroup struct {
	mock.Mock
	GroupPath    string
	EndpointPath string
	Handler      []gin.HandlerFunc
	basePath     string
}

func (m *MockRouterGroup) Group(path string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	m.Called(path)
	m.GroupPath = path
	engine := gin.Default()
	return engine.Group("v1")
}

func (m *MockRouterGroup) GET(path string, handlerFunc ...gin.HandlerFunc) gin.IRoutes {
	m.Called(path, handlerFunc)
	m.EndpointPath = path
	m.Handler = handlerFunc
	return nil
}

func (m *MockRouterGroup) SetBasePath() {
	m.basePath = ""
}

type MockHandlerFactory struct {
	mock.Mock
}

func (m *MockHandlerFactory) Get(path string) ping.IHandler {
	m.Called(path)
	return new(MockPingHandler)
}

type MockPingHandler struct {
	mock.Mock
}

func (m *MockPingHandler) Process(c ping.Context) {
	m.Called(c)
}

type MockPusher struct {
	mock.Mock
}

func (m *MockPusher) Push(target string, opts *http.PushOptions) error {
	m.Called(target, opts)
	return nil
}

type MockResponseWriter struct {
	mock.Mock
}

func (m *MockResponseWriter) Header() http.Header {
	m.Called()
	return http.Header{}
}

func (m *MockResponseWriter) Write(bytes []byte) (int, error) {
	m.Called(bytes)
	return 0, nil
}

func (m *MockResponseWriter) WriteHeader(statusCode int) {
	m.Called(statusCode)
}

func (m *MockResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	m.Called()
	return nil, nil, nil
}

func (m *MockResponseWriter) Flush() {
	m.Called()
}

func (m *MockResponseWriter) CloseNotify() <-chan bool {
	m.Called()
	return nil
}

func (m *MockResponseWriter) Status() int {
	m.Called()
	return 0
}

func (m *MockResponseWriter) Size() int {
	m.Called()
	return 0
}

func (m *MockResponseWriter) WriteString(content string) (int, error) {
	m.Called()
	return 0, nil
}

func (m *MockResponseWriter) Written() bool {
	m.Called()
	return true
}

func (m *MockResponseWriter) WriteHeaderNow() {
	m.Called()
}

func (m *MockResponseWriter) Pusher() http.Pusher {
	m.Called()
	return new(MockPusher)
}

type RouterTestSuite struct {
	suite.Suite
}

func (suite *RouterTestSuite) TestRegisterGroupRegistersGroupAndEndpoints() {
	mockRouterGroup := new(MockRouterGroup)
	mockRouterGroup.SetBasePath()
	mockRouterGroup.On("Group", mock.Anything).Return(mockRouterGroup)
	mockRouterGroup.On("GET", mock.Anything, mock.Anything)
	mockPingHandler := new(MockPingHandler)
	mockPingHandler.On("Process", mock.Anything)
	mockHandlerFactory := new(MockHandlerFactory)
	mockHandlerFactory.On("Get", mock.Anything).Return(mockPingHandler)
	registrar := Registrar{HandlerFactory: mockHandlerFactory}
	registrar.RegisterGroup(mockRouterGroup)
	assert.Equal(suite.T(), "v1", mockRouterGroup.GroupPath)
	assert.Equal(suite.T(), "ping", mockRouterGroup.EndpointPath)
	assert.NotNil(suite.T(), mockRouterGroup.Handler)
	var handler []gin.HandlerFunc
	handler = mockRouterGroup.Handler
	context := new(gin.Context)
	mockResponseWriter := new(MockResponseWriter)
	mockResponseWriter.On("WriteHeader", mock.Anything)
	mockResponseWriter.On("Header")
	mockResponseWriter.On("Write", mock.Anything)
	context.Writer = mockResponseWriter
	handler[0](context)
	mockPingHandler.AssertCalled(suite.T(), "Ping", mock.Anything)
}

func TestRouterTestSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}
