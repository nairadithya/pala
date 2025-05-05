package api

import "errors"

type VoterService interface {
	New(request NewVoterRequest) error
	GetVoterInfo(voterID int) (Voter, error)
}

type VoterRepository interface {
	CreateVoter(w NewVoterRequest) error
	GetVoter(voterID int) (Voter, error)
}

type voterService struct {
	storage VoterRepository
}

func NewVoterService(voteRepo VoterRepository) VoterService {
	return &voterService{storage: voteRepo}
}

func (w *voterService) New(request NewVoterRequest) error {
	err := w.storage.CreateVoter(request)

	if err != nil {
		return err
	}

	return nil
}
