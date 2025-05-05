package api

import (
	"errors"
	"strings"
)

// UserService contains the methods of the user service
type VoterService interface {
	New(user NewVoterRequest) error
}

type VoterRepository interface {
	CreateVoter(NewVoterRequest) error
}

type voterService struct {
	storage VoterRepository
}

func NewVoterService(voterRepo VoterRepository) VoterService {
	return &voterService{storage: voterRepo}
}

func (v *voterService) New(voter NewVoterRequest) error {
	if voter.FirstName == "" {
		return errors.New("voter service - first name required")
	}

	if voter.LastName == "" {
		return errors.New("voter service - last name required")
	}

	voter.FirstName = strings.ToLower(voter.FirstName)
	voter.LastName = strings.TrimSpace(voter.LastName)

	err := v.storage.CreateVoter(voter)

	if err != nil {
		return err
	}

	return nil
}
