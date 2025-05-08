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
			user.GET("/:id", s.GetVoterInfo())
			user.GET("/:id/voted", s.VoterHasVoted())
		}

		vote := v1.Group("/vote")
		{
			vote.POST("", s.CreateVote())
		}

		candidates := v1.Group("/candidates")
		{
			candidates.GET("", s.GetAllCandidateInfo())
		}

		parties := v1.Group("/parties")
		{
			parties.GET("", s.GetAllParties())
		}
	}

	return router
}
