package middlewareService

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MiddlewareMockService struct {
	mock.Mock
}

func (s *MiddlewareMockService) DeserializeUser() gin.HandlerFunc {
	ret := s.Called()
	return ret.Get(0).(gin.HandlerFunc)
}

func (s *MiddlewareMockService) CheckAdmin() gin.HandlerFunc {
	ret := s.Called()
	return ret.Get(0).(gin.HandlerFunc)
}
