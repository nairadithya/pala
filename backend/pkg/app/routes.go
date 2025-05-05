package app

import "github.com/gin-gonic/gin"

func (s *Server) Routes() *gin.Engine {
	router := s.router

	v1 := router.Group("/v1/api")
	{
		v1.GET("/status", s.ApiStatus())
		user := v1.Group("/voter")
		{
			user.POST("", s.CreateVoter())
		}

		vote := v1.Group("/vote")
		{
			vote.POST("", s.CreateVote())
		}
	}

	return router
}
