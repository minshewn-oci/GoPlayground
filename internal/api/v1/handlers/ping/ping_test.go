package ping

import (
	"GoPlayground/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PingTestSuite struct {
	suite.Suite
}

func (suite *PingTestSuite) TestPingReturnsResponse() {
	mockContext := new(mocks.MockContext)
	mockContext.On("JSON", mock.Anything, mock.Anything)
	pingHandler := new(PingHandler)
	pingHandler.Process(mockContext)
	assert.Equal(suite.T(), PingResponse{Message: "pong"}, mockContext.Response)
}

func TestPingTestSuite(t *testing.T) {
	suite.Run(t, new(PingTestSuite))
}
