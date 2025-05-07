package api

import "errors"

type PartyService interface {
	New(request NewPartyRequest) error
	GetPartyInfo(partyID int) (Party, error)
	GetAllParties() ([]Party, error)
}

type PartyRepository interface {
	CreateParty(request NewPartyRequest) error
	GetParty(partyID int) (Party, error)
	GetAllParties() ([]Party, error)
}

type partyService struct {
	storage PartyRepository
}

func NewPartyService(repo PartyRepository) PartyService {
	return &partyService{storage: repo}
}

func (p *partyService) New(request NewPartyRequest) error {
	if request.PartyName == "" {
		return errors.New("party: name is required")
	}

	return p.storage.CreateParty(request)
}

func (p *partyService) GetPartyInfo(partyID int) (Party, error) {
	return p.storage.GetParty(partyID)
}

func (p *partyService) GetAllParties() ([]Party, error) {
	return p.storage.GetAllParties()
}
