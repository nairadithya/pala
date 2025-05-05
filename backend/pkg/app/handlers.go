package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pala/backend/pkg/api"
)

func (s *Server) ApiStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		response := map[string]string{
			"status": "success",
			"data":   "voter API running smoothly",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) CreateVoter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newVoter api.NewVoterRequest
		err := c.ShouldBindJSON(&newVoter)

		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.voterService.New(newVoter)

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "new voter created",
		}

		c.JSON(http.StatusOK, response)
	}
}

func (s *Server) CreateVote() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		var newVote api.NewVoteRequest
		err := c.ShouldBindJSON(&newVote)

		if err != nil {
			log.Printf("handler error: %v", err)
			c.JSON(http.StatusBadRequest, nil)
			return
		}

		err = s.newVoteService.New(newVote)

		if err != nil {
			log.Printf("service error: %v", err)
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		response := map[string]string{
			"status": "success",
			"data":   "new voter created",
		}

		c.JSON(http.StatusOK, response)
	}
}
