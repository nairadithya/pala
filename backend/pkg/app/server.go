package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"pala/backend/pkg/api"
)

type Server struct {
	router       *gin.Engine
	voterService api.VoterService
	voteService  api.VoteService
}

func NewServer(router *gin.Engine, userService api.VoterService, weightService api.VoteService) *Server {
	return &Server{
		router:       router,
		voterService: userService,
		voteService:  weightService,
	}
}

func (s *Server) Run() error {
	r := s.Routes()

	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
