package api

type VoterService interface {
	New(request NewVoterRequest) (int, error)
	GetVoterInfo(voterID int) (Voter, error)
	HasVoted(voterID int) (bool, error)
}

type VoterRepository interface {
	CreateVoter(w NewVoterRequest) (int, error)
	GetVoter(voterID int) (Voter, error)
	HasVoted(voterID int) (bool, error)
}

type voterService struct {
	storage VoterRepository
}

func NewVoterService(voteRepo VoterRepository) VoterService {
	return &voterService{storage: voteRepo}
}

func (w *voterService) New(request NewVoterRequest) (int, error) {
	voterID, err := w.storage.CreateVoter(request)
	if err != nil {
		return 0, err
	}

	return voterID, nil
}

func (w *voterService) GetVoterInfo(voterID int) (Voter, error) {
	voter, err := w.storage.GetVoter(voterID)

	if err != nil {
		return Voter{}, err
	}

	return voter, nil
}

func (w *voterService) HasVoted(voterID int) (bool, error) {
	hasVoted, err := w.storage.HasVoted(voterID)
	if err != nil {
		return false, err
	}

	return hasVoted, nil
}
